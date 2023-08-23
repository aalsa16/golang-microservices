package database

import (
	"fmt"
	"time"

	"github.com/aalsa16/golang-microservices/authentication/types"
	"github.com/google/uuid"
)

func (d *SqlServer) SaveUser(req *types.SignUpRequest) (signupUser *types.SaveResponse, err error) {
	id := uuid.New().String()

	result, err := d.DB.Exec(fmt.Sprintf("INSERT INTO `users` (`email`, `password`, `uuid`, `token`) VALUES ('%s', '%s', '%s', '')", req.Email, req.Password, id))

	if err != nil {
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()

	var row types.User
	query := "SELECT id, email, password, uuid, token, created_at FROM users WHERE id = ?"
	err = d.DB.QueryRow(query, lastInsertID).Scan(&row.ID, &row.Email, &row.Password, &row.Uuid, &row.Token, &row.CreatedAt)

	fmt.Println(err)

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
		Token:     "",
		CreatedAt: unixTimestamp,
	}, nil
}

func (d *SqlServer) GetUser(email string) (user types.User, err error) {
	return types.User{}, nil
}
