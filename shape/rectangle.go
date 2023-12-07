package shape

type Rectangle struct {
	Width  float32
	Length float32
}

func (rectangel *Rectangle) Area() float32 {
	return rectangel.Width * rectangel.Length
}

func (rectangel *Rectangle) Perimeter() float32 {
	return 2 * (rectangel.Width + rectangel.Length)
}
