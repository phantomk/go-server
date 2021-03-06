package util

import (
	"flag"
	"github.com/axetroy/go-fs"
	"github.com/joho/godotenv"
	"os"
	"path"
)

var (
	Test    bool   // 当前是否是测试环境
	Env     string // 当前的运行环境
	RootDir string // 当前运行的二进制所在的目录
)

func LoadEnv() (err error) {
	var envFilePath = ".env"
	isRunInTest := flag.Lookup("test.v") != nil

	Test = isRunInTest
	Env = os.Getenv("GO_ENV")

	if isRunInTest {
		RootDir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "axetroy", "go-server")
		envFilePath = path.Join(RootDir, envFilePath)
	}

	if !fs.PathExists(envFilePath) {
		return
	}

	err = godotenv.Load(envFilePath)
	return
}
