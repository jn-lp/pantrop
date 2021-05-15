package authdto

type LoginRequest struct {
  Email string `json:"email"`
  Pass  string `json:"pass"`
}
