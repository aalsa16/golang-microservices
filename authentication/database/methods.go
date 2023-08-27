package database

import (
	"time"

	"github.com/aalsa16/golang-microservices/authentication/types"
	"github.com/google/uuid"
)

func (d *SqlServer) SaveUser(req *types.SignUpRequest) (signupUser *types.SaveResponse, err error) {
	id := uuid.New().String()

	query := "INSERT INTO `users` (`email`, `password`, `uuid`) VALUES (?, ?, ?)"

	result, err := d.DB.Exec(query, req.Email, req.Password, id)

	if err != nil {
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()

	var row types.User
	query = "SELECT id, email, password, uuid, created_at FROM users WHERE id = ?"
	err = d.DB.QueryRow(query, lastInsertID).Scan(&row.ID, &row.Email, &row.Password, &row.Uuid, &row.CreatedAt)

	if err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02 15:04:05", row.CreatedAt)

	if err != nil {
		return nil, err
	}

	unixTimestamp := t.Unix()

	return &types.SaveResponse{
		ID:        lastInsertID,
		Email:     req.Email,
		Uuid:      id,
		CreatedAt: unixTimestamp,
	}, nil
}

func (d *SqlServer) GetUser(email string) (user types.User, err error) {
	var row types.User
	query := "SELECT id, email, password, uuid, created_at FROM users WHERE email = ?"
	err = d.DB.QueryRow(query, email).Scan(&row.ID, &row.Email, &row.Password, &row.Uuid, &row.CreatedAt)

	if err != nil {
		return user, err
	}

	return row, nil
}

func (d *SqlServer) SaveRefreshToken(token string, uuid string) error {
	query := "UPDATE users SET refresh_token = ? WHERE uuid = ?"

	_, err := d.DB.Exec(query, token, uuid)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}
