package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"
//go:generate sh -c "cd ../sql && sqlboiler --wipe psql -o generated"

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
	proto "github.com/pojntfx/miza-backend/pkg/proto/generated"
	models "github.com/pojntfx/miza-backend/pkg/sql/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	protobuf "github.com/golang/protobuf/proto"
)

// Todos manages todos
type Todos struct {
	proto.UnimplementedTodosServer
	DB   *sql.DB
	NATS *nats.Conn
}

// getNamespaceFromContext returns the namespace from the context
func (t *Todos) getNamespaceFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("could not parse metadata")
	}

	namespace := md.Get("x-miza-namespace")
	if len(namespace) == 0 || namespace[0] == "" {
		return "", errors.New("no namespace specified")
	}

	return namespace[0], nil
}

// Create creates a todo
func (t *Todos) Create(ctx context.Context, req *proto.NewTodo) (*proto.Todo, error) {
	ns, err := t.getNamespaceFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	todoCount, err := models.Todos(qm.Where(models.TodoColumns.Namespace+"= ?", ns)).Count(context.Background(), t.DB)
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not get todos")
	}

	todo := &models.Todo{
		Title:     req.GetTitle(),
		Body:      req.GetBody(),
		Namespace: ns,
		Index:     todoCount + 1,
	}

	if err := todo.Insert(context.Background(), t.DB, boil.Infer()); err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not create todo")
	}

	protoTodo := &proto.Todo{
		ID:    int64(todo.ID),
		Title: todo.Title,
		Body:  todo.Body,
		Index: todo.Index,
	}

	natsTodo, err := protobuf.Marshal(protoTodo)
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not marshal created todo")
	}

	if err := t.NATS.Publish(fmt.Sprintf("todos:create:%s", ns), natsTodo); err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not published created todo")
	}

	return protoTodo, nil
}

// List lists all todos
func (t *Todos) List(ctx context.Context, req *empty.Empty) (*proto.TodoList, error) {
	ns, err := t.getNamespaceFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	todos, err := models.Todos(
		qm.Where(models.TodoColumns.Namespace+"= ?", ns),
		qm.OrderBy(models.TodoColumns.Index),
	).All(context.Background(), t.DB)
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
			Index: todo.Index,
		})
	}

	return &proto.TodoList{
		Todos: outTodos,
	}, nil
}

// Get gets one todo
func (t *Todos) Get(ctx context.Context, req *proto.TodoID) (*proto.Todo, error) {
	ns, err := t.getNamespaceFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	todo, err := models.Todos(qm.Where(models.TodoColumns.Namespace+"= ?", ns), qm.Where(models.TodoColumns.ID+"= ?", req.GetID())).One(context.Background(), t.DB)
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
		Index: todo.Index,
	}, nil
}

// Update updates one todo
func (t *Todos) Update(ctx context.Context, req *proto.Todo) (*proto.Todo, error) {
	ns, err := t.getNamespaceFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	todo, err := models.Todos(qm.Where(models.TodoColumns.Namespace+"= ?", ns), qm.Where(models.TodoColumns.ID+"= ?", req.GetID())).One(context.Background(), t.DB)
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
		Index: todo.Index,
	}, nil
}

// Delete deletes one todo
func (t *Todos) Delete(ctx context.Context, req *proto.TodoID) (*proto.Todo, error) {
	ns, err := t.getNamespaceFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not begin transaction")
	}

	todo, err := models.Todos(qm.Where(models.TodoColumns.Namespace+"= ?", ns), qm.Where(models.TodoColumns.ID+"= ?", req.GetID())).One(context.Background(), t.DB)
	if err == sql.ErrNoRows {
		log.Println(err.Error())

		return nil, status.Errorf(codes.NotFound, "could not find todo")
	}
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not get todo")
	}

	// Recalculate indexes for todos which have a higher index
	{
		todosToUpdate, err := models.Todos(
			qm.Where(models.TodoColumns.Namespace+"= ?", ns),
			qm.Where(models.TodoColumns.Index+">= ?", todo.Index),
		).All(context.Background(), t.DB)
		if err != nil {
			log.Println(err.Error())

			return nil, status.Errorf(codes.Unknown, "could not get todos")
		}

		for _, todoToUpdate := range todosToUpdate {
			todoToUpdate.Index = todoToUpdate.Index - 1

			if _, err := todoToUpdate.Update(context.Background(), t.DB, boil.Infer()); err != nil {
				if err != nil {
					log.Println(err.Error())

					return nil, status.Errorf(codes.Unknown, "could not update todo")
				}
			}
		}
	}

	if _, err := todo.Delete(context.Background(), t.DB); err != nil {
		if err != nil {
			log.Println(err.Error())

			return nil, status.Errorf(codes.Unknown, "could not delete todo")
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not commit transaction")
	}

	return &proto.Todo{
		ID:    int64(todo.ID),
		Title: todo.Title,
		Body:  todo.Body,
		Index: todo.Index,
	}, nil
}

