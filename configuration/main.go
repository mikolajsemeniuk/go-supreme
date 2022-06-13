package configuration

type Configuration interface {
	Load() error
}
