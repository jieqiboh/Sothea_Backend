package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_patientHttpDelivery "github.com/jieqiboh/sothea_backend/patient/delivery/http"
	_patientPostgresRepository "github.com/jieqiboh/sothea_backend/patient/repository/postgres"
	_patientUseCase "github.com/jieqiboh/sothea_backend/patient/usecase"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	// Initialize global variables
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)
	dbSslMode := viper.GetString(`database.sslmode`)

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbSslMode)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// You might want to check the connection here to handle errors
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	patientRepo := _patientPostgresRepository.NewPostgresPatientRepository(db)
	patientUseCase := _patientUseCase.NewPatientUsecase(patientRepo, 2*time.Second)
	_patientHttpDelivery.NewPatientHandler(router, patientUseCase)
	router.Run("localhost:9090")
}
