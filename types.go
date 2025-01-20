package skypiea

import "log"

type Skypiea struct {
	AppName  string
	Debug    bool
	Version  string
	RootPath string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

type initPaths struct {
	rootPath    string
	folderNames []string
}
