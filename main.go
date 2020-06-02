package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	proto "github.com/pojntfx/miza-backend/pkg/proto/generated"
	"github.com/pojntfx/miza-backend/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	laddr := flag.String("laddr", ":3069", "Listen address")
	dbhost := flag.String("dbhost", "localhost", "Database host")
	dbport := flag.String("dbport", "5432", "Database port")
	dbusr := flag.String("dbusr", "miza-backend", "Database user")
	dbpass := flag.String("dbpass", "miza-backend", "Database password")
	dbname := flag.String("dbname", "miza-backend", "Database name")
	busaddr := flag.String("busaddr", "localhost:4222", "Bus address")
	bususr := flag.String("bususr", "miza-backend", "Bus user")
	buspass := flag.String("buspass", "miza-backend", "Bus password")
	flag.Parse()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%v port=%v sslmode=disable user=%v password=%v dbname=%v", *dbhost, *dbport, *dbusr, *dbpass, *dbname))
	if err != nil {
		log.Fatal(err)
	}

	nats, err := nats.Connect(fmt.Sprintf("nats://%v", *busaddr), nats.UserInfo(*bususr, *buspass))
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", *laddr)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	tsvc := services.Todos{DB: db, NATS: nats}
	proto.RegisterTodosServer(srv, &tsvc)

	log.Println("starting server")

	srv.Serve(lis)
}
