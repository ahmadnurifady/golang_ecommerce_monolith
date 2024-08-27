package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-gorm/internal/config"
	"golang-gorm/internal/delivery/handler"
	"golang-gorm/internal/provider/db"
	"golang-gorm/internal/repository"
	"golang-gorm/internal/usecase"
	"log"
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
	err := s.engine.Run(s.host)
	if err != nil {
		log.Fatal("server can't run")
	}

}

func NewServer() *Server {

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	db, err := db.NewConnectDatabase(cfg)
	if err != nil {
		fmt.Println(err)
	}

	repo := repository.NewRepositoryUser(db.Conn())
	uc := usecase.NewUsecaseUser(repo)
	engine := gin.Default()

	return &Server{repo: repo, uc: uc, cfg: cfg, engine: engine}

}
