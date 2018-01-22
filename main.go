package main

import (
	"fmt"

	"github.com/a8uhnf/dir-cleanup/cmds"
)

func main() {
	fmt.Println("Started main function.")

	// cmds.NewDeleteCmd()
	cmds.RootCmd()
}
