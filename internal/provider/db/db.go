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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		c.cfg.Host, c.cfg.User, c.cfg.Password, c.cfg.Name, c.cfg.Port, c.cfg.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open connection: %v", err)
	}

	// Verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	c.db = db
	return nil
}

func (c *connectDatabase) Conn() *gorm.DB {
	if c.db == nil {
		if err := c.openConn(); err != nil {
			panic(fmt.Sprintf("failed to connect database: %v", err))
		}
	}
	return c.db
}

func NewConnectDatabase(cfg *config.Config) (ConnectDatabase, error) {
	conn := &connectDatabase{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}
	return conn, nil
}
