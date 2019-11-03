package encryptor

import (
	"github.com/huacnlee/gobackup/config"
	"github.com/huacnlee/gobackup/logger"
	"github.com/spf13/viper"
	"os"
)

// Base encryptor
type Base struct {
	model       config.ModelConfig
	viper       *viper.Viper
	archivePath string
}

// Context encryptor interface
type Context interface {
	perform() (encryptPath string, err error)
}

func newBase(archivePath string, model config.ModelConfig) (base Base) {
	base = Base{
		archivePath: archivePath,
		model:       model,
		viper:       model.EncryptWith.Viper,
	}
	return
}

// Run compressor
func Run(archivePath string, model config.ModelConfig) (encryptPath string, err error) {
	base := newBase(archivePath, model)
	var ctx Context
	switch model.EncryptWith.Type {
	case "openssl":
		ctx = &OpenSSL{Base: base}
	default:
		encryptPath = archivePath
		return
	}

	logger.Info("------------ Encryptor -------------")

	logger.Info("=> Encrypt | " + model.EncryptWith.Type)

	encryptPath, err = ctx.perform()
	if err != nil {
		remove(encryptPath)
		return
	}

	remove(archivePath)

	logger.Info("->", encryptPath)
	logger.Info("------------ Encryptor -------------\n")

	return
}

func remove(filePath string)  {
	err := os.Remove(filePath)
	if err != nil {
		logger.Warn("removal of archive file failed:", err)
	}
}