package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type LogItemRequest struct {
	Tenantid  string
	Sessionid string
}
type AddLogRequest struct {
	Tenantid  string
	Sessionid string
	Message   string
	Time      int64
	Type      string
}

func (r *Repository) GetLogItem(req LogItemRequest) ([]LogItem, error) {
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

// Add new action item log.
func (r *Repository) AddLogItem(req AddLogRequest) error {
	newLogEntry := map[string]types.AttributeValue{
		prictllog_message: &types.AttributeValueMemberS{Value: req.Message},
		prictllog_time:    &types.AttributeValueMemberN{Value: fmt.Sprint(req.Time)},
		prictllog_type:    &types.AttributeValueMemberS{Value: req.Type},
	}
	updateExpression := "SET #tenantid = :tenantid, #sessionid = :sessionid, #logdata = list_append(if_not_exists(#logdata, :empty_list), :new_log_entry)"
	expressionAttributeValues := map[string]types.AttributeValue{
		":tenantid":      &types.AttributeValueMemberS{Value: req.Tenantid},
		":sessionid":     &types.AttributeValueMemberS{Value: req.Sessionid},
		":empty_list":    &types.AttributeValueMemberL{Value: []types.AttributeValue{}},
		":new_log_entry": &types.AttributeValueMemberL{Value: []types.AttributeValue{&types.AttributeValueMemberM{Value: newLogEntry}}},
	}
	expressionAttributeNames := map[string]string{"#logdata": prictllog_logdata, "#tenantid": prictllog_tenantid, "#sessionid": prictllog_session}

	item := &LogItem{Tenantid: req.Tenantid, Sessionid: req.Sessionid}
	_, err := r.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("prictllog"),
		Key:                       item.GetlogKey(),
		UpdateExpression:          &updateExpression,
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributeNames,
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		return errors.WithMessage(err, "AddLogItem: UpdateItem")
	}
	return err
}
