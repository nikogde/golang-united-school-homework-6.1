package golang_united_school_homework

import (
	"errors"
	"reflect"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity == len(b.shapes) {
	  return errors.New("it goes out of range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
	  return nil, errors.New("index went out of the range")
	}
	  for index, shape := range b.shapes {
		if i == index {
		  return shape, nil
		}
	  }
	  return nil, errors.New("index went out of the range")
  }

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, err
  }

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	val, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}

	b.shapes[i] = shape
	
	return val, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum_perimeter := 0.0
	
	for _, shape := range b.shapes {
	  sum_perimeter += shape.CalcPerimeter()
	}
  return sum_perimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum_area := 0.0
	
	for _, shape := range b.shapes {
	  sum_area += shape.CalcArea()
	}
  return sum_area

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var shapes []Shape
	
	for _, shape := range b.shapes {
	  circle := &Circle{}
	  if reflect.TypeOf(circle) != reflect.TypeOf(shape){
      shapes = append(shapes, shape)
	  }
  }
	
  if len(b.shapes) == len(shapes) {
	  return errors.New("element with type 'Circle' aren't exist in the list")
  }
  
  b.shapes = shapes
  return nil
}
