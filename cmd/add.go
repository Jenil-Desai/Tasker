package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Jenil-Desai/Tasker/options"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Deploy new mission",
	Long:  "Deploy a new mission to your task radar. Each task can include priority level and due date.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newTask, err := options.AddTask(args[0])
		if err != nil {
			color.Red("%v", err)
			return
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 5, ' ', 0)

		fmt.Fprintln(w, "ID\tTask\tCreated On\tStatus")
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", newTask.Id, newTask.Task, newTask.CreatedOn, newTask.Status)

		w.Flush()

		color.Green("\n🎯 Mission [%v] deployed successfully to command center", newTask.Id)
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}
