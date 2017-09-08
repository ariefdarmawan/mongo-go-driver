package ops

import (
	"context"

	"github.com/10gen/mongo-go-driver/bson"
	"github.com/10gen/mongo-go-driver/yamgo/internal"
	"github.com/10gen/mongo-go-driver/yamgo/options"
)

// Delete executes an delete command with a given set of delete documents and options.
func Delete(ctx context.Context, s *SelectedServer, ns Namespace, deleteDocs []bson.D,
	result interface{}, options ...options.DeleteOption) error {

	if err := ns.validate(); err != nil {
		return err
	}

	command := bson.D{
		{Name: "delete", Value: ns.Collection},
		{Name: "deletes", Value: deleteDocs},
	}

	for _, option := range options {
		command.AppendElem(option.DeleteName(), option.DeleteValue())
	}

	// TODO GODRIVER-27: write concern

	err := runMustUsePrimary(ctx, s, ns.DB, command, result)
	if err != nil {
		return internal.WrapError(err, "failed to execute delete")
	}

	return nil
}
