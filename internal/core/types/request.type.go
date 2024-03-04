package types

type Request struct {
	Body          func(obj any) error
	GetParam      func(key string) string
	GetQueryParam func(key string) string
	GetHeader     func(key string) string
}
