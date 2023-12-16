package dynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type LogRequest struct {
	Tenantid  string
	Sessionid string
}

func (r *Repository) GetLogItem(req LogRequest) ([]LogItem, error) { //TODO: Check if req.filter_options.field valid
	partitionKeyValue := formatKey(prictlactions_PK_prefix, "<tenantid>", req.Tenantid)
	sortKeyValue := formatKey(prictlactions_SK_prefix, "<tenantid>", req.Tenantid, "<sessionid>", req.Sessionid)

	queryRes, err := r.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName: aws.String(prictllog),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :partitionKeyValue and %s = :sortKeyvalue",
			prictllog_partitionkeyname, prictllog_sortkeyname)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":partitionKeyValue": &types.AttributeValueMemberS{Value: partitionKeyValue},
			":sortKeyvalue":      &types.AttributeValueMemberS{Value: sortKeyValue},
		},
	})
	if err != nil {
		return nil, errors.WithMessage(err, "GetLogItem:Query")
	}
	var a []LogItem
	err = attributevalue.UnmarshalListOfMaps(queryRes.Items, &a)
	if err != nil {
		fmt.Printf("UnmarshalListOfMaps: %s", err)
		return nil, errors.WithMessage(err, "GetLogItem:UnmarshalMap")
	}
	return a, nil
}

// Add new action item.
func (r *Repository) AddLogItem() error {

	newLogEntry := map[string]types.AttributeValue{
		"message": &types.AttributeValueMemberS{Value: "New log entry 3"},
		"time":    &types.AttributeValueMemberN{Value: fmt.Sprint(time.Now().Unix())},
		"type":    &types.AttributeValueMemberS{Value: "info"},
	}
	updateExpression := "SET #logdata = list_append(if_not_exists(#logdata, :empty_list), :new_log_entry)"
	expressionAttributeValues := map[string]types.AttributeValue{
		":new_log_entry": &types.AttributeValueMemberL{Value: []types.AttributeValue{&types.AttributeValueMemberM{Value: newLogEntry}}},
		":empty_list":    &types.AttributeValueMemberL{Value: []types.AttributeValue{}},
	}
	expressionAttributeNames := map[string]string{"#logdata": "logdata"}

	entityidval := "tenantid#5901153D5DAD4DEB84F6E6D72FCA42B1"
	relationidval := "PRICTLACTIONS:tenantid#5901153D5DAD4DEB84F6E6D72FCA42B1#sessionid#999a261e-2513-4475-85f5-4bd78dcfb5ef"
	_, err := r.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("prictllog"),
		Key:                       map[string]types.AttributeValue{"EntityID": &types.AttributeValueMemberS{Value: entityidval}, "RelationID": &types.AttributeValueMemberS{Value: relationidval}},
		UpdateExpression:          &updateExpression,
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributeNames,
		ReturnValues:              types.ReturnValueAllNew,
	})

	if err != nil {
		return errors.WithMessage(err, "AddActionItem: PutItem")
	}
	return err
}
