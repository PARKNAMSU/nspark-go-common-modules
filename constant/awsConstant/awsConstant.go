package awsConstant

type AwsRegions struct {
	NorthVirginia string
	Ohaio         string
	California    string
	Oregon        string
	Mumbai        string
	Osaka         string
	Seoul         string
	Singapore     string
	Sydney        string
	Tokyo         string
	CentalCanada  string
}

var (
	Regions AwsRegions = AwsRegions{
		NorthVirginia: "us-east-1",
		Ohaio:         "us-east-2",
		California:    "us-west-1",
		Oregon:        "us-west-2",
		Mumbai:        "ap-south-1",
		Osaka:         "ap-northeast-3",
		Seoul:         "ap-northeast-2",
		Singapore:     "ap-southeast-1",
		Sydney:        "ap-southeast-2",
		Tokyo:         "ap-northeast-1",
		CentalCanada:  "ca-central-1",
	}
)
