package models

import (
	"time"
)


// модель юзера
type User struct {
	ID			uint 		`gorm:"not null; primarykey; type: BIGINT" json:"id"`
	Username	string 		`gorm:"not null; size:20; unique" json:"username"`
	Password	string 		`gorm:"not null; size:128" json:"-"`
	FIO 		string 		`gorm:"not null; size:50" json:"fio"`
	CreatedAt	time.Time 	`gorm:"not null; autoCreateTime" json:"-"`
	JobTitle 	string 		`gorm:"null; size:50" json:"job_title"`
	
	// ассоциация задач, для которых юзер является исполнителем
	Tasks []Task `gorm:"many2many:boards_task_executors;joinForeignKey:customuser_id;JoinReferences:task_id" json:"-"`
}
// изменение названия таблицы в БД
func (User) TableName() string {
  return "users_customuser"
}
