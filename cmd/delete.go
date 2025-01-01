package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/Jenil-Desai/Tasker/options"
	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.ParseUint(args[0], 10, 32)
		newTask, err := options.DeleteTask(id)
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
	rootCmd.AddCommand(deleteTaskCmd)
}
