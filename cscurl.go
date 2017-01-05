package main

import (
	"fmt"
	cs "github.com/cloudshare/go-sdk/cloudshare"
	"github.com/urfave/cli"
	neturl "net/url"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "method, m",
			Value: "get",
			Usage: "HTTP method (get|post|put|delete)",
		},
		cli.StringFlag{
			Name:   "api-key",
			Value:  "",
			Usage:  "CloudShare API key",
			EnvVar: "CLOUDSHARE_API_KEY",
		},
		cli.StringFlag{
			Name:   "api-id",
			Value:  "",
			Usage:  "CloudShare API ID",
			EnvVar: "CLOUDSHARE_API_ID",
		},
		cli.StringFlag{
			Name:  "data, d",
			Value: "",
			Usage: "JSON document",
		},
	}

	app.Action = func(c *cli.Context) error {
		apiKey := c.String("api-key")
		apiID := c.String("api-id")
		if apiKey == "" {
			return fmt.Errorf("api-key must be set")
		}

		if apiID == "" {
			return fmt.Errorf("api-id must be set")
		}

		if c.NArg() < 0 {
			cli.ShowAppHelp(c)
			return fmt.Errorf("Expecting URL argument")
		}

		url := c.Args().Get(0)

		method := c.String("method")

		client := &cs.Client{
			APIKey: apiKey,
			APIID:  apiID,
		}

		data := c.String("data")
		parsed, err := neturl.Parse(url)
		if err != nil {
			return err
		}
		query := parsed.Query()
		path := strings.Replace(parsed.Path, "api/v3/", "", 1)

		response, err := client.Request(method, path, &query, &data)
		if err != nil {
			return err

		}

		fmt.Println(string(response.Body))
		return nil
	}

	app.Run(os.Args)
}