package componentstree

import (
	"github.com/hashicorp/terraform/internal/terraform-ng/internal/ngaddrs"
)

// Node represents a single node in a components tree. Each node corresponds
// with a single component group.
type Node struct {
	// Parent refers to the parent node of this node, or nil if this is the
	// root node of the tree.
	Parent *Node

	// Root refers to the root node of the tree. It's a self-reference when
	// inside the root node of the tree already.
	Root *Node

	// CallPath is the sequence of static component group calls leading to
	// this node. For the root node in a tree, this has length zero.
	CallPath []ngaddrs.ComponentGroupCall
}
