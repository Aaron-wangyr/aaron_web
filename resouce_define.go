package aaronweb

import (
	"context"

	"github.com/Aaron-wangyr/aaron-json"
)

// IResouceManager defines the interface for resource management(manager one or more resource instances), in db layer, this is usually a table or collection.
type IResouceManager interface {
	ResourceName() string
	ResourceNamePlural() string

	HasName() bool
	ValidateName(name string) bool

	ValidateListConditions(ctx context.Context, query *aaronjson.JsonObject) error
}

// IResource defines the interface for a resource instance, in db layer, this is usually a row or document.
type IResource interface {
}
