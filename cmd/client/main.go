package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/williamhgough/comman-do/rpc/todo"
)

var todoCount int32 = 1

func main() {
	action := flag.String("command", "list", "Command to interact with Todos. create, get, list, delete")
	title := flag.String("title", "", "The todo title")
	id := flag.Int("id", 1, "The ID of the todo you wish to retrieve")
	flag.Parse()

	client := todo.NewCommandoProtobufClient("http://localhost:8080", &http.Client{})

	switch *action {
	case "create":
		_, err := client.MakeTodo(context.Background(), &todo.Todo{
			ID:       todoCount + 1,
			Title:    *title,
			Complete: false,
		})
		if err != nil {
			log.Println("could not create todo:", err)
		}
	case "list":
		res, err := client.AllTodos(context.Background(), &todo.Empty{})
		if err != nil {
			log.Println("could not fetch todos:", err)
		}
		for i, t := range res.Todos {
			log.Printf("%d. %s [%v]", i+1, t.Title, t.Complete)
		}
	case "get":
		todo, err := client.GetTodo(context.Background(), &todo.TodoQuery{
			Id: int32(*id),
		})
		if err != nil {
			log.Println("could not fetch todo:", err)
		}

		log.Printf("%d. %s [%v]", todo.ID, todo.Title, todo.Complete)
	case "delete":
		_, err := client.RemoveTodo(context.Background(), &todo.TodoQuery{
			Id: int32(*id),
		})
		if err != nil {
			log.Println("could not remove todo:", err)
		}
	default:
		log.Println("invalid command, please try again.")
	}
}
