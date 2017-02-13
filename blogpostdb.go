package main

import (
	//"fmt"

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

	db := InitDb()
	defer db.Close()

	var posts []Post
	db.Find(&posts)
	//	for i, _ := range posts {
	//		fmt.Printf("i = %d\n", i)
	//		fmt.Printf("posts[%d].Post_id %d ", i, posts[i].Post_id)
	//	}
	return posts
}

func CreateBlogPostRecord(bp PostCreate) Post {
	// get connection to db
	db := InitDb()
	defer db.Close()

	// get number of records in posts table so post_id is unique.
	var count int
	var posts []Post
	db.Find(&posts)
	for i, _ := range posts {
		count = i
	}

	count++
	//	fmt.Printf("count = %d\n", count)
	//	fmt.Printf("bp.Title = %s\n", bp.Title)
	//	fmt.Printf("bp.Body = %s\n", bp.Body)

	//testing the ORM here...
	blogPostRecord := Post{Post_id: count + 1, Title: bp.Title, Body: bp.Body}

	db.Save(&blogPostRecord)

	return blogPostRecord
}

//func insert(original []Post, position int, value int) []Post {
//	l := len(original)
//	target := original
//	if cap(original) == l {
//		target = make([]Post, l+1, l+10)
//		copy(target, original[:position])
//	} else {
//		target = append(target, -1)
//	}
//	copy(target[position+1:], original[position:])
//	target[position] = value
//	return target
//}
