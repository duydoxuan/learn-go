package k8scli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli tool",
	Short: "The app executes kubectl command line",
	Long:  "The app executes kubectl command line for print hello-world and get pod",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 && args[0] == "kubectl" {
			switch args[1] {
			case "get":
				if len(args) > 2 {
					execKubectl(args)
				}
			case "duydo":
				fmt.Println("Hello World")
			default:
				fmt.Println("command not found")
			}
		} else {
			fmt.Println("incorrect input")
		}
	},
}

func execKubectl(args []string) {
	cmdArgs := args[1:]
	cmd := exec.Command("kubectl", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("error executing kubectl %v\n", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
