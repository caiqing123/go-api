package main

import (
	"api/commands"
	_ "api/configor"
	_ "api/di"
	_ "api/dotenv"
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
)

// @title api
// @version 1.0
// @description api desc
// @host 127.0.0.1:8022
// @BasePath /pay/
func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
