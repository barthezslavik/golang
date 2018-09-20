package main

import (
	"github.com/kataras/iris/_examples/tutorial/vuejs-todo-mvc/src/todo"
	"github.com/kataras/iris/_examples/tutorial/vuejs-todo-mvc/src/web/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"

	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.StaticWeb("/", "./public")

	sess := sessions.New(sessions.Config{
		Cookie: "iris_session",
	})

	ws := websocket.New(websocket.Config{})
	todosRouter := app.Party("/todos")
	todosRouter.Any("/iris-ws.js", websocket.ClientHandler())

	todosApp := mvc.New(todosRouter)

	todosApp.Register(
		todo.NewMemoryService(),
		sess.Start,
		ws.Upgrade,
	)

	todosApp.Handle(new(controllers.TodoController))
	app.Run(iris.Addr(":1234"), iris.WithoutVersionChecker)
}
