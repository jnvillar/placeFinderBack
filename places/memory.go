package places

type MemoryBackend struct {
	places []*Place
}

func NewMemoryBackend() *MemoryBackend{
	return &MemoryBackend{
		places:[]*Place{},
	}
}

func (b *MemoryBackend) List() []*Place {
	return b.places
}

func (b *MemoryBackend) Create(place *Place) error {
	b.places = append(b.places, place)
	return nil
}