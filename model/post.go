package model

import (
	"strings"
	"time"
)

const (
	STATUS_NORMAL = 1
	STATUS_LITTER = 2
)

type Post struct {
	Id        int64  `db:"id"`
	Title     string `db:"title,size:500"`
	Content   string `db:"content"`
	CatId     int64  `db:"cat_id"`
	Pubdate   int64  `db:"pubdate"`
	Tags      string `db:"tags"`
	Uid       int32  `db:"uid"`
	Status    int32  `db:"status,size:2"`
	ViewCount int64  `db:"view_count"`
	Utime     int64  `db:"utime"`
	Ctime     int64  `db:"ctime"`
}

func NewPost(title string, content string, catId int64, pubdate int64, tags string, viewCount int64, status int32, utime int64, uid int32) *Post {
	title = strings.TrimSpace(title)
	tags = strings.TrimSpace(tags)

	return &Post{
		Title:     title,
		Content:   content,
		CatId:     catId,
		Pubdate:   pubdate,
		Tags:      tags,
		Uid:       uid,
		Status:    status,
		ViewCount: viewCount,
		Utime:     utime,
		Ctime:     time.Now().Unix(),
	}
}

func (this *Post) InsertPost() error {
	return Dbmap.Insert(this)
}
