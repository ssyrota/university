package core

type LazyData[T any] func() (*T, error)
