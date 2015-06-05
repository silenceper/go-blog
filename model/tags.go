package model

import (
	"time"
)

type Tags struct {
	Id    int64  `db:"id"`
	Name  string `db:"name,size:500"`
	Count int64  `db:"count"`
	Utime int64  `db:"utime"`
	Ctime int64  `db:"ctime"`
}

func NewTags(name string, count, utime int64) *Tags{
	return &Tags{
		Name:  name,
		Count: count,
		Utime: utime,
		Ctime: time.Now().Unix(),
	}
}

//始终能够获取到id 没有则新增
func GetTagIdByNameAlways(name string) (int64, error) {
	id, err := GetTagIdByName(name)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		//insert
		tagsObj := NewTags(name, 0, time.Now().Unix())
		err = Dbmap.Insert(tagsObj)
		if err!=nil{
			return 0,err
		}
		return tagsObj.Id,nil
	} else {
		return id, nil
	}
}

//根据name 获取tagId
func GetTagIdByName(name string) (int64, error) {
	return Dbmap.SelectInt("select id from tags where name=?", name)
}

//更新count数
func UpdateTagCountById(id int64)error{
	_,err:=Dbmap.Exec("update tags set count=`count`+1 where id=?",id)
	return err
}

//检测是否存在
func CheckTagExists(name string) bool {
	count, err := Dbmap.SelectInt("select count(*) from tags where name=?", name)
	if err == nil && count == 0 {
		return false
	}
	return true
}
