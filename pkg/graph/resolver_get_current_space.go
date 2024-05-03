package graph

func (r *Resolver) GetCurrentSpace() (*Space, error) {
	return &Space{name: "aaa"}, nil
}
