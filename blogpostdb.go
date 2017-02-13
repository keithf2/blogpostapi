package main

import (
	//	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

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

func GetAllBlogPostRecords() []Post {
	// get connection to db
	db := InitDb()
	defer db.Close()

	//find and return all blog posts
	var posts []Post
	db.Find(&posts)

	return posts
}

func CreateBlogPostRecord(pc PostCreate) Post {
	// get connection to db
	db := InitDb()
	defer db.Close()

	// get number of records in posts table so post_id is unique.
	var posts []Post
	db.Find(&posts)
	count := 0
	var id int
	for i, _ := range posts {
		count += i
		id++
	}

	//creating the blog post record
	blogPostRecord := Post{Post_id: id + 1, Title: pc.Title, Body: pc.Body}

	db.Save(&blogPostRecord)

	return blogPostRecord
}
