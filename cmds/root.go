package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCmd() {
	rootCmd := &cobra.Command{
		Use:   "hnfctl",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("---------........")
			visitDowaloads()
			watchDownloadFolder()
		},
	}
	rootCmd.AddCommand(NewDeleteCmd())
	rootCmd.Execute()
}
