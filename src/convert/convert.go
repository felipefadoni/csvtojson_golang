package convert

import (
	"github.com/urfave/cli/v2"
)

func Convert() *cli.App {
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
			Action: ConvertCsvToJson,
		},
		{
			Name:  "json",
			Usage: "Converter JSON em CSV",
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
			Action: ConvertJsonToCsv,
		},
	}

	return app
}
