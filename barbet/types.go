package barbet

import (
  "time"

  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Username   string  `json:"username"`
  Password   string  `json:"-"`
  AvatarUrl  string  `json:"avatar_url"`
  GravatarId string  `json:"gravatar_id"`
  Url        string  `json:"url"`
  Type       string  `json:"type"`
  SiteAdmin  bool    `json:"site_admin"`
  Trips      []Trip  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  Attending  []*Trip `gorm:"many2many:attendants;"`
  Friends    []*User `gorm:"many2many:friends"`
}

type Trip struct {
  gorm.Model
  UserID   uint
  Title    string    `json:"title"`
  StartAt  time.Time `json:"start_at"`
  PitStops bool      `json:"pit_stops"`
  Tempo    string    `json:"tempo"`
  Users    []*User   `gorm:"many2many:attendants;"`
}
