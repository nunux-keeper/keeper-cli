package cmd

import (
	"errors"
	"io"
	"net/url"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
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

func NewCmdCreateJob(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var opts createJobOptions
	cmd := &cobra.Command{
		Use:   "job <name> ",
		Short: "Create a job (ADMIN)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateJob(f, out, cmd, &opts, args)
		},
	}

	flags := cmd.Flags()
	flags.VarP(&opts.params, "param", "p", "Job parameter (key=val)")
	return cmd
}

func runCreateJob(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, opts *createJobOptions, args []string) error {
	if len(args) != 1 {
		return errors.New("Name required")
	}

	name := args[0]

	c, err := f.Client()
	if err != nil {
		return err
	}

	res, err := c.CreateJob(name, opts.params.Get())
	if err != nil {
		return err
	}

	tmpl, err := template.New("job").Parse(jobTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, res)
	return err
}
