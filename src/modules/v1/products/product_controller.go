package products

import ("net/http"
		"fmt"
		"encoding/json")

type product_ctrl struct {
	repo *product_repo
}

func NewCtrl(rep *product_repo) *product_ctrl {
	return &product_ctrl{rep}
}

func (rep *product_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	
	data, err := rep.repo.FindAll()

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}