package internal

// ReadinessCheck defines an interface for components that perform readiness checks.
// Implementations of this interface should provide a Check method that returns an error if the component is not ready.
type ReadinessCheck interface {
	// Check performs a readiness check and returns an error if the component is not ready
	Check() error
}
