package site

const (
	VideoType Type = iota + 1
	PlaylistType
	UserType
)

type Type int

type Site interface {
	IsSuitable(url string) bool
	GetType(url string) Type

	GetPlaylist(url string) Playlist
	GetVideo(url string) Video
	GetUser(url string) User
}

type Playlist interface {
	GetName() (string, error)
	GetVideos() ([]Video, error)
}

type Video interface {
	GetName() (string, error)
}

type User interface {
	GetName() (string, error)
	GetPlaylists() ([]Playlist, error)
	GetVideos() ([]Video, error)
}
