package osr

import (
	"encoding/binary"
	"fmt"
	"os"
)

const modTargetPractice int32 = 8388608

func (o *OsrModel) Decode() error {
	file, err := os.Open(o.dir)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	if err := binary.Read(file, binary.LittleEndian, &o.format.Header); err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	if o.format.ReplayInfo.BeatmapMD5, err = readString(file); err != nil {
		return fmt.Errorf("read beatmap md5: %w", err)
	}

	if o.format.ReplayInfo.Username, err = readString(file); err != nil {
		return fmt.Errorf("read username: %w", err)
	}

	if o.format.ReplayInfo.ReplayMD5, err = readString(file); err != nil {
		return fmt.Errorf("read replay md5: %w", err)
	}

	if err := binary.Read(file, binary.LittleEndian, &o.format.Stats); err != nil {
		return fmt.Errorf("read stats: %w", err)
	}

	if o.format.LifeBarGraph, err = readString(file); err != nil {
		return fmt.Errorf("read life bar graph: %w", err)
	}

	if err := binary.Read(file, binary.LittleEndian, &o.format.Timestamp); err != nil {
		return fmt.Errorf("read timestamp: %w", err)
	}

	if err := binary.Read(file, binary.LittleEndian, &o.format.CompressedSize); err != nil {
		return fmt.Errorf("read compressed size: %w", err)
	}

	if err := o.readReplayData(file); err != nil {
		return fmt.Errorf("read replay data: %w", err)
	}

	if err := binary.Read(file, binary.LittleEndian, &o.format.Additional.ScoreID); err != nil {
		return fmt.Errorf("read score id: %w", err)
	}

	if o.format.Stats.Mods&modTargetPractice != 0 {
		if err := binary.Read(file, binary.LittleEndian, &o.format.Additional.TargetPracticeAccuracy); err != nil {
			return fmt.Errorf("read target practice accuracy: %w", err)
		}
	}

	return nil
}