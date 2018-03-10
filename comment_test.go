package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
	"golang-web-services-sample/comment/commentEntitie"
	"encoding/json"
	"strconv"
	"bytes"
	"fmt"
)

// $ go test -v
func TestNewApp(t *testing.T) {
	app := NewApp()
	e := httptest.New(t, app)

	comment := commentEntitie.Comment{
		Content:"content",
		Author:"author" }

	id, _ := strconv.Atoi(e.POST("/comment").WithJSON(comment).Expect().
		Status(httptest.StatusOK).Body().Raw())


	comment = commentEntitie.Comment{
		Id: id,
		Content:"content",
		Author:"author" }

	jsonComment, _ := json.Marshal(comment)

	var pathBuffer bytes.Buffer

	pathBuffer.WriteString("/comment/")
	pathBuffer.WriteString(string(id))

	fmt.Printf(pathBuffer.String())

//	e.GET(pathBuffer.String()).Expect().
//		Status(httptest.StatusOK).Body().Equal(string(jsonComment))

}