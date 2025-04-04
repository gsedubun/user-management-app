package controllers

import (
	"net/http"
	"strconv"
	"user-management-app/models"
	"user-management-app/views"
)

type UserController struct {
	Model *models.UserModel
	Views *views.View
}

func NewUserController(model *models.UserModel, views *views.View) *UserController {
	return &UserController{
		Model: model,
		Views: views,
	}
}

func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := uc.Model.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Users []models.User
	}{
		Title: "All Users",
		Users: users,
	}

	uc.Views.Render(w, "user/index", data)
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uc.Model.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		User  *models.User
	}{
		Title: "User Details",
		User:  user,
	}

	uc.Views.Render(w, "user/show", data)
}

func (uc *UserController) CreateForm(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Create New User",
	}

	uc.Views.Render(w, "user/create", data)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.Model.Create(
		r.FormValue("firstname"),
		r.FormValue("lastname"),
		r.FormValue("email"),
		r.FormValue("address"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (uc *UserController) EditForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uc.Model.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		User  *models.User
	}{
		Title: "Edit User",
		User:  user,
	}

	uc.Views.Render(w, "user/edit", data)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.Model.Update(
		id,
		r.FormValue("firstname"),
		r.FormValue("lastname"),
		r.FormValue("email"),
		r.FormValue("address"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = uc.Model.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
