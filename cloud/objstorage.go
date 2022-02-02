package cloud

type ObjectStore interface {
	Client() interface{}
}
