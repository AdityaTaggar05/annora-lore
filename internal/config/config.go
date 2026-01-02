package config

import "time"

type Config struct {
	Server ServerConfig
	Neo4j  Neo4jConfig
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getDuration("SERVER_READ_TIMEOUT", 5*time.Second),
			WriteTimeout: getDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
		},

		Neo4j: Neo4jConfig{
			URI:      mustGetEnv("NEO4J_URI"),
			User:     getEnv("NEO4J_USER", "neo4j"),
			Password: getEnv("NEO4J_PASS", "password"),
		},
	}
}
