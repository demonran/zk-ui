package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"zk-ui/server"
)

func main() {

	var rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd, args)
			server.NewService()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
