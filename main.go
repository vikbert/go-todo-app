package main

import (
	"github.com/vikbert/go-todo-app/api/router"
	"github.com/vikbert/go-todo-app/dbconfig"
	"github.com/vikbert/go-todo-app/repository"
)

func main() {
	db := dbconfig.ConnectPostgreSQL()
	todoRepo := repository.NewTodoOrmRepository(db)

	// ------------------------------------------------
	// Context 5x
	// ------------------------------------------------

	// Gin
	//router.GinRouter(todoRepo)

	// fiber
	//router.FiberRouter(todoRepo)

	// echo
	//router.EchoRouter(todoRepo)

	// iris
	//router.IrisRouter(todoRepo)

	// FastHttp + fastHttp/router
	//router.FastHttpRouter(todoRepo)

	// ------------------------------------------------
	// net/php 3x
	// ------------------------------------------------

	// chi
	//router.ChiRouter(todoRepo)

	// Mux (Gorilla)
	router.MuxRouter(todoRepo)

	// HttpRouter
	//router.HttpRouterRouter(todoRepo)
}
