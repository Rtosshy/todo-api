package repository

import (
	"fmt"
	"todo-api/model"
	"todo-api/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userID types.UserID) error
	GetTaskById(task *model.Task, userID types.UserID, taskID types.TaskID) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userID types.UserID, taskID types.TaskID) error
	DeleteTask(userID types.UserID, taskID types.TaskID) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userID types.UserID) error {
	if err := tr.db.Joins("User").Where("user_id=?", userID).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userID types.UserID, taskID types.TaskID) error {
	if err := tr.db.Joins("User").Where("user_id=?", userID).First(task, taskID).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userID types.UserID, taskID types.TaskID) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskID, userID).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userID types.UserID, taskID types.TaskID) error {
	result := tr.db.Where("id=? AND user_id=?", taskID, userID).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
