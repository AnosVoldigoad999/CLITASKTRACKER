package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Args:  cobra.ExactArgs(1),
	Short: "deleting a task",
	Long:  `deleting a task: ctt delete 1`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0]) //convert id passed to int
		if err == nil {
			RetrieveAndDelete(id)
		} else {
			fmt.Println("Kindly provide a valid id")
		}
	},
}

func RetrieveAndDelete(id int) {
	file, err := os.ReadFile("tasks.json")
	count := 0 //to update the id of the appended items
	var updated []Task
	if err == nil {
		json.Unmarshal(file, &tasks)
		for _, item := range tasks {
			if item.ID != id {
				newItem := Task{
					ID:          count, //update the id based on the order of entry
					DESCRIPTION: item.DESCRIPTION,
					STATUS:      item.STATUS,
					CREATEDAT:   item.CREATEDAT,
					UPDATEDAT:   item.UPDATEDAT,
				}
				updated = append(updated, newItem) //append
				count++                            //increase count to match the previous item's id
			}

		}
		file, err := json.MarshalIndent(updated, "", " ")
		if err == nil {
			os.WriteFile("tasks.json", file, 0644)
			fmt.Printf("Task %d deleted", id)
		} else {
			fmt.Println("Something bad happened, exiting...")
		}
	} else {
		fmt.Println("File not found")
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
