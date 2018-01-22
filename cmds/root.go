package cmds

import (
	"github.com/spf13/cobra"
)

func RootCmd() {
	rootCmd := &cobra.Command{
		Use:   "hnfctl",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("---------")
		},
	}
	rootCmd.Execute()
}
