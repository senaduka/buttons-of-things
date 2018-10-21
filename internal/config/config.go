package config

type Config struct {
  Version int
  Buttons []struct {
    Label string
    Id string
  }
}
