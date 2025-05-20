package models

type Config struct {
	TodoFileName string `json:"todo_file_name"`
}

func DefaultConfig() *Config {
	return &Config{
		TodoFileName: ".todo",
	}
} 