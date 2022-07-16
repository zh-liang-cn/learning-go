package model

import (
	"github.com/zh-liang-cn/learning-go/02-a-tiny-wiki/common"
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Path    string
	Title   string
	Content string
}

func GetPageByPath(path string) (Page, error) {
	page := Page{}
	tx := common.GetDB().Where("path=?", path).First(&page)

	return page, tx.Error
}

func SavePage(path string, title string, content string) (page *Page) {
	page = &Page{Path: path, Title: title, Content: content}
	common.GetDB().Create(page)
	return
}
