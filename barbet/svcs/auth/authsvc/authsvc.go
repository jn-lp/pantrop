package authsvc

import (
  "context"
  "errors"
  "time"

  "github.com/form3tech-oss/jwt-go"
  "github.com/jn-lp/pantrop/barbet/svcs/auth"
  "github.com/jn-lp/pantrop/barbet/svcs/auth/authdto"
  "github.com/jn-lp/pantrop/barbet/svcs/users"
)

type service struct {
  // repository usersrepo.Repository
  users users.Service
}

func New(
// r usersrepo.Repository,
  us users.Service,
) auth.Service {
  return &service{
    users: us,
  }
}

func (s service) Login(ctx context.Context, req *authdto.LoginRequest) (res interface{}, err error) {
  if req.Email != "eugene@lepei.co" || req.Pass != "1234" {
    return nil, errors.New("unauthorized")
  }

  token := jwt.New(jwt.GetSigningMethod("HS256"))

  claims := token.Claims.(jwt.MapClaims)
  claims["name"] = "John Doe"
  claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

  // FIXME:
  t, err := token.SignedString([]byte("secret"))
  if err != nil {
    return nil, err
  }

  return t, nil
}
