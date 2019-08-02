package main

import (
	"fmt"
	"flag"
	"github.com/dmitry-udod/codes_go/app/router"
	"github.com/dmitry-udod/codes_go/app/services"
	"github.com/dmitry-udod/codes_go/cmd"
	. "github.com/dmitry-udod/codes_go/logger"
	"os"
)

var importFop = flag.String("import-fop", "", "Import FOPs from file")
var importLegalEntity = flag.String("import-legal-entity", "", "Import legal entities from file")
var importTerrorists = flag.String("import-terrorist", "", "Import terrorist list from file")
var generateSiteMap = flag.String("generate-site-map", "", "Generate sitemap files from file")

func main() {
	flag.Parse()

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

		if *importFop != "" {
			Log.Info("Run import FOP command")
			cmd.ImportFop(file)
			finishFileProcess(filePath)
			return true
		}

		if *importLegalEntity != "" {
			Log.Info("Run import legal entity command")
			cmd.ImportLegalEntity(file)
			finishFileProcess(filePath)
			return true
		}

		if *importTerrorists != "" {
			Log.Info("Run import terrorists command")
			cmd.ImportLegalEntity(file)
			finishFileProcess(filePath)
			return true
		}

		if *generateSiteMap != "" {
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