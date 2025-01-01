package options

import (
	"context"
	"fmt"

	"github.com/Jenil-Desai/Tasker/db"
)

type TaskListResponse struct {
	Id        uint
	Task      string
	Status    bool
	CreatedOn string
}

func ListTask() ([]TaskListResponse, error) {
	client := db.NewClient()
	defer client.Prisma.Disconnect()

	err := client.Prisma.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	ctx := context.Background()

	tasks, err := client.Task.FindMany(
		db.Task.ID.Not(0),
	).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	var taskList []TaskListResponse
	for _, task := range tasks {
		taskList = append(taskList, TaskListResponse{
			Id:        uint(task.ID),
			Task:      task.Task,
			Status:    task.Status,
			CreatedOn: task.CreatedAt.Format("02 Jan 2006"),
		})
	}

	return taskList, nil
}
