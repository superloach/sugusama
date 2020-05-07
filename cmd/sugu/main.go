package main

import (
	"flag"

	"github.com/superloach/minori"
	sugu "github.com/superloach/sugusama"
)

var (
	user = flag.String("user", "", "instagram username")
	pass = flag.String("pass", "", "instagram password")
)

var Log = minori.GetLogger("sugu")

func main() {
	_f := "main"
	flag.Parse()

	if *user == "" {
		Log.Fatal(_f, "please provide user")
	}
	if *pass == "" {
		Log.Fatal(_f, "please provide pass")
	}

	client, err := sugu.NewClient(nil)
	if err != nil {
		Log.Fatal(_f, err)
	}

	login, err := client.Login(*user, *pass)
	if err != nil {
		Log.Fatal(_f, err)
	}
	Log.Infof(_f, "logged in as %s", login.UserID)

	shared, additional, err := client.Home()
	if err != nil {
		Log.Fatal(_f, err)
	}
	Log.Infof(_f, "%#v", shared.Config.Viewer)
	Log.Info(_f, additional)

	stories, err := client.Stories(nil)
	if err != nil {
		Log.Error(_f, err)
	} else {
		Log.Infof(_f, "%d stories available", len(stories.Stories()))
	}

	activity, err := client.Activity()
	if err != nil {
		Log.Error(_f, err)
	} else {
		a, err := activity.Activity()
		if err != nil {
			Log.Fatal(_f, err)
		}

		Log.Infof(_f, "%d comment likes", a.CommentLikes)
		Log.Infof(_f, "%d comments", a.Comments)
		Log.Infof(_f, "%d likes", a.Likes)
		Log.Infof(_f, "%d relationships", a.Relationships)
		Log.Infof(_f, "%d user tags", a.UserTags)
	}

	broadcasts, err := client.Broadcasts()
	if err != nil {
		Log.Error(_f, err)
	} else {
		Log.Infof(_f, "%d broadcasts", len(broadcasts.Broadcasts))
	}

	directBadgeCount, err := client.DirectBadgeCount(true)
	if err != nil {
		Log.Error(_f, err)
	} else {
		Log.Infof(_f, "%d direct messages", directBadgeCount.BadgeCount)
	}
}
