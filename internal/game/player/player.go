package player

type Player struct {
	Name string `json:"name"`
}

func New(name string) Player {
	return Player{
		Name: name,
	}
}
