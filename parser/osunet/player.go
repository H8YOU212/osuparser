package osunet


// OsuAPIUsersResponse is useful for list-based user responses:
// GET /api/v2/users?ids[]=...
type OsuAPIUsersResponse struct {
	Users []OsuAPIUser `json:"users"`
}

// OsuAPIUser is the main user payload for:
// GET /api/v2/users/{user}
// GET /api/v2/me
// and embedded "user" objects in other endpoints.
type OsuAPIUser struct {
	ID                int                   `json:"id"`
	Username          string                `json:"username"`
	CountryCode       string                `json:"country_code"`
	AvatarURL         string                `json:"avatar_url"`
	DefaultGroup      string                `json:"default_group"`
	IsOnline          bool                  `json:"is_online"`
	IsSupporter       bool                  `json:"is_supporter"`
	LastVisit         string                `json:"last_visit"`
	ProfileColour     *string               `json:"profile_colour"`
	Statistics        *OsuAPIUserStatistics `json:"statistics"`
	StatisticsRuleset string                `json:"statistics_rulesets"`
}

// OsuAPIUserStatistics is nested in OsuAPIUser. It is used when user
// statistics are included in the response (usually full user profile).
type OsuAPIUserStatistics struct {
	Level                  OsuAPIUserLevel       `json:"level"`
	GlobalRank             *int                  `json:"global_rank"`
	RankedScore            int64                 `json:"ranked_score"`
	HitAccuracy            float64               `json:"hit_accuracy"`
	PlayCount              int                   `json:"play_count"`
	PlayTime               int                   `json:"play_time"`
	TotalScore             int64                 `json:"total_score"`
	TotalHits              int64                 `json:"total_hits"`
	MaximumCombo           int                   `json:"maximum_combo"`
	ReplaysWatchedByOthers int                   `json:"replays_watched_by_others"`
	IsRanked               bool                  `json:"is_ranked"`
	PP                     float64               `json:"pp"`
	CountryRank            *int                  `json:"country_rank"`
	GradeCounts            OsuAPIUserGradeCounts `json:"grade_counts"`
}

// OsuAPIUserLevel is part of user statistics:
// GET /api/v2/users/{user}
type OsuAPIUserLevel struct {
	Current  int     `json:"current"`
	Progress float64 `json:"progress"`
}

// OsuAPIUserGradeCounts is part of user statistics:
// GET /api/v2/users/{user}
type OsuAPIUserGradeCounts struct {
	SSH int `json:"ssh"`
	SS  int `json:"ss"`
	SH  int `json:"sh"`
	S   int `json:"s"`
	A   int `json:"a"`
}

// OsuAPIUserScoresResponse is a wrapper for score lists when endpoint returns
// object form. Useful for:
// GET /api/v2/users/{user}/scores/{type}
type OsuAPIUserScoresResponse struct {
	Scores []OsuAPIScore `json:"scores"`
}

// OsuAPIScore is the main score payload for:
// GET /api/v2/users/{user}/scores/{type}
// GET /api/v2/beatmaps/{beatmap}/scores
// GET /api/v2/scores/{mode}/{score}
type OsuAPIScore struct {
	ID                    int64                     `json:"id"`
	UserID                int                       `json:"user_id"`
	Accuracy              float64                   `json:"accuracy"`
	MaxCombo              int                       `json:"max_combo"`
	Passed                bool                      `json:"passed"`
	Perfect               bool                      `json:"perfect"`
	PP                    *float64                  `json:"pp"`
	Rank                  string                    `json:"rank"`
	Score                 int64                     `json:"score"`
	Statistics            OsuAPIScoreStatistics     `json:"statistics"`
	CreatedAt             string                    `json:"created_at"`
	BestID                *int64                    `json:"best_id"`
	Mode                  string                    `json:"mode"`
	ModeInt               int                       `json:"mode_int"`
	Replay                bool                      `json:"replay"`
	CurrentUserAttributes OsuAPIScoreUserAttributes `json:"current_user_attributes"`
	Beatmap               *OsuAPIBeatmap            `json:"beatmap"`
	Beatmapset            *OsuAPIBeatmapset         `json:"beatmapset"`
	User                  *OsuAPIUser               `json:"user"`
	Weight                *OsuAPIScoreWeight        `json:"weight"`
}

// OsuAPIScoreStatistics stores hit-result counters inside OsuAPIScore.
type OsuAPIScoreStatistics struct {
	Count300  int `json:"count_300"`
	Count100  int `json:"count_100"`
	Count50   int `json:"count_50"`
	CountMiss int `json:"count_miss"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
}

// OsuAPIScoreUserAttributes is nested in score payload for current-user flags.
type OsuAPIScoreUserAttributes struct {
	Pin *OsuAPIPinAttributes `json:"pin"`
}

// OsuAPIPinAttributes describes pinned-score metadata in score payloads.
type OsuAPIPinAttributes struct {
	IsPinned bool `json:"is_pinned"`
	ScoreID  int  `json:"score_id"`
}

// OsuAPIScoreWeight is used in "best scores" style endpoints:
// GET /api/v2/users/{user}/scores/best
type OsuAPIScoreWeight struct {
	Percentage float64 `json:"percentage"`
	PP         float64 `json:"pp"`
}
