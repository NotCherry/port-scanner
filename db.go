package main

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Targets = List{list: make(map[string]Target)}

type List struct {
	mu   sync.Mutex
	list map[string]Target
}

type Target struct {
	gorm.Model
	IP    string
	Ports []Port
}

type Port struct {
	gorm.Model
	Port     int
	Filtered bool
	TargetID uint
}

func SaveOutput() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(sqlite.Open(dir+"/scan.db"), &gorm.Config{})
	db = db.Session(&gorm.Session{CreateBatchSize: 1000})

	if err != nil {
		log.Fatal("failed to connect database")
		return
	}

	values := []Target{}
	for _, value := range Targets.list {
		values = append(values, value)
	}

	db.AutoMigrate(&Target{}, &Port{})

	if len(values) > 0 {
		db.Create(&values)
	}
}
