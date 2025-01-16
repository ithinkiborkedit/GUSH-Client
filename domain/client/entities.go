package client

type LocalPlayer struct {
	ID   string
	Name string
	Room string
}

type LocalRoom struct {
	ID          string
	Name        string
	Description string
}

type World struct {
	Player *LocalPlayer
	Rooms  map[string]*LocalRoom
}

func NewWorld() *World {
	return &World{
		Rooms: make(map[string]*LocalRoom),
	}
}
