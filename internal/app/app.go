package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"app/internal/database"
	"app/internal/services"
	"app/internal/transport"
	"database/sql"
)

func Run() error {
	db, err := sql.Open("postgres", "postgres://tls_monitoring:tls_monitoring_password@db:5432/tls_monitoring?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	dbCert := database.NewDDBCertificateTLB(db)
	serviseCert := services.NewCertificateTLSService(dbCert)
	handlerCert := transport.NewHandler(serviseCert)

	router := gin.Default()
	handlerCert.Register(router)
	transport.RegisterSwagger(router)

	return router.Run(":8080")
}
