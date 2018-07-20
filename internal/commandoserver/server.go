package commandoserver

import (
	"context"

	"github.com/twitchtv/twirp"
	pb "github.com/williamhgough/comman-do/rpc/todo"
)

// Server implements Commando interface
type Server struct{}

var todos = map[int32]*pb.Todo{
	1: &pb.Todo{ID: 1, Title: "hello world", Complete: false},
}

// MakeTodo implements the command interface method MakeTodo
func (s *Server) MakeTodo(ctx context.Context, todo *pb.Todo) (*pb.Empty, error) {
	if todo.ID < 1 {
		return nil, twirp.InvalidArgumentError("ID:", "No ID provided")
	}
	if todo.Title == "" {
		return nil, twirp.InvalidArgumentError("Title:", "No title given")
	}
	todos[todo.ID] = todo
	return &pb.Empty{}, nil
}

// GetTodo implements the command interface method GetTodo
func (s *Server) GetTodo(ctx context.Context, q *pb.TodoQuery) (*pb.Todo, error) {
	if q.Id < 1 {
		return nil, twirp.InvalidArgumentError("ID:", "Must be greater than zero")
	}
	return todos[q.Id], nil
}

// AllTodos implements the command interface method AllTodos
func (s *Server) AllTodos(ctx context.Context, e *pb.Empty) (*pb.TodoResponse, error) {
	t := []*pb.Todo{}
	for _, v := range todos {
		t = append(t, v)
	}
	if len(t) < 1 {
		return nil, twirp.InternalError("No todos found!")
	}
	return &pb.TodoResponse{
		Todos: t,
	}, nil
}

// RemoveTodo implements the command interface method RemoveTodo
func (s *Server) RemoveTodo(ctx context.Context, q *pb.TodoQuery) (*pb.Empty, error) {
	if q.Id < 1 {
		return nil, twirp.InvalidArgumentError("ID:", "Must be greater than zero")
	}
	delete(todos, q.Id)
	return nil, nil
}
