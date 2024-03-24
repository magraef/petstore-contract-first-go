package internal

type PetsRepository interface {
	Repository
	PetsResource
	ReadinessCheck
}

type Repository interface {
	Close()
}
