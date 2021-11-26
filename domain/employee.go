package domain

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID           uint64    `gorm:"column:id;primary_key;autoIncrement"`
	Name         string    `gorm:"column:name"`
	DepartmentID uint64    `gorm:"column:department_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func FindById(id uint64) Employee {
	dsn := "host=localhost port=5432 user=practice dbname=practice password=practice"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var employee Employee
	db.Find(&employee, id)
	return employee
}
