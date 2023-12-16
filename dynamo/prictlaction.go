package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type QueryRequest struct {
	Tenantid      string
	FilterOptions FilterOption //  List for filtering (optional)
	PageSize      int32        // Number of results per page (optional)
	PageToken     string       // Token to retrieve the next page (optional)
}
type FilterOption struct {
	Field string // Field to filter by (e.g., "action")
	Value string // Value to filter for (e.g., "export")
}

// Get data by tenantid without filters. QueryRequest should contains: tenantid,pagesize and pagetoken.
func (r *Repository) GetActionItemsByTenant(req QueryRequest) (string, []ActionItem, error) {
	EntityID := formatKey(prictlactions_PK_prefix, "<tenantid>", req.Tenantid)
	var pageToken map[string]types.AttributeValue
	var err error
	if len(req.PageToken) > 5 { //Ignore empty object: {}
		pageToken, err = Deserialize(req.PageToken)
		if err != nil {
			return "", nil, errors.WithMessage(err, "GetDataByTenant:Deserialize page token")
		}
	}
	queryRes, err := r.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName: aws.String(prictlactions),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :partitionKeyValue",
			prictlactions_partitionkeyname)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":partitionKeyValue": &types.AttributeValueMemberS{Value: EntityID},
		},
		ScanIndexForward:  aws.Bool(false),
		ExclusiveStartKey: pageToken,
		Limit:             aws.Int32(PageSize(req.PageSize)),
	})
	if err != nil {
		return "", nil, errors.WithMessage(err, "GetDataByTenant:Query")
	}
	var a []ActionItem
	err = attributevalue.UnmarshalListOfMaps(queryRes.Items, &a)
	if err != nil {
		return "", nil, errors.WithMessage(err, "GetDataByTenant:UnmarshalMap")
	}
	newPageToken, err := Serialize(queryRes.LastEvaluatedKey)
	if err != nil {
		return "", nil, errors.WithMessage(err, "GetDataByTenant:Serialize page token")
	}
	return newPageToken, a, nil
}

// Get data by tenantid with additional filters. QueryRequest should contains: tenantid,pagesize and pagetoken.
// filterOptions.field could  be one of: actionType, status or userName
func (r *Repository) GetFilteredActionItems(req QueryRequest) (string, []ActionItem, error) { //TODO: Check if req.filter_options.field valid
	indexName := fmt.Sprintf("tenantid%s-sessionid-index", req.FilterOptions.Field)
	partitionKeyName := fmt.Sprintf("tenantid%s", req.FilterOptions.Field)
	partitionKeyValue := fmt.Sprintf("tenantid#%s#%s#%s", req.Tenantid, req.FilterOptions.Field, req.FilterOptions.Value)
	var pageToken map[string]types.AttributeValue
	var err error
	if len(req.PageToken) > 5 { //Ignore empty object: {}
		pageToken, err = Deserialize(req.PageToken)
		if err != nil {
			return "", nil, errors.WithMessage(err, "GetFilteredData:Deserialize page token")
		}
	}
	queryRes, err := r.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName: aws.String(prictlactions),
		IndexName: aws.String(indexName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :partitionKeyValue",
			partitionKeyName)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":partitionKeyValue": &types.AttributeValueMemberS{Value: partitionKeyValue},
		},
		ScanIndexForward:  aws.Bool(true),
		ExclusiveStartKey: pageToken,
		Limit:             aws.Int32(PageSize(req.PageSize)),
	})
	if err != nil {
		return "", nil, errors.WithMessage(err, "GetFilteredData:Query")
	}
	var a []ActionItem
	err = attributevalue.UnmarshalListOfMaps(queryRes.Items, &a)
	if err != nil {
		return "", nil, errors.WithMessage(err, "GetFilteredData:UnmarshalMap")
	}
	newPageToken, err := Serialize(queryRes.LastEvaluatedKey)
	if err != nil {
		return "", nil, errors.WithMessage(err, "GetFilteredData:Serialize page token")
	}
	return newPageToken, a, nil
}

// Add new action item.
func (r *Repository) AddActionItem(action *ActionItem) error {
	action.EntityID = formatKey(prictlactions_PK_prefix, "<tenantid>", action.Tenantid)
	action.RelationID = formatKey(prictlactions_SK_prefix, "<tenantid>", action.Tenantid, "<sessionid>", action.Sessionid)
	action.TenantidActionType = formatKey(prictlactions_TenantidActionType_prefix, "<tenantid>", action.Tenantid, "<actiontype>", action.Actiontype)
	action.TenantidStatus = formatKey(prictlactions_TenantidStatus_prefix, "<tenantid>", action.Tenantid, "<status>", action.Status)
	action.TenantidUsername = formatKey(prictlactions_TenantidUsername_prefix, "<tenantid>", action.Tenantid, "<username>", action.Username)

	item, err := attributevalue.MarshalMap(action)
	if err != nil {
		return errors.WithMessage(err, "AddActionItem: MarshalMap")
	}
	_, err = r.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(prictlactions), Item: item,
	})
	if err != nil {
		return errors.WithMessage(err, "AddActionItem: PutItem")
	}
	return err
}

// Update action item. tActionItem param must include: tenantid, sessionid, status, finishedtime
func (r *Repository) UpdateActionItem(action ActionItem) error {
	tenantidStatus := formatKey(prictlactions_TenantidStatus_prefix, "<tenantid>", action.Tenantid, "<status>", action.Status)
	update := expression.Set(expression.Name("finishedtime"), expression.Value(action.Finishedtime))
	update.Set(expression.Name("status"), expression.Value(action.Status))
	update.Set(expression.Name("tenantidstatus"), expression.Value(tenantidStatus))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return errors.WithMessage(err, "UpdateActionItem: Build expression for update")
	} else {
		_, err := r.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName:                 aws.String(prictlactions),
			Key:                       action.GetActionKey(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			return errors.WithMessage(err, "UpdateActionItem: UpdateItem")
		}
	}
	return nil
}
