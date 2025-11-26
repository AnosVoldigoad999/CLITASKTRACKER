package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// progressCmd represents the progress command
var progressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err == nil {
			RetrieveAndUpdateInProgress(id)
		} else {
			fmt.Println("Something went wrong, exiting...")
		}
	},
}

func RetrieveAndUpdateInProgress(id int) {
	file, err := os.ReadFile("tasks.json")
	isIn := false
	if err == nil {
		json.Unmarshal(file, &tasks)
		//taskToUpdate := tasks[id]
		for index, item := range tasks {
			if id == item.ID {
				tasks[index].STATUS = "in-progress"
				tasks[index].UPDATEDAT = time.Now().Format("2006-01-02 15:04:05")
				isIn = true
			}
		}
		file, err = json.MarshalIndent(tasks, "", " ")
		if err == nil {
			os.WriteFile("tasks.json", file, 0644)
			if isIn {
				fmt.Printf("Status of Task %d updated to 'in-progress'", id)
			} else {
				fmt.Println("Task not found, exiting...")
			}
		} else {
			fmt.Println("Updating task failed, exiting...")
		}
	} else {
		fmt.Println("There are no tasks to update, exiting...")
	}
}

func init() {
	rootCmd.AddCommand(progressCmd)
}
