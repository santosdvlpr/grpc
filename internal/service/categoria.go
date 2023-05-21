package service

import (
	"context"
	"io"

	"github.com/santosdvlpr/grpc/internal/database"
	"github.com/santosdvlpr/grpc/internal/pb"
)

type CategoriaService struct {
	pb.UnimplementedCategoriaServiceServer
	CategoriaDB database.Categoria
}

func NewCategoriaService(categoriaDB database.Categoria) *CategoriaService {
	return &CategoriaService{
		CategoriaDB: categoriaDB,
	}
}
func (c *CategoriaService) CreateCategoria(ctx context.Context, in *pb.CreateCategoriaRequest) (*pb.Categoria, error) {
	categoria, err := c.CategoriaDB.Create(in.Nome, in.Descricao)
	if err != nil {
		return nil, err
	}
	categoriaResponse := &pb.Categoria{
		Id: categoria.ID, Nome: categoria.Nome, Descricao: categoria.Descricao,
	}
	return categoriaResponse, nil
}

func (c CategoriaService) ListaCategorias(ctx context.Context, in *pb.Blank) (*pb.CategoriaList, error) {
	categorias, err := c.CategoriaDB.FindAll()
	if err != nil {
		panic(err)
	}
	var categoriasResponse []*pb.Categoria
	for _, categoria := range categorias {
		categoriaResponse := &pb.Categoria{
			Id: categoria.ID, Nome: categoria.Nome, Descricao: categoria.Descricao,
		}
		categoriasResponse = append(categoriasResponse, categoriaResponse)
	}
	return &pb.CategoriaList{Categorias: categoriasResponse}, nil
}

func (c *CategoriaService) GetCategoria(ctx context.Context, in *pb.CategoriaGetRequest) (*pb.Categoria, error) {
	categoria, err := c.CategoriaDB.FindById(in.Id)
	if err != nil {
		panic(err)
	}
	categoriaResponse := &pb.Categoria{
		Id: categoria.ID, Nome: categoria.Nome, Descricao: categoria.Descricao,
	}
	return categoriaResponse, nil
}

func (c *CategoriaService) CreateCategoriaStream(stream pb.CategoriaService_CreateCategoriaStreamServer) error {
	categorias := &pb.CategoriaList{}
	for {
		categoria, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categorias)
		}
		if err != nil {
			return err
		}
		categoriaResult, err := c.CategoriaDB.Create(categoria.Nome, categoria.Descricao)
		if err != nil {
			return err
		}
		categorias.Categorias = append(categorias.Categorias, &pb.Categoria{
			Id: categoriaResult.ID, Nome: categoriaResult.Nome, Descricao: categoriaResult.Descricao,
		})
	}
}

func (c *CategoriaService) CreateCategoriaStreamBidirectional(stream pb.CategoriaService_CreateCategoriaStreamBidirectionalServer) error {
	for {
		categoria, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		categoriaResult, err := c.CategoriaDB.Create(categoria.Nome, categoria.Descricao)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Categoria{
			Id: categoriaResult.ID, Nome: categoriaResult.Nome, Descricao: categoriaResult.Descricao,
		})
		if err != nil {
			return err
		}
	}
}
