package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleffmpeg"
	"layeh.com/gumble/gumbleutil"

	_ "layeh.com/gumble/opus"
)

func main() {
	var stream *gumbleffmpeg.Stream

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	gumbleutil.Main(gumbleutil.AutoBitrate, gumbleutil.Listener{
		Connect: func(e *gumble.ConnectEvent) {
			user := e.Client.Users.Find("steve")
			user.SetComment("Send message ... KAFFEE ...")
		},

		TextMessage: func(e *gumble.TextMessageEvent) {
			t := time.Now()
			var tageszeit string
			switch {
			case t.Hour() > 5 && t.Hour() < 12:
				tageszeit = "Guten Morgen!"
			case t.Hour() >= 12 && t.Hour() <= 13:
				tageszeit = "Mahlzeit!"
			case t.Hour() > 13 && t.Hour() < 17:
				tageszeit = "Hallo!"
			default:
				tageszeit = "Guten Abend!"
			}
			r := strings.NewReplacer(
				"Monday", "Montag",
				"Tuesday", "Dienstag",
				"Wednesday", "Mittwoch",
				"Thursday", "Donnerstag",
				"Friday", "Freitag",
				"Saturday", "Samstag",
				"Sunday", "Sonntag")

			if e.Sender.Channel.Name != "Küche" {
				e.Sender.Send(e.Sender.Name + "! Komme doch bitte in die Küche!")
				return
			}

			if e.Sender == nil {
				return
			}
			lowerS, lowerWord := strings.ToLower(e.Message), strings.ToLower("kaffee")
			ok := strings.Contains(lowerS, lowerWord)
			if !ok {
				return
			}
			e.Sender.Send(tageszeit + " " + e.Sender.Name + "! Dein Kaffee wird nun kredenzt ... ")
			if stream != nil && stream.State() == gumbleffmpeg.StatePlaying {
				return
			}
			stream = gumbleffmpeg.New(e.Client, gumbleffmpeg.SourceFile("./kaffee.mp3"))
			if err := stream.Play(); err != nil {
				fmt.Printf("%s\n", err)
			}
			stream.Wait()
			weekday := time.Now().Weekday()
			wochentag := r.Replace(weekday.String())
			e.Sender.Send("Moege Dein Kaffee heute stark und Dein " + wochentag + " kurz sein!")
		},
	})
}
