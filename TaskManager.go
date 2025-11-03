package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Id              int
	TaskName        string
	TaskDescription string
}

func (t Task) String() string {
	return fmt.Sprintf("Task %d -> %sDescription: %s",
		t.Id, t.TaskName, t.TaskDescription)
}

func CreateTask() Task {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o nome da tarefa: ")
	taskNameFromInput, _ := reader.ReadString('\n')
	fmt.Print("Digite a descrição da tarefa: ")
	taskDescriptionFromInput, _ := reader.ReadString('\n')
	task := Task{1, taskNameFromInput, taskDescriptionFromInput}
	return task
}

func main() {
	task := CreateTask()
	fmt.Print(task.String())
}
