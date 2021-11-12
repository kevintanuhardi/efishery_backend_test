package cmd

import (
	"fmt"

	"github.com/kevintanuhardi/efishery_backend_test/ports/rest"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Running rest service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rest.Application(
			&rest.Config{
				Cfg:         cfg,
				// GormStarter: adapter.NewGormStarter(),
			},
		); err != nil {
			panic(fmt.Errorf("cannot start rest server: %w", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}