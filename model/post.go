package model

import(
)

type Post struct{
    Id int64 `db:"id"`
    Title string `db:"title,size:500"`
    Content string `db:"content"`
    Pubdate int64 `db:"pubdate"`
    Tags string `db:"tags"`
    Uid int64 `db:"uid"`
    Utime int64 `db:"utime"`
    Status int32 `db:"status,size:2"`
    ViewCount int64 `db:"view_count"`
}

