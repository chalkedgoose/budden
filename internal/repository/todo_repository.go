package repository

import (
	"context"
	"log"
	"lupa/calba/budden/gen/dbstore"
	"lupa/calba/budden/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Todo = dbstore.Todo

// TodoRepository provides methods to interact with the Todo database.
type TodoRepository struct {
	dbpool *pgxpool.Pool
}

// NewTodoRepository initializes a new TodoRepository with a database connection.
func NewTodoRepository() (*TodoRepository, error) {
	cfg := config.LoadConfig()
	connString := cfg.Database.ConnectionString

	log.Println("Connecting to the database...")
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return &TodoRepository{dbpool: dbpool}, nil
}

// Close terminates the database connection.
func (r *TodoRepository) Close() {
	log.Println("Closing the database connection...")
	r.dbpool.Close()
	log.Println("Database connection closed.")
}

// ListTodos retrieves all todos from the database.
func (r *TodoRepository) ListTodos(ctx context.Context) ([]Todo, error) {
	queries := dbstore.New(r.dbpool)
	return queries.ListTodos(ctx)
}

// CreateTodo adds a new todo to the database.
func (r *TodoRepository) CreateTodo(ctx context.Context, todo Todo) error {
	tx, err := r.dbpool.Begin(ctx)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			log.Println("Rolling back transaction due to an error.")
			tx.Rollback(ctx)
		}
	}()

	queries := dbstore.New(tx)
	err = queries.CreateTodo(ctx, dbstore.CreateTodoParams{
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
	})
	if err != nil {
		log.Printf("Failed to create todo: %v", err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	return nil
}
