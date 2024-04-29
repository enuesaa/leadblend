package graph

type CreatePlanetArgs struct {
	Name string
}
func (r *Resolver) CreatePlanet(args CreatePlanetArgs) (*string, error) {
	id := "a"
	return &id, nil
}
