package gerar

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/urfave/cli/v2"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando para gerar converter CSV em JSON"
	app.Usage = "Converter CSV em JSON"

	app.Commands = []*cli.Command{
		{
			Name:  "csv",
			Usage: "Converter CSV em JSON",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "origin",
					Value: "",
				},
				&cli.StringFlag{
					Name:  "output",
					Value: "",
				},
			},
			Action: convertCsvToJson,
		},
	}

	return app
}

func convertCsvToJson(c *cli.Context) error {
	origin := c.String("origin")
	if origin == "" {
		log.Fatal("É necessário informar o caminho do arquivo CSV")
	}

	_, err := os.Stat(origin)
	if err != nil {
		log.Fatal("Arquivo não encontrado")
	}

	csvFile, _ := os.Open(origin)
	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	json := gabs.New()

	for idxRec, record := range records {
		if idxRec > 0 {
			newObj := gabs.New()
			for idx, parameter := range records[0] {
				newObj.Set(record[idx], parameter)
			}
			json.ArrayAppend(newObj.Data(), "data")
		}
	}

	output := c.String("output")
	jsonFile, err := os.Create(output)
	jsonFile.Write(json.BytesIndent("", "  "))
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	return nil
}
