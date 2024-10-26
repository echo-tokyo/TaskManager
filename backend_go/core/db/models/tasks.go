package models

import (
	"time"
)


// модель доски
type Board struct {
	ID			uint 		`gorm:"not null; primarykey; type: BIGINT" json:"id"`
	Title 		string 		`gorm:"not null; size:50" json:"title"`
	CreatedAt	time.Time 	`gorm:"not null; autoCreateTime:true" json:"created_at"`
	// ассоциация задач
	Tasks     	[]Task    	`gorm:"foreignKey:BoardId" json:"tasks"`
}
// изменение названия таблицы в БД
func (Board) TableName() string {
  return "boards_board"
}


// модель задачи
type Task struct {
	ID			uint 		`gorm:"not null; primarykey; type: BIGINT" json:"id"`
	Title 		string 		`gorm:"not null; size:50" json:"title"`
	TaskDesc	string 		`gorm:"not null; type: LONGTEXT" json:"task_desc"`
	Status 		string 		`gorm:"not null; size:20; default:backlog" json:"status"`
	BoardId 	int 		`gorm:"not null; type: BIGINT; index" json:"board_id"`
	CreatedAt	time.Time 	`gorm:"not null; autoCreateTime:true" json:"created_at"`
	// ассоциация c доской, к которой принадлежат задачи
	Board     	Board     	`gorm:"foreignKey:BoardId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	// ассоциация исполнителей задач
	Executors []User `gorm:"many2many:boards_task_executors;joinForeignKey:task_id;JoinReferences:customuser_id" json:"executors"`
}
func (Task) TableName() string {
  return "boards_task"
}


// // модель связи задачи и исполнителя
// type TaskExecutor struct {
// 	ID				uint 		`gorm:"not null; primarykey; type: BIGINT" json:"id"`
// 	TaskId			string 		`gorm:"not null; type: BIGINT; unique" json:"taskId"`
// 	CustomuserId 	string 		`gorm:"not null; type: BIGINT; unique" json:"customuserId"`

// 	// ассоциация c Task
// 	// Task     		Task     	`gorm:"foreignKey:TaskId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
// 	// ассоциация c User
// 	// User     		User     	`gorm:"foreignKey:TaskId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
// }
// func (TaskExecutor) TableName() string {
//   return "boards_task_executors"
// }
