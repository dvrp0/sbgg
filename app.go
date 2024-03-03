package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/m7shapan/njson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
)

type RegistryData struct {
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

type Profile struct {
	IsDarkMode     bool    `json:"isDarkMode"`
	IsLeagueThemed bool    `json:"isLeagueThemed"`
	UserTrophies   int     `json:"userTrophies"`
	RankedPlayed   int     `json:"rankedPlayed"`
	RankedWon      int     `json:"rankedWon"`
	Matches        []Match `json:"matches"`
}

// App struct
type App struct {
	ctx          context.Context
	registryData RegistryData
	profile      Profile
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	registry, err := a.GetRegistryData()
	if err == nil {
		a.registryData = *registry
	}

	profile, err := a.GetProfile()
	if err == nil {
		a.profile = *profile
	}

	a.monitorRegistryData()
}

func (a *App) GetRegistryData() (*RegistryData, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Paladin Studios\Stormbound`, registry.ALL_ACCESS)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	names, err := key.ReadValueNames(0)
	if err != nil {
		return nil, err
	}

	var analytics string
	for _, name := range names {
		if strings.HasPrefix(name, "MIRAGE_ANALYTICS_DATA") {
			analytics = name
			break
		}
	}

	b, _, err := key.GetBinaryValue(analytics)
	if err != nil {
		return nil, err
	}

	var data RegistryData
	if err = njson.Unmarshal(b[:len(b)-1], &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (a *App) getProfilePath() (dir string, path string, err error) {
	dir, err = os.UserConfigDir()
	if err != nil {
		return "", "", err
	}
	log.Print(a.registryData)
	dir = filepath.Join(dir, "sbgg")
	path = filepath.Join(dir, a.registryData.UserId+".json")

	return dir, path, nil
}

func (a *App) updateProfile() {
	a.profile.UserTrophies = a.registryData.UserTrophies
	a.profile.RankedPlayed = a.registryData.RankedPlayed
	a.profile.RankedWon = a.registryData.RankedWon
}

func (a *App) saveProfile(data *Profile) error {
	a.profile = *data

	_, path, err := a.getProfilePath()
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(*data, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(bytes); err != nil {
		return err
	}

	return nil
}

func (a *App) GetProfile() (*Profile, error) {
	var profile Profile

	dir, path, err := a.getProfilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); err == nil {
		log.Print("Profile exists")

		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(data, &profile); err != nil {
			return nil, err
		}

		// There are some untracked games that did not play in this device
		// Matches can't be tracked if games were played in another device and trophies were not updated too
		if profile.RankedPlayed == a.registryData.RankedPlayed && profile.RankedWon == a.registryData.RankedWon && profile.UserTrophies != a.registryData.UserTrophies {
			profile.Matches = append(profile.Matches, Match{
				Date:         time.Now().Format("2006-01-02 15:04:05"),
				TrophiesFrom: profile.UserTrophies,
				TrophiesTo:   a.registryData.UserTrophies,
			})
			a.updateProfile()
			a.saveProfile(&profile)
		} else if profile.RankedPlayed != a.registryData.RankedPlayed || profile.RankedWon != a.registryData.RankedWon { // User played untracked games in this device
			profile.Matches = append(profile.Matches, Match{
				Date:           time.Now().Format("2006-01-02 15:04:05"),
				TrophiesFrom:   profile.UserTrophies,
				TrophiesTo:     a.registryData.UserTrophies,
				UntrackedWins:  a.registryData.RankedWon - profile.RankedWon,
				UntrackedLoses: (a.registryData.RankedPlayed - a.registryData.RankedWon) - (profile.RankedPlayed - profile.RankedWon),
			})
			a.updateProfile()
			a.saveProfile(&profile)
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Print("Profile does not exist")

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}

		profile = Profile{
			IsDarkMode:     false,
			IsLeagueThemed: true,
			UserTrophies:   a.registryData.UserTrophies,
			RankedPlayed:   a.registryData.RankedPlayed,
			RankedWon:      a.registryData.RankedWon,
			Matches:        []Match{},
		}
		a.updateProfile()
		a.saveProfile(&profile)
	}

	return &profile, nil
}

func (a *App) monitorRegistryData() {
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
				data, err := a.GetRegistryData()

				if err == nil && a.registryData != *data {
					a.registryData = *data

					// There is a new finished game
					if a.registryData.RankedPlayed != a.profile.RankedPlayed && a.registryData.UserTrophies != a.profile.UserTrophies {
						is_won := a.registryData.UserTrophies > a.profile.UserTrophies

						a.profile.Matches = append(a.profile.Matches, Match{
							Date:         time.Now().Format("2006-01-02 15:04:05"),
							Turns:        a.registryData.GameTurns,
							Won:          is_won,
							TrophiesFrom: a.profile.UserTrophies,
							TrophiesTo:   a.registryData.UserTrophies,
						})
						a.updateProfile()
						if is_won {
							a.profile.RankedWon++ // Manually update RankedWon because it is updated later than UserTrophies
						}
						a.saveProfile(&a.profile)
					}

					runtime.EventsEmit(a.ctx, "dataChanged")
				}
			}
		}()
	}
}
