package main

import (
	"os"
	"time"

	gerar "github.com/felipefadoni/csvtojson_golang/src/convert"
	log "github.com/sirupsen/logrus"
)

func main() {
	startTime := time.Now()
	log.Info("Start CLI! - ", startTime)

	application := gerar.Convert()

	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	endTime := time.Now()
	timeSince := time.Since(startTime)

	log.Info("Finished! - ", endTime, " - ", timeSince)
}
