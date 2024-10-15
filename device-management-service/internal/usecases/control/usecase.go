package control

type UseCase interface {
}

type Storage interface {
}

type Implementation struct {
	storage Storage
}

func New(storage Storage) *Implementation {
	return &Implementation{storage: storage}
}
