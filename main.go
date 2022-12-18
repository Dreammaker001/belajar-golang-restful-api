package main

import (
	"dreammaker001/belajar-golang-restful-api/app"
	"dreammaker001/belajar-golang-restful-api/controller"
	"dreammaker001/belajar-golang-restful-api/helper"
	"dreammaker001/belajar-golang-restful-api/middleware"
	"dreammaker001/belajar-golang-restful-api/repository"
	"dreammaker001/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
