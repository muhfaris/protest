package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.sicepat.tech/muhfaris/protest/controller"
)

var (
	// Used for flags.
	url string

	rootCmd = &cobra.Command{
		Use:   "protest",
		Short: "producer testing",
		RunE: func(cmd *cobra.Command, args []string) error {

			// initialize
			// p := controller.NewProtest(url)
			// p.Parse()

			controller.Menu()

			return nil
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(screen)
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "the address of source data")
}
