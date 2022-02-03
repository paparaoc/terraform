package tfcomponents

import (
	"github.com/hashicorp/terraform/internal/terraform-ng/internal/ngaddrs"
	"github.com/hashicorp/terraform/internal/tfdiags"
)

type ComponentGroup struct {
	Name string

	DeclRange tfdiags.SourceRange
}

func (c *ComponentGroup) CallAddr() ngaddrs.ComponentGroupCall {
	return ngaddrs.ComponentGroupCall{Name: c.Name}
}
