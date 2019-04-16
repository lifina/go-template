package config

type DBConfig struct {
	Host     string `validate:"ip,omitempty"`
	Port     int64  `validate:"min=0,max=65535,omitempty"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}
