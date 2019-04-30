package smartbalancer

type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}