// Reorder reorders a todo
func (t *Todos) Reorder(ctx context.Context, req *proto.TodoReorder) (*proto.Todo, error) {
	ns, err := t.getNamespaceFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not begin transaction")
	}

	todo, err := models.Todos(qm.Where(models.TodoColumns.Namespace+"= ?", ns), qm.Where(models.TodoColumns.ID+"= ?", req.GetID())).One(context.Background(), t.DB)
	if err == sql.ErrNoRows {
		log.Println(err.Error())

		return nil, status.Errorf(codes.NotFound, "could not find todo")
	}
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not get todo")
	}

	oldIndex := todo.Index
	todo.Index += req.GetOffset()

	// Recalculate the other indexes
	{
		var todosToUpdate models.TodoSlice
		if todo.Index > oldIndex {
			// The new index is higher than the old one; recount the indexes in between
			// the new index and the old index
			todosToUpdate, err = models.Todos(
				qm.Where(models.TodoColumns.Namespace+"= ?", ns),
				qm.Where(models.TodoColumns.Index+">= ?", oldIndex),
				qm.Where(models.TodoColumns.Index+"<= ?", todo.Index),
			).All(context.Background(), t.DB)
			if err != nil {
				log.Println(err.Error())

				return nil, status.Errorf(codes.Unknown, "could not get todos")
			}
		} else {
			// The new index is lower than the old one; recount the indexes in between
			// the new index and the old index
			todosToUpdate, err = models.Todos(
				qm.Where(models.TodoColumns.Namespace+"= ?", ns),
				qm.Where(models.TodoColumns.Index+">= ?", todo.Index),
				qm.Where(models.TodoColumns.Index+"<= ?", oldIndex),
			).All(context.Background(), t.DB)
			if err != nil {
				log.Println(err.Error())

				return nil, status.Errorf(codes.Unknown, "could not get todos")
			}
		}

		for _, todoToUpdate := range todosToUpdate {
			if todo.Index > oldIndex {
				todoToUpdate.Index = todoToUpdate.Index - 1
			} else {
				todoToUpdate.Index = todoToUpdate.Index + 1
			}

			if _, err := todoToUpdate.Update(context.Background(), t.DB, boil.Infer()); err != nil {
				if err != nil {
					log.Println(err.Error())

					return nil, status.Errorf(codes.Unknown, "could not update todo")
				}
			}
		}
	}

	// Update the new index
	if _, err := todo.Update(context.Background(), t.DB, boil.Infer()); err != nil {
		if err != nil {
			log.Println(err.Error())

			return nil, status.Errorf(codes.Unknown, "could not update todo")
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Unknown, "could not commit transaction")
	}

	return &proto.Todo{
		ID:    int64(todo.ID),
		Title: todo.Title,
		Body:  todo.Body,
		Index: todo.Index,
	}, nil
}

// SubscribeToChanges subscribes to all changes
func (t *Todos) SubscribeToChanges(req *empty.Empty, srv proto.Todos_SubscribeToChangesServer) error {
	ns, err := t.getNamespaceFromContext(srv.Context())
	if err != nil {
		return status.Errorf(codes.Unknown, err.Error())
	}

	createTodosChan := make(chan *nats.Msg, 64)

	if _, err = t.NATS.ChanSubscribe(fmt.Sprintf("todos:create:%s", ns), createTodosChan); err != nil {
		log.Println(err)

		return status.Errorf(codes.Unknown, "could not subscribe to created todo changes")
	}

	for msg := range createTodosChan {
		var todo proto.Todo
		if err := protobuf.Unmarshal(msg.Data, &todo); err != nil {
			log.Println("could not unmarshal created todo change", err.Error())
		}

		if err := srv.Send(&proto.TodoChange{Type: proto.TodoChangeType_CREATE, Todo: &todo}); err != nil {
			log.Println("could not send created todo change", err.Error())
		}
	}

	return status.Errorf(codes.Unknown, "could not send todo changes")
}
