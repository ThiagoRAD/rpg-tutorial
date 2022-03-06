package entities

type Enemy struct {
	*Drawable
	FollowsPlayer bool
}
