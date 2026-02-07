package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"project.com/todo/internal/auth"
	"project.com/todo/internal/config"
	"project.com/todo/internal/database"
	"project.com/todo/internal/todo"

	"github.com/go-chi/cors"
)

func main() {
	// 1. LOADING ENVIRONMENT VARIABLES
	myConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to Load .env")
	}

	// 2. CONNECTING POSTGRES(DATABASE)
	myDb, err := database.ConnectDB(myConfig)
	if err != nil {
		log.Fatalf("Failed to Connect to Database: %v", err)
	}
	defer myDb.Close()

	// 3. INITIALIZING FIREBASE
	if err := auth.InitFirebase(myConfig.FirebaseCredBase64); err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// 4 CREATING ROUTE
	myRouter := chi.NewRouter()

	myRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutes
	}))

	// 5.INITIALIZE TODO COMPONENTS
	repo := todo.NewRepository(myDb)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	// 6. SETTING UP ROUTES
	myRouter.Post("/todo", handler.CreateTodo)
	myRouter.Get("/todo", handler.GetAllTodos)
	myRouter.Put("/todo", handler.UpdateTodo)
	myRouter.Delete("/todo", handler.DeleteTodo)

	// 5. DUMMY TEST ROUTE
	myRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// 6. FINALLY, STARTING THE SERVER:
	log.Println("Server running on Port:", myConfig.Port)
	log.Println("DB_HOST:", myConfig.DBHost)
	log.Println("DB_PORT:", myConfig.DBPort)
	log.Println("DB_NAME:", myConfig.DBName)
	log.Println("DB_USER:", myConfig.DBUser)
	log.Println("DB_PASSWORD set:", myConfig.DBPassword != "")

	http.ListenAndServe(":"+myConfig.Port, myRouter)

}
