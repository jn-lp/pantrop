package users

import (
  "context"

  "github.com/jn-lp/pantrop/barbet"
)

type Service interface {
  ListUsers(ctx context.Context) (*[]barbet.User, error)
  CreateUser(ctx context.Context, user *barbet.User) (*barbet.User, error)
  GetUser(ctx context.Context, username string) (*barbet.User, error)
  UpdateUser(ctx context.Context, user *barbet.User) (*barbet.User, error)
  DeleteUser(ctx context.Context, username string) error
}
