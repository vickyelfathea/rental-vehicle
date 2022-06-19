package vehicles

import ("gorm.io/gorm"
"errors"
"carRent/src/database/gorm/models"
// "fmt"
)

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *vehicle_repo {
	return &vehicle_repo{grm}

}

func (r *vehicle_repo) FindAll() (*models.Vehicles, error) {
	var vehicle models.Vehicles

	result := r.db.Order("id desc").Find(&vehicle)

	if result.Error != nil {
		return nil, errors.New("gak bisa ambil data")
	}	

	return &vehicle, nil
}

func (r *vehicle_repo) Add(data *models.Vehicle) (*models.Vehicle, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return data, nil
}

func (r *vehicle_repo) Update(data *models.Vehicle) (*models.Vehicle, error) {
	// var vehicle = vehicle{ID: data.ID}

	result := r.db.Save(&data)
	
	// fmt.Println(data)

	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return data, nil
}

func (r *vehicle_repo) Delete(id int) (*models.Vehicle, error) {
	var vehicle = models.Vehicle{ID: id}

	result := r.db.Delete(&vehicle)
	
	// fmt.Println(data)

	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return &vehicle, nil
}

func (r *vehicle_repo) FindId(id int) (*models.Vehicle, error) {
	var vehicle = models.Vehicle{ID: id}

	result := r.db.Find(&vehicle)
	
	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return &vehicle, nil
}