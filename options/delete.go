package options

import (
	"context"
	"fmt"

	"github.com/Jenil-Desai/Tasker/db"
)

type DeleteTaskResponse struct {
	Id        uint
	Task      string
	Status    bool
	CreatedOn string
}

func DeleteTask(taskId uint64) (DeleteTaskResponse, error) {
	client := db.NewClient()
	client.Prisma.Connect()

	ctx := context.Background()

	deletedTask, err := client.Task.FindUnique(
		db.Task.ID.Equals(int(taskId)),
	).Delete().Exec(ctx)

	if err != nil {
		return DeleteTaskResponse{}, fmt.Errorf("failed to add task: %w", err)
	}

	deletedTaskResponse := DeleteTaskResponse{
		Id:        uint(deletedTask.ID),
		Task:      deletedTask.Task,
		Status:    deletedTask.Status,
		CreatedOn: deletedTask.CreatedAt.Format("02 Jan 2006"),
	}

	return deletedTaskResponse, nil
}
