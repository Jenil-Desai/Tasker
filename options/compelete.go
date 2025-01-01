package options

import (
	"context"
	"fmt"

	"github.com/Jenil-Desai/Tasker/db"
)

type CompeleteTaskResponse struct {
	Id        uint
	Task      string
	Status    bool
	CreatedOn string
}

func CompeleteTask(taskId uint) (CompeleteTaskResponse, error) {
	client := db.NewClient()
	defer client.Prisma.Disconnect()

	err := client.Prisma.Connect()
	if err != nil {
		return CompeleteTaskResponse{}, fmt.Errorf("failed to connect to the database: %w", err)
	}

	ctx := context.Background()

	updatedTask, err := client.Task.FindUnique(
		db.Task.ID.Equals(int(taskId)),
	).Update(
		db.Task.Status.Set(true),
	).Exec(ctx)

	if err != nil {
		return CompeleteTaskResponse{}, fmt.Errorf("failed to mark task: %w", err)
	}

	compeleteTaskResponse := CompeleteTaskResponse{
		Id:        uint(updatedTask.ID),
		Task:      updatedTask.Task,
		Status:    updatedTask.Status,
		CreatedOn: updatedTask.CreatedAt.Format("02 Jan 2006"),
	}

	return compeleteTaskResponse, nil
}
