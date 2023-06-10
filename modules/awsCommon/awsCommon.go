package awsCommon

import (
	"context"

	"github.com/PARKNAMSU/nspark-go-common-modules/modules/common"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/smithy-go/logging"
)

type AwsConfigWithOpts struct {
	Region  *string
	Profile *string
	Logger  *logging.Logger
}

func InitConfig(cfg *aws.Config, opt AwsConfigWithOpts) error {
	var funcs []func(*config.LoadOptions) error = []func(*config.LoadOptions) error{}

	if opt.Region != nil {
		funcs = append(funcs, config.WithRegion(*opt.Region))
	}
	if opt.Profile != nil {
		funcs = append(funcs, config.WithSharedConfigProfile(*opt.Profile))
	}
	if opt.Logger != nil {
		funcs = append(funcs, config.WithLogger(*opt.Logger))
	}

	getCfg, err := config.LoadDefaultConfig(context.TODO(), funcs...)

	if common.CheckError(err) {
		return err
	}

	cfg = &getCfg

	return nil
}
