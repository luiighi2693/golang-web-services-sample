package main

import (
	_ "github.com/lib/pq"
)

func retrieveComment(id int) (comment Comment, err error)  {
	comment = Comment{}
	err = Db.QueryRow("select id, content, author from comments where id = $1",
		id).Scan(&comment.Id, &comment.Content, &comment.Author)
	return
}

func (comment *Comment) createComment() (err error) {
	statement := "insert into comments (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(comment.Content, comment.Author).Scan(&comment.Id)
	return
}

func (comment *Comment) updateComment() (err error) {
	_, err = Db.Exec("update comments set content = $2, author = $3 where id = $1",
		comment.Id, comment.Content, comment.Author)
	return
}

func (comment *Comment) deleteComment() (err error) {
	_, err = Db.Exec("delete from comments where id = $1", comment.Id)
	return
}