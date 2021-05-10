package users

import (
  "context"

  "github.com/jn-lp/pantrop/barbet"
  "github.com/jn-lp/pantrop/barbet/svcs/users/usersrepo"
)

type Service interface {
  ListUsers(ctx context.Context) (*[]barbet.User, error)
  CreateUser(ctx context.Context, user *barbet.User) (*barbet.User, error)
  GetUser(ctx context.Context, username string) (*barbet.User, error)
  UpdateUser(ctx context.Context, user *barbet.User) (*barbet.User, error)
  DeleteUser(ctx context.Context, username string) error
}

type service struct {
  repository usersrepo.Repository
}

func New(r usersrepo.Repository) Service {
  return &service{
    repository: r,
  }
}

func (s *service) ListUsers(ctx context.Context) (*[]barbet.User, error) {
  return s.repository.ListUsers(ctx)
}

func (s *service) CreateUser(ctx context.Context, user *barbet.User) (*barbet.User, error) {
  return s.repository.CreateUser(ctx, user)
}

func (s *service) GetUser(ctx context.Context, username string) (*barbet.User, error) {
  return s.repository.GetUser(ctx, username)
}

func (s *service) UpdateUser(ctx context.Context, user *barbet.User) (*barbet.User, error) {
  return s.repository.UpdateUser(ctx, user)
}

func (s *service) DeleteUser(ctx context.Context, username string) error {
  return s.repository.DeleteUser(ctx, username)
}
