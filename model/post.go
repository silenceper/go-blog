package model

import(
)

type Post struct{
    Id int64 `db:"id"`
    Title string `db:"title,size:500"`
    Content string `db:"content"`
	CatId	int32 `db:"cat_id"`
    Pubdate int32 `db:"pubdate"`
    Tags string `db:"tags"`
    Uid int32 `db:"uid"`
    Status int32 `db:"status,size:2"`
    ViewCount int64 `db:"view_count"`
	Utime int32 `db:"utime"`
	Ctime int32 `db:"ctime"`
}

