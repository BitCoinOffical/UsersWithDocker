package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"main.go/internal/models"
	"main.go/internal/storage"
)

type Handler struct {
	storage *storage.DataBase
}

func NewHandler(db *storage.DataBase) *Handler {
	return &Handler{storage: db}
}

func (h *Handler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.storage.GetUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var u models.User
	err = h.storage.GetUser(id, u)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	err := h.storage.CreateUser(u)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&u)

}
func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	idstr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.storage.UpdateUser(id, u)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&u)

}
func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.storage.DeleteUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("User Deleted")
}
