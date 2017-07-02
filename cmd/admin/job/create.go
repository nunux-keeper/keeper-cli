package job

import (
	"errors"
	"net/url"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/cli"
)

var jobTmpl = `Job details:
 ID {{.Id}}
`

type createJobOptions struct {
	params paramsFlag
}

type paramsFlag struct {
	params url.Values
}

func (pf *paramsFlag) String() string {
	return pf.params.Encode()
}

func (pf *paramsFlag) Set(value string) error {
	p := strings.Split(value, "=")
	if len(p) != 2 {
		return errors.New("Bad parameter definition")
	}
	if pf.params == nil {
		pf.params = url.Values{}
	}
	pf.params.Add(p[0], p[1])
	return nil
}

func (pf *paramsFlag) Type() string {
	return "params"
}

func (pf *paramsFlag) Get() url.Values {
	return pf.params
}

func newCreateCommand() *cobra.Command {
	var opts createJobOptions
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a job",
		RunE: func(cc *cobra.Command, args []string) error {
			return runCreateCommand(cc, &opts, args)
		},
	}

	flags := cmd.Flags()
	flags.VarP(&opts.params, "param", "p", "Job parameter (key=val)")
	return cmd
}

func runCreateCommand(cmd *cobra.Command, opts *createJobOptions, args []string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	if len(args) != 1 {
		return errors.New("Name required")
	}

	name := args[0]

	res, err := kli.API.CreateJob(name, opts.params.Get())
	if err != nil {
		return err
	}

	tmpl, err := template.New("job").Parse(jobTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, res)
	return err
}
