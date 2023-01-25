package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"sync"

	"test/test/internal/database"
	"test/test/internal/domain"
	"test/test/internal/leveledlog"
)

var logger = leveledlog.NewLogger(os.Stdout, leveledlog.LevelAll, true)

func main() {
	err := run()
	if err != nil {
		logger.Fatal(err, debug.Stack())
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn         string
		automigrate bool
	}
}

type application struct {
	config config
	db     *database.DB
	wg     sync.WaitGroup

	WtfService domain.WtfService
}

func run() error {
	var cfg config
	parseFlags(&cfg)

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config: cfg,
		db:     db,

		WtfService: &database.WtfService{DB: db.DB},
	}

	return app.serveHTTP()
}

func parseFlags(cfg *config) {
	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:4444", "base URL for the application")
	flag.IntVar(&cfg.httpPort, "http-port", 4444, "port to listen on for HTTP requests")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "dev:dev@localhost:5432/dev?sslmode=disable", "postgreSQL DSN")
	flag.BoolVar(&cfg.db.automigrate, "db-automigrate", true, "run migrations on startup")

	flag.Parse()
}

func (app *application) backgroundTask(fn func()) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				logger.Error(fmt.Errorf("%s", err), debug.Stack())
			}
		}()

		fn()
	}()
}
