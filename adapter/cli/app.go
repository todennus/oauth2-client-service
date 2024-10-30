package cli

import (
	"github.com/spf13/cobra"
	"github.com/todennus/oauth2-client-service/adapter/cli/seed"
)

var Command = &cobra.Command{
	Use:   "cli",
	Short: "The Todennus OAuth2 Client CLI",
}

func init() {
	Command.AddCommand(seed.Command)
}
