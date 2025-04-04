package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"user-management-app/controllers"
	"user-management-app/models"
	"user-management-app/views"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize models and views
	userModel := &models.UserModel{DB: db}
	view := views.NewView()

	// Initialize controllers
	userController := controllers.NewUserController(userModel, view)

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	})
	http.HandleFunc("/users", userController.Index)
	http.HandleFunc("/users/show", userController.Show)
	http.HandleFunc("/users/create", userController.CreateForm)
	http.HandleFunc("/users/create/post", userController.Create)
	http.HandleFunc("/users/edit", userController.EditForm)
	http.HandleFunc("/users/update", userController.Update)
	http.HandleFunc("/users/delete", userController.Delete)

	// Start server
	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
