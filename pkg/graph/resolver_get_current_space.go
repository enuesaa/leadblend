package graph

func (r *Resolver) GetCurrentSpace() (*Space, error) {
	spaceName := r.repos.Using

	return &Space{name: spaceName}, nil
}
