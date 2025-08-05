package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/smolyaninov/go-task-tracker-cli/internal/domain"
	"github.com/smolyaninov/go-task-tracker-cli/internal/repository"
	"github.com/smolyaninov/go-task-tracker-cli/internal/service"
)

func main() {
	repo := repository.NewJSONRepository("tasks.json")

	tasks, err := repo.Load()
	if err != nil {
		fmt.Println("Error load tasks:", err)
		os.Exit(1)
	}

	s := service.NewServiceWithData(tasks)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("\ttaskcli add \"description\"")
		fmt.Println("\ttaskcli list [status]")
		fmt.Println("\ttaskcli update <id> \"new description\"")
		fmt.Println("\ttaskcli delete <id>")
		fmt.Println("\ttaskcli mark-in-progress <id>")
		fmt.Println("\ttaskcli mark-done <id>")
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "add":
		requireArgs(args, 2, "add \"description\"")

		description := strings.TrimSpace(strings.Join(args[1:], " "))

		task, err := s.Add(description)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		mustSave(repo, s.GetAll())

		fmt.Printf("Task added: [%d] %s\n", task.ID, task.Description)
	case "list":
		var filter *domain.Status

		if len(args) == 2 {
			status := domain.Status(args[1])
			filter = &status
		}

		tasks := s.List(filter)
		if len(tasks) == 0 {
			fmt.Println("Empty")
			return
		}

		for _, t := range tasks {
			fmt.Printf("[%d] %-20s (%s)\n", t.ID, t.Description, t.Status)
		}
	case "update":
		requireArgs(args, 3, "update <id> <new description>")

		id := mustID(args)

		newDescription := strings.TrimSpace(strings.Join(args[2:], " "))
		err = s.Update(id, newDescription)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		mustSave(repo, s.GetAll())

		fmt.Println("Task updated")
	case "delete":
		requireArgs(args, 2, "delete <id>")

		id := mustID(args)

		err = s.Delete(id)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		mustSave(repo, s.GetAll())

		fmt.Println("Task deleted")
	case "mark-in-progress":
		requireArgs(args, 2, "mark-in-progress <id>")

		id := mustID(args)

		err = s.ChangeStatus(id, domain.StatusInProgress)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		mustSave(repo, s.GetAll())

		fmt.Println("Task marked as in-progress")
	case "mark-done":
		requireArgs(args, 2, "mark-done <id>")

		id := mustID(args)

		err = s.ChangeStatus(id, domain.StatusDone)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		mustSave(repo, s.GetAll())

		fmt.Println("Task marked as done")
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func mustID(args []string) int {
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID")
		os.Exit(1)
	}

	return id
}

func mustSave(repo repository.TaskRepository, tasks []domain.Task) {
	if err := repo.Save(tasks); err != nil {
		fmt.Println("Save error", err)
		os.Exit(1)
	}
}

func requireArgs(args []string, expected int, usage string) {
	if len(args) < expected {
		fmt.Println("Usage:", usage)
		os.Exit(1)
	}
}
