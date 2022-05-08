package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rs/xid"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id xid.ID, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id xid.ID) (*ent.User, error) {
	client := ent.FromContext(ctx).User
	record, err := client.Query().Where(User.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to querying user with id %q: %v", id, err)
	}

	deleteErr := client.DeleteOne(record).Exec(ctx)
	if deleteErr != nil {
		return nil, fmt.Errorf("failed to delete user with id %q: %v", id, deleteErr)
	}
	return record, nil
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}
