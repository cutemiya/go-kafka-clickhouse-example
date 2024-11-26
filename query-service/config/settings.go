package config

type Settings struct {
	Port       int    `json:"port"`
	Clickhouse string `json:"clickhouse"`
}
