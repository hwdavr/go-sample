package main

import (
	"github.com/hwdavr/perfmon/delivery/http"
	"github.com/hwdavr/perfmon/domain"
	"github.com/hwdavr/perfmon/infrastructure/database"
	"github.com/hwdavr/perfmon/internal/log"
)

func main() {
	log.Setup()
	db := database.NewPsql("prod")
	use := domain.NewUsecase(db)
	svr := http.NewServer(use)
	svr.Start()
}
