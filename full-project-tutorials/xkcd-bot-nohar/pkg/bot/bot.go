package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/neuware/xkcd-bot/pkg/xkcd"
)

// Config is the bot configuration
type Config struct {
	Token       string
	IndexPath   string
	WatchPeriod time.Duration
}

// Run runs the discord bot
func Run(cfg *Config) error {
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return fmt.Errorf("couldn't create discordgo session: %w", err)
	}
	dg.ShouldReconnectOnError = true

	dg.AddHandler(onMessageCreate)
	dg.AddHandler(onGuildCreate)
	dg.AddHandler(onGuildDelete)

	err = dg.Open()
	if err != nil {
		return fmt.Errorf("couldn't open connection: %w", err)
	}
	defer dg.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	log.Println("up and running, press Ctrl-C to exit.")

	registerHealthChecks(dg)
	xkcd.SetWatchPeriod(cfg.WatchPeriod)
	startIndexing(cfg.IndexPath)

	<-sc
	return nil
}

func onReady(dg *discordgo.Session, r *discordgo.Ready) {
	log.Println("listening for commands on", len(r.Guilds), "guilds")
	guildsCount.Set(float64(len(r.Guilds)))
}

func onGuildCreate(dg *discordgo.Session, g *discordgo.GuildCreate) {
	log.Printf("joined guild %s (%s)", g.ID, g.Name)
	guildsCount.Inc()
}

func onGuildDelete(dg *discordgo.Session, g *discordgo.GuildDelete) {
	log.Printf("left guild %s", g.ID)
	guildsCount.Dec()
}
