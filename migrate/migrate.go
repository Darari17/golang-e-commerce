package migrate

import (
	"fmt"

	"github.com/Darari17/golang-e-commerce/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("DB Connection is nil")
	}

	err := db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN CREATE TYPE order_status AS ENUM ('pending', 'completed', 'canceled'); END IF; END $$;").Error
	if err != nil {
		return fmt.Errorf("failed to create enum order_status: %v", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
	)
	return err
}
