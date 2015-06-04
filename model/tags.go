package model

import(
	
)

type Tags struct{
	Id int64 `db:"id"`
	Name string `db:"name,size:500"`
	Count int64 `db:"count"`
	Utime int32 `db:"utime"`
	Ctime int32 `db:"ctime"`
}