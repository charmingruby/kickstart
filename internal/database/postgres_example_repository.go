package database

import (
	"github.com/charmingruby/kickstart/internal/domain/example"
	"github.com/jmoiron/sqlx"
)

const (
	createExample   = "create example"
	findExampleByID = "find example by id"
)

func exampleQueries() map[string]string {
	return map[string]string{
		createExample: `INSERT INTO examples
		(id, name)
		VALUES ($1, $2)
		RETURNING *`,
		findExampleByID: `SELECT * FROM examples 
		WHERE id = $1`,
	}
}

func NewPostgresExampleRepository(db *sqlx.DB) (*PostgresExampleRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range exampleQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				newPreparationErr(queryName, "example", err)
		}

		stmts[queryName] = stmt
	}

	return &PostgresExampleRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type PostgresExampleRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PostgresExampleRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			newStatementNotPreparedErr(queryName, "example")
	}

	return stmt, nil
}

func (r *PostgresExampleRepository) Store(e *example.Example) error {
	stmt, err := r.statement(createExample)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(
		e.ID,
		e.Name,
	); err != nil {
		return err
	}

	return nil
}

func (r *PostgresExampleRepository) FindByID(id string) (*example.Example, error) {
	stmt, err := r.statement(findExampleByID)
	if err != nil {
		return nil, err
	}

	var example example.Example
	if err := stmt.Get(&example, id); err != nil {
		println(err.Error())
		return nil, err
	}

	return &example, nil
}
