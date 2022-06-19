package models

import ("time"
// "gorm.io/gorm"
)

// gorm.Model definition
type Vehicletype struct {
	ID       	  int           `gorm:"primaryKey"`
	Type		  string		 `json:"name"`
	Createdat 	  time.Time		 `json:"createdat"`
	Updatedat 	  time.Time      `json:"updatedat"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
  }

  type Vehiclestypes []Vehicletype