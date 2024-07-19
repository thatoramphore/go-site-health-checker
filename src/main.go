package main

import  (
	"fmt"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "HealthChecker",
		Usage: "A tool that checks if a website is running or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aleases: []string{"d"}
				Usage: "Domain name to check.",
				Required: true,
			},
			&clie.StringFlag{
				Name: "port",
				Aliases: []string{"p"}
				Usage: "Port number to check.",
				Required: false,
			}
		},
		Action: func (c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}