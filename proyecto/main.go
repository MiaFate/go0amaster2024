package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	desc        string
	isCompleted bool
}

type TaskList struct {
	tasks []Task
}

// metodo para agregar tareas
func (l *TaskList) addTask(t Task) {
	l.tasks = append(l.tasks, t)
}

// metodo para marcar una tarea como completada
func (l *TaskList) setCompleted(index int) {
	l.tasks[index-1].isCompleted = true
}

// metodo para editar tarea
func (l *TaskList) editTask(index int, t Task) {
	l.tasks[index-1] = t
}

// metodo para eliminar una tarea
func (l *TaskList) deleteTask(index int) {
	l.tasks = append(l.tasks[:index-1], l.tasks[index:]...)
}

func main() {
	// instancia de lista de tareas
	list := TaskList{}

	// instancia de bufio para entrada de datos
	read := bufio.NewReader(os.Stdin)
	for {

		var option int
		fmt.Println("Select an option:\n",
			"1. Add a new task\n",
			"2. Set task as completed\n",
			"3. Edit task\n",
			"4. Delete task\n",
			"5. Exit")

		fmt.Println("Enter number: ")
		_, _ = fmt.Scanln(&option)

		switch option {
		case 1:
			var t Task
			fmt.Println("Enter task's name: ")
			t.name, _ = read.ReadString('\n')
			fmt.Println("Enter task's description: ")
			t.desc, _ = read.ReadString('\n')
			list.addTask(t)
			fmt.Println("task added")
		case 2:
			var index int
			fmt.Println("Enter task's index: ")
			_, _ = fmt.Scanln(&index)
			list.setCompleted(index)
			fmt.Println("task marked as completed")
		case 3:
			var index int
			var t Task
			fmt.Print("Enter task's index: ")
			_, _ = fmt.Scanln(&index)
			fmt.Print("Enter task's name:")
			t.name, _ = read.ReadString('\n')
			fmt.Print("Enter task's description: ")
			t.desc, _ = read.ReadString('\n')
			list.editTask(index, t)
			fmt.Println("task edited")
		case 4:
			var index int
			fmt.Print("Enter task's index: ")
			_, _ = fmt.Scanln(&index)
			list.deleteTask(index)
			fmt.Println("task deleted")
		case 5:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option!")
		}

		// listar todas las tareas
		fmt.Println("Lista de tareas:")
		fmt.Println("================")
		for index, task := range list.tasks {
			fmt.Printf("%d. %s - %s - Completado: %t\n", index+1, task.name, task.desc, task.isCompleted)
		}

	}
}
