package web

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zhtheyun/fibonacci/lib/config"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type contextKey int

const (
	// RequestTimeOut is the duration(second) to wait for a request finishes.
	RequestTimeOut = 60

	configKey contextKey = iota
	startedAtKey
)

//FibCache used for cache the fibonacci numbers to increase the performance.
type FibCache struct {
	// Cached fib numbers
	Data []string

	// Cached numbers
	Numbers uint64

	//FixME: This variable is try to keep the value to allow app continue calculate the fib number if exceed cache list.
	Start big.Int

	Next big.Int
}

//A global variable allow us the fibonacci result instead of evaluate every time
//the size is configurable based on environment variable.
var fibCache *FibCache

// Server represents the Server instance.
type Server struct {
	// When started.
	startedAt time.Time

	// Configuration
	config config.Config
}

// NewServer build a Server instance using the config.
func NewServer(c config.Config) (*Server, error) {
	//Init Fib cache to help increase the performance instead of calculate every time

	fibCache = new(FibCache)

	data, start, next := c.Generator.Generate(*big.NewInt(0), *big.NewInt(1), c.CachedNumbers)
	fibCache.Data = data
	fibCache.Start = start
	fibCache.Next = next
	fibCache.Numbers = c.CachedNumbers

	logrus.Infof("Init cache. the size is %d", c.CachedNumbers)

	return &Server{
		startedAt: time.Now(),
		config:    c,
	}, nil
}

// AddContext used to add the server info the request.
func (s *Server) AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), configKey, s.config)
		ctx = context.WithValue(ctx, startedAtKey, s.startedAt)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Run used to make the server instance listen and accept connections.
func (s *Server) Run(ctx context.Context) error {

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	router.
		Path("/fibonacci").
		HandlerFunc(FibonacciHandler).
		Methods("GET").
		Queries("numbers", "{numbers:[0-9]+}").
		Name("Fibonaccis")

	router.HandleFunc("/heartbeat", HeartBeatHandler)

	logrus.Info(
		fmt.Sprintf("RESTful service run on port %d, version: %s", s.config.Port, s.config.Version),
	)
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.AddContext(router),
	}
	go func() {
		if err := svr.ListenAndServe(); err != nil {
			logrus.Error(
				fmt.Sprintf("Error occurs when starting http server. Error: %s", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "fail to shut down the server")
	}
	logrus.Info("RESTful service exits")

	return nil
}
