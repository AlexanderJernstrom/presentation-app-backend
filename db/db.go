package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

type config struct{
	Secret string `mapstructure:"SECRET"`
	DbHost string `mapstructure:"DB_HOST"`
	DbUser string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName string `mapstructure:"DB_NAME"`
}

func LoadEnv() (envs config){
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	var configuration config


	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}


	return configuration
}


func Connect() {
	fmt.Println(os.Getenv("SECRET"))
	envs := LoadEnv()

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", envs.DbHost, envs.DbUser, envs.DbPassword, envs.DbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold: time.Second,   // Slow SQL threshold
		  LogLevel:      logger.Info, // Log level
		  Colorful:      false,         // Disable color
		},
	  )

	Db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err)
	}
	
	//Db.AutoMigrate(&models.User{}, &models.Presentation{}, &models.Slide{}, &models.Element{})

	
	
}

