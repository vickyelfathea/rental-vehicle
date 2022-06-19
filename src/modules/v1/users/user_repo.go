package users

import ("gorm.io/gorm"
"errors"
"carRent/src/database/gorm/models"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *user_repo {
	return &user_repo{grm}

}

func (r *user_repo) FindAll() (*models.Users, error) {
	var user models.Users

	result := r.db.Order("id desc").Find(&user)

	if result.Error != nil {
		return nil, errors.New("gak bisa ambil data")
	}	

	return &user, nil
}

func (r *user_repo) Add(data *models.User) (*models.User, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return data, nil
}

func (r *user_repo) Update(data *models.User) (*models.User, error) {
	// var user = User{ID: data.ID}

	result := r.db.Save(&data)
	
	// fmt.Println(data)

	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return data, nil
}

func (r *user_repo) Delete(id int) (*models.User, error) {
	var user = models.User{ID: id}

	result := r.db.Delete(&user)
	
	// fmt.Println(data)

	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return &user, nil
}

func (r *user_repo) FindId(id int) (*models.User, error) {
	var user = models.User{ID: id}

	result := r.db.Find(&user)
	
	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return &user, nil
}

func (r *user_repo) FindByUsername(username string) (*models.User, error) {
	var data models.User

	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return nil, errors.New("username tidak ditemukan")
	}
	return &data, nil
}