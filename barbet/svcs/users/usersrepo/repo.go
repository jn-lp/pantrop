package usersrepo

import (
  "context"
  "fmt"

  _ "github.com/lib/pq"
  "gorm.io/gorm"

  "github.com/jn-lp/pantrop/barbet"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
  ListUsers(ctx context.Context) (*[]barbet.User, error)
  CreateUser(ctx context.Context, user *barbet.User) (*barbet.User, error)
  GetUser(ctx context.Context, username string) (*barbet.User, error)
  UpdateUser(ctx context.Context, user *barbet.User) (*barbet.User, error)
  DeleteUser(ctx context.Context, username string) error
}

type repository struct {
  db *gorm.DB
}

func New(db *gorm.DB) Repository {
  return &repository{
    db: db,
  }
}

func (r repository) ListUsers(ctx context.Context) (*[]barbet.User, error) {
  var users []barbet.User
  r.db.WithContext(ctx).Find(&users)

  return &users, nil
}

func (r repository) CreateUser(ctx context.Context, user *barbet.User) (*barbet.User, error) {
  r.db.WithContext(ctx).Create(user)

  return user, nil
}

func (r repository) GetUser(ctx context.Context, username string) (*barbet.User, error) {
  var user barbet.User
  r.db.WithContext(ctx).First(&user, username)

  if user.ID == 0 {
    return nil, fmt.Errorf("trip not found")
  }

  return &user, nil
}

func (r repository) UpdateUser(ctx context.Context, user *barbet.User) (*barbet.User, error) {
  r.db.WithContext(ctx).Save(&user)

  return user, nil
}

func (r repository) DeleteUser(ctx context.Context, username string) error {
  r.db.WithContext(ctx).Delete(username)

  return nil
}
