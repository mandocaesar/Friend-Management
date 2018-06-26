package config

//ServerConfiguration : configuration model for server settings
type ServerConfiguration struct {
	Mode            string
	Addr            string
	LogDuration     int
	ShutdownTimeout int
	BaseURL         string
}
