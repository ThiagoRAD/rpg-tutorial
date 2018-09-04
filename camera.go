package main

type Camera struct {
	X, Y float64
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) FollowTarget(targetX, targetY float64, screenWidth, screenHeight float64) {
	c.X = screenWidth/2.0 - targetX
	c.Y = screenHeight/2.0 - targetY
}

func (c *Camera) Constrain(tilemapWidthPixels, tilemapHeightPixels, screenWidth, screenHeight float64) {
	if c.X > 0 {
		c.X = 0
	}
	if c.Y > 0 {
		c.Y = 0
	}
	if c.X < screenWidth-tilemapWidthPixels {
		c.X = screenWidth - tilemapWidthPixels
	}
	if c.Y < screenHeight-tilemapHeightPixels {
		c.Y = screenHeight - tilemapHeightPixels
	}
}
