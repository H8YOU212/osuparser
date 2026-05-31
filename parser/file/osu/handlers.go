package osu

import (
	"strconv"
	"strings"
)

func (s *Spinner) handlespinner(line string) {
	s.Endtime, _ = strconv.Atoi(line)
}

func (s *Slider) handleslider(line []string) {
	if len(line) < 3 {
		return
	}

	parts := strings.Split(line[0], "|")
	if len(parts) == 0 {
		return
	}

	s.CurveType = parts[0]

	pointspart := parts[1:]

	for _, points := range pointspart {
		splitedpoints := strings.Split(points, ":")
		if len(splitedpoints) < 2 {
			continue
		}

		x, xerr := strconv.Atoi(splitedpoints[0])
		y, yerr := strconv.Atoi(splitedpoints[1])
		if xerr != nil || yerr != nil {
			continue
		}

		s.CurvePoints = append(s.CurvePoints, curvepoints{X: x, Y: y})
	}

	s.Slides, _ = strconv.Atoi(line[1])
	s.length, _ = strconv.ParseFloat(line[2], 64)
}
