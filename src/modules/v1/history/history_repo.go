package history

import ("gorm.io/gorm"
"errors"
"carRent/src/database/gorm/models"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *history_repo {
	return &history_repo{grm}

}

func (r *history_repo) FindAll() (*models.Histories, error) {
	var history models.Histories
	result := r.db.Order("id desc").Preload("User").Preload("Vehicle").Find(&history)

	if result.Error != nil {
		return nil, errors.New("gak bisa ambil data")
	}	

	return &history, nil
}

func (r *history_repo) FindId(id int) (*models.Results, error) {
	var history models.History
	var result models.Results

	data := r.db.Where("histories.userid = ?", id).Model(&history).Select("users.firstname, name, vehicles.city, type").Joins("inner join vehicles on histories.vehicleid = vehicles.id").Joins("inner join users on histories.userid = users.id").Scan(&result)

	if data.Error != nil {
		return nil, errors.New("gak bisa ambil data")
	}	

	return &result, nil
}

func (r *history_repo) FindPopular() (*models.Populars, error) {
	var history models.Histories
	var popular models.Populars

	data := r.db.Order("COUNT(vehicleid) DESC").Model(&history).Select("vehicleid, COUNT(vehicleid) AS MOST_FREQUENT").Group("vehicleid").Preload("Vehicle").Find(&popular)

	if data.Error != nil {
		return nil, errors.New("gak bisa ambil data")
	}	

	return &popular, nil
}

func (r *history_repo) Add(data *models.History) (*models.History, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal nyimpen data")
	}	

	return data, nil
}

func (r *history_repo) Update(data *models.History) (*models.History, error) {
	result := r.db.Save(&data)

	if result.Error != nil {
		return nil, errors.New("gagal nyimpen data")
	}	

	return data, nil
}

func (r *history_repo) Delete(id int) (*models.History, error) {
	var history = models.History{ID: id}

	result := r.db.Delete(&history)

	if result.Error != nil {
		return nil, errors.New("gagagl nyimpen data")
	}	

	return &history, nil
}