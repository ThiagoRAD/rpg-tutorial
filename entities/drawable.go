package entities

type Drawable struct {
	*Sprite
	*Position
}

type Position struct {
	X, Y float64
}
