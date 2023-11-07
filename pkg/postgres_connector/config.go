package postgres_connector

type Postgres struct {
	Name           string `yaml:"name"`     //`env-required:"true"  env:"POSTGRES_DB"`
	Password       string `yaml:"password"` //`env-required:"true"  env:"POSTGRES_PASSWORD"`
	MaxConnections int    `yaml:"max_cons"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	Timeout        int    `yaml:"conn_timeout"`
}
