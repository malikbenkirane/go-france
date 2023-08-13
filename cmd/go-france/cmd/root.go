package cmd

import "github.com/spf13/cobra"

var (
	flagFirstnamesFile string
	flagLastnamesFile  string
	flagGenOutput      string
)

func init() {
	cmdRoot.AddCommand(cmdGenerate)
	cmdGenerate.Flags().StringVarP(&flagFirstnamesFile, "firstnames", "f", "data/prenoms_sorted_by_count.csv", "firstnames csv file")
	cmdGenerate.Flags().StringVarP(&flagLastnamesFile, "lastnames", "l", "data/noms_sorted_by_count.csv", "lastnames csv file")
	cmdGenerate.Flags().StringVarP(&flagGenOutput, "output", "o", "names.gen.go", "generated go file destination")
}

func Execute() error {
	return cmdRoot.Execute()
}

var cmdRoot = &cobra.Command{
	Use:   "go-france",
	Short: "go-france toolkit",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}
