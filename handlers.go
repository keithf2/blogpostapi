//
//  Handlers for blogpostapi
//
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Please use http://localhost:8080/post to POST to blog database.\nPlease use http://localhost:8080/posts to GET all blog posts.\n")
}

//
//  GET
//  curl -i http://localhost:8080/posts
//
func GetAllBlogPosts(w http.ResponseWriter, r *http.Request) {

	var posts []Post
	posts = GetAllBlogPostRecords()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

//
// POST
// curl -X POST -d "{\"title\": \"My test post\", \"body\": \"This is my test blog post\"}"  http://localhost:8080/post
//
func CreateBlogPost(w http.ResponseWriter, r *http.Request) {

	//limiting how much we read.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	//After a successful read, time to unmarshal the body into our PostCreate struct
	var blogPost PostCreate
	if err := json.Unmarshal(body, &blogPost); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//Calling the DB layer
	bp := CreateBlogPostRecord(blogPost)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(bp); err != nil {
		panic(err)
	}
}
