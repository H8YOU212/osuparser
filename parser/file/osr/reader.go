package osr

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/ulikunitz/xz/lzma"
)

func readString(r io.Reader) (string, error) {
	var marker byte
	if err := binary.Read(r, binary.LittleEndian, &marker); err != nil {
		return "", fmt.Errorf("read marker: %w", err)
	}

	switch marker {
	case 0x00:
		return "", nil
	case 0x0B:
		// продолжаем
	default:
		return "", fmt.Errorf("invalid string marker: 0x%X", marker)
	}

	length, err := readULEB128(r)
	if err != nil {
		return "", fmt.Errorf("read uleb128: %w", err)
	}

	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", fmt.Errorf("read string bytes: %w", err)
	}

	return string(buf), nil
}

func readULEB128(r io.Reader) (uint32, error) {
	var result uint32
	var shift uint

	for {
		var b [1]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return 0, fmt.Errorf("read uleb128 byte: %w", err)
		}

		if shift >= 32 {
			return 0, fmt.Errorf("uleb128 overflow")
		}

		result |= uint32(b[0]&0x7F) << shift

		if b[0]&0x80 == 0 {
			break
		}

		shift += 7
	}

	if result > math.MaxInt32 {
		return 0, fmt.Errorf("uleb128 value too large: %d", result)
	}

	return result, nil
}

func (o *OsrModel) readReplayData(r io.Reader) error {
	compressed := make([]byte, o.format.CompressedSize)
	if _, err := io.ReadFull(r, compressed); err != nil {
		return fmt.Errorf("read compressed bytes: %w", err)
	}

	lzmaReader, err := lzma.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return fmt.Errorf("create lzma reader: %w", err)
	}

	decompressed, err := io.ReadAll(lzmaReader)
	if err != nil {
		return fmt.Errorf("decompress lzma: %w", err)
	}

	o.format.ReplayData = string(decompressed)
	return nil
}