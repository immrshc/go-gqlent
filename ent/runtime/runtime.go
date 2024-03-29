// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"
	"time"

	"github.com/immrshc/go-gqlent/ent/card"
	"github.com/immrshc/go-gqlent/ent/group"
	"github.com/immrshc/go-gqlent/ent/schema"
	"github.com/immrshc/go-gqlent/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescExpired is the schema descriptor for expired field.
	cardDescExpired := cardFields[1].Descriptor()
	// card.DefaultExpired holds the default value on creation for the expired field.
	card.DefaultExpired = cardDescExpired.Default.(time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}

const (
	Version = "v0.13.0"                                         // Version of ent codegen.
	Sum     = "h1:DclxWczaCpyiKn6ZWVcJjq1zIKtJ11iNKy+08lNYsJE=" // Sum of ent codegen.
)
