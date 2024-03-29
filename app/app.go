package app

import (
	"fmt"
	"context"
	"strings"
	// "database/sql"
	"net/http"
	"net/url"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

type UserRow struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
}

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", indexHandler)

	var err error
	db, err = connectDB()
	if err != nil {
		e.Logger.Fatalf("failed to connect db: %v", err)
		return
	}

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

// indexHandler is a request handler for "/".
func indexHandler(c echo.Context) error {
	ctx := context.Background()
	name, _ := url.QueryUnescape(c.QueryParam("name"))
	if name == "" {
		return c.String(http.StatusOK, "Plz set \"name\" parameter.")
	}
	query := "SELECT id, name FROM users WHERE name LIKE '%" + name + "%' AND is_visible = true;"
	query = strings.Replace(query, "%27", "'", -1)
	query = strings.Replace(query, "%22", "\"", -1)
	fmt.Printf("{query: \"%s\"}\n", query)
	users := []UserRow{}
	if err := db.SelectContext(ctx, &users, query); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	
	return c.JSON(http.StatusOK, users)
}