package persistence

import (
	"database/sql"
	"myapp/domain"
	"myapp/domain/repository"

	"gorm.io/gorm"
)

type userPersistence struct {
	DB *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.IUserRepository {
	return &userPersistence{DB: db}
}

func (up *userPersistence) Insert(name string, email string) error {
	stmt, err := up.DB.Raw("INSERT INTO users (name, email) VALUES (?, ?)", name, email).Rows()
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func (up *userPersistence) GetByUserID(ID string) (*domain.User, error) {
	stmt := up.DB.Raw("SELECT * FROM users WHERE id = ?", ID).Row()
	return convertToDomainUser(stmt)
}

func convertToDomainUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
