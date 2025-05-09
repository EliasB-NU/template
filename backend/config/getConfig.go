package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database struct {
		Postgres struct {
			Host     string `yaml:"Host"`
			Port     int    `yaml:"Port"`
			User     string `yaml:"User"`
			Password string `yaml:"Password"`
			Database string `yaml:"Database"`
			Timezone string `yaml:"Timezone"`
		} `yaml:"Postgres"`
		Valkey struct {
			Host     string `yaml:"Host"`
			Port     int    `yaml:"Port"`
			User     string `yaml:"User"`
			Password string `yaml:"Password"`
			Database int    `yaml:"Database"`
		} `yaml:"Valkey"`
	} `yaml:"Database"`

	Notifications struct {
		Mail struct {
			Host        string `yaml:"Host"`
			Port        int    `yaml:"Port"`
			User        string `yaml:"User"`
			Password    string `yaml:"Password"`
			SenderEmail string `yaml:"SenderEmail"`
		}
	} `yaml:"Notifications"`

	Storage struct {
		S3 struct {
			Bucket string `yaml:"Bucket"`
			Region string `yaml:"Region"`
			APIKey string `yaml:"ApiKey"`
		}
	} `yaml:"Storage"`
}

// GetConfig command line arguments:
//   - dev: Loads the config.yaml
//   - prod: Loads from os variables
func GetConfig() *Config {
	var config Config

	if os.Args[1] == "dev" {
		file, err := os.Open("config.yaml")
		if err != nil {
			log.Fatalf("Error opening config.yaml: %v\n", err)
		}
		defer file.Close()

		yamlParser := yaml.NewDecoder(file)
		err = yamlParser.Decode(&config)
		if err != nil {
			log.Fatalf("Error parsing config.yaml: %v\n", err)
		}
	} else if os.Args[1] == "prod" {
		// Database Env
		config.Database.Postgres.Host = os.Getenv("POSTGRES_HOST")
		config.Database.Postgres.Port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		config.Database.Postgres.User = os.Getenv("POSTGRES_USER")
		config.Database.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
		config.Database.Postgres.Database = os.Getenv("POSTGRES_DATABASE")
		config.Database.Postgres.Timezone = os.Getenv("POSTGRES_TIMEZONE")
		config.Database.Valkey.Host = os.Getenv("VALKEY_HOST")
		config.Database.Valkey.Port, _ = strconv.Atoi(os.Getenv("VALKEY_PORT"))
		config.Database.Valkey.User = os.Getenv("VALKEY_USER")
		config.Database.Valkey.Password = os.Getenv("VALKEY_PASSWORD")
		config.Database.Valkey.Database, _ = strconv.Atoi(os.Getenv("VALKEY_DATABASE"))
		// Notifications Env
		config.Notifications.Mail.Host = os.Getenv("MAIL_HOST")
		config.Notifications.Mail.Port, _ = strconv.Atoi(os.Getenv("MAIL_PORT"))
		config.Notifications.Mail.User = os.Getenv("MAIL_USER")
		config.Notifications.Mail.Password = os.Getenv("MAIL_PASSWORD")
		config.Notifications.Mail.SenderEmail = os.Getenv("MAIL_SENDER_EMAIL")
		// Storage
		config.Storage.S3.Bucket = os.Getenv("S3_BUCKET")
		config.Storage.S3.Region = os.Getenv("S3_REGION")
		config.Storage.S3.APIKey = os.Getenv("S3_API_KEY")
	} else {
		log.Fatalln("Invalid argument.")
		return nil
	}

	return &config
}
