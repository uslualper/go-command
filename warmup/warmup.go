package warmup

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"go-command/pkg/warmup"
	"go-command/pkg/warmup/services"
)

var maxWorker int
var maxRequestInTime int
var maxRunTime int // second

var command = &cobra.Command{
	Use:     "warmup",
	Aliases: []string{"rev"},
	Short:   "Warmup",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run Warmup")

		siteMap, err := (&services.SiteMap{}).GetSiteMap(args[0])

		if err != nil {
			fmt.Println("command error: ", err, ", url:", args[0])
			return
		}

		fmt.Println("max-worker:", maxWorker)
		fmt.Println("max-request-in-time:", maxRequestInTime)
		fmt.Println("max-run-time:", maxRunTime)

		(&warmup.Sitemap{
			MaxWorker:        maxWorker,
			MaxRequestInTime: maxRequestInTime,
			MaxRunTime:       time.Second * time.Duration(maxRunTime),
		}).WarmUp(siteMap)
	},
}

func init() {

	command.Flags().IntVarP(&maxWorker, "max-worker", "w", 10, "Maximum number of workers")
	command.Flags().IntVarP(&maxRequestInTime, "max-request-in-time", "r", 600, "Maximum number of requests in time")
	command.Flags().IntVarP(&maxRunTime, "max-run-time", "s", 60, "The time limit for the specified limit in seconds")

	rootCmd.AddCommand(command)
}
