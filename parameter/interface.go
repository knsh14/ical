package parameter

type Container map[TypeName][]Base

type Base interface {
	implementParameter()
}
