package infrastructures

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env has environment stored
type Env struct {
	ServerPort    string
	BrokerAddress string
}

// NewEnv creates a new environment
func NewEnv() *Env {
	env := Env{}
	return &env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env.ServerPort = os.Getenv("ServerPort")
	env.BrokerAddress = os.Getenv("BrokerAddress")
}
