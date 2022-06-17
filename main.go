package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/kjain0073/go-Todo/tasks"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	var httpAddr = flag.String("http", ":8081", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "tasks",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *mgo.Database
	{
		// to connect with DB
		sess, err := mgo.Dial(tasks.HostName)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		sess.SetMode(mgo.Monotonic, true)
		db = sess.DB(tasks.DbName)
	}

	flag.Parse()
	ctx := context.Background()
	var srv tasks.Service
	{
		repository := tasks.NewRepo(db, logger)

		srv = tasks.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := tasks.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := tasks.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)

}
