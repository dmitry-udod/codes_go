package main

import (
	"fmt"
	"github.com/dmitry-udod/codes_go/cmd"
	. "github.com/dmitry-udod/codes_go/logger"
	"github.com/dmitry-udod/codes_go/services"
	"os"
)


func main() {
	fmt.Println("Hello")
	InitLogger()
	services.InitElasticSearchClient()
	checkArgs()
}

func checkArgs() {
	if (len(os.Args) > 2) {
		if (os.Args[1] == `--import-fop` && os.Args[2] != "") {
			Log.Info("Run import FOP command")
			cmd.ImportFop(os.Args[2])
		}
	}
}