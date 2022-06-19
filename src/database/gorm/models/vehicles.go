package models

import ("time"
// "gorm.io/gorm"
)

// gorm.Model definition
type Vehicle struct {
	ID       	  int           `gorm:"primaryKey"`
	Name		  string		 `json:"name"`
	City	 	  string		 `json:"city"`
	Type	 	  string		 `json:"type"`
	Createdat 	  time.Time		 `json:"createdat"`
	Updatedat 	  time.Time      `json:"updatedat"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
  }

  type Vehicles []Vehicle