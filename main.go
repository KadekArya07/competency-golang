package main

import (
	"competency/config"
	"competency/controller"
	"competency/dao"
	"competency/server"
	"competency/service"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	g := initDb()
	url := e.Group("/api")

	dao.SetDao(g)
	go server.SetProto()
	service.SetService(g)
	config.Connected()
	controller.SetInit(e)
	controller.SetCompetency(url)
	controller.SetLov(url)
	controller.SetJob(url)
	e.Logger.Fatal(e.Start(":1234"))
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}
