package dailymotion

import (
	"github.com/fionera/video-grabber/site"
)

type video struct {
	id string
	title string
}

func (v video) GetName() (string, error) {
	return v.title, nil
}

var (
	_ site.Video = (*video)(nil)
)

