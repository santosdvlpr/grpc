package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Curso struct {
	db          *sql.DB
	ID          string
	Nome        string
	Descricao   string
	CategoriaID string
}

func NewCurso(db *sql.DB) *Curso {
	return &Curso{db: db}
}

// método que cria e retorna um curso
func (c *Curso) Create(nome string, descricao string, categoriaID string) (*Curso, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO curso (id, nome, descricao, categoria_id) VALUES($1,$2,$3,$4)",
		id, nome, descricao, categoriaID)
	if err != nil {
		return nil, err
	}
	return &Curso{
		ID:          id,
		Nome:        nome,
		Descricao:   descricao,
		CategoriaID: categoriaID,
	}, nil
}

// método que retorna lista de cursos
func (c *Curso) FindAll() ([]Curso, error) {
	rows, err := c.db.Query("SELECT id,nome,descricao FROM curso")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cursos := []Curso{}
	for rows.Next() {
		var id, nome, descricao string
		if err := rows.Scan(&id, &nome, &descricao); err != nil {
			return nil, err
		}
		cursos = append(cursos, Curso{ID: id, Nome: nome, Descricao: descricao})
	}
	return cursos, nil
}

// método que retorna lista de cursos de uma dada categoria
func (c *Curso) FindByCategoriaID(categoriaID string) ([]Curso, error) {
	rows, err := c.db.Query("SELECT id, nome, descricao, categoria_id FROM curso WHERE categoria_id = $1 ", categoriaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cursos := []Curso{}
	for rows.Next() {
		var id, nome, descricao, categoriaID string
		if err := rows.Scan(&id, &nome, &descricao, &categoriaID); err != nil {
			return nil, err
		}
		cursos = append(cursos, Curso{ID: id, Nome: nome, Descricao: descricao, CategoriaID: categoriaID})
	}
	return cursos, nil
}
