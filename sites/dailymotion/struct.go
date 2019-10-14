package dailymotion

type userStruct struct {
	ID         string `json:"id"`
	Screenname string `json:"screenname"`
}

type userPlaylistsStruct struct {
	Explicit bool `json:"explicit"`
	HasMore  bool `json:"has_more"`
	Limit    int  `json:"limit"`
	List     []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Owner string `json:"owner"`
	} `json:"list"`
	Page  int `json:"page"`
	Total int `json:"total"`
}

type playlistVideosStruct struct {
	Explicit bool `json:"explicit"`
	HasMore  bool `json:"has_more"`
	Limit    int  `json:"limit"`
	List     []struct {
		Channel string `json:"channel"`
		ID      string `json:"id"`
		Owner   string `json:"owner"`
		Title   string `json:"title"`
	} `json:"list"`
	Page  int `json:"page"`
	Total int `json:"total"`
}
