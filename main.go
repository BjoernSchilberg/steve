package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleffmpeg"
	"layeh.com/gumble/gumbleutil"

	_ "layeh.com/gumble/opus"
)

var daysToGerman = map[time.Weekday]string{
	time.Monday:    "Montag",
	time.Tuesday:   "Dienstag",
	time.Wednesday: "Mittwoch",
	time.Thursday:  "Donnerstag",
	time.Friday:    "Freitag",
	time.Saturday:  "Samstag",
	time.Sunday:    "Sonntag",
}

func daytimeGreet(t time.Time) string {
	switch {
	case t.Hour() > 5 && t.Hour() < 12:
		return "Guten Morgen!"
	case t.Hour() >= 12 && t.Hour() <= 13:
		return "Mahlzeit!"
	case t.Hour() > 13 && t.Hour() < 17:
		return "Hallo!"
	default:
		return "Guten Abend!"
	}
}

func main() {
	var stream *gumbleffmpeg.Stream

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	gumbleutil.Main(gumbleutil.AutoBitrate, gumbleutil.Listener{
		Connect: func(e *gumble.ConnectEvent) {
			user := e.Client.Users.Find("steve")
			user.SetComment("Send message ... KAFFEE ...<br>Source: <a href='https://github.com/BjoernSchilberg/steve/'>https://github.com/BjoernSchilberg/steve</a>")
		},

		TextMessage: func(e *gumble.TextMessageEvent) {
			now := time.Now()

			if e.Sender.Channel.Name != "Küche" {
				e.Sender.Send(e.Sender.Name + "! Komme doch bitte in die Küche!")
				return
			}

			if e.Sender == nil || !strings.Contains(strings.ToLower(e.Message), "kaffee") {
				return
			}

			tageszeit := daytimeGreet(now)
			e.Sender.Send(tageszeit + " " + e.Sender.Name + "! Dein Kaffee wird nun kredenzt ... ")
			if stream != nil && stream.State() == gumbleffmpeg.StatePlaying {
				return
			}
			stream = gumbleffmpeg.New(e.Client, gumbleffmpeg.SourceFile("./kaffee.mp3"))
			if err := stream.Play(); err != nil {
				fmt.Printf("%s\n", err)
			}
			stream.Wait()

			wochentag := daysToGerman[now.Weekday()]
			e.Sender.Send("Moege Dein Kaffee heute stark und Dein " + wochentag + " kurz sein!")

			r := rand.Float64()
			if r > 2.0/3 {
				e.Sender.Send("KAFFEESATZ LEEREN")
			}
			if r > 5.0/6 {
				e.Sender.Send("WASSERTANK FUELLEN")
			}
			if r > 6.0/7 {
				e.Sender.Send("BOHNEN FUELLEN")
			}
			if r > 7.0/8 {
				e.Sender.Send("REINIGUNG STARTEN")
			}
		},
	})
}
