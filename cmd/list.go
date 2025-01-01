package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Jenil-Desai/Tasker/options"
	"github.com/spf13/cobra"
)

var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "Scan active missions",
	Long:  "Survey your battlefield - displays all active missions with their status, priority, and timeline. Filter options available",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		allTasks, err := options.ListTask()
		if err != nil {
			return
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 5, ' ', 0)

		fmt.Fprintln(w, "ID\tTask\tCreated On\tStatus")
		for _, task := range allTasks {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", task.Id, task.Task, task.CreatedOn, task.Status)
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}
