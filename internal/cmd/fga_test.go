package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/workos/workos-go/v4/pkg/fga"
)

func TestBuildDecisionTreeList(t *testing.T) {
	t.Run("leaf-preserves-check-details", func(t *testing.T) {
		node := fga.DecisionTreeNode{
			Check: fga.WarrantCheck{
				ResourceType: "document",
				ResourceId:   "123",
				Relation:     "viewer",
				Subject: fga.Subject{
					ResourceType: "group",
					ResourceId:   "editors",
					Relation:     "member",
				},
			},
			Policy:         "allow",
			ProcessingTime: 2 * time.Millisecond,
		}

		rendered := buildDecisionTreeList(node).String()
		require.Equal(t, "└── document:123#viewer@group:editors#member - allow (2ms)", rendered)
	})

	t.Run("nested-checks-preserve-tree-branches", func(t *testing.T) {
		node := fga.DecisionTreeNode{
			Check: fga.WarrantCheck{ResourceType: "document", ResourceId: "root"},
			Children: []fga.DecisionTreeNode{
				{Check: fga.WarrantCheck{ResourceType: "document", ResourceId: "first"}},
				{Check: fga.WarrantCheck{ResourceType: "document", ResourceId: "last"}},
			},
		}

		rendered := buildDecisionTreeList(node).String()
		require.Equal(t, "└── document:root#@: (0ms)\n  ├── document:first#@: (0ms)\n  └── document:last#@: (0ms)", rendered)
	})
}
