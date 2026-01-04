package handler

type CommandArguments struct {
	Path string
}

type CommandFlags struct {
	HumanizeSize    bool
	ShowHiddenFiles bool
}

type PathSizeHandler struct {
	Arguments CommandArguments
	Flags     CommandFlags
}
