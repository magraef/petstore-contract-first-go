package internal

// PetsRepository defines an interface that aggregates the behavior of a repository for pets,
// ensuring that it includes methods from Repository, PetsResource, and ReadinessCheck interfaces.
type PetsRepository interface {
	Repository
	PetsResource
}

// Repository defines an interface for repositories in general, including ReadinessCheck interface and a Close method.
// It serves as a basic contract for any repository implementations to implement.
type Repository interface {
	ReadinessCheck
	// Close closes the repository, releasing any resources it holds
	Close()
}
