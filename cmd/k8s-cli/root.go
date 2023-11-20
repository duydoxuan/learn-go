package k8scli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() {
	// root command for the plugin
	rootCmd := &cobra.Command{
		Use:   "kubectl",
		Short: "A sample kubectl plugin that prints 'hello-world' when called with 'duydo'",
	}
	// funtion handle 'duydo' command
	handleddoCmd := func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] != "duydo" {
			fmt.Println("invalid argument.")
			return
		}
		fmt.Println("hello world")
	}
	// create a subcommand
	ddoCmd := &cobra.Command{
		Use:   "duydo",
		Short: "Print 'hello-world'",
		Run:   handleddoCmd,
	}
	// add subcom to rootcom
	rootCmd.AddCommand(ddoCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing  plugin", err)
	}
}
