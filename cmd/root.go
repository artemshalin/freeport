package cmd

import (
	"log/slog"
	"os"

	"github.com/artemshalin/freeport/internal/app/core"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals // This is Cobra default pattern.
var rootCmd = &cobra.Command{
	Use:   "freeport",
	Short: "app retrieving free port on host",
	Long:  `App retrieving free port on host`,
	Run: func(_ *cobra.Command, _ []string) {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		slog.SetDefault(logger)

		p, err := core.GetFreePort()
		if err != nil {
			logger.Error("an error occurred while trying to get free port", slog.String("wrapped error", err.Error()))
		}

		logger.Info("done", slog.Int("freeport is", p))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
