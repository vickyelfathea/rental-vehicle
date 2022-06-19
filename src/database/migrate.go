package database

import ("github.com/spf13/cobra"
		"github.com/go-gormigrate/gormigrate/v2"
		"gorm.io/gorm"
		"carRent/src/database/gorm/models"
		"log"
	)

var MigrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "database migration",
	RunE: dbMigrate,
}

func dbMigrate (cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
			&models.Vehicle{},
			&models.History{},
			&models.Vehiclestypes{},
		)

		if err != nil {
			return err
		}

		return nil
	})

	if err := m.Migrate(); err != nil {
		return err
	}
	log.Fatal("init schema successfully")
	return nil
}