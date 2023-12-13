package dynamodb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

const (
	PK_EntityID   = "EntityID"
	SK_RelationID = "RelationID"
)

type Repository struct {
	client *dynamodb.Client
}

func NewRepository() (*Repository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, errors.WithMessage(err, "NewRepository: LoadDefaultConfig")
	}
	dynamoClient := dynamodb.NewFromConfig(cfg)

	return &Repository{
		client: dynamoClient,
	}, nil
}

func formatKey(key string, patternValues ...string) string {
	if len(patternValues)%2 == 1 || len(patternValues) == 0 {
		return ""
	}
	replacer := strings.NewReplacer(patternValues...)

	return replacer.Replace(key)
}

func Serialize(input map[string]types.AttributeValue) (string, error) {
	var inputMap map[string]interface{}
	err := attributevalue.UnmarshalMap(input, &inputMap)
	if err != nil {
		return "", err
	}
	bytesJSON, err := json.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	output := base64.StdEncoding.EncodeToString(bytesJSON)
	return output, nil
}

func Deserialize(input string) (map[string]types.AttributeValue, error) {
	bytesJSON, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	outputJSON := map[string]interface{}{}
	err = json.Unmarshal(bytesJSON, &outputJSON)
	if err != nil {
		return nil, err
	}
	return attributevalue.MarshalMap(outputJSON)
}

func PageSize(pageSize int32) int32 {
	if pageSize < 1 {
		return 50
	} else {
		return pageSize
	}

}
