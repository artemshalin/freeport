package cmd

import (
	"errors"
	"log/slog"
	"os"

	"github.com/artemshalin/freeport/internal/app/core"

	"github.com/spf13/cobra"
)

var fromPort int
var toPort int
var useIPv6 bool

//nolint:gochecknoglobals // This is Cobra default pattern.
var rootCmd = &cobra.Command{
	Use:   "freeport",
	Short: "CLI utility that prints the first available port to stdout",
	Long: `CLI utility that prints the first available port to stdout. 
It is possible to specify a range that the utility will check.`,
	Run: func(_ *cobra.Command, _ []string) {
		addrPrefix := "127.0.0.1"
		if useIPv6 {
			addrPrefix = "::1"
		}

		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
		}))
		slog.SetDefault(logger)

		p, err := core.CheckRange(fromPort, toPort, addrPrefix)
		if err != nil {
			if errors.Is(err, core.ErrNoFreePort) {
				logger.Info(core.ErrNoFreePort.Error())
				return
			}

			logger.Error(err.Error())
			return
		}

		print(p)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&fromPort, "from", "f", 3000, "start scan with this port")
	rootCmd.PersistentFlags().IntVarP(&toPort, "to", "t", 8000, "end scan with this port")
	rootCmd.PersistentFlags().BoolVar(&useIPv6, "ipv6", false, "use ip v6")
}
