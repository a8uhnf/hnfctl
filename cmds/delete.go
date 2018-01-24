package cmds

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	wlog "gopkg.in/dixonwille/wlog.v2"

"github.com/daviddengcn/go-colortext"
"github.com/dixonwille/wmenu"
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

	menu := wmenu.NewMenu("What is your favorite food?")
	menu.AddColor(wlog.Color{Code: ct.Green}, wlog.Color{Code: ct.Yellow}, wlog.Color{Code: ct.Magenta}, wlog.Color{Code: ct.Yellow})
	menu.Action(func(opts []wmenu.Opt) error {
		fmt.Println(opts[0].Text + " is your favorite food.")
		return nil
	})
	menu.Option("Pizza", nil, true, nil)
	menu.Option("Ice Cream", nil, false, nil)
	menu.Option("Tacos", nil, false, func(o wmenu.Opt) error {
		fmt.Println("Tacos are great")
		return nil
	})
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
