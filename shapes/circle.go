package shapes

import "fmt"

type Circle struct {
	X int
	Y int
	R int
}

func CreateCircle(x int, y int, r int) (Circle, error) {
	if r <= 0 {
		return Circle{}, fmt.Errorf("radius needs to be larger 0")
	}
	return Circle{X: x, Y: y, R: r}, nil
}

func (circle *Circle) CountContainedPoints(xCoords []int, yCoords []int) (int, error) {
	nrPts := 0
	if xCoords != nil && len(xCoords) != 0 {
		if yCoords != nil && len(yCoords) != 0 {
			if len(xCoords) == len(yCoords) {
				for i := 0; i < len(xCoords); i++ {
					nrPts = circle.Contains(xCoords, yCoords, i, nrPts)
				}
				return nrPts, nil
			} else {
				return 0, fmt.Errorf("not every provided X coordinate has a matching Y coordinate")
			}
		} else {
			return 0, fmt.Errorf("Y coordinates are empty")
		}
	} else {
		return 0, fmt.Errorf("X coordinates are empty")
	}
}

func (circle *Circle) Contains(xCoords []int, yCoords []int, i int, nrPts int) int {
	result := (xCoords[i]-circle.X)*(xCoords[i]-circle.X)+(yCoords[i]-circle.Y)*(yCoords[i]-circle.Y) <= circle.R*circle.R

	if result {
		nrPts++
	}

	return nrPts
}

func (circle *Circle) Resize(r int) {
	circle.R = r
}

func (circle *Circle) Format() string {
	return fmt.Sprintf("(%d, %d), %d", circle.X, circle.Y, circle.R)
}

func (circle *Circle) MoveTo(x int, y int) {
	circle.X = x
	circle.Y = y
}
