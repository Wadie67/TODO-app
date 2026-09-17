package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() error {
	var err error

	db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.AutoMigrate(&todo{})
}

func addTask(task todo) error {
	return db.Create(&task).Error
}

func getTasks() ([]todo, error) {
	var tasks []todo

	result := db.Find(&tasks)

	return tasks, result.Error
}

func deleteTask(name string) error {
	return db.Where("name = ?", name).Delete(&todo{}).Error
}

func clearTasks() error {
	return db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&todo{}).Error
}

func completeTask(name string) error {
	return db.Model(&todo{}).
		Where("name = ?", name).
		Update("completed", true).Error
}
