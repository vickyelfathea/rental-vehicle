package products

import ("time"
// "gorm.io/gorm"
)

// gorm.Model definition
type Product struct {
	ID        uint           `gorm:"primaryKey"`
	Name 	  string		 `json:"name"`
	Price	  string		 `json:"price"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
  }

  type Products []Product