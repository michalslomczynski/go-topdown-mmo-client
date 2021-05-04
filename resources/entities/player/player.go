package player

const (
	None = iota
	Left
	Right
	Up
	Down
)

type Position struct {
	X               int16
	Y               int16
	MovingDirection int8
}

type Player struct {
	ID              int16
	X               int16
	Y               int16
	MovingDirection int8
	Level           int16
	Outfit          int16

	/* local resources */
	Ux   float64
	Uy   float64
	Name string
}

var Hero Player
