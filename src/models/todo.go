package models

import (
	"Starter/config"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	UserID uint64
	Title  string
}

func (todo *Todo) TableName() string {
	return "todo"
}

func CreateTodo(todo *Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodo(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}

	return nil
}

func GetTodoById(todo *Todo, id string) (err error) {
	if err = config.DB.Where("user_id = ?", id).First(todo).Error; err != nil {
		return err
	}
	fmt.Println(todo)
	return nil
}

func UpdateTodo(todo *Todo, id string) (err error) {
	err = config.DB.Model(todo).Where("user_id = ?", id).Update(Todo{
		Title: "Detail todo in here",
	}).Error

	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Data not found")
	}

	fmt.Println(gorm.IsRecordNotFoundError(err))
	return
}

func DeleteTodo(todo *Todo, id string) (err error) {
	if err = config.DB.Where("user_id = ?", id).Delete(todo).Error; err != nil {
		return err
	}

	return nil
}
