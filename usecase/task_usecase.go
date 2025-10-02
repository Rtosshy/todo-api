package usecase

import (
	"todo-api/model"
	"todo-api/repository"
	"todo-api/types"
	"todo-api/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userID types.UserID) ([]model.TaskResponse, error)
	GetTaskById(userID types.UserID, taskID types.TaskID) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userID types.UserID, taskID types.TaskID) (model.TaskResponse, error)
	UpdateTaskStatus(task model.Task, userID types.UserID, taskID types.TaskID) (model.TaskResponse, error)
	DeleteTask(userID types.UserID, taskID types.TaskID) error
}

func toTaskResponse(t model.Task) model.TaskResponse {
	return model.TaskResponse{
		ID:        t.ID,
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userID types.UserID) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userID); err != nil {
		return nil, err
	}
	resTasks := make([]model.TaskResponse, len(tasks))
	for i, v := range tasks {
		resTasks[i] = toTaskResponse(v)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userID types.UserID, taskID types.TaskID) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := toTaskResponse(task)
	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := toTaskResponse(task)
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userID types.UserID, taskID types.TaskID) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTask(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := toTaskResponse(task)
	return resTask, nil
}

func (tu *taskUsecase) UpdateTaskStatus(task model.Task, userID types.UserID, taskID types.TaskID) (model.TaskResponse, error) {
	if err := tu.tr.UpdateTaskStatus(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := toTaskResponse(task)
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userID types.UserID, taskID types.TaskID) error {
	if err := tu.tr.DeleteTask(userID, taskID); err != nil {
		return err
	}
	return nil
}
