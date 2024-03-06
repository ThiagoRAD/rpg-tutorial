package main

type Camera struct {
	X, Y float64
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) FollowTarget(targetX, targetY float64, screenWidth, screenHeight int) {
	c.X = float64(screenWidth)/2.0 - targetX
	c.Y = float64(screenHeight)/2.0 - targetY
}
