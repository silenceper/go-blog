package service

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"../model"
)

//添加文章
func WritePost(title, content, tagStr string, pubdate int64, catId int64, status int32) error {
	//判断catId是否存在
	exists := model.CheckCatExistsById(catId)
	if !exists {
		return errors.New("不存在该栏目分类")
	}
	//根据tag取tagid
	tagsArr := strings.Split(tagStr, ",")
	var tagIdArr []string
	for _, tag := range tagsArr {
		id, err := model.GetTagIdByNameAlways(tag)
		if err != nil {
			return err
		}
		tagIdArr = append(tagIdArr, strconv.FormatInt(id, 10))

		//更新
		model.UpdateTagCountById(id)
	}
	tagIdStr := strings.Join(tagIdArr, ",")

	//添加文章
	post := model.NewPost(title, content, catId, pubdate, tagIdStr, 0, model.STATUS_NORMAL, time.Now().Unix(), 1)
	err := post.InsertPost()
	if err != nil {
		return err
	}
	return nil
}
