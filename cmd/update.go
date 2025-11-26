/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [task]",
	Args:  cobra.ExactArgs(2),
	Short: "Updating a task",
	Long:  `Updating a task: ctt update 1 "new task"`,
	Run: func(cmd *cobra.Command, args []string) {
		idToUpdate, err := strconv.Atoi(args[0])
		newTask := args[1]
		if err != nil {
			fmt.Println("Did you input a valid id?")
		} else {
			RetrieveAndUpdateTask(idToUpdate, newTask)
		}

	},
}

func RetrieveAndUpdateTask(id int, task string) {
	file, err := os.ReadFile("tasks.json")
	if err == nil {
		json.Unmarshal(file, &tasks)
		for index, item := range tasks {
			if id == item.ID {
				tasks[index].DESCRIPTION = task
				tasks[index].UPDATEDAT = time.Now().Format("2006-01-02 15:04:05")
			}
		}
		file, err = json.MarshalIndent(tasks, "", " ")
		if err == nil {
			os.WriteFile("tasks.json", file, 0644)
			fmt.Printf("Task %d updated", id)
		} else {
			fmt.Println("Updating task failed, exiting...")
		}
	} else {
		fmt.Println("There are no tasks to update, exiting...")
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
