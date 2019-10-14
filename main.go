package main

import (
	"flag"
	"io"
	"os"

	"github.com/fionera/video-grabber/site"
	"github.com/fionera/video-grabber/sites"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		logrus.Errorf("You need to pass the url as argument")
		return
	}

	url := flag.Arg(0)
	
	logrus.SetLevel(logrus.TraceLevel)
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, file))

	for _, siteAdapter := range sites.Sites {
		if !siteAdapter.IsSuitable(url) {
			continue
		}

		switch siteAdapter.GetType(url) {
		case site.VideoType:
			siteAdapter.GetVideo(url)
		case site.PlaylistType:
			siteAdapter.GetPlaylist(url)
		case site.UserType:
			user := siteAdapter.GetUser(url)

			playlists, err := user.GetPlaylists()
			if err != nil {
				logrus.Fatal(err)
			}

			for _, p := range playlists {
				videos, err := p.GetVideos()
				if err != nil {
					logrus.Fatal(err)
				}

				for _, v := range videos {
					s, err := v.GetName()
					if err != nil {
						logrus.Fatal(err)
					}

					logrus.Info(s)
				}
			}

		}
	}
}
