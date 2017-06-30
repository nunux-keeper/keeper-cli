package commands

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/admin"
	"github.com/nunux-keeper/keeper-cli/cmd/document"
	"github.com/nunux-keeper/keeper-cli/cmd/label"
	"github.com/nunux-keeper/keeper-cli/cmd/login"
	"github.com/nunux-keeper/keeper-cli/cmd/logout"
	"github.com/nunux-keeper/keeper-cli/cmd/profile"
	"github.com/nunux-keeper/keeper-cli/cmd/trash"
	"github.com/nunux-keeper/keeper-cli/cmd/version"
	"github.com/spf13/cobra"
)

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, kCli *cli.KeeperCLI) {
	cmd.AddCommand(
		// version
		version.NewCommand(kCli),

		// login
		login.NewCommand(kCli),

		// logout
		logout.NewCommand(kCli),

		// profile
		profile.NewCommand(kCli),

		// label
		label.NewCommand(kCli),

		// document
		document.NewCommand(kCli),

		// trash
		trash.NewCommand(kCli),

		// admin
		admin.NewCommand(kCli),
	)
}
