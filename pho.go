package pho

import (
	"github.com/gravestench/pho/geom/circle"
	"github.com/gravestench/pho/geom/rectangle"
)

type Pho struct {
	*Geom
}

type Geom struct {
	Circle *circle.Namespace
	Rectangle *rectangle.Namespace
}

func New() *Pho {
	return &Pho{
		&Geom{
			Circle: &circle.Namespace{},
			Rectangle: &rectangle.Namespace{},
		},
	}
}
