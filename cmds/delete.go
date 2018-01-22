package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Delete Command...")
			getTheInput()
		},
	}
}

func getTheInput() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(s)

	
}
