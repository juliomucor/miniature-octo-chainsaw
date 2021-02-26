package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliomucor/miniature-octo-chainsaw/domain"
	"strconv"
)

type task domain.Task

var tasks = []*task{
	{
		Id:          1,
		Description: "doing some work 1",
		Duration:    45,
		Next:        2,
		Done:        false,
	},
	{
		Id:          2,
		Description: "doing some work 2",
		Duration:    13,
		Next:        -1,
		Done:        false,
	},
}

// get all tasks
func getTasks(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"tasks": tasks,
		},
	})
}

// create a task
func createTask(c *fiber.Ctx) error {
	var t task
	err := c.BodyParser(&t)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cannot parse JSON",
		})
	}

	t.Id = len(tasks) + 1
	t.Next = -1
	tasks = append(tasks, &t)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"Id": t.Id,
		},
	})

}

// update task is a fully replace of the task with the same id, a partial update=PATCH
func updateTask(c *fiber.Ctx) error {
	var ut task
	err := c.BodyParser(&ut)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cannot parse JSON",
		})
	}

	for i, t := range tasks {
		if t.Id == ut.Id {
			tasks = append(tasks[:i], &ut)
			tasks = append(tasks, tasks[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"task": t,
				},
			})

		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Task not found",
	})
}

func getTask(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cannot parse Id",
		})
	}

	// find task and return, this replaces DB
	for _, t := range tasks {
		if t.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"task": t,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Task not found",
	})
}

func deleteTask(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cannot parse Id",
		})
	}
	for i, t := range tasks {
		if t.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"task": t,
				},
			})

		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Task not found",
	})
}

func TaskRoute(route fiber.Router) {
	route.Get("", getTasks)
	route.Get("/:id", getTask)
	route.Post("", createTask)
	route.Put("", updateTask)
	route.Delete("/:id", getTask)
}
