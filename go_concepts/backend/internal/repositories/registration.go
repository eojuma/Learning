package repositories

import (
	"backend-001/internal/models"
	"context"
	"database/sql"
	"errors"
	"strings"
)

type newsqlRepository struct {
	db *sql.DB
}

func (r *newsqlRepository) Create(ctx context.Context, u *models.User) error {
	const query = `INSERT INTO users(id,name,username,email_address,phone_number,id_number,hashed_password) VALUES(?,?,?,?,?,?,?)`

	_, err := r.db.ExecContext(ctx, query, u.ID.String(), u.Name, u.Username, u.EmailAddress, u.PhoneNumber, u.IDNumber, u.HashedPassword)

	if err != nil {
		if strings.Contains(err.Error(), "users.email_address") {
			return models.ErrEmailTaken
		} else if strings.Contains(err.Error(), "users.username") {
			return models.ErrUsernameTaken
		} else if strings.Contains(err.Error(), "users.phone_number") {
			return models.ErrPhoneNumberTaken
		} else if strings.Contains(err.Error(), "users.id_number") {
			return models.ErrIDNumberTaken
		} else {

			return err
		}
	}
	return nil
}

func (r *newsqlRepository) GetByEmail(ctx context.Context,email string)(*models.User,error){

const query = `SELECT id,name,username,email_address,phone_number,id_number,hashed_password FROM users WHERE email_address = ?`
var e models.User

err:=r.db.QueryRowContext(ctx,query,email).Scan(&e.ID,&e.Name,&e.Username,&e.EmailAddress,&e.PhoneNumber,&e.IDNumber,&e.HashedPassword)

if err !=nil{
	if errors.Is(err,sql.ErrNoRows){
		return nil,models.ErrUserNotFound
	}
}
	return &e,nil
}