package products

import ("gorm.io/gorm"
"errors"
)

type product_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *product_repo {
	return &product_repo{grm}

}

func (r *product_repo) FindAll() (*Products, error) {
	var product Products

	result := r.db.Find(&product)

	if result.Error != nil {
		return nil, errors.New("gak bisa ambil data")
	}	

	return &product, nil
}