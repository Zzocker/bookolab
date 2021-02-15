package core

// export provide concret value of all the cores to handlers

// these core will be built using Build function
var (
	uCore UserCore
	bCore BookCore
)

// GetUserCore will return user core to handlers
func GetUserCore() UserCore {
	return uCore
}

func GetBookCore() BookCore {
	return bCore
}