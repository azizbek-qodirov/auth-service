package main

import (
	"auth-service/api"
	"auth-service/api/handlers"
	"auth-service/config"
	"auth-service/config/logger"
	"auth-service/postgresql"
	"auth-service/service"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	cf := config.Load()
	logger := logger.NewLogger(basepath, cf.LOG_PATH)
	em := config.NewErrorManager(logger)

	conn, err := postgresql.ConnectDB(&cf)
	em.CheckErr(err)
	defer conn.Close()

	us := service.NewUserService()

	handler := handlers.NewHandler()
	api.NewRouter()

}
