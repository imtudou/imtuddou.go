package models

import (
	"2.gowebsample/chapter21_todo/todo_api/dao"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string
	Status bool
}

/*

modeTodo.go 这个文件的所有操作都放这里
*/

func Create(todo *Todo) (err error) {
	return dao.DB.Create(&todo).Error
}

func Update(todo *Todo) (err error) {
	return dao.DB.Save(&todo).Error
}

// 如果你的模型包含了 gorm.DeletedAt字段 当调用Delete时，GORM并不会从数据库中删除该记录，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录
func Delete(id uint) (err error) {
	return dao.DB.Delete(&Todo{}, id).Error
}

func Get(id uint) (todo *Todo, err error) {
	todo.ID = id
	if dao.DB.First(&todo).Error != nil {
		return nil, err
	}
	return todo, nil
}

func GetList() (todos []*Todo, err error) {
	if dao.DB.Find(&todos).Error != nil {
		return nil, err
	}
	return todos, nil
}
