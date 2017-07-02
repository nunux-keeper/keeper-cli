package commands

import (
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
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		// version
		version.NewCommand(),

		// login
		login.NewCommand(),

		// logout
		logout.NewCommand(),

		// profile
		profile.NewCommand(),

		// label
		label.NewCommand(),

		// document
		document.NewCommand(),

		// trash
		trash.NewCommand(),

		// admin
		admin.NewCommand(),
	)
}
