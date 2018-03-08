package commentModel

import (
	"database/sql"
	_ "github.com/lib/pq"
	"irisProject/comment/commentEntitie"
)

var Db *sql.DB

func init()  {
	var err error
	Db, err = sql.Open("postgres",
		"postgres://postgres:postgres@localhost/admin?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func FindById(id int) (comment commentEntitie.Comment, err error){
	comment = commentEntitie.Comment{}
	err = Db.QueryRow("select id, content, author from comments where id = $1",
		id).Scan(&comment.Id, &comment.Content, &comment.Author)
	return comment, nil
}
