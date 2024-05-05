package main

import (
	"os"
	"time"
	"vjshi/parser"

	"github.com/urfave/cli/v2"
)

func init() {
	cli.HelpFlag = &cli.BoolFlag{Name: "Help"}
	cli.VersionFlag = &cli.BoolFlag{Name: "print-version"}
}

func main() {
	app := cli.App{
		Name:     "vjshi Sales Analysis",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Commands: []*cli.Command{
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "fetch page once, for test using.",
				Action: func(ctx *cli.Context) error {
					return parser.GrabRecentSales()
				},
			},
		},
	}
	app.Run(os.Args)

	// s, err := gocron.NewScheduler()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = s.NewJob(gocron.DurationJob(5*time.Second), gocron.NewTask(func() {
	// 	errTask := parser.GrabRecentSales()
	// 	fmt.Println(errTask)
	// }))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// s.Start()

	// // block until you are ready to shut down
	// done := make(chan os.Signal, 1)
	// signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	// fmt.Println("Press ctrl+c to continue..")
	// <-done
}
