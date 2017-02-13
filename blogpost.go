package main

type BlogPost struct {
	Id      string    `json:"post_id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
}

type BlogPosts []BlogPost
