package cloudformation

// AWSAppSyncDataSource_HttpConfig AWS CloudFormation Resource (AWS::AppSync::DataSource.HttpConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html
type AWSAppSyncDataSource_HttpConfig struct {

	// Endpoint AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html#cfn-appsync-datasource-httpconfig-endpoint
	Endpoint string `json:"Endpoint,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncDataSource_HttpConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::DataSource.HttpConfig"
}
