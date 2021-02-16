package core

// export provide concret value of all the cores to handlers

// these core will be built using Build function
var (
	uCore UserCore
	bCore BookCore
	iCore ImageCore
	cCore CommentCore
	tCore TokenCore
)

// GetUserCore will return user core to handlers
func GetUserCore() UserCore {
	return uCore
}

func GetBookCore() BookCore {
	return bCore
}

func GetImageCore() ImageCore {
	return iCore
}

func GetCommentCore() CommentCore {
	return cCore
}

func GetTokenCore() TokenCore {
	return tCore
}
