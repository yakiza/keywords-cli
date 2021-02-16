package builder

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

type BuilderOptions struct {
	service BuilderService
	log zap.Logger
}

func (b BuilderOptions) CreateDirectory(name string) error {
	cwd, err := os.Getwd()
	b.log.Info("Creating directory...")
	if err != nil {
		return err
	}

	fmt.Println(cwd)

	os.Mkdir(name, os.ModeDir)
	return  nil
}