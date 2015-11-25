package director

type Director interface {
	Status() string
}
type director struct {
	target string
}

func NewDirector(target string) Director {
	return &director{target}
}

func (d *director) Status() string {
	return d.target
}
