package storage

import (
	"fmt"

	"main.go/internal/models"
)

func (db *DataBase) GetUsers() (*[]models.User, error) {
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()
	users := []models.User{}
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return &users, nil
}

func (db *DataBase) GetUser(id int, u models.User) error {
	err := db.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return err
	}
	return nil
}
func (db *DataBase) CreateUser(u models.User) error {
	err := db.DB.QueryRow("INSERT INTO users (name, email) VALUES($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}
func (db *DataBase) UpdateUser(id int, u models.User) error {
	_, err := db.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) DeleteUser(id int) error {
	_, err := db.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
