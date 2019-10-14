package sites

import (
	"github.com/fionera/video-grabber/site"
	"github.com/fionera/video-grabber/sites/dailymotion"
)

var Sites = []site.Site{
	dailymotion.NewSite(),
}