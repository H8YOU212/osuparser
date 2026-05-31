package osu

type Commonosu struct {
	Path    string	`json:"path"`
	Osufile Beatmap	`json:"osufile"`
}

type Beatmap struct {
	General      General
	Editor       Editor
	Metadata     Metadata
	Difficulty   Difficulty
	Events       []Event
	TimingPoints []TimingPoint
	HitObjects   []HitObject
	Colours      []Colour
}

type General struct {
	AudioFilename        string
	AudioLeadIn          int
	AudioHash            string
	PreviewTime          int
	Countdown            int
	SampleSet            string
	StackLeniency        float64
	Mode                 int
	LetterboxInBreaks    int
	EpilepsyWarning      int
	WidescreenStoryboard int
}

type Editor struct {
	Bookmarks       string
	DistanceSpacing float64
	BeatDivisor     int
	GridSize        int
	TimelineZoom    float64
}

type Metadata struct {
	Title         string
	TitleUnicode  string
	Artist        string
	ArtistUnicode string
	Creator       string
	Version       string
	Source        string
	Tags          string
	BeatmapID     int
	BeatmapSetID  int
}

type Difficulty struct {
	HP               float64
	OD               float64
	AR               float64
	CS               float64
	SliderMultiplier float64
	SliderTickRate   float64
}

type Event struct {
	Eventtype   any
	StartTime   int
	EvemtParams any
}

type TimingPoint struct {
	Time        int
	BeatLength  float64
	Meter       int
	SampleSet   int
	SampleIndex int
	Volume      int
	Uninherited int
	Effects     int
}

type Colour struct {
	Combo               []string
	SliderTrackOverride string
	SliderBorder        string
}

type HitObject struct {
	X            int
	Y            int
	Time         int
	Type         int
	HitSound     int
	ObjectParams ObjectParams
	HitSample    HitSample
}

type HitSample struct {
	normalSet     int
	additionalSet int
	index         int
	volume        int
	filename      string
}

type ObjectParams struct {
	Circle  *Circle
	Slider  *Slider
	Spinner *Spinner
}

type Circle struct {
}

type Slider struct {
	CurveType   string
	CurvePoints []curvepoints
	Slides      int
	length      float64
	edgeSounds  []int
	edgeSets    []string
}

type Spinner struct {
	Endtime int
}

type curvepoints struct {
	X int
	Y int
}
