package config

type Config struct {
	WebServer struct {
		Port int `json:"Port"`
		GIN  struct {
			ReleaseMode bool `json:"ReleaseMode"`
			UseLogger   bool `json:"UseLogger"`
			UseRecovery bool `json:"UseRecovery"`
		} `json:"GIN"`
	} `json:"WebServer"`
	DB struct {
		Host     string `json:"Host"`
		User     string `json:"User"`
		Password string `json:"Password"`
		Port     int    `json:"Port"`
	} `json:"DB"`
	Cache struct {
		Host     string `json:"Host"`
		Password string `json:"Password"`
		Port     int    `json:"Port"`
	} `json:"Cache"`
}
