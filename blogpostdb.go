package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var blogPosts BlogPosts

type Post struct { //not sure I'll need these mappings...commented out for now.
	Post_id int    //`gorm:"AUTO_INCREMENT" form:"post_id" json:"post_id"`
	Title   string //`gorm:"not null" form:"title" json:"title"`
	Body    string //`gorm:"not null" form:"body" json:"body"`
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

	// TBD return empty for now
	return BlogPosts{}
}

func CreateBlogPostRecord(bp BlogPost) Post {
	// get connection to db
	db := InitDb()
	defer db.Close()

	// get number of records in posts table so post_id is unique.
	var count int
	var posts []Post
	db.Find(&posts)
	for i, _ := range posts {
		count++
	}

	//fmt.Printf("count = %d\n", count)

	//testing the ORM here...
	blogPostRecord := Post{Post_id: count + 1, Title: "Trial Run", Body: "This is a trial run"}

	db.Save(&blogPostRecord)

	return blogPostRecord
}
