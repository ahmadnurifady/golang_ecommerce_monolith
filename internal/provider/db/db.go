package db

import (
	"fmt"
	"golang-gorm/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectDatabase interface {
	Conn() *gorm.DB
}

type connectDatabase struct {
	db  *gorm.DB
	cfg *config.Config
}

func (c *connectDatabase) openConn() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", c.cfg.Host, c.cfg.User, c.cfg.Password, c.cfg.Name, c.cfg.Port, c.cfg.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	c.db = db
	return nil
}

func (c *connectDatabase) Conn() *gorm.DB {
	return c.db
}

func NewConnectDatabase(cfg *config.Config) (ConnectDatabase, error) {
	conn := &connectDatabase{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}
	return conn, nil
}
