package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Directorio donde buscar archivos
	dir, err := os.Getwd()
	workingDir := filepath.Join(dir, "pb")
	// Recorrer los archivos en el directorio
	err = filepath.Walk(workingDir, walkingFunction)
	if err != nil {
		fmt.Println("Error al recorrer el directorio:", err)
	}
}

func makeMock(filePath string) {
	// Abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crear un nuevo conjunto de tokens
	fs := token.NewFileSet()

	// Analizar el archivo
	f, err := parser.ParseFile(fs, filePath, file, parser.ParseComments)
	if err != nil {
		fmt.Println("Error al analizar el archivo:", err)
		return
	}

	err = generateTemplate(f, filePath)

	if err != nil {
		return
	}
}

func generateTemplate(f *ast.File, path string) error {
	newPath := ""
	packageAddress := ""
	dirsSplit := strings.Split(path, "/")
	for i, d := range dirsSplit {
		if d == "go" {
			newPath = strings.Join(dirsSplit[i+1:len(dirsSplit)-1], "/")
			packageAddress = strings.Join(dirsSplit[i+1:len(dirsSplit)-1], "/")
		}

	}
	fileName := dirsSplit[len(dirsSplit)-1]
	baseDir, _ := os.Getwd()
	nfilePath := filepath.Join(baseDir, "mocks", newPath)
	packageToImport := fmt.Sprintf("%s/%s", "truenorth", packageAddress)
	err := os.MkdirAll(nfilePath, os.ModePerm)
	if err != nil {
		return err
	}

	implementationFile, err := os.Create(filepath.Join(nfilePath, fileName))
	if err != nil {
		return err
	}
	defer implementationFile.Close()
	err = setUpPackage(f, implementationFile)
	if err != nil {
		return err
	}
	err = setUpImports(f, implementationFile, packageToImport)
	if err != nil {
		return err
	}
	// Recorrer las declaraciones en el archivo
	for _, decl := range f.Decls {
		// Verificar si es una declaración de interfaz
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.InterfaceType); ok {
						// Si es una declaración de interfaz, generar implementación
						interfaceName := typeSpec.Name.Name
						interfaceType := typeSpec.Type.(*ast.InterfaceType)
						if strings.Contains(interfaceName, "Client") && !strings.Contains(interfaceName, "Unsafe") {
							generateImplementation(implementationFile, interfaceName, interfaceType.Methods.List, f.Name.Name)
						}
					}
				}
			}
		}
	}
	fmt.Printf("Implementación generada en %s\n", fileName)
	return nil
}

func setUpPackage(f *ast.File, implementationFile *os.File) error {
	_, err := implementationFile.WriteString(fmt.Sprintf("package %s\n\n", f.Name.Name))
	return err
}

func setUpImports(f *ast.File, implementationFile *os.File, packageToImport string) error {
	_, err := implementationFile.WriteString("import (\n")
	if err != nil {
		return err
	}
	_, err = implementationFile.WriteString("\t\"context\"\n\n")
	if err != nil {
		return err
	}
	_, err = implementationFile.WriteString("\t\"google.golang.org/grpc\"\n\n")
	_, err = implementationFile.WriteString("\t\"github.com/stretchr/testify/mock\"\n\n")
	if err != nil {
		return err
	}
	_, err = implementationFile.WriteString(fmt.Sprintf("\t%s \"%s\"\n", f.Name.Name, packageToImport))
	if err != nil {
		return err
	}
	_, err = implementationFile.WriteString(")\n\n")
	if err != nil {
		return err
	}
	return nil
}

func generateImplementation(implementationFile *os.File, interfaceName string, methods []*ast.Field, importFile string) {
	// Escribir estructura de implementación
	implementationFile.WriteString(fmt.Sprintf("type %sImpl struct {\n\tmock.Mock\n}\n\n", interfaceName))
	reqType := ""
	resType := ""
	// Recorrer los métodos de la interfaz y generar implementación
	for _, method := range methods {
		funcType, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		methodName := method.Names[0].Name
		if !strings.Contains(methodName, "mustEmbed") {

			implementationFile.WriteString(fmt.Sprintf("func (impl *%sImpl) %s(", interfaceName, methodName))
			for i, param := range funcType.Params.List {
				if i > 0 {
					implementationFile.WriteString(", ")
				}

				if _, ok := param.Type.(*ast.SelectorExpr); ok {
					implementationFile.WriteString(fmt.Sprintf("ctx context.Context"))
					continue
				}
				if sel, ok := param.Type.(*ast.StarExpr); ok {
					pkg := sel.X.(*ast.Ident)
					reqType = fmt.Sprintf("*%s.%s", importFile, pkg.Name)
					implementationFile.WriteString(fmt.Sprintf("req %s", reqType))
				}
			}
			implementationFile.WriteString(fmt.Sprintf("opts ...grpc.CallOption"))
			implementationFile.WriteString(") ")
			if funcType.Results != nil {
				implementationFile.WriteString("(")
				for _, result := range funcType.Results.List {
					if _, ok := result.Type.(*ast.Ident); ok {
						implementationFile.WriteString("error")
					}
					if sel, ok := result.Type.(*ast.StarExpr); ok {
						pkg := sel.X.(*ast.Ident)
						resType = fmt.Sprintf("*%s.%s", importFile, pkg.Name)
						implementationFile.WriteString(fmt.Sprintf("%s, ", resType))
					}
				}
				implementationFile.WriteString(")")
			}
			implementationFile.WriteString(
				fmt.Sprintf(
					`{
	ret := impl.Called(ctx, req)
	if len(ret) == 0 {
		panic("no return value specified for %s")
	}
	var r0 %s
	if rf, ok := ret.Get(0).(func(context.Context, %s) (%s, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, %s) %s); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(%s)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, %s) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

`, methodName, resType, reqType, resType, reqType, resType, resType, reqType))
		}
	}
}

func walkingFunction(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Error al acceder al directorio:", err)
		return nil
	}
	if !info.IsDir() && filepath.Ext(path) == ".go" {
		// Procesar archivo si es un archivo Go
		fmt.Println("Procesando archivo:", path)
		if strings.Contains(path, "grpc") {
			makeMock(path)
		}
	} else if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range entries {
			newWorkingDir := filepath.Join(path, e.Name())
			err = filepath.Walk(newWorkingDir, walkingFunction)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
