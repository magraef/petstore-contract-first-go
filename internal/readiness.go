package internal

type ReadinessCheck interface {
	Check() error
}
