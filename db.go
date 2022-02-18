package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Targets = make(map[string]Target)

type Target struct {
	gorm.Model
  IP string
	Ports []Port
}

type Port struct {
	gorm.Model
	Port int
	TargetID uint
}

func Save_to_db() {
	db, err := gorm.Open(sqlite.Open("scan.db"), &gorm.Config{})

	if err != nil {
    panic("failed to connect database")
  }

	values := []Target{}
	for _, value := range Targets {
			values = append(values, value)
	}

	db.AutoMigrate(&Target{}, &Port{})
	db.Create(&values)

}