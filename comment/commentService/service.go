package commentService

import (
	"irisProject/comment/commentModel"
	"irisProject/comment/commentEntitie"
)

func FindById(id int) (commentEntitie.Comment, error) {
	var comment commentEntitie.Comment
	comment, err := commentModel.FindById(id)
	if err != nil {
		return comment, err
	}
	return comment, err
}