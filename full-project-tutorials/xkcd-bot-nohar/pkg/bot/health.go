package bot

import (
	"fmt"
	"net/url"
	"time"

	"github.com/InVisionApp/go-health/v2"
	"github.com/InVisionApp/go-health/v2/checkers"
	"github.com/bwmarrin/discordgo"
)

// HealthCheck is the global health checker
var HealthCheck = health.New()

type discordCheck struct {
	ses *discordgo.Session
}

func (c *discordCheck) Status() (interface{}, error) {
	_, err := c.ses.User(c.ses.State.User.ID)
	if err != nil {
		return nil, fmt.Errorf("Couldn't fetch self data: %w", err)
	}
	return nil, nil
}

func registerHealthChecks(dg *discordgo.Session) error {
	url, _ := url.Parse("https://xkcd.com")
	xkcdCheck, _ := checkers.NewHTTP(&checkers.HTTPConfig{
		URL:    url,
		Method: "HEAD",
	})

	HealthCheck.AddChecks([]*health.Config{
		{
			Name:     "xkcd-check",
			Checker:  xkcdCheck,
			Interval: 5 * time.Second,
			Fatal:    true,
		},
		{
			Name:     "discord-check",
			Checker:  &discordCheck{dg},
			Interval: 5 * time.Second,
			Fatal:    true,
		},
	})
	return HealthCheck.Start()
}
