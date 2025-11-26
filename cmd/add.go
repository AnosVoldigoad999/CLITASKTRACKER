package cmd

import (
	"fmt"

	"encoding/json"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		RetrieveAndCallAddTask(args[0])
	},
}

type Task struct {
	ID          int    `json:"id"`
	DESCRIPTION string `json:"description"`
	STATUS      string `json:"status"`
	CREATEDAT   string `json:"createdAt"`
	UPDATEDAT   string `json:"updatedAt"`
}

var tasks []Task

func RetrieveAndCallAddTask(task string) {
	file, err := os.ReadFile("tasks.json")
	if err == nil {
		json.Unmarshal(file, &tasks)
	} else {
		fmt.Println("The tasks file is missing, creating one...")
	}
	AddTask(task, len(tasks))
}

func AddTask(task string, id int) {

	tasks = append(tasks, Task{ID: id,
		DESCRIPTION: task,
		STATUS:      "todo",
		CREATEDAT:   time.Now().Format("2006-01-02 15:04:05"),
		UPDATEDAT:   time.Now().Format("2006-01-02 15:04:05"),
	})

	updated, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("MAYDAY MAYDAY", err)
	} else {
		fmt.Println("Task Added successfully with id:", id)
	}
	os.WriteFile("tasks.json", updated, 0644)

}

func init() {
	rootCmd.AddCommand(addCmd)
}
