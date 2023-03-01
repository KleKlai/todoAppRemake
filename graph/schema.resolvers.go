package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/KleKlai/todoappremake/graph/model"
	database "github.com/KleKlai/todoappremake/infra"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	db, err := database.Open()

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	svc := NewTodoService(NewTodoRepository(db))
	res, err := svc.AddTodo(&input)

	if err != nil {
		return nil, fmt.Errorf("Failed to add todo: %v", err)
	}

	return res, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	db, err := database.Open()

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	svc := NewTodoService(NewTodoRepository(db))
	res, err := svc.AddUser(&input)

	if err != nil {
		return nil, fmt.Errorf("Failed to add user: %v", err)
	}

	return res, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	db, err := database.Open()

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	svc := NewTodoService(NewTodoRepository(db))
	res, err := svc.DeleteTodo(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to delete todo: %v", err)
	}

	return res, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	db, err := database.Open()

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	svc := NewTodoService(NewTodoRepository(db))
	res, err := svc.DeleteUser(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to delete user: %v", err)
	}

	return res, nil
}

// UpdateTodoStatus is the resolver for the updateTodoStatus field.
func (r *mutationResolver) UpdateTodoStatus(ctx context.Context, input model.UpdateTodoStatusInput) (*model.UpdateTodoStatus, error) {
	panic(fmt.Errorf("not implemented: UpdateTodoStatus - updateTodoStatus"))
}

// UpdateTodoTask is the resolver for the updateTodoTask field.
func (r *mutationResolver) UpdateTodoTask(ctx context.Context, input model.UpdateTodoTaskInput) (*model.UpdateTodoTask, error) {
	panic(fmt.Errorf("not implemented: UpdateTodoTask - updateTodoTask"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todo - todo"))
}

// GetTodoByStatus is the resolver for the getTodoByStatus field.
func (r *queryResolver) GetTodoByStatus(ctx context.Context, userID string, status bool) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodoByStatus - getTodoByStatus"))
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
