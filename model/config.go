package model

var AwsConfig AwsConfiguration

type AwsConfiguration struct {
	Bucket          string
	Region          string
	AccessKey       string
	SecretAccessKey string
}
