package places

type Place struct {
	Latitude  string
	Longitude string
	Name      string
}

type Backend interface {
	List() []*Place
	Create(place *Place) error
}

type Factory struct {
	Backend
}

func NewFactory() *Factory{
	backend := NewMemoryBackend()
	factory := &Factory{Backend:backend}
	return factory
}