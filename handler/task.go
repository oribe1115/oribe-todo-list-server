package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/oribe1115/oribe-todo-list-server/model"
)

func CreateTableHandler(c echo.Context) error {
	err := model.CreateTable()

	if err != nil {
		return c.String(http.StatusInternalServerError, "faild to craete table\n")
	}
	return c.String(http.StatusCreated, "tasks table created!\n")
}

func AddNewTaskHandler(c echo.Context) error {
	taskFromClient := model.TaskToAdd{}
	c.Bind(&taskFromClient)

	err := model.AddNewTask(taskFromClient)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to add")
	}

	tasklist, err := model.AllTasks()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get task list")
	}

	return c.JSON(http.StatusCreated, tasklist)
}

func GetAllTasksHandler(c echo.Context) error {
	tasklist, err := model.AllTasks()

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to get task list")
	}

	return c.JSON(http.StatusOK, tasklist)
}

func ChangeStatusFandler(c echo.Context) error {
	task := model.TaskForClient{}
	c.Bind(&task)
	err := model.ChangeStatus(task)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "faild to change status")
	}
	return c.String(http.StatusOK, "task is updated")
}
