package postgres_repository

import (
	"github.com/charmingruby/kickstart/internal/example/database/postgres_repository/mapper"
	"github.com/charmingruby/kickstart/internal/example/domain/entity"
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
				NewPreparationErr(queryName, "example", err)
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
			NewStatementNotPreparedErr(queryName, "example")
	}

	return stmt, nil
}

func (r *PostgresExampleRepository) Store(e *entity.Example) error {
	stmt, err := r.statement(createExample)
	if err != nil {
		return err
	}

	mappedEntity := mapper.DomainExampleToPostgres(*e)

	if _, err := stmt.Exec(
		mappedEntity.ID,
		mappedEntity.Name,
	); err != nil {
		return err
	}

	return nil
}

func (r *PostgresExampleRepository) FindByID(id string) (*entity.Example, error) {
	stmt, err := r.statement(findExampleByID)
	if err != nil {
		return nil, err
	}

	var example mapper.PostgresExample
	if err := stmt.Get(&example, id); err != nil {
		return nil, err
	}

	mappedExample := mapper.PostgresExampleToDomain(example)

	return &mappedExample, nil
}
