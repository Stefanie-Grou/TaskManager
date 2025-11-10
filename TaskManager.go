package main

import (
	"GoLang/models"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	pessoa := models.NovaPessoa("João", 25)
	fmt.Printf("Nome: %s, Idade: %d\n", pessoa.Nome, pessoa.Idade)

	task := models.CreateTask("Wash the dishes", "Laundry machine is full", true)
	fmt.Printf("Task Name: %s // Task Description: %s\n // Status: %t", task.TaskName, task.TaskDescription, task.Done)
}
