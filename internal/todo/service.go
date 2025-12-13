package todo

import "errors"

type Service struct{
	repo *Repository
}

func NewService(repo *Repository) *Service{
	return &Service{repo: repo}
}


// CREATE TODO
func (s *Service) Create(todo *Todo) error{
	if todo.Title == "" {
   return errors.New("title field cannot be empty")
		}
	return s.repo.CreateTodo(todo)
}

// GET ALL TODO
 func (s *Service) GetAll(userID string)([]Todo, error){
	return s.repo.GetAllTodos(userID)
 }

//  UPDATE TODO
func (s *Service) Update(todo *Todo) error{
	if todo.Title ==""{
		return errors.New("title field cannot be empty")
	}
	return s.repo.UpdateTodo(todo)
}

// DELETE TODO
func (s * Service) Delete(id int, userID string) error{
	return s.repo.DeleteTodo(id, userID)
}