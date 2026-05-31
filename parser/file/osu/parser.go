package osu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (o *Commonosu) parse() error {
	file, err := os.Open(o.Path)
	if err != nil {
		return fmt.Errorf("osu parse open file: %w", err)
	}
	defer file.Close()

	section := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "[") {
			section = line
			continue
		}

		switch section {

		case "[General]":
			o.Osufile.parseGeneral(line)
		case "[Editor]":
			o.Osufile.parseEditor(line)
		case "[Metadata]":
			o.Osufile.parseMetadata(line)
		case "[Difficulty]":
			o.Osufile.parseDifficulty(line)
		case "[TimingPoints]":
			o.Osufile.parseTimingPoints(line)
		case "[HitObjects]":
			o.Osufile.parseHitObjects(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("osu parse scanner: %w", err)
	}

	return nil
}

func (o *Commonosu) ParseFile(path string) error {
	if path != "" {
		o.Path = path
	}
	return o.parse()
}

func splitKV(line string) (string, string, bool) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return "", "", false
	}

	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])
	return key, val, true
}

func (b *Beatmap) parseGeneral(line string) {
	key, val, ok := splitKV(line)
	if !ok {
		return
	}

	var err error
	switch key {
	case "AudioFilename":
		b.General.AudioFilename = val
	case "AudioLeadIn":
		b.General.AudioLeadIn, _ = strconv.Atoi(val)
	case "PreviewTime":
		b.General.PreviewTime, _ = strconv.Atoi(val)
	case "Countdown":
		b.General.Countdown, _ = strconv.Atoi(val)
	case "SampleSet":
		b.General.SampleSet = val
	case "StackLeniency":
		b.General.StackLeniency, _ = strconv.ParseFloat(val, 64)
	case "Mode":
		b.General.Mode, _ = strconv.Atoi(val)
	case "LetterboxInBreaks":
		b.General.LetterboxInBreaks, _ = strconv.Atoi(val)
	case "EpilepsyWarning":
		b.General.EpilepsyWarning, _ = strconv.Atoi(val)
	case "WidescreenStoryboard":
		b.General.WidescreenStoryboard, _ = strconv.Atoi(val)
	}
	_ = err
}

func (b *Beatmap) parseEditor(line string) {
	key, val, ok := splitKV(line)
	if !ok {
		return
	}

	var err error
	switch key {
	case "Bookmarks":
		b.Editor.Bookmarks = val
	case "DistanceSpacing":
		b.Editor.DistanceSpacing, _ = strconv.ParseFloat(val, 64)
	case "BeatDivisor":
		b.Editor.BeatDivisor, _ = strconv.Atoi(val)
	case "GridSize":
		b.Editor.GridSize, _ = strconv.Atoi(val)
	case "TimelineZoom":
		b.Editor.TimelineZoom, _ = strconv.ParseFloat(val, 64)
	}
	_ = err
}

func (b *Beatmap) parseMetadata(line string) {
	key, val, ok := splitKV(line)
	if !ok {
		return
	}

	var err error
	switch key {
	case "Title":
		b.Metadata.Title = val
	case "TitleUnicode":
		b.Metadata.TitleUnicode = val
	case "Artist":
		b.Metadata.Artist = val
	case "ArtistUnicode":
		b.Metadata.ArtistUnicode = val
	case "Creator":
		b.Metadata.Creator = val
	case "Version":
		b.Metadata.Version = val
	case "Source":
		b.Metadata.Source = val
	case "Tags":
		b.Metadata.Tags = val
	case "BeatmapID":
		b.Metadata.BeatmapID, _ = strconv.Atoi(val)
	case "BeatmapSetID":
		b.Metadata.BeatmapSetID, _ = strconv.Atoi(val)
	}
	_ = err
}

func (b *Beatmap) parseDifficulty(line string) {
	key, val, ok := splitKV(line)
	if !ok {
		return
	}

	var err error
	switch key {
	case "HPDrainRate":
		b.Difficulty.HP, _ = strconv.ParseFloat(val, 64)
	case "CircleSize":
		b.Difficulty.CS, _ = strconv.ParseFloat(val, 64)
	case "OverallDifficulty":
		b.Difficulty.OD, _ = strconv.ParseFloat(val, 64)
	case "ApproachRate":
		b.Difficulty.AR, _ = strconv.ParseFloat(val, 64)
	case "SliderMultiplier":
		b.Difficulty.SliderMultiplier, _ = strconv.ParseFloat(val, 64)
	case "SliderTickRate":
		b.Difficulty.SliderTickRate, _ = strconv.ParseFloat(val, 64)
	}
	_ = err
}

func (b *Beatmap) parseTimingPoints(line string) {
	parts := strings.Split(line, ",")
	if len(parts) < 8 {
		return
	}

	var err error
	newobj := TimingPoint{}
	newobj.Time, _ = strconv.Atoi(parts[0])
	newobj.BeatLength, _ = strconv.ParseFloat(parts[1], 64)
	newobj.Meter, _ = strconv.Atoi(parts[2])
	newobj.SampleSet, _ = strconv.Atoi(parts[3])
	newobj.SampleIndex, _ = strconv.Atoi(parts[4])
	newobj.Volume, _ = strconv.Atoi(parts[5])
	newobj.Uninherited, _ = strconv.Atoi(parts[6])
	newobj.Effects, _ = strconv.Atoi(parts[7])

	b.TimingPoints = append(b.TimingPoints, newobj)
	_ = err
}

func (b *Beatmap) parseHitObjects(line string) {
	splited := strings.Split(line, ",")
	if len(splited) < 5 {
		return
	}

	newobj := HitObject{}

	newobj.X, _ = strconv.Atoi(splited[0])
	newobj.Y, _ = strconv.Atoi(splited[1])
	newobj.Time, _ = strconv.Atoi(splited[2])
	newobj.Type, _ = strconv.Atoi(splited[3])
	newobj.HitSound, _ = strconv.Atoi(splited[4])

	if newobj.Type&2 != 0 {
		if len(splited) >= 8 {
			newobj.ObjectParams.Slider = &Slider{}
			newobj.ObjectParams.Slider.handleslider(splited[5:8])
		}
	} else if newobj.Type&8 != 0 {
		if len(splited) >= 7 {
			newobj.ObjectParams.Spinner = &Spinner{}
			newobj.ObjectParams.Spinner.handlespinner(splited[5])
		}
	}

	b.HitObjects = append(b.HitObjects, newobj)
}
