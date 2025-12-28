package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"app/internal/database"
	"app/internal/services"
	"app/internal/transport"
	"database/sql"
)

func Run() error {
	dsn := buildPostgresdataSourceName()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	dbCert := database.NewDataBase(db)
	serviceCert := services.NewService(dbCert)
	handlerCert := transport.NewHandler(serviceCert)

	router := gin.Default()
	handlerCert.Register(router)
	transport.RegisterSwagger(router)

	return router.Run(":8080")
}

func buildPostgresdataSourceName() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "postgres"),
	)
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
