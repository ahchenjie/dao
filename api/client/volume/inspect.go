package volume

import (
	"golang.org/x/net/context"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/api/client/inspect"
	"github.com/docker/docker/cli"
	"github.com/spf13/cobra"
)

type inspectOptions struct {
	format string
	names  []string
}

func newInspectCommand(dockerCli *client.DockerCli) *cobra.Command {
	var opts inspectOptions

	cmd := &cobra.Command{
		Use:   "inspect [OPTIONS] VOLUME [VOLUME...]",
		Short: "显示一个或多个存储卷的详细信息",
		Long:  inspectDescription,
		Args:  cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.names = args
			return runInspect(dockerCli, opts)
		},
	}

	cmd.Flags().StringVarP(&opts.format, "format", "f", "", "基于指定的Go语言模板格式化命令输出内容")

	return cmd
}

func runInspect(dockerCli *client.DockerCli, opts inspectOptions) error {
	client := dockerCli.Client()

	ctx := context.Background()

	getVolFunc := func(name string) (interface{}, []byte, error) {
		i, err := client.VolumeInspect(ctx, name)
		return i, nil, err
	}

	return inspect.Inspect(dockerCli.Out(), opts.names, opts.format, getVolFunc)
}

var inspectDescription = `
Returns information about one or more volumes. By default, this command renders
all results in a JSON array. You can specify an alternate format to execute a
given template is executed for each result. Go's https://golang.org/pkg/text/template/
package describes all the details of the format.

`
