package bunclient

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	_ "github.com/lib/pq"
)

var dbInstance *bun.DB
var once sync.Once

type DatabaseConfig struct {
	Host       string `mapstructure:"DB_HOST"`
	Port       string `mapstructure:"DB_PORT"`
	User       string `mapstructure:"DB_USER"`
	Password   string `mapstructure:"DB_PASSWORD"`
	Name       string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"DB_SSL_MODE"`
	URL        string `mapstructure:"DB_URL"`
	Debug      bool   `mapstructure:"DB_DEBUG"`
	DebugLevel int    `mapstructure:"DB_DEBUG_LEVEL"`
}

func InitDB(config *DatabaseConfig) *bun.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			config.User, config.Password, config.Host, config.Port, config.Name, config.SSLMode)

		sqlDb, err := sql.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}

		dbInstance = bun.NewDB(sqlDb, pgdialect.New())
		dbInstance.AddQueryHook(
			bundebug.NewQueryHook(
				bundebug.WithEnabled(config.Debug),
				bundebug.WithVerbose(config.DebugLevel == 2),
			),
		)
	})

	return dbInstance
}

// GetConn return database connection instance
func GetConn() *bun.DB {
	return dbInstance
}
