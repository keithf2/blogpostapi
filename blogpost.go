package main

type Post struct {
	Post_id int    `json:"post_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type PostCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
