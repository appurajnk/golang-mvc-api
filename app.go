package main

import (
	"os"

	"github.com/iris-contrib/middleware/cors"
	R "github.com/mirzafaizan/gom-api/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})
	app.Use(crs)
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panicsR
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET Default Endpoint
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Welcome to my API"})
	})

	// API endpoints
	app.AllowMethods(iris.MethodOptions)
	api := app.Party("/api")
	{
		api.Post("/signup", R.CreateUser)
		api.Post("/login", R.GetUser)
		api.Get("/getusers", R.GetAllUsers)
		api.Get("/users/{msisdn}", R.GetUser)
		api.Get("/loans", R.GetAllLoans)
		api.Get("/loantotal", R.CountLoans)
	}

	app.Listen(os.Getenv("PORT"))
}
