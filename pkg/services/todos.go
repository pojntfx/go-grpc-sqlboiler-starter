package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"
//go:generate sh -c "cd ../sql && sqlboiler --wipe psql -o generated"

import (
	"context"
	"database/sql"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	proto "github.com/pojntfx/miza-backend/pkg/proto/generated"
	models "github.com/pojntfx/miza-backend/pkg/sql/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Todos manages todos
type Todos struct {
	proto.UnimplementedTodosServer
	DB *sql.DB
}

// Create creates a todo
func (t *Todos) Create(ctx context.Context, req *proto.NewTodo) (*proto.Todo, error) {
	todo := &models.Todo{
		Title: req.GetTitle(),
		Body:  req.GetBody(),
	}

	if err := todo.Insert(context.Background(), t.DB, boil.Infer()); err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not create todo")
	}

	return &proto.Todo{
		ID:    int64(todo.ID),
		Title: todo.Title,
		Body:  todo.Body,
	}, nil
}

// List lists all todos
func (t *Todos) List(ctx context.Context, req *empty.Empty) (*proto.TodoList, error) {
	todos, err := models.Todos().All(context.Background(), t.DB)
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not get todos")
	}

	outTodos := []*proto.Todo{}
	for _, todo := range todos {
		outTodos = append(outTodos, &proto.Todo{
			ID:    int64(todo.ID),
			Title: todo.Title,
			Body:  todo.Body,
		})
	}

	return &proto.TodoList{
		Todos: outTodos,
	}, nil
}

// Get gets one todo
func (t *Todos) Get(ctx context.Context, req *proto.TodoID) (*proto.Todo, error) {
	todo, err := models.FindTodo(context.Background(), t.DB, int(req.GetID()))
	if err == sql.ErrNoRows {
		log.Println(err.Error())

		return nil, status.Errorf(codes.NotFound, "could not find todo")
	}
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not get todo")
	}

	return &proto.Todo{
		ID:    int64(todo.ID),
		Title: todo.Title,
		Body:  todo.Body,
	}, nil
}

// Update updates one todo
func (t *Todos) Update(ctx context.Context, req *proto.Todo) (*proto.Todo, error) {
	todo, err := models.FindTodo(context.Background(), t.DB, int(req.GetID()))
	if err == sql.ErrNoRows {
		log.Println(err.Error())

		return nil, status.Errorf(codes.NotFound, "could not find todo")
	}
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not get todo")
	}

	if req.GetTitle() != "" {
		todo.Title = req.GetTitle()
	}

	if req.GetBody() != "" {
		todo.Body = req.GetBody()
	}

	if _, err := todo.Update(context.Background(), t.DB, boil.Infer()); err != nil {
		if err != nil {
			log.Println(err.Error())

			return nil, status.Errorf(codes.Unknown, "could not update todo")
		}
	}

	return &proto.Todo{
		ID:    int64(todo.ID),
		Title: todo.Title,
		Body:  todo.Body,
	}, nil
}

// Delete deletes one todo
func (t *Todos) Delete(ctx context.Context, req *proto.TodoID) (*proto.Todo, error) {
	todo, err := models.FindTodo(context.Background(), t.DB, int(req.GetID()))
	if err == sql.ErrNoRows {
		log.Println(err.Error())

		return nil, status.Errorf(codes.NotFound, "could not find todo")
	}
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not get todo")
	}

	if _, err := todo.Delete(context.Background(), t.DB); err != nil {
		if err != nil {
			log.Println(err.Error())

			return nil, status.Errorf(codes.Unknown, "could not delete todo")
		}
	}

	return &proto.Todo{
		ID:    int64(todo.ID),
		Title: todo.Title,
		Body:  todo.Body,
	}, nil
}
