package config

type Config struct {
	Port            int    `json:"port"`
	CacheExpiration string `json:"cacheExpiration"`
}
