package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Jenil-Desai/Tasker/options"
	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newTask, err := options.AddTask(args[0])
		if err != nil {
			return
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 1, ' ', 0)

		fmt.Fprintln(w, "ID\tTask\tCreated On\tStatus")
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", newTask.Id, newTask.Task, newTask.CreatedOn, newTask.Status)

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}
