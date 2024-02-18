package main

import (
	"context"
	"strings"

	"github.com/m7shapan/njson"
	"golang.org/x/sys/windows/registry"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type UserData struct {
	UserId                 string `njson:"UserId._content"`
	Username               string `njson:"Username._content"`
	UserTrophies           int    `njson:"UserTrophies._content"`
	UserRank               int    `njson:"UserRank._content"`
	UserLevel              int    `njson:"UserLevel._content"`
	TimeMatchmakingStarted string `njson:"TimeMatchmakingStarted._content"`
	GameTurns              int    `njson:"GameTurns._content"`
	TimeMatchStarted       string `njson:"TimeMatchStarted._content"`
	RankedPlayed           int    `njson:"RankedPlayed._content"`
	RankedWon              int    `njson:"RankedWon._content"`
}

func (a *App) GetUserData() UserData {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Paladin Studios\Stormbound`, registry.ALL_ACCESS)
	if err != nil {
		return UserData{}
	}
	defer key.Close()

	names, err := key.ReadValueNames(0)
	if err != nil {
		return UserData{}
	}

	var analytics string
	for _, name := range names {
		if strings.HasPrefix(name, "MIRAGE_ANALYTICS_DATA") {
			analytics = name
		}
	}

	b, _, err := key.GetBinaryValue(analytics)
	if err != nil {
		return UserData{}
	}

	var data UserData
	if e := njson.Unmarshal(b[:len(b)-1], &data); e != nil {
		return UserData{}
	}

	return data
}
