package graph

type QueryIslandArgs struct {
	PlanetId string
}
func (r *Resolver) Islands(args QueryIslandArgs) ([]*Island, error) {
	return make([]*Island, 0), nil
}
