package structs

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Redius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func Area(width float64, height float64) float64 {
	return width * height
}
