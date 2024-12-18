package repositories

import (
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func UserDBInstance(d *gorm.DB) domain.IUserRepository {
	return &userRepo{
		db: d,
	}
}

func (repo *userRepo) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepo) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepo) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *userRepo) Delete(id uint) error {
	return repo.db.Delete(&models.User{}, id).Error
}

func (repo *userRepo) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
