package main

import (
	"fmt"
	router2 "github.com/dmitry-udod/codes_go/app/router"
	"github.com/dmitry-udod/codes_go/app/services"
	"github.com/dmitry-udod/codes_go/cmd"
	. "github.com/dmitry-udod/codes_go/logger"
	"os"
)


func main() {
	fmt.Println("Hello")
	InitLogger()
	services.InitElasticSearchClient()
	if ! isCliCommand() {
		router := router2.SetupRouter()
		router.Run()
	}
}

func isCliCommand() bool {
	if (len(os.Args) > 2) {
		if (os.Args[1] == `--import-fop` && os.Args[2] != "") {
			Log.Info("Run import FOP command")
			cmd.ImportFop(os.Args[2])
			return true
		}
	}

	return false
}