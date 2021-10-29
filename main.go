package main

import (
	"AltaStore/api"
	"AltaStore/api/middleware"

	// Controller
	cateController "AltaStore/api/v1/category"
	productController "AltaStore/api/v1/product"
	purchaseController "AltaStore/api/v1/purchasereceiving"

	// Service
	cateService "AltaStore/business/category"
	loggerService "AltaStore/business/logger"
	productService "AltaStore/business/product"
	purchaseService "AltaStore/business/purchasereceiving"

	// Repository
	cateRepository "AltaStore/modules/category"
	loggerRepo "AltaStore/modules/logger"
	productRepository "AltaStore/modules/product"
	purchaseRepository "AltaStore/modules/purchasereceiving"
	purchaseDetailRepository "AltaStore/modules/purchasereceivingdetail"

	"context"
	"fmt"
	"time"

	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"AltaStore/config"
	"AltaStore/modules/migration"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabaseConnection(cfg *config.ConfigApp) *gorm.DB {
	stringConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUsername, cfg.DbPassword, cfg.DbName,
	)
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func newMongoDBConnection(cfg *config.ConfigApp) *mongo.Database {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.MongoUsername, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort),
	)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	return client.Database(cfg.MongoDbName)
}

func main() {
	// retrieves application configuration and returns common values when there is a problem
	config := config.GetConfig()

	// Open mongodb logger
	mongoConnection := newMongoDBConnection(config)

	// Register repository
	logrRepo := loggerRepo.NewRepository(mongoConnection)

	// Register service
	logeService := loggerService.NewService(logrRepo)

	// Register logs
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(logeService)

	// open database server base session
	dbConnection := newDatabaseConnection(config)

	// Initiate Respository Category
	categoryRepo := cateRepository.NewRepository(dbConnection)

	// Initiate Service Category
	// categoryService := cateService.NewService(adminService, categoryRepo)
	categoryServc := cateService.NewService(categoryRepo)

	// Initiate Controller Category
	cateControl := cateController.NewController(categoryServc)

	// Initiate Respository Product
	productRepo := productRepository.NewRepository(dbConnection)

	// Initiate Service Product
	// ProductService := productService.NewService(adminService, categoryService, product)
	ProductServc := productService.NewService(categoryServc, productRepo)

	// Initiate Controller Product
	productControl := productController.NewController(ProductServc)

	// initiate Purchase Receiving repository
	purchase := purchaseRepository.NewRepository(dbConnection)
	purchaseDetail := purchaseDetailRepository.NewRepository(dbConnection)

	// initiate Purchase Receiving service
	// purchaseService := purchaseService.NewService(adminService, purchase, purchaseDetail)
	purchaseService := purchaseService.NewService(purchase, purchaseDetail)

	// initiate Purchase Receiving controller
	purchaseControl := purchaseController.NewController(purchaseService)

	// create echo http
	e := echo.New()

	// Register API Path and Controller
	api.RegisterPath(e,
		cateControl,
		productControl,
		purchaseControl,
	)

	lock := make(chan error)

	go func(lock chan error) {
		address := fmt.Sprintf(":%d", config.AppPort)
		lock <- e.Start(address)
	}(lock)

	time.Sleep(1 * time.Millisecond)
	middleware.MakeLogEntry(nil).Info(fmt.Sprintf("Application Start In Port => ::%d", config.AppPort))

	err := <-lock
	if err != nil {
		middleware.MakeLogEntry(nil).Panic("Shutdown Echo Service")
	}
}
