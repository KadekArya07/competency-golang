package config

import (
	"competency/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var tables = []interface{}{
	&model.LovCompetency{},
	&model.Competency{},
	&model.Behaviour{},
	&model.Proficiency{},
	&model.Training{},
	&model.TrainingDetail{},
	&model.Concern{},
	&model.NonTraining{},
	&model.JobCompetency{},
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "competency_hr"
	sslmode  = "disable"
)

func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
