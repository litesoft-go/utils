package paths

import "fmt"

type Path struct {
	parent *Path
	leg    string
}

//noinspection GoUnusedExportedFunction
func New(leg string) *Path {
	return &Path{leg: leg}
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
	return p.String() + "."
}

func (p *Path) String() string {
	return p.parent.prefix() + p.leg
}

func (p *Path) Negative(pathLeg string, currentValue interface{}) error {
	return p.Plus(pathLeg).
		ErrorOf("negative", currentValue)
}

func (p *Path) Zero(pathLeg string, currentValue interface{}) error {
	return p.Plus(pathLeg).
		ErrorOf("zero", currentValue)
}

func (p *Path) Missing(pathLeg string, currentValue interface{}) error {
	return p.Plus(pathLeg).
		ErrorOf("missing", currentValue)
}

func (p *Path) LeadingTrailingWhitespace(pathLeg, currentValue string) error {
	return p.Plus(pathLeg).
		ErrorOf("leading or trailing whitespace", "'"+currentValue+"'")
}
