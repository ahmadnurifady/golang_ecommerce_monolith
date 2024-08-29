package server

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"golang-gorm/internal/delivery/manager"
	"log"

	"github.com/gin-gonic/gin"
	"golang-gorm/internal/config"
	"golang-gorm/internal/delivery/handler"
	"golang-gorm/internal/provider/db"
	"golang-gorm/internal/provider/migration"
)

type Server struct {
	repo   manager.RepoManager
	uc     manager.UsecaseManager
	engine *gin.Engine
	host   string
	cfg    *config.Config
	migra  migration.ModelMigration
}

func (s *Server) setupHandler() {
	group := s.engine.Group("/api/v1")

	s.engine.LoadHTMLGlob("template/*")
	handler.NewHandlerUser(s.uc.UserManager(), group).Route()
	handler.NewHandlerAuth(group).Route()
}

func (s *Server) setupGorm() {
	err := s.migra.Migrate()
	if err != nil {
		fmt.Println(err.Error())
	}

	//err = s.migra.InputData()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

}

func (s *Server) setupGoAuth() {
	store := sessions.NewCookieStore([]byte("key"))

	gothic.Store = store

	goth.UseProviders(google.New(s.cfg.ClientID, s.cfg.ClientSecret, "http://localhost:8080/api/v1/auth/google/callback"))
}

func (s *Server) Run() {
	s.setupHandler()
	s.setupGorm()
	s.setupGoAuth()
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

	migra, err := migration.NewModelMigration(dbConn.Conn())
	if err != nil {
		return nil
	}

	repo := manager.NewRepoManager(dbConn)
	uc := manager.NewUsecaseManager(repo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		repo:   repo,
		uc:     uc,
		host:   host,
		migra:  migra,
		cfg:    cfg,
		engine: engine,
	}
}
