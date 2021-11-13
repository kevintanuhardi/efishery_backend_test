package config

// Config struct generate
type Config struct {
	Port struct {
		HTTP int `json:"http"`
	} `json:"ports"`
	JWT_SECRET string `json:"jwt_secret"`
}
