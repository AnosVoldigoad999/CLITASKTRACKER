package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Marking a task as done",
	Long:  `Marking a task as done: ctt mark-done 1`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0]) //convert to int
		if err == nil {
			RetrieveAndUpdateDone(id)
		} else {
			fmt.Println("Something went wrong, exiting...")
		}
	},
}

func RetrieveAndUpdateDone(id int) {
	file, err := os.ReadFile("tasks.json")
	isIn := false
	if err == nil {
		json.Unmarshal(file, &tasks)
		for index, item := range tasks {
			if id == item.ID {
				tasks[index].STATUS = "done"
				tasks[index].UPDATEDAT = time.Now().Format("2006-01-02 15:04:05")
				isIn = true
				break
			}
		}
		file, err = json.MarshalIndent(tasks, "", " ")
		if err == nil {
			os.WriteFile("tasks.json", file, 0644)
			if isIn {
				fmt.Printf("Status of Task %d updated to 'done'", id)
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
	rootCmd.AddCommand(doneCmd)

}
