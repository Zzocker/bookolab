package config

// ApplicationConf represents all the configuration for this application
type ApplicationConf struct {
	Cores CoreConfigs `yaml:"cores"`
}

// CoreConfigs represents configuration for each core
type CoreConfigs struct {
	User UserCoreConfig `yaml:"user"`
}

// UserCoreConfig represents configuration for user core
type UserCoreConfig struct {
	UserStore DatastoreConf `yaml:"userstore"`
}

// DatastoreConf represents configuration of a datastore (like mongo/redis)
type DatastoreConf struct {
	URL        string `yaml:"URL"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	Database   string `yaml:"Database"`
	Collection string `yaml:"Collection"`
}
