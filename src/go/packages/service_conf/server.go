package service_conf

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	database2 "truenorth/packages/database"
	"truenorth/packages/logger"
	"truenorth/packages/utils"
)

var (
	log           = logger.GetLogger()
	IgnoredTraces = []string{
		"grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo",
		"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo",
		"grpc.health.v1.Health/Check",
		"/grpc.health.v1.Health/Check",
	}
)

func serverInterceptor(s *utils.ServerStartUp) func(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		appName := os.Getenv("APP_NAME")
		newCtx := context.WithValue(ctx, "appName", appName)
		log.Infow(newCtx, "Request started",
			"Method requested:", info.FullMethod,
			"Request: ", req,
			"Started at", start.UTC().String())

		h, err := handler(newCtx, req)

		log.Infow(newCtx, "Request ended ",
			"Method: ", info.FullMethod,
			"Duration: ", time.Since(start),
			"Error: ", err)

		return h, err
	}
}

func withServerUnaryInterceptor(s *utils.ServerStartUp) grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor(s))
}

func getServerOpts(s *utils.ServerStartUp) []grpc.ServerOption {
	opts := []grpc.ServerOption{
		withServerUnaryInterceptor(s),
	}

	return opts
}

type FuncServer = func(s *utils.ServerStartUp)

func WithDatabase() FuncServer {
	return func(s *utils.ServerStartUp) {
		s.InitDatabase = true
	}
}

func WithMigrations() FuncServer {
	return func(s *utils.ServerStartUp) {
		s.RunMigrations = true
	}
}
func WithPort(port string) FuncServer {
	return func(s *utils.ServerStartUp) {
		s.Port = port
	}
}

func InitServer(fs ...FuncServer) (net.Listener, *grpc.Server) {
	ctx := context.TODO()
	debug := os.Getenv("DEBUG") == "true"
	ls := GetLocalServerInstance()
	ls.Debug = debug
	for _, f := range fs {
		f(ls)
	}
	if ls.InitDatabase {
		err := database2.Init()
		if err != nil {
			log.Fatal(ctx, "cannot connect to db: ", err)
		}
	}

	if ls.InitDatabase && ls.RunMigrations {
		database2.RunMigrations(ctx)
	} else if !ls.InitDatabase && ls.RunMigrations {
		log.Fatal(ctx, "Run WithDatabase for running migrations")
	}

	opts := getServerOpts(ls)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", ls.Port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(opts...)

	return listener, s
}
