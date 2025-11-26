package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [filter]",
	Args:  cobra.MaximumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			RetrieveAndListTasks(args[0])
		} else {
			RetrieveAndListTasks("")
		}
	},
}

func RetrieveAndListTasks(filter string) {
	file, err := os.ReadFile("tasks.json")
	if err == nil {
		json.Unmarshal(file, &tasks)

		if filter == "" {
			for index, item := range tasks {
				fmt.Printf("Id-%d %s \n", index, item.DESCRIPTION)
			}
		} else {
			for index, item := range tasks {
				if tasks[index].STATUS == filter {
					fmt.Printf("Id-%d %s \n", index, item.DESCRIPTION)
				}
			}
		}

	} else {
		fmt.Println("There are no tasks to Print")
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
