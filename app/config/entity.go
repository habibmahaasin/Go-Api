package config

type Conf struct {
	App struct {
		Port   string `env:"APP_PORT"`
		Secret string `env:"APP_SECRET"`
	}
	Database struct {
		Host string `env:"DATABASE_HOST"`
		Name string `env:"DATABASE_NAME"`
		User string `env:"DATABASE_USER"`
		Pass string `env:"DATABASE_PASSWORD"`
		Port string `env:"DATABASE_PORT"`
	}
}
