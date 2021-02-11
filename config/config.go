package config

// DatastoreConf represents configuration of a datastore (like mongo/redis)
type DatastoreConf struct {
	URL        string `yaml:"URL"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	Database   string `yaml:"Database"`
	Collection string `yaml:"Collection"`
}
