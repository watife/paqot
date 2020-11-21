package main

import (
	"fmt"
	"github.com/fakorede-bolu/deliva/api/courier"
	"github.com/fakorede-bolu/deliva/api/customer"
	"github.com/fakorede-bolu/deliva/api/jobs"
	"github.com/fakorede-bolu/deliva/pkg/logs"
	"github.com/fakorede-bolu/deliva/pkg/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)
import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}


func main ()  {
	databaseName := os.Getenv("DATABASE_NAME")
	databaseUser := os.Getenv("DATABASE_USER")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	appAddr := ":" + os.Getenv("PORT")

	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, databaseUser, databaseName, databasePassword)

	db := openDB(dbConn)

	r := mux.NewRouter()

	//COURIER
	courierRepo := courier.NewCourierPg(db)
	courierService := courier.NewService(courierRepo)
	courier.MakeCourierHandlers(r, courierService)

	//CUSTOMER
	customerRepo := customer.NewCustomerPg(db)
	customerService := customer.NewService(customerRepo)
	customer.MakeCustomerHandlers(r, customerService)

	//JOBS
	JobsRepo := jobs.NewJobsPg(db)
	jobService := jobs.NewService(JobsRepo, courierService, customerService)
	jobs.MakeJobsHandlers(r, jobService)

	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middleware.SecureHeaders)
	r.Use(middleware.LogRequest)
	r.Use(middleware.RecoverPanic)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         appAddr,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logs.ErrorLog,
	}
	logs.InfoLog.Printf("Starting server on %s", appAddr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}



func openDB(database string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err)
	}
	if err := migrateDB(db); err != nil {
		panic(err)
	}
	return db
}
func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(courier.Courier{}, customer.Customer{}, jobs.Jobs{})
}


