package sandbox

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jmoiron/sqlx"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", indexHandler)

	db, err := connectDB()
	if err != nil {
		e.Logger.Fatalf("failed to connect db: %v", err)
		return
	}

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// getEnv fetches environment variables. If it doesn't exist, returns default value.
func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

// connectDB connects MySQL database.
func connectDB() (*sqlx.DB, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = getEnv("MYSQL_DB_HOST", "127.0.0.1") + ":" + getEnv("MYSQL_DB_PORT", "3306")
	config.User = getEnv("MYSQL_DB_USER", "foo")
	config.Passwd = getEnv("MYSQL_DB_PASSWORD", "bar")
	config.DBName = getEnv("MYSQL_DB_NAME", "someservice")
	config.ParseTime = true
	dsn := config.FormatDSN()
	return sqlx.Open("mysql", dsn)
}

func indexHandler(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	query := "SELECT name FROM users WHERE id ='" + id "';"
	ret, err := adminDB.ExecContext(ctx, query)
	
	return c.String(http.StatusOK, ret)
}