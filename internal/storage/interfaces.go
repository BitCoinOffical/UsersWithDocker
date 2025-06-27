package storage

import "main.go/internal/models"

type UserStorage interface {
	GetUsers() (*[]models.User, error)
	GetUser(id int, u models.User) error
	CreateUser(u models.User) error
	UpdateUser(id int, u models.User) error
	DeleteUser(id int) error
}
