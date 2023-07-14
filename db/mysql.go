package db

import (
	"fmt"
	"log"
	"os"
	"sharingvision_backendtest/entity"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabaseMysqlConnection() *gorm.DB {
	err := initToYmal()
	if err != nil {
		fmt.Println("error", err.Error())
		panic("Failed read environment yaml")
	}

	dbUser := viper.GetString("database.mysql.user")
	dbPass := viper.GetString("database.mysql.password")
	dbHost := viper.GetString("database.mysql.host")
	dbName := viper.GetString("database.mysql.dbname")
	dbPort := viper.GetString("database.mysql.port")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			IgnoreRecordNotFoundError: false, // Ignore ErrRecordNotFound error for logger
			// SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel: logger.Silent, // Log level
			Colorful: true,          // Disable color
			// Logger: logger.Default.LogMode(logger.Info),
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Println("error", err.Error())
		panic("Failed to connect database mysql")
	}

	if !db.Migrator().HasTable("posts") {
		if err := db.Migrator().CreateTable(&entity.Post{}); err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
		fmt.Println("success create table post")
	}

	fmt.Println("successfully connect to mysql")
	return db
}

func CloseDatabaseMysSqlConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()

}

func initToYmal() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("sharingvision")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
