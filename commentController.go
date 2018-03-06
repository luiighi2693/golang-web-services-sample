package main

import (
	"net/http"
	"strconv"
	"path"
	"encoding/json"
)

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func handleRequestComment(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = getComment(w, r)
	case "POST":
		err = createComment(w, r)
	case "PUT":
		err = updateComment(w, r)
	case "DELETE":
		err = deleteComment(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getComment(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	comment, err := retrieveComment(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&comment, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func createComment(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var comment Comment
	json.Unmarshal(body, &comment)
	err = comment.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
func updateComment(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	comment, err := retrieveComment(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &comment)
	err = comment.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func deleteComment(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	comment, err := retrieveComment(id)
	if err != nil {
		return
	}
	err = comment.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

