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
	Region      string      `json:"region"`
	SNS         SNS         `json:"SNS"`
	DynamoDB    DynamoDB    `json:"DynamoDB"`
	S3          S3          `json:"S3"`
	Rekognition Rekognition `json:"Rekognition"`
}

// DynamoDB contains test parameters for DynamoDB
type DynamoDB struct {
	Endpoint     string `json:"endpoint"`
	PkgTableName string `json:"pkg_table_name"`
	CmdTableName string `json:"cmd_table_name"`
	PrimaryKey   string `json:"primary_key"`
}

// SNS contains test parameters for SNS
type SNS struct {
	TargetArn string `json:"target_arn"`
}

// S3 contains test parameters for S3
type S3 struct {
	Bucket      string `json:"bucket"`
	SourceImage string `json:"source_image"`
}

// Rekognition contains test parameters for Rekognition
type Rekognition struct {
	Region       string `json:"region"`
	CompareFaces struct {
		Similarity  float64 `json:"similarity"`
		SourceImage string  `json:"source_image"`
		TargetImage string  `json:"target_image"`
	} `json:"compare_faces"`
	DetectFaces struct {
		SourceImage string `json:"source_image"`
	} `json:"detect_faces"`
	DetectText struct {
		SourceImage string `json:"source_image"`
	} `json:"detect_text"`
}

// Get returns Configuration leaded from configuration file
func Get() (conf Configuration, err error) {

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), confFileName)

	err = gonfig.GetConf(filePath, &conf)

	return

}
