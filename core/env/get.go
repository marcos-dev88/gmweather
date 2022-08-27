package env

import (
	"os"
	"strconv"
)

type EnvTypes interface {
	String() string
	Int() int
}

type getEnv struct {
	Name string
}

func Get(EnvName string) EnvTypes {
	return getEnv{Name: EnvName}
}

func getEnvStr(envName string) string {
	return os.Getenv(envName)
}

func (g getEnv) String() string {
	return getEnvStr(g.Name)
}

func (g getEnv) Int() int {
	e := getEnvStr(g.Name)
	v, err := strconv.Atoi(e)
	if err != nil {
		panic(err)
	}
	return v
}
