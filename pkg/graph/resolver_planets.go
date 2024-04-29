package graph

func (*Resolver) Planets() ([]Planet, error) {
	list := make([]Planet, 0)
	list = append(list, Planet{})
	return list, nil
}
