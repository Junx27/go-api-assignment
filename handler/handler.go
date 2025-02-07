package handler

import (
	"encoding/json"
	"go-api-assignment/model"
	"go-api-assignment/repository"
	"net/http"
	"strconv"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}
func (h *UserHandler) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if user.Age <= 0 {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	existingUser, _ := h.repo.GetUserByName(user.Name)
	if existingUser.ID != 0 {
		http.Error(w, "User with this name already exists", http.StatusBadRequest)
		return
	}

	err = h.repo.AddUser(user)
	if err != nil {
		http.Error(w, "Error adding user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
