package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Please use http://localhost:8080/post to POST to blog database.\nPlease use http://localhost:8080/posts to GET all blog posts.\n")
}

//
//  curl -i http://localhost:8080/posts
//
func GetAllBlogPosts(w http.ResponseWriter, r *http.Request) {
	blogposts := BlogPosts{
		BlogPost{Id: "1", Title: "My first post", Body: "This is my first blog post"},
		BlogPost{Id: "2", Title: "My second post", Body: "This is my second blog post"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(blogposts); err != nil {
		panic(err)
	}
}

//
//putting the json in a file fixes the glob error.
//curl -d @json.txt -H "Content-Type: application/json" -i  http://localhost:8080/post
//
func CreateBlogPost(w http.ResponseWriter, r *http.Request) {

	var bp = BlogPost{Id: "1", Title: "My first post", Body: "This is my first blog post"}
	CreateBlogPostRecord(bp)
	fmt.Fprintf(w, "TBD\n")
}
