package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Categoria struct {
	db        *sql.DB // todas as interfaces para trabalhar com database
	ID        string
	Nome      string
	Descricao string
}

// Construtor: retorna um ponteiro para uma categoria vazia. Apenas com a conexão
func NewCategoria(db *sql.DB) *Categoria {
	return &Categoria{db: db}
}

// método que cria e retorna uma categoria
func (c *Categoria) Create(nome string, descricao string) (Categoria, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT  INTO categoria (id, nome, descricao) VALUES ($1, $2, $3)", id, nome, descricao)
	if err != nil {
		return Categoria{}, err
	}
	return Categoria{ID: id, Nome: nome, Descricao: descricao}, nil
}

// método que retorna ula lista de categorias
func (c *Categoria) FindAll() ([]Categoria, error) {
	rows, err := c.db.Query("SELECT * FROM categoria")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categorias := []Categoria{}
	for rows.Next() {
		var id, nome, descricao string
		if err := rows.Scan(&id, &nome, &descricao); err != nil {
			return nil, err
		}
		categorias = append(categorias, Categoria{ID: id, Nome: nome, Descricao: descricao})
	}
	return categorias, nil
}

func (c *Categoria) FindById(id string) (Categoria, error) {
	var nome, descricao string
	err := c.db.QueryRow("SELECT nome, descricao FROM categoria WHERE id = $1 ", id).Scan(&nome, &descricao)
	if err != nil {
		return Categoria{}, err
	}
	return Categoria{ID: id, Nome: nome, Descricao: descricao}, nil
}

func (c *Categoria) FindByCursoId(cursoID string) (Categoria, error) {
	var id, nome, descricao string
	err := c.db.QueryRow("SELECT c.id, c.nome, c.descricao FROM categoria c JOIN curso co ON c.id = co.categoria_id").Scan(&id, &nome, &descricao)
	if err != nil {
		return Categoria{}, err
	}
	return Categoria{ID: id, Nome: nome, Descricao: descricao}, nil
}
