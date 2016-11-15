package main

import (
	"github.com/adamar/delta-server/delta"
	"github.com/adamar/delta-server/models"
)

func main() {

	DBi := delta.SetupGormDB()
	//if err != nil {
	//	panic(err)
	//}

	DBi.AutoMigrate(&models.Event{})

}
