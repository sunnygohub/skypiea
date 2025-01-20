package skypiea

import "github.com/joho/godotenv"

// const version = "1.0.0"

func (s *Skypiea) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := s.Init(pathConfig)
	if err != nil {
		return err
	}

	err = s.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (s *Skypiea) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := s.createDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}
