package util

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Point struct {
	X float32
	Y float32
	Z float32
}

const (
	ModeModel    = 1
	ModelPredict = 2
)

func ParsePoint(line string) (*Point, error) {
	axes := strings.Split(line, "\t")
	if len(axes) != 3 {
		return nil, errors.New("invalid coordinate number")
	}

	p := new(Point)
	val, err := strconv.ParseFloat(axes[0], 10)
	if err != nil {
		return nil, err
	}
	p.X = float32(val)
	val, err = strconv.ParseFloat(axes[1], 10)
	if err != nil {
		return nil, err
	}
	p.Y = float32(val)
	val, err = strconv.ParseFloat(axes[2], 10)
	if err != nil {
		return nil, err
	}
	p.Z = float32(val)

	return p, nil
}
