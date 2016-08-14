package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// documentsCmd represents the documents command
var documentsCmd = &cobra.Command{
	Use:   "documents",
	Short: "Get documents.",
	RunE:  documentsRun,
}

func documentsRun(cmd *cobra.Command, args []string) error {
	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	documents, err := kClient.GetDocuments()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tTitle\tContent-Type\tOrigin\tDate\t")
	for _, doc := range documents {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", doc.Id, doc.Title, doc.ContentType, doc.Origin, doc.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil

}

func init() {
	RootCmd.AddCommand(documentsCmd)
}
