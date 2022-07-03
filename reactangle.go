package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rec Rectangle) CalcPerimeter() float64 {
	return 2 * (rec.Height + rec.Weight)
}

func (rec Rectangle) CalcArea() float64 {
	return rec.Height * rec.Weight
}
