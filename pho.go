package pho

import (
	"github.com/gravestench/pho/geom/circle"
	"github.com/gravestench/pho/geom/line"
	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/geom/rectangle"
)

type Pho struct {
	Geom *Geom
}

type Geom struct {
	Circle    *circle.Namespace
	Rectangle *rectangle.Namespace
	Line      *line.Namespace
	Point     *point.Namespace
}

func New() *Pho {
	return &Pho{
		&Geom{
			Circle:    &circle.Namespace{},
			Rectangle: &rectangle.Namespace{},
			Line:      &line.Namespace{},
			Point:     &point.Namespace{},
		},
	}
}
