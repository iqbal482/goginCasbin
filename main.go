package main

import (
	"goginCasbin/model"
	"goginCasbin/routes"
	"goginCasbin/seed"
)

func main(){
	db, _ := model.DBConnection()
	seed.Load(db)
	routes.SetupRoutes(db)
}
