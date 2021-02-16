package config

// ApplicationConf represents all the configuration for this application
type ApplicationConf struct {
	Cores CoreConfigs `yaml:"cores"`
}

// CoreConfigs represents configuration for each core
type CoreConfigs struct {
	User    UserCoreConfig  `yaml:"user"`
	Book    BookCoreConfig  `yaml:"book"`
	Image   ImageCoreConf   `yaml:"image"`
	Comment CommentCoreConf `yaml:"comment"`
	Token   TokenCoreConf   `yaml:"token"`
}

// UserCoreConfig represents configuration for user core
type UserCoreConfig struct {
	UserStore DatastoreConf `yaml:"userstore"`
}

// BookCoreConfig represents configuration for book core
type BookCoreConfig struct {
	BookStore DatastoreConf `yaml:"bookstore"`
}

// ImageCoreConf represents configuration for image core
type ImageCoreConf struct {
	ImageStore DatastoreConf `yaml:"imagestore"`
}

// CommentCoreConf represents configuration for comment core
type CommentCoreConf struct {
	CommentStore DatastoreConf `yaml:"commentstore"`
}

// TokenCoreConf represents configuration for token core
type TokenCoreConf struct {
	TokenStore DatastoreConf `yaml:"tokenstore"`
}

// DatastoreConf represents configuration of a datastore (like mongo/redis)
type DatastoreConf struct {
	URL        string `yaml:"URL"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	Database   string `yaml:"Database"`
	Collection string `yaml:"Collection"`
}
