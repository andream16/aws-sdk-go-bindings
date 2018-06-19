// Package configuration allows reading configuration parameters from a given configuration file
package configuration

import (
	"github.com/tkanos/gonfig"
	"path"
	"path/filepath"
	"runtime"
)

const confFileName = "configuration.json"

// Configuration contains parameters used in multiple parts of the code base
type Configuration struct {
	Region    string `json:"region"`
	SNS SNS `json:"SNS"`
	DynamoDB DynamoDB `json:"DynamoDB"`
}

// DynamoDB contains test parameters for DynamoDB
type DynamoDB struct {
	Endpoint string `json:"endpoint"`
	TableName string `json:"table_name"`
	PrimaryKey string `json:"primary_key"`
}

// SNS contains test parameters for SNS
type SNS struct {
	TargetArn string `json:"target_arn"`
}

// Get returns Configuration leaded from configuration file
func Get() (conf Configuration, err error) {

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), confFileName)

	err = gonfig.GetConf(filePath, &conf)

	return

}
