package options

import (
	"context"
	"fmt"

	"github.com/Jenil-Desai/Tasker/db"
)

type AddTaskResponse struct {
	Id        uint
	Task      string
	Status    bool
	CreatedOn string
}

func AddTask(task string) (AddTaskResponse, error) {
	client := db.NewClient()
	defer client.Prisma.Disconnect()

	err := client.Prisma.Connect()
	if err != nil {
		return AddTaskResponse{}, fmt.Errorf("failed to connect to the database: %w", err)
	}

	ctx := context.Background()

	newtask, err := client.Task.CreateOne(
		db.Task.Task.Set(task),
	).Exec(ctx)

	if err != nil {
		return AddTaskResponse{}, fmt.Errorf("failed to add task: %w", err)
	}

	addTaskResponse := AddTaskResponse{
		Id:        uint(newtask.ID),
		Task:      newtask.Task,
		Status:    newtask.Status,
		CreatedOn: newtask.CreatedAt.Format("02 Jan 2006"),
	}

	return addTaskResponse, nil
}
