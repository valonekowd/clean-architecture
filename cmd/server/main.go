package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/oklog/run"

	httpDelivery "github.com/valonekowd/clean-architecture/adapter/delivery/http"
	"github.com/valonekowd/clean-architecture/adapter/endpoint"
	"github.com/valonekowd/clean-architecture/adapter/formatter"
	"github.com/valonekowd/clean-architecture/adapter/repository"
	"github.com/valonekowd/clean-architecture/config"
	"github.com/valonekowd/clean-architecture/infrastructure/auth"
	"github.com/valonekowd/clean-architecture/infrastructure/datastore/sql"
	"github.com/valonekowd/clean-architecture/infrastructure/log/zap"
	"github.com/valonekowd/clean-architecture/infrastructure/validator/playground"
	"github.com/valonekowd/clean-architecture/usecase"
)

func main() {
	// load config
	cfg, err := config.Create()
	if err != nil {
		log.Fatalf("env=dev method=config.Create err=%v", err)
	}

	// setup logging
	logger, err := zap.NewLogger(cfg.IsProd())
	if err != nil {
		log.Fatalf("log=zap method=zap.NewLogger err=%v", err)
	}

	// setup primary database
	db, err := sql.Connect(
		sql.WithDriverName(cfg.Datastore.Primary.DriverName),
		sql.WithHost(cfg.Datastore.Primary.Host),
		sql.WithPort(cfg.Datastore.Primary.Port),
		sql.WithUsername(cfg.Datastore.Primary.Username),
		sql.WithPassword(cfg.Datastore.Primary.Password),
		sql.WithDBName(cfg.Datastore.Primary.DBName),
	)
	if err != nil {
		logger.Fatal("datastore", "primary", "driver", cfg.Datastore.Primary.DriverName, "method", "sql.Connect", "err", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Log("datastore", "primary", "driver", cfg.Datastore.Primary.DriverName, "method", "db.Close", "err", err)
		} else {
			logger.Log("datastore", "primary", "driver", cfg.Datastore.Primary.DriverName, "method", "db.Close", "msg", "success")
		}
	}()

	// setup authentication
	jwtCfg, err := auth.NewJWTConfig(
		auth.WithKeyFunc(cfg.Auth.JWT.Secret),
		auth.WithIssuer(cfg.Auth.JWT.Issuer),
		auth.WithSigningMethod(cfg.Auth.JWT.Algorithm),
	)
	if err != nil {
		logger.Fatal("datastore", "primary", "driver", cfg.Datastore.Primary.DriverName, "method", "sql.Connect", "err", err)
	}

	authCfg := auth.Config{
		JWT: jwtCfg,
		// OAuth...
	}

	// setup service
	var (
		validator  = playground.NewValidator(cfg.Validation.Playground.TagName, logger)
		repository = repository.New(
			repository.WithSQLDB(db),
			// elasticsearch
			// redis
		)
		formatter   = formatter.New(authCfg, logger)
		usecase     = usecase.New(repository, formatter, logger)
		endpoint    = endpoint.MakeServerEndpoint(usecase, authCfg, logger)
		httpHandler = httpDelivery.NewHTTPHandler(endpoint, validator, logger)
	)

	httpAddr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)

	var g run.Group
	// setup shutdown hook
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-ctx.Done():
				return ctx.Err()
			}
		}, func(error) {
			cancel()
		})
	}
	// setup server
	{
		var ln net.Listener
		g.Add(func() error {
			ln, err := net.Listen("tcp", httpAddr)
			if err != nil {
				return err
			}

			logger.Log("transport", "HTTP", "addr", httpAddr, "msg", "listening")
			return http.Serve(ln, httpHandler)
		}, func(error) {
			logger.Fatal("transport", "HTTP", "method", "net.Listen", "err", err)
			ln.Close()
		})
	}

	// start all
	logger.Log("exit", g.Run())
}
