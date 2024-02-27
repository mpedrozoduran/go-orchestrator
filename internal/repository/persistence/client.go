package persistence

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	appConfig "github.com/mpedrozoduran/go-orchestrator/internal/config"
	"log"
)

type DbClient struct {
	svc *dynamodb.Client
}

func NewDbClient(appConfig appConfig.Config) DbClient {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: appConfig.Database.Url}, nil
			}),
		))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return DbClient{svc: dynamodb.NewFromConfig(cfg)}
}
