package osr

type Parser interface {
	Decode() error
}

type OsrModel struct {
	dir     string
	dst     string
	format  format
	isLaser bool
}

type format struct {
	Header         Header
	ReplayInfo     ReplayInfo
	Stats          Stats
	LifeBarGraph   string
	Timestamp      int64
	CompressedSize int32
	ReplayData     string
	Additional     AdditionalData
}

type Header struct {
	Mode    byte
	Version int32
}

type ReplayInfo struct {
	BeatmapMD5 string
	Username   string
	ReplayMD5  string
}

type Stats struct {
	Count300      int16
	Count100      int16
	Count50       int16
	CountGeki     int16
	CountKatu     int16
	CountMiss     int16
	TotalScore    int32
	GreatestCombo int16
	Perfect       byte
	Mods          int32
}

type AdditionalData struct {
	ScoreID                int64
	TargetPracticeAccuracy float64
}

type Precompressed struct {
	// убрали Precompressed []byte — не нужен
}