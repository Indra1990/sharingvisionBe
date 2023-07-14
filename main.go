package main

import (
	"context"
	"sharingvision_backendtest/app"
	"sharingvision_backendtest/db"

	"gorm.io/gorm"
)

var (
	dbmysql *gorm.DB = db.SetupDatabaseMysqlConnection()
)

func main() {
	defer db.CloseDatabaseMysSqlConnection(dbmysql)
	ctx := context.Background()

	server := app.NewServer(dbmysql)
	server.Start(ctx)
	server.Stop(ctx)
	return
}
