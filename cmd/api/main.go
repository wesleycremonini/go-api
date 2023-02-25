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
	domainName string
	email      string
	db         struct {
		dsn         string
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

	db, err := database.New(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config: cfg,
		db:     db,

		WtfService: &database.WtfService{DB: db.DB},
	}

	return app.run()
}

func parseFlags(cfg *config) {
	flag.StringVar(&cfg.domainName, "domain-name", "localhost", "base URL for the application")
	flag.StringVar(&cfg.email, "email", "example@email.com", "application TLS certificate e-mail")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "dev:dev@localhost:5432/go_db?sslmode=disable", "postgreSQL DSN")

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
