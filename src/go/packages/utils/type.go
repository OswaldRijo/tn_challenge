package utils

type ServerStartUp struct {
	InitDatabase  bool
	RunMigrations bool
	Port          string
	RunMetrics    bool
	Debug         bool
}
