package models

import ("time"
// "gorm.io/gorm"
)

// gorm.Model definition
type History struct {
	ID       	  int           `gorm:"primaryKey"`
	User          User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Userid		  int			
	Vehicle       Vehicle		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Vehicleid	  int			
	City          string         `json:"city"`
	Datestart 	  string		 `json:"datestart"`
	Datefinish	  string		 `json:"datefinish"`
	Createdat 	  time.Time		 `json:"createdat"`
	Updatedat 	  time.Time      `json:"updatedat"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
  }

//   type Vehicle struct {
// 	ID       	  int           `gorm:"primaryKey"`
// 	Name		  string		 `json:"name"`
// 	City	 	  string		 `json:"city"`
// 	Type	 	  string		 `json:"type"`
// 	Createdat 	  time.Time		 `json:"createdat"`
// 	Updatedat 	  time.Time      `json:"updatedat"`
// 	// Histories     []History
// 	// DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

//   type User struct {
// 	ID       	  int           `gorm:"primaryKey"`
// 	Firstname 	  string		 `json:"firstname"`
// 	Lastname	  string		 `json:"lastname"`
// 	City	 	  string		 `json:"city"`
// 	Createdat 	  time.Time		 `json:"created_at"`
// 	Updateat 	  time.Time      `json:"updated_at"`
// 	// Histories     []History
// 	// DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

  type Result struct {
	ID int 
	Firstname string `json:"Customer name"`
	Name  string `json:"Vehicle name"`
	City  string
	Type  string
  }

  type Popular struct {
	Vehicle       Vehicle		`gorm:"foreignkey:Vehicleid";references:vehicleid`
	Vehicleid	  int			 `gorm:"index"`
  }

  type Histories []History
  type Results []Result
  type Populars []Popular