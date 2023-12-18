package types
type Item struct {
	Type interface{}
	Value  interface{}
	HasCHildren bool
	Children []Item
}