package main

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var blogPosts BlogPosts

type Post struct {
	Post_id int
	Title   string
	Body    string
}

func InitDb() *gorm.DB {
	// Opening file
	db, err := gorm.Open("sqlite3", "./blog.db")
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}

	return db
}

func GetAllBlogPostRecords(id int) BlogPosts {

	// return empty for now
	return BlogPosts{}
}

func CreateBlogPostRecord(bp BlogPost) Post {
	db := InitDb()
	defer db.Close()

	blogPostRecord := Post{Title: "Trial Run", Body: "This is a trial run"}

	db.Save(&blogPostRecord)

	var post Post

	db.Where("title = ?", "Trial Run").Find(&post)

	return post
}
