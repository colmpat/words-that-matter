package db

import (
	"github.com/colmpat/words-that-matter/pkg/models"
)

func (DB *DB) CreateUser(user models.User) (models.User, error) {
	result, err := DB.NewInsert().
		Model(&user).
		Exec(DB.Ctx)
	if err != nil {
		return user, err
	}

	resID, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.ID = int(resID)
	return user, nil
}

func (DB *DB) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}

	err := DB.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(DB.Ctx, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
