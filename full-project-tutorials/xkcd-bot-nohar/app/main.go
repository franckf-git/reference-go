package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/InVisionApp/go-health/v2/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/neuware/xkcd-bot/pkg/bot"
	"gitlab.com/neuware/xkcd-bot/pkg/xkcd"
)

const (
	defaultIndexPath      = "./xkcd.index"
	defaultPrometheusPort = 2112
)

var (
	token          string
	indexPath      string
	watchPeriod    time.Duration
	prometheusPort int
)

func init() {
	var err error

	flag.StringVar(&token, "token", os.Getenv("BOT_TOKEN"), "Bot token")

	p := os.Getenv("INDEX_PATH")
	if p == "" {
		p = defaultIndexPath
	}
	flag.StringVar(&indexPath, "index", p, "Index path")

	var period time.Duration
	if period, err = time.ParseDuration(os.Getenv("WATCH_PERIOD")); err != nil {
		period = xkcd.DefaultWatchPeriod
	}
	flag.DurationVar(&watchPeriod, "period", period, "Watch period")

	var port int
	if port, err = strconv.Atoi(os.Getenv("PROMETHEUS_PORT")); err != nil {
		port = defaultPrometheusPort
	}
	flag.IntVar(&prometheusPort, "port", port, "Prometheus port")
	flag.Parse()
}

func main() {
	// Start serving metrics to
	go func() {
		log.Printf("serving prometheus metrics on port %d", prometheusPort)
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/health", handlers.NewJSONHandlerFunc(bot.HealthCheck, nil))
		http.ListenAndServe(fmt.Sprintf(":%d", prometheusPort), nil)
	}()
	err := bot.Run(&bot.Config{
		Token:       token,
		IndexPath:   indexPath,
		WatchPeriod: watchPeriod,
	})
	if err != nil {
		log.Fatal(err)
	}
}
