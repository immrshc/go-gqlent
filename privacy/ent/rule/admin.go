package rule

import (
	"context"

	"entgo.io/ent/privacy"
	"github.com/immrshc/go-gqlent/privacy/viewer"
)

func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if v := viewer.FromContext(ctx); v == nil {
			return privacy.Denyf("viewer-context is missing")
		}
		return privacy.Skip
	})
}

func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if v := viewer.FromContext(ctx); v.Admin() {
			return privacy.Allow
		}
		return privacy.Skip
	})
}
