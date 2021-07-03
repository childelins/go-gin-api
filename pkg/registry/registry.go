package registry

type Registry interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}
