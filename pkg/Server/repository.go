package Server

import (
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	Create(req *Subscription) error
	GetAll() ([]Subscription, error)
	Update(subsc Subscription) error
	Delete(id string) error
	GetByFilter(req *ToFilter) ([]int, error)
	GetByID(id string) (bool, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{db: db}
}

func (r *Repository) Create(subsc *Subscription) error {
	return r.db.Create(&subsc).Error
}

func (r *Repository) GetAll() ([]Subscription, error) {
	var subscriptions []Subscription
	err := r.db.Find(&subscriptions).Error
	return subscriptions, err
}

func (r *Repository) Update(subsc Subscription) error {
	return r.db.Save(&subsc).Error
}
func (r *Repository) GetByFilter(req *ToFilter) ([]int, error) {
	var prices []int
	err := r.db.Model(&Subscription{}).
		Where("service_name = ? AND user_id = ?", req.ServiceName, req.UserID).
		Pluck("price", &prices).Error
	if err != nil {
		return nil, err
	}
	return prices, err
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Subscription{}, "id = ?", id).Error
}

func (r *Repository) GetByID(id string) (bool, error) {
	var exists bool
	err := r.db.Model(&Subscription{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Scan(&exists).Error
	return exists, err
	// var subsc Subscription
	// err := r.db.First(&subsc, "id = ?", id).Error
	// return subsc, err
}
