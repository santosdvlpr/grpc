package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/santosdvlpr/grpc/internal/database"
	"github.com/santosdvlpr/grpc/internal/pb"
	"github.com/santosdvlpr/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	/**  Passos de criação do serviço */

	//instacia banco e abre uma conexao
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	categoriaDb := database.NewCategoria(db) // e abre uma conexao
	defer db.Close()
	// instancia o serviço e um servidor
	categoriaService := service.NewCategoriaService(*categoriaDb)
	grpcServer := grpc.NewServer()

	// Registra o serviço no serveidor
	pb.RegisterCategoriaServiceServer(grpcServer, categoriaService)

	// Registra o server para processar suas próprias requisições
	reflection.Register(grpcServer)

	//abrir uma porta tcp
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	/**  Passos de iteração com o serviço */

}
