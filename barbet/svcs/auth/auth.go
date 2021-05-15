package auth

import (
  "context"

  "github.com/jn-lp/pantrop/barbet/svcs/auth/authdto"
)

type Service interface {
  Login(ctx context.Context, req *authdto.LoginRequest) (res interface{}, err error)
}
