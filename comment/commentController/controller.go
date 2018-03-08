package commentController

import (
	"github.com/kataras/iris"
	"golang-web-services-sample/comment/commentService"
	"golang-web-services-sample/comment/commentEntitie"
)

func SetEndpoints(app *iris.Application)  {
	app.Get("/comment/{id:int min(1)}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		comment, _ := commentService.FindById(id)

		ctx.JSON(comment)
	})

	app.Get("/comment/", func(ctx iris.Context) {
		comments, _ := commentService.FindAll()
		ctx.JSON(comments)
	})

	app.Post("/comment/", func(ctx iris.Context) {
		comment :=commentEntitie.Comment{}
		err := ctx.ReadJSON(&comment)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		}else{
			commentService.Create(comment)
			ctx.Writef("Comment created")
		}
	})

	app.Put("/comment/", func(ctx iris.Context) {
		comment :=commentEntitie.Comment{}
		err := ctx.ReadJSON(&comment)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		}else{
			commentService.Update(comment)
			ctx.Writef("Comment created")
		}
	})

	app.Delete("/comment/{id:int min(1)}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		commentService.Delete(id)
	})
}