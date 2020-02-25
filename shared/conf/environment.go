package conf

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const (
	AZCK_ZOOKEEPER_SERVERS_ENV = "AZCK_ZOOKEEPER_SERVERS"
)

func LoadEnvFile(filePath string) error {
	return godotenv.Load(filePath)
}

func GetInt64Var(envVar string) (int64, error) {
	return strconv.ParseInt(os.Getenv(envVar), 10, 64)
}

func GetStringVar(envVar string) (string, error) {
	envVarValue := os.Getenv(envVar)
	if envVarValue == "" {
		return envVarValue, errors.Errorf("env. var [%s] is empty", envVar)
	}

	return envVarValue, nil
}
