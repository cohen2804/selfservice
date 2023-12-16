package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	//prictlactions table
	prictllog = "prictllog"
	//PK Prefix
	prictllog_PK_prefix = "tenantid#<tenantid>"
	// SK Prefix
	prictllog_SK_prefix = "PRICTLACTIONS:tenantid#<tenantid>#sessionid#<sessionid>"

	// keys
	prictllog_partitionkeyname = "EntityID"
	prictllog_sortkeyname      = "RelationID"
)

type LogEntry struct {
	Message string `dynamodbav:"message"`
	Time    int64  `dynamodbav:"time"`
	Type    string `dynamodbav:"type"`
}

type LogItem struct {
	EntityID   string `dynamodbav:"EntityID"`
	RelationID string `dynamodbav:"RelationID"`
	//LogData    []map[string]LogEntry `dynamodbav:"logdata"`
	Tenantid  string     `dynamodbav:"tenantid"`
	Sessionid string     `dynamodbav:"sessionid"`
	LogData   []LogEntry `dynamodbav:"logdata"`
}

func (a LogItem) GetlogKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		PK_EntityID:   &types.AttributeValueMemberS{Value: formatKey(prictlactions_PK_prefix, "<tenantid>", fmt.Sprint(a.Tenantid))},
		SK_RelationID: &types.AttributeValueMemberS{Value: formatKey(prictlactions_SK_prefix, "<tenantid>", a.Tenantid, "<sessionid>", fmt.Sprint(a.Sessionid))},
	}
}
