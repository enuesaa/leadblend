package graph

func (r *Resolver) Planets() ([]*Planet, error) {
	list := make([]*Planet, 0)
	list = append(list, &Planet{
		id: "aaa",
		name: "a",
	})
	return list, nil
}
