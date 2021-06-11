package db

import (
	"flag"
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
var IsProduction *bool

type config struct{
	Secret string `mapstructure:"SECRET"`
	DbHost string `mapstructure:"DB_HOST"`
	DbUser string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName string `mapstructure:"DB_NAME"`
	DbURL string `mapstructure:"DB_URL"`
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
	
	IsProduction = flag.Bool("production", true, "decides if the project is in development or in production")

	var connectionString string 
	if *IsProduction == true {
		connectionString = 	os.Getenv("DATABASE_URL")
	} else {
		envs := LoadEnv()
		connectionString = envs.DbURL
	}

	
	//fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", envs.DbHost, envs.DbUser, envs.DbPassword, envs.DbName)

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

