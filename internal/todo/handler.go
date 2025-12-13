package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"project.com/todo/internal/auth"
)


type Handler struct{
	service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}

// CREATE TODO
func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request){
	userID, err := getUserIDFromHeader(r)
	if err != nil{
		 http.Error(w, err.Error(), http.StatusUnauthorized)
		 return
	}

	var todo Todo
	if err:= json.NewDecoder(r.Body).Decode(&todo); err!= nil{
		 http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		 return
	}

	todo.UserID = userID

	if err := h.service.Create(&todo); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// GET ALL TODO
func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request){
	userID, err := getUserIDFromHeader(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	todos, err := h.service.GetAll(userID)
	if err != nil{
		http.Error(w, "Failed to fetch details", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(todos)
}

// UPDATE TODO
func(h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request){
	userID, err := getUserIDFromHeader(r)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var todo Todo

	if err:= json.NewDecoder(r.Body).Decode(&todo); err != nil{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	todo.UserID = userID
	if err := h.service.Update(&todo); err!= nil{
		http.Error(w, "Failed to Update", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)

}

// DELETE TODO
func(h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request){
	userID, err := getUserIDFromHeader(r)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	idStr := r.URL.Query().Get("id")
	
	id, err := strconv.Atoi(idStr)
	if err != nil{
	  http.Error(w,"Invalid ID", http.StatusBadRequest)
	  return
	}

	if err := h.service.Delete(id, userID); err!= nil{
		http.Error(w, "Failed to Delete TODO", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode([]byte ("Deleted Successfully"))
}

// EXTRACT FIREBASE TOKEN
func getUserIDFromHeader(r *http.Request) (string, error){
	token := r.Header.Get("Authorization")
	if token == ""{
		return "", http.ErrNoCookie
	}

	parse, err := auth.VerifyIdToken(token)
	if err != nil{
		return "", nil
	}

	return parse.UID, nil
}
