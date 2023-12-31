package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type LocalGitUser struct {
	UserEmail string `json:"user_email"`
	UserName  string `json:"user_name"`
	Tag       string `json:"tag"`
}

func (localGitUser *LocalGitUser) GetSummary() string {
	return fmt.Sprintf("%s / %s (%s)", localGitUser.UserEmail, localGitUser.UserName, localGitUser.Tag)
}

func LoadLocalGitData() *[]LocalGitUser {
	configFilePath := GetLocalConfigPath()

	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileBytes := make([]byte, fileInfo.Size())
	bytesRead, err := file.Read(fileBytes)
	if err != nil {
		log.Fatal(err)
	}

	var configData []LocalGitUser

	err = json.Unmarshal(fileBytes[:bytesRead], &configData)
	if err != nil {
		log.Fatal(err)
	}

	return &configData
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetLocalConfigPath() string {
	const DEFAULT_CONFIG_FILE_NAME = "local-git-help-config.json"

	currentOsUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(currentOsUser.HomeDir, DEFAULT_CONFIG_FILE_NAME)
}
