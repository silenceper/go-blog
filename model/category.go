package model

import (
	"time"
)

type Category struct {
	Id          int32  `db:"id"`
	Name        string `db:"name,size:500"`
	Description string `db:"description,size:500"`
	Utime       int64  `db:"utime"`
	Ctime       int64  `db:"ctime"`
}

func NewCategory(name, description string, utime int64) *Category {
	return &Category{
		Name:        name,
		Description: description,
		Utime:       utime,
		Ctime:       time.Now().Unix(),
	}
}

func CheckCatExistsById(id int64) bool {
	count,err:=Dbmap.SelectInt("select count(*) from category where id=?",id)
	if err==nil && count==0{
		return false
	}
	return true
}
