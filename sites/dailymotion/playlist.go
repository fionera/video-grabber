package dailymotion

import (
	"encoding/json"
	"fmt"
	"github.com/fionera/video-grabber/site"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type playlist struct {
	name string
	id string
}

func (p playlist) GetName() (string, error) {
	panic("implement me")
}

func (p playlist) GetVideos() ([]site.Video, error) {
	var videos []site.Video
	var page = 1

requestPage:
	logrus.Infof("Requesting video page %d for Playlist %s", page, p.name)
	resp, err := http.Get(fmt.Sprintf("https://api.dailymotion.com/playlist/%s/videos?limit=100&page=%d", p.id, page))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pvs := playlistVideosStruct{}
	err = json.Unmarshal(data, &pvs)
	if err != nil {
		return nil, err
	}

	for _, v := range pvs.List {
		videos = append(videos, video{
			id:    v.ID,
			title: v.Title,
		})
	}

	if pvs.HasMore {
		page++
		goto requestPage
	}

	return videos, nil
}

var (
	_ site.Playlist = (*playlist)(nil)
)

