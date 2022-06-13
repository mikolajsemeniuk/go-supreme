package configuration

type EnvConfiguration struct {
	Listen string
}

func (c *EnvConfiguration) Load() error {
	c.Listen = ":8080"
	return nil
}
