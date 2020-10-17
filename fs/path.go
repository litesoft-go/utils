package fs

import "fmt"

type Path struct {
	parent *Path
	leg    string
}

func (p *Path) Plus(leg string) *Path {
	if leg == "" {
		return p
	}
	return &Path{parent: p, leg: leg}
}

func (p *Path) ErrorOf(why string, currentValue interface{}) error {
	return fmt.Errorf("invalid %s %s of: %v", p.String(), why, currentValue)
}

func (p *Path) prefix() string {
	if p == nil {
		return ""
	}
	return p.String() + PathSeparator
}

func (p *Path) String() string {
	return p.parent.prefix() + p.leg
}
