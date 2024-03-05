package warmup

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "warmup",
	Short: "Warmup is a CLI",
	Long:  `Warmup is a CLI application that can be used to warmup your servers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run Warmup")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
