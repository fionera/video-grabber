package dailymotion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fionera/video-grabber/site"
	"github.com/sirupsen/logrus"
)

type user string

func (u user) GetPlaylists() ([]site.Playlist, error) {
	var playlists []site.Playlist
	var page = 1

requestPage:
	logrus.Infof("Requesting playlists page %d for User %s", page, u)
	resp, err := http.Get(fmt.Sprintf("https://api.dailymotion.com/playlists?owner=%s&limit=100&page=%d", string(u), page))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ups := userPlaylistsStruct{}
	err = json.Unmarshal(data, &ups)
	if err != nil {
		return nil, err
	}

	for _, v := range ups.List {
		playlists = append(playlists, playlist{
			name: v.Name,
			id:   v.ID,
		})
	}

	if ups.HasMore {
		page++
		goto requestPage
	}

	return playlists, nil
}

func (u user) GetVideos() ([]site.Video, error) {
	return nil, nil
}

func (u user) GetName() (string, error) {
	logrus.Infof("Requesting User %s", u)
	resp, err := http.Get("https://api.dailymotion.com/user/" + string(u))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	us := userStruct{}
	err = json.Unmarshal(data, &us)
	if err != nil {
		return "", err
	}

	return us.Screenname, nil
}

var (
	_ site.User = (*user)(nil)
)
