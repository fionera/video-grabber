package dailymotion

import (
	"regexp"

	"github.com/fionera/video-grabber/site"
)

var (
	DailymotionRegexp = regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.)?dailymotion\.com\/(?P<name>[0-9A-Za-z]+)(?:\/playlists)?$`)
)

func NewSite() site.Site {
	return &dailymotionSite{}
}

type dailymotionSite struct {
}

func (d *dailymotionSite) GetPlaylist(url string) site.Playlist {
	panic("implement me")
}

func (d *dailymotionSite) GetVideo(url string) site.Video {
	panic("implement me")
}

func (d *dailymotionSite) GetUser(url string) site.User {
	s := DailymotionRegexp.FindStringSubmatch(url)

	if len(s) == 1 {
		return nil
	}

	return user(s[1])
}

func (d *dailymotionSite) GetType(url string) site.Type {
	return site.UserType
}


func (d dailymotionSite) IsSuitable(u string) bool {
	return DailymotionRegexp.MatchString(u)
}

var (
	_ site.Site = (*dailymotionSite)(nil)
)
