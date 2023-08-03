package main

type Point struct {
	X int
	Y int
}

func (p Point) MovablePoints(length int) []Point {
	tmpPoints := []Point{}
	tmpPoints = append(
		tmpPoints,
		Point{p.X + 1, p.Y + 2},
		Point{p.X - 1, p.Y + 2},
		Point{p.X + 1, p.Y - 2},
		Point{p.X - 1, p.Y - 2},
		Point{p.X + 2, p.Y + 1},
		Point{p.X - 2, p.Y + 1},
		Point{p.X + 2, p.Y - 1},
		Point{p.X - 2, p.Y - 1},
	)

	points := []Point{}
	for _, v := range tmpPoints {
		if v.InRange(length) {
			points = append(points, v)
		}
	}

	return points
}

func (p Point) InRange(length int) bool {
	return 0 <= p.X && p.X < length && 0 <= p.Y && p.Y < length
}
