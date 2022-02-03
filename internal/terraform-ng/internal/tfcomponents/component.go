package tfcomponents

import (
	"github.com/hashicorp/terraform/internal/terraform-ng/internal/ngaddrs"
	"github.com/hashicorp/terraform/internal/tfdiags"
)

type Component struct {
	Name string

	DeclRange tfdiags.SourceRange
}

func (c *Component) CallAddr() ngaddrs.ComponentCall {
	return ngaddrs.ComponentCall{Name: c.Name}
}
