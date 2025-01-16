package client

type PlayerRepository interface {
	GetLocalPlayer() (*LocalPlayer, error)
	SaveLocalPlayer(*LocalPlayer) error
}

type WorldRepository interface {
	GetWorld() (*World, error)
	SaveWorld(*World) error
}
