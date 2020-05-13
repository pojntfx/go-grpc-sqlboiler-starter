package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	proto "github.com/pojntfx/go-grpc-sqlboiler-starter/pkg/proto/generated"
	"github.com/pojntfx/go-grpc-sqlboiler-starter/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate sqlboiler --wipe psql -o pkg/sql/generated

func main() {
	laddr := flag.String("laddr", ":3069", "Listen address")
	dbhost := flag.String("dbhost", "localhost", "Database host")
	dbport := flag.String("dbport", "5432", "Database port")
	dbusr := flag.String("dbusr", "go-grpc-sqlboiler-starter", "Database user")
	dbpass := flag.String("dbpass", "go-grpc-sqlboiler-starter", "Database password")
	dbname := flag.String("dbname", "go-grpc-sqlboiler-starter", "Database name")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%v port=%v sslmode=disable user=%v password=%v dbname=%v", *dbhost, *dbport, *dbusr, *dbpass, *dbname))
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", *laddr)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	tsvc := services.Todos{DB: db}
	proto.RegisterTodosServer(srv, &tsvc)

	log.Println("starting server")

	srv.Serve(lis)
}