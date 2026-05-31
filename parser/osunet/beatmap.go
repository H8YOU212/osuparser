package osunet


func (c *Client) GetBeatmapset() *OsuAPIBeatmapset {
	var res OsuAPIBeatmapset

	req, err := sampleGet()
	if err != nil{
		return nil
	}

	c.Do(req, res)

	return &res
}

func (c *Client) DownloadBeatmapset() {
	
}

// OsuAPIBeatmap is the main beatmap payload for:
// GET /api/v2/beatmaps/{beatmap}
// and embedded beatmap objects in score/user responses.
type OsuAPIBeatmap struct {
	ID            int                     `json:"id"`
	BeatmapsetID  int                     `json:"beatmapset_id"`
	Mode          string                  `json:"mode"`
	ModeInt       int                     `json:"mode_int"`
	Status        string                  `json:"status"`
	TotalLength   int                     `json:"total_length"`
	UserID        int                     `json:"user_id"`
	Version       string                  `json:"version"`
	Accuracy      float64                 `json:"accuracy"`
	AR            float64                 `json:"ar"`
	BPM           float64                 `json:"bpm"`
	Convert       bool                    `json:"convert"`
	CountCircles  int                     `json:"count_circles"`
	CountSliders  int                     `json:"count_sliders"`
	CountSpinners int                     `json:"count_spinners"`
	CS            float64                 `json:"cs"`
	DeletedAt     *string                 `json:"deleted_at"`
	Drain         float64                 `json:"drain"`
	HitLength     int                     `json:"hit_length"`
	IsScoreable   bool                    `json:"is_scoreable"`
	LastUpdated   string                  `json:"last_updated"`
	Passcount     int64                   `json:"passcount"`
	Playcount     int64                   `json:"playcount"`
	Ranked        int                     `json:"ranked"`
	URL           string                  `json:"url"`
	Checksum      *string                 `json:"checksum"`
	MaxCombo      *int                    `json:"max_combo"`
	Failtimes     *OsuAPIBeatmapFailtimes `json:"failtimes"`
	Beatmapset    *OsuAPIBeatmapset       `json:"beatmapset"`
}

// OsuAPIBeatmapFailtimes is nested in beatmap payload:
// GET /api/v2/beatmaps/{beatmap}
type OsuAPIBeatmapFailtimes struct {
	Fail []int `json:"fail"`
	Exit []int `json:"exit"`
}

// OsuAPIBeatmapset is the main beatmapset payload for:
// GET /api/v2/beatmapsets/{beatmapset}
// and embedded "beatmapset" objects in scores/beatmaps.
type OsuAPIBeatmapset struct {
	ID                    int                            `json:"id"`
	Artist                string                         `json:"artist"`
	ArtistUnicode         string                         `json:"artist_unicode"`
	DifficultyRating      float64                        `json:"difficulty_rating"`
	Covers                OsuAPICovers                   `json:"covers"`
	Creator               string                         `json:"creator"`
	FavouriteCount        int                            `json:"favourite_count"`
	PlayCount             int64                          `json:"play_count"`
	PreviewURL            string                         `json:"preview_url"`
	Source                string                         `json:"source"`
	Status                string                         `json:"status"`
	Title                 string                         `json:"title"`
	TitleUnicode          string                         `json:"title_unicode"`
	UserID                int                            `json:"user_id"`
	Video                 bool                           `json:"video"`
	BPM                   float64                        `json:"bpm"`
	CanBeHyped            bool                           `json:"can_be_hyped"`
	DeletedAt             *string                        `json:"deleted_at"`
	DiscussionEnabled     bool                           `json:"discussion_enabled"`
	DiscussionLocked      bool                           `json:"discussion_locked"`
	IsScoreable           bool                           `json:"is_scoreable"`
	LastUpdated           string                         `json:"last_updated"`
	LegacyThreadURL       string                         `json:"legacy_thread_url"`
	NominationsSummary    OsuAPINominationsSummary       `json:"nominations_summary"`
	Ranked                int                            `json:"ranked"`
	RankedDate            *string                        `json:"ranked_date"`
	Storyboard            bool                           `json:"storyboard"`
	SubmittedDate         string                         `json:"submitted_date"`
	Tags                  string                         `json:"tags"`
	Availability          OsuAPIBeatmapsetAvailability   `json:"availability"`
	HasFavourited         bool                           `json:"has_favourited"`
	Beatmaps              []OsuAPIBeatmap                `json:"beatmaps"`
	Converts              []OsuAPIBeatmap                `json:"converts"`
	CurrentNominations    []OsuAPICurrentNomination      `json:"current_nominations"`
	CurrentUserAttributes OsuAPIBeatmapsetUserAttributes `json:"current_user_attributes"`
	Description           OsuAPIBeatmapsetDescription    `json:"description"`
	Genre                 OsuAPIBeatmapsetGenreLanguage  `json:"genre"`
	Language              OsuAPIBeatmapsetGenreLanguage  `json:"language"`
	Hype                  *OsuAPIBeatmapsetHype          `json:"hype"`
	Ratings               []int                          `json:"ratings"`
	RecentFavourites      []OsuAPIUser                   `json:"recent_favourites"`
	User                  *OsuAPIUser                    `json:"user"`
}

// OsuAPICovers describes cover image URLs from beatmapset endpoints.
type OsuAPICovers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}

// OsuAPINominationsSummary is nested in beatmapset payload.
type OsuAPINominationsSummary struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

// OsuAPIBeatmapsetAvailability is nested in beatmapset payload.
type OsuAPIBeatmapsetAvailability struct {
	DownloadDisabled bool    `json:"download_disabled"`
	MoreInformation  *string `json:"more_information"`
}

// OsuAPICurrentNomination is nested in beatmapset payload.
type OsuAPICurrentNomination struct {
	UserID   int      `json:"user_id"`
	Rulesets []string `json:"rulesets"`
}

// OsuAPIBeatmapsetUserAttributes holds current user permissions for beatmapset.
type OsuAPIBeatmapsetUserAttributes struct {
	CanDelete          bool `json:"can_delete"`
	CanEditMetadata    bool `json:"can_edit_metadata"`
	CanEditTags        bool `json:"can_edit_tags"`
	CanHype            bool `json:"can_hype"`
	CanLove            bool `json:"can_love"`
	CanRemoveFromLoved bool `json:"can_remove_from_loved"`
	IsWatching         bool `json:"is_watching"`
	NewHypeTime        *int `json:"new_hype_time"`
	RemainingHype      *int `json:"remaining_hype"`
}

// OsuAPIBeatmapsetDescription is nested in beatmapset payload.
type OsuAPIBeatmapsetDescription struct {
	Description string `json:"description"`
	BBCode      string `json:"bbcode"`
}

// OsuAPIBeatmapsetGenreLanguage is used for both genre and language
// nested objects in beatmapset responses.
type OsuAPIBeatmapsetGenreLanguage struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// OsuAPIBeatmapsetHype is nested in beatmapset payload.
type OsuAPIBeatmapsetHype struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}
