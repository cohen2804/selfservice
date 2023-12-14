package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	//prictlactions table
	prictlactions = "prictlactions"
	//PK Prefix
	prictlactions_PK_prefix = "tenantid#<tenantid>"
	// SK Prefix
	prictlactions_SK_prefix = "PRICTLACTIONS:tenantid#<tenantid>#sessionid#<sessionid>"
	//Attributes Prefixes
	prictlactions_TenantidActionType_prefix = "tenantid#<tenantid>#actiontype#<actiontype>"
	prictlactions_TenantidStatus_prefix     = "tenantid#<tenantid>#status#<status>"
	prictlactions_TenantidUsername_prefix   = "tenantid#<tenantid>#username#<username>"

	// keys
	prictlactions_partitionkeyname = "EntityID"
	prictlactions_sortkeyname      = "RelationID"
	// prictlactions attr
	// prictlactions_tenantid           = "tenantid"
	Prictlactions_actiontype = "actiontype"
	// prictlactions_finishedtime       = "finishedtime"
	// prictlactions_sessionid          = "sessionid"
	// prictlactions_startedtime        = "startedtime"
	// prictlactions_status             = "status"
	// prictlactions_subject            = "subject"
	// prictlactions_username           = "username"
	// prictlactions_tenantidactiontype = "tenantidactiontype"
	// prictlactions_tenantidstatus     = "tenantidstatus"
	// prictlactions_tenantidusername   = "tenantidusername"
)

type ActionItem struct {
	EntityID           string `dynamodbav:"EntityID"`
	RelationID         string `dynamodbav:"RelationID"`
	Tenantid           string `dynamodbav:"tenantid"`
	Actiontype         string `dynamodbav:"actiontype"`
	Finishedtime       int64  `dynamodbav:"finishedtime"`
	Sessionid          string `dynamodbav:"sessionid"`
	StartedTime        int64  `dynamodbav:"startedtime"`
	Status             string `dynamodbav:"status"`
	Subject            string `dynamodbav:"subject"`
	TenantidActionType string `dynamodbav:"tenantidactiontype"`
	TenantidStatus     string `dynamodbav:"tenantidstatus"`
	TenantidUsername   string `dynamodbav:"tenantidusername"`
	Username           string `dynamodbav:"username"`
}

func (a ActionItem) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		PK_EntityID:   &types.AttributeValueMemberS{Value: formatKey(prictlactions_PK_prefix, "<tenantid>", fmt.Sprint(a.Tenantid))},
		SK_RelationID: &types.AttributeValueMemberS{Value: formatKey(prictlactions_SK_prefix, "<tenantid>", a.Tenantid, "<sessionid>", fmt.Sprint(a.Sessionid))},
	}
}
