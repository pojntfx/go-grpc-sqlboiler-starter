package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"database/sql"
	"log"

	proto "github.com/pojntfx/go-grpc-sqlboiler-starter/pkg/proto/generated"
	models "github.com/pojntfx/go-grpc-sqlboiler-starter/pkg/sql/generated"
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
