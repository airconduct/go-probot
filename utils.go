package probot

func ToPointer[T any](s T) *T {
	return &s
}
