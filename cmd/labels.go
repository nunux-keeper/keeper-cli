package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// labelsCmd represents the labels command
var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "Get labels.",
	RunE:  labelsRun,
}

func init() {
	RootCmd.AddCommand(labelsCmd)
}

func labelsRun(cmd *cobra.Command, args []string) error {
	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	labels, err := kClient.GetLabels()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tLabel\tColor\tDate\t")
	for _, label := range labels {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", label.Id, label.Label, label.Color, label.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
