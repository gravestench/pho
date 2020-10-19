package point

type PointNamespace interface {
	New(x, y float64) *Point
	GetCentroid(points []*Point, out *Point) *Point
	Ceil(p *Point) *Point
	Floor(p *Point) *Point
	Clone(p *Point) *Point
	CopyFrom(source, dest *Point) *Point
	Equals(a, b *Point) bool
	GetMagnitude(p *Point) float64
	GetMagnitudeSquared(p *Point) float64
	Interpolate(a, b *Point, t float64, out *Point) *Point
	Invert(p *Point) *Point
	Negative(p *Point) *Point
	Project(a, b, out *Point) *Point
	SetMagnitude(p *Point, magnitude float64) *Point
}

type Namespace struct{}

func (*Namespace) New(x, y float64) *Point {
	return New(x, y)
}

func (*Namespace) GetCentroid(points []*Point, out *Point) *Point {
	return GetCentroid(points, out)
}

func (*Namespace) Ceil(p *Point) *Point {
	return Ceil(p)
}

func (*Namespace) Floor(p *Point) *Point {
	return Floor(p)
}

func (*Namespace) Clone(p *Point) *Point {
	return Clone(p)
}

func (*Namespace) CopyFrom(source, dest *Point) *Point {
	return CopyFrom(source, dest)
}

func (*Namespace) Equals(a, b *Point) bool {
	return Equals(a, b)
}

func (*Namespace) GetMagnitude(p *Point) float64 {
	return GetMagnitude(p)
}

func (*Namespace) GetMagnitudeSquared(p *Point) float64 {
	return GetMagnitudeSquared(p)
}

func (*Namespace) Interpolate(a, b *Point, t float64, out *Point) *Point {
	return Interpolate(a, b, t, out)
}

func (*Namespace) Invert(p *Point) *Point {
	return Invert(p)
}

func (*Namespace) Negative(p *Point) *Point {
	return Negative(p)
}

func (*Namespace) Project(a, b, out *Point) *Point {
	return Project(a, b, out)
}

func (*Namespace) SetMagnitude(p *Point, magnitude float64) *Point {
	return SetMagnitude(p, magnitude)
}
