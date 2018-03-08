package commentController

import (
	"github.com/kataras/iris"
	"irisProject/comment/commentService"
)

func SetEndpoints(app *iris.Application)  {
	app.Get("/comment/{id:int min(1)}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		comment, _ := commentService.FindById(id)

		ctx.JSON(comment)
	})
}