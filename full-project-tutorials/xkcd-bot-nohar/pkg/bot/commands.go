package bot

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/neuware/xkcd-bot/pkg/index"
	"gitlab.com/neuware/xkcd-bot/pkg/xkcd"
)

const prefix = "xkcd!"

func getCommand(dg *discordgo.Session, content string) (string, bool) {
	mention := "<@" + dg.State.User.ID + "> "
	nmention := "<@!" + dg.State.User.ID + "> "

	switch {
	case strings.HasPrefix(content, prefix):
		return content[len(prefix):], true
	case strings.HasPrefix(content, mention):
		return content[len(mention):], true
	case strings.HasPrefix(content, nmention):
		return content[len(nmention):], true
	}
	return "", false
}

func onMessageCreate(dg *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == dg.State.User.ID {
		return
	}

	var cmd string
	var ok bool
	var err error
	var guild string

	channel, err := dg.Channel(m.ChannelID)
	if err != nil {
		log.Print("failed to fetch channel information:", err)
		return
	}
	if channel.Type == discordgo.ChannelTypeDM {
		// Channel is a DM, the message is the command
		cmd, ok = m.Content, true
		guild = "DM"
	} else {
		// Channel is a public one, prefix or mention is required
		cmd, ok = getCommand(dg, m.Content)
		guild = m.GuildID
	}

	if !ok {
		return
	}

	switch cmd {
	case "help":
		dg.ChannelMessageSend(
			m.ChannelID,
			fmt.Sprint(
				"Commands:"+
					"\n* `latest`: Show the latest xkcd"+
					"\n* `random`: Pick a random xkcd"+
					"\n* `help`: Well, you just typed it"+
					"\n\nOtherwise, I'll show you the xkcd strip that best matches your query."+
					" You can give me a comic number or its title, or search by keywords,"+
					" and I can even handle `/regular expressions/`."+
					" I also support advanced syntax. e.g: `video games -ferret`"+
					" searches for a comic that talks about video games,"+
					" but not the one with the ferret.",
			),
		)
	case "latest":
		err = showXKCD(dg, m.ChannelID, xkcd.Latest)
	case "random":
		err = showRandomXKCD(dg, m.ChannelID)
	default:
		res := index.Search(cmd)
		if len(res) > 0 {
			err = showXKCD(dg, m.ChannelID, res[0].ComicNum)
		} else {
			err = fmt.Errorf("search for %q didn't yield any result", cmd)
		}
	}
	if err != nil {
		dg.MessageReactionAdd(m.ChannelID, m.ID, "🤷")
		log.Printf("while processing command %q: %s", cmd, err)
	}
	cmdsProcessed.WithLabelValues(guild).Inc()
}

func showRandomXKCD(dg *discordgo.Session, channelID string) error {
	return showXKCD(dg, channelID, fmt.Sprint(rand.Int()%xkcd.LatestNum()+1))
}

func showXKCD(dg *discordgo.Session, channelID, xkcdID string) error {
	c, err := xkcd.Get(xkcdID)
	if err != nil {
		return err
	}
	date, _ := c.Date()
	_, err = dg.ChannelMessageSendEmbed(
		channelID,
		&discordgo.MessageEmbed{
			URL:         c.URL(),
			Type:        discordgo.EmbedTypeImage,
			Title:       c.SafeTitle,
			Description: c.Alt,
			Image:       &discordgo.MessageEmbedImage{URL: c.Img},
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("#%d (%s)", c.Num, date.Format("2006-01-02")),
			},
		},
	)
	return err
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
