package main

import (
	"context"
	"strings"
	"syscall"

	"github.com/m7shapan/njson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
)

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

type Match struct {
	Date           string `json:"date"`
	Turns          int    `json:"turns"`
	UntrackedWins  int    `json:"untrackedWins"`
	UntrackedLoses int    `json:"untrackedLoses"`
	Won            bool   `json:"won"`
	Streak         int    `json:"streak"`
	TrophiesFrom   int    `json:"trophiesFrom"`
	TrophiesTo     int    `json:"trophiesTo"`
}

// App struct
type App struct {
	ctx      context.Context
	userData UserData
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.userData = a.GetUserData()
	a.monitorUserData()
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

func (a *App) monitorUserData() {
	var regNotifyChangeKeyValue *syscall.Proc

	if advapi32, err := syscall.LoadDLL("advapi32.dll"); err == nil {
		if proc, err := advapi32.FindProc("RegNotifyChangeKeyValue"); err == nil {
			regNotifyChangeKeyValue = proc
		} // else
	}

	if regNotifyChangeKeyValue != nil {
		go func() {
			key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Paladin Studios\Stormbound`, syscall.KEY_NOTIFY)
			if err != nil {
				return
			}

			for {
				regNotifyChangeKeyValue.Call(uintptr(key), 0, 0x00000001|0x00000004, 0, 0)
				data := a.GetUserData()

				if a.userData != data {
					runtime.EventsEmit(a.ctx, "userDataChanged", data)
				}
			}
		}()
	}
}
