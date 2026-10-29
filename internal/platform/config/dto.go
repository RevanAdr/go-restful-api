package config

import "strconv"

type Env struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

func ConvertPort(p string) int {
	result, err := strconv.Atoi(p)
	if err != nil {
		panic("could not load your port")
	}
	return result
}
