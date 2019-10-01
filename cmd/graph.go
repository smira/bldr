package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/talos-systems/bldr/internal/pkg/solver"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Graph dependencies between pkgs",
	Long: `This command outputs 'dot' formatted DAG of dependencies
starting from target to all the dependencies.
`,
	Run: func(cmd *cobra.Command, args []string) {
		loader := solver.FilesystemPackageLoader{
			Root:    pkgRoot,
			Context: options.GetVariables(),
		}
		packages, err := solver.NewPackages(&loader)
		if err != nil {
			log.Fatal(err)
		}
		graph, err := packages.Resolve(options.Target)
		if err != nil {
			log.Fatal(err)
		}
		graph.DumpDot(os.Stdout)
	},
}

func init() {
	graphCmd.Flags().StringVarP(&options.Target, "target", "t", "", "Target image to build")
	graphCmd.MarkFlagRequired("target")
	rootCmd.AddCommand(graphCmd)
}
