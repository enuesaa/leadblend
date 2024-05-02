package graph

type Event struct {
	page     string
	userName string
}

func (p *Event) Page() string {
	return p.page
}

func (p *Event) UserName() string {
	return p.userName
}
