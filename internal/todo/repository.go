package todo

import (
	"database/sql"
	
)


type Repository struct{
	DB *sql.DB
}

func NewRepository(db *sql.DB) * Repository {
	return &Repository{DB:db}
}

// CREATE TODO
func (r *Repository) CreateTodo(todo *Todo) error {
	query := `INSERT INTO todos (title, completed, user_id, created_at)
	           VALUES ($1, $2, $3, NOW())
			   RETURNING id, created_at`

	return r.DB.QueryRow(query, todo.Title, todo.Completed, todo.UserID).Scan(&todo.ID, &todo.CreatedAt)
}

// GET ALL TODO
func (r *Repository) GetAllTodos(userID string)([]Todo, error) {
	query := `SELECT id, title, completed, user_id, created_at
			  FROM todos
			  WHERE user_id = $1
			  ORDER BY created_at DESC`
	
	rows, err := r.DB.Query(query, userID)
	if err!= nil{
		return nil, err
	}
	defer rows.Close()

	var todo []Todo

	for rows.Next(){
		var t Todo
		if err := rows.Scan(
			&t.ID,
		    &t.Title,
			&t.Completed,
			&t.UserID,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		todo = append(todo, t)
	}

	return todo, nil
	
}


// UPDATE AN EXISTING TODO

func (r *Repository) UpdateTodo(todo *Todo) error{
	query := `UPDATE todos
			  SET title = $1, completed = $2
			  WHERE id=$3 AND user_id = $4`

   _, err := r.DB.Exec(query, 
						todo.Title,
						todo.Completed,
						todo.ID,
						todo.UserID)
						
					return err
}


// DELETE TODO
func (r *Repository) DeleteTodo(todoID int, userID string) error {
	query := `DELETE FROM todos
				WHERE id=$1 AND user_id = $2`

	_, err := r.DB.Exec(query, todoID, userID)
	return err
}
