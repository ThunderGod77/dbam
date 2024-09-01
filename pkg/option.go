package pkg

type Option[T any] struct {
	Valid bool
	Value T
}
