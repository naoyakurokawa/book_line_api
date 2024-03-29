package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env        string   `env:"ENV" envDefault:"dev"`
	Port       int      `env:"PORT" envDefault:"80"`
	DBHost     string   `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int      `env:"DB_PORT" envDefault:"33306"`
	DBUser     string   `env:"DB_USER" envDefault:"book_line"`
	DBPassword string   `env:"DB_PASSWORD" envDefault:"book_line"`
	DBName     string   `env:"DB_NAME" envDefault:"book_line"`
	RedisHost  string   `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort  int      `env:"REDIS_PORT" envDefault:"36379"`
	Cors       []string `env:"Cors" envDefault:"http://localhost:3000" envSeparator:","`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
