package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang-gorm/internal/config"
	"golang-gorm/internal/delivery/handler"
	"golang-gorm/internal/provider/db"
	"golang-gorm/internal/provider/migration"
	"golang-gorm/internal/repository"
	"golang-gorm/internal/usecase"
)

type Server struct {
	repo   repository.RepositoryUser
	uc     usecase.UsecaseUser
	engine *gin.Engine
	host   string
	cfg    *config.Config
}

func (s *Server) setupHandler() {
	group := s.engine.Group("/api/v1")
	handler.NewHandlerUser(s.uc, group).Route()
}

func (s *Server) Run() {
	s.setupHandler()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatalf("server can't run: %v", err.Error())
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dbConn, err := db.NewConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = migration.NewModelMigration(dbConn.Conn()).Migrate(dbConn.Conn())
	if err != nil {
		return nil
	}

	repo := repository.NewRepositoryUser(dbConn.Conn())
	uc := usecase.NewUsecaseUser(repo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		repo: repo,
		uc:   uc,
		host: host,

		cfg:    cfg,
		engine: engine,
	}
}
