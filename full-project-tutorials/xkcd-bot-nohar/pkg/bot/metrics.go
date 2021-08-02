package bot

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	guildsCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "xkcd_bot_guilds_total",
		Help: "The number of guilds where the bot is present",
	})

	cmdsProcessed = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "xkcd_bot_processed_commands",
			Help: "The number of processed commands",
		},
		[]string{"guild"},
	)
)
