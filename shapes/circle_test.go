package shapes_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"softwaredesign-workshop-go/shapes"
	"testing"
)

func Test_Should_Always_Have_A_Radius_Larger_0(t *testing.T) {
	_, err := shapes.CreateCircle(0, 0, 0)
	assert.Equal(t, fmt.Errorf("radius needs to be larger 0"), err)

	_, err = shapes.CreateCircle(0, 0, -1)

	assert.Equal(t, fmt.Errorf("radius needs to be larger 0"), err)
}

func Test_Should_Handle_Coordinates_With_Unequal_Length(t *testing.T) {
	circle, _ := shapes.CreateCircle(0, 0, 2)

	xCoords := []int{2, 3, 4, -12, -20}
	yCoords := []int{8, 20, 15, -4}

	_, err := circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, fmt.Errorf("not every provided X coordinate has a matching Y coordinate"), err)
}

func Test_Should_Handle_Empty_X_Coordinates(t *testing.T) {
	var xCoords []int
	yCoords := []int{8, 20, 15, -4}
	circle, _ := shapes.CreateCircle(0, 0, 20)

	_, err := circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, fmt.Errorf("X coordinates are empty"), err)
}

func Test_Should_Handle_Empty_Y_Coordinates(t *testing.T) {
	circle, _ := shapes.CreateCircle(0, 0, 2)

	xCoords := []int{8, 20, 15, -4}
	var yCoords []int

	_, err := circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, fmt.Errorf("Y coordinates are empty"), err)
}

func Test_Should_Count_Contained_Points(t *testing.T) {
	circle, _ := shapes.CreateCircle(5, -5, 10)

	xCoords := []int{2, 1, 3, 8, 4, -12, -20, -4}
	yCoords := []int{8, 1, 20, -4, 15, -4, -20, -4}

	cnt, _ := circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, 3, cnt)
}

func Test_Should_Count_Contained_Points_When_Resized_For_New_Area(t *testing.T) {
	circle, _ := shapes.CreateCircle(0, 0, 20)

	xCoords := []int{2, 3, 4, -12, -20}
	yCoords := []int{8, 20, 15, -4, -20}

	cnt, _ := circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, 3, cnt)

	circle.Resize(40)
	cnt, _ = circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, 5, cnt)

	circle.Resize(1)
	cnt, _ = circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, 0, cnt)
}

func Test_Should_Count_Contained_Points_When_Moved_To_Another_Location(t *testing.T) {
	circle, _ := shapes.CreateCircle(0, 0, 20)

	xCoords := []int{2, 3, 4, -12, -20}
	yCoords := []int{8, 20, 15, -4, -20}

	cnt, _ := circle.CountContainedPoints(xCoords, yCoords)

	assert.Equal(t, 3, cnt)

	circle.MoveTo(-11, -20)
	cnt, _ = circle.CountContainedPoints(xCoords, yCoords)
	assert.Equal(t, 2, cnt)
}

func Test_Should_Format_Location_And_Radius_As_String(t *testing.T) {
	circle, _ := shapes.CreateCircle(1, 4, 7)

	assert.Equal(t, "(1, 4), 7", circle.Format())
}
