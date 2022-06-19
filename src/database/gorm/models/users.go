package models

import ("time"
// "gorm.io/gorm"
)

// gorm.Model definition
type User struct {
	ID       	  int           `gorm:"primaryKey"`
	Firstname 	  string		 `json:"firstname"`
	Lastname	  string		 `json:"lastname"`
	Username	  string		 `json:"username"`
	Password	  string		 `json:"password"`
	City	 	  string		 `json:"city"`
	Email 		  string 		 `validate:"email"`
	Createdat 	  time.Time		 `json:"created_at"`
	Updateat 	  time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
  }

  type Users []User