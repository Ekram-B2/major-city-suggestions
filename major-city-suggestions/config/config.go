package config

type Config interface {
	LoadConfiguration() (Config, error)
}
