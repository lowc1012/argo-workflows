package cache

import (
	"context"
	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

type artifactDriverCache struct {
	name string
}

func NewArtifactDriverCache(n string) MemoizationCache {
	return &artifactDriverCache{
		name: n,
	}
}

func (c *artifactDriverCache) Load(ctx context.Context, key string) (*Entry, error) {
	return nil, nil
}

func (c *artifactDriverCache) Save(ctx context.Context, key string, nodeId string, value *wfv1.Outputs) error {
	return nil
}
