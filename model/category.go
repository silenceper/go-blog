package model

import(
	
)

type Category struct{
	Id int32 `db:"id"`
	Name string `db:"name,size:500"`
	Description string `db:"description,size:500"`
	Utime int32 `db:"utime"`
	Ctime int32 `db:"ctime"`
}