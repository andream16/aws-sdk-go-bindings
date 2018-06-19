// Package configuration allows reading configuration parameters from a given configuration file
package configuration

import (
	"runtime"
	"path"
	"path/filepath"
	"github.com/tkanos/gonfig"
)

const confFileName = "configuration.json"

type Configuration struct {
	Region string `json:"region"`
	TargetArn string `json:"target_arn"`
}

// Get returns Configuration leaded from configuration file
func Get() (conf Configuration, err error) {

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), confFileName)

	err = gonfig.GetConf(filePath, &conf)

	return

}
