package utils

import (
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog/config"
)

const (
	// TemplatesRepoURL is the URL for files in nuclei-templates repository
	TemplatesRepoURL = "https://github.com/projectdiscovery/nuclei-templates/blob/main/"
)

var configData *config.Config

func init() {
	configData, _ = config.ReadConfiguration()
}

// TemplatePathURL returns the Path and URL for the provided template
func TemplatePathURL(fullPath string) (string, string) {
	var templateDirectory string
	if configData != nil && configData.TemplatesDirectory != "" && strings.HasPrefix(fullPath, configData.TemplatesDirectory) {
		templateDirectory = configData.TemplatesDirectory
	} else {
		return "", ""
	}

	finalPath := strings.TrimPrefix(strings.TrimPrefix(fullPath, templateDirectory), "/")
	templateURL := TemplatesRepoURL + finalPath
	return finalPath, templateURL
}

// GetDefaultTemplatePath on default settings
func GetDefaultTemplatePath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "nuclei-templates"), nil
}
