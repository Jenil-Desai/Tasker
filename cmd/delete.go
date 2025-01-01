package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/Jenil-Desai/Tasker/options"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Abort mission",
	Long:  "Remove a mission from your command center. Requires mission ID for precise target elimination.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.ParseUint(args[0], 10, 32)
		newTask, err := options.DeleteTask(uint(id))
		if err != nil {
			color.Red("%v", err)
			return
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 5, ' ', 0)

		fmt.Fprintln(w, "ID\tTask\tCreated On\tStatus")
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", newTask.Id, newTask.Task, newTask.CreatedOn, newTask.Status)

		w.Flush()

		color.Yellow("\n⚠️ Mission [%v] has been aborted and cleared from radar", newTask.Id)
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
}
