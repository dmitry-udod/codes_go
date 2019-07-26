package main

import (
	"fmt"
	"github.com/dmitry-udod/codes_go/app/router"
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
		r := router.SetupRouter()
		r.Run()
	}
}

func isCliCommand() bool {
	if len(os.Args) > 2 {
		filePath := os.Args[2]
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			msg := fmt.Sprintf("File %s NOT found", filePath)
			fmt.Println(msg)
			Log.Fatal(msg);
		}

		file, err := os.Open(filePath)
		if err != nil {
			Log.Fatal(err)
		}
		defer file.Close()

		Log.Info("Start processing file: " + filePath);

		if os.Args[1] == `--import-fop` {
			Log.Info("Run import FOP command")
			cmd.ImportFop(file)
			finishFileProcess(filePath)
			return true
		}

		if os.Args[1] == `--import-legal-entity` {
			Log.Info("Run import legal entity command")
			cmd.ImportLegalEntity(file)
			finishFileProcess(filePath)
			return true
		}

		if os.Args[1] == "--generate-site-map" {
			Log.Info("Start site map generation")
			cmd.GenerateSiteMap(file)
			Log.Info("Finish site map generation")
			return true
		}
	}

	return false
}

func finishFileProcess(filePath string) {
	Log.Info("Finish processing file: " + filePath);
}