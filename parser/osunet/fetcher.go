package osunet

import "net/http"

type BeatmapFetcher interface {
	GetBeatmaps() *http.Request
	GetBeatmap() *http.Request
	GetBeatmapSet() *http.Request
	GetBeatmapSets() *http.Request
}

type PlayerFetcher interface {
	GetUser() *http.Request
	GetUsers() *http.Request
}
