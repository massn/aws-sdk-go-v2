// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package health

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Information about an entity that is affected by a Health event.
type AffectedEntity struct {
	_ struct{} `type:"structure"`

	// The 12-digit AWS account number that contains the affected entity.
	AwsAccountId *string `locationName:"awsAccountId" type:"string"`

	// The unique identifier for the entity. Format: arn:aws:health:entity-region:aws-account:entity/entity-id
	// . Example: arn:aws:health:us-east-1:111222333444:entity/AVh5GGT7ul1arKr1sE1K
	EntityArn *string `locationName:"entityArn" type:"string"`

	// The URL of the affected entity.
	EntityUrl *string `locationName:"entityUrl" type:"string"`

	// The ID of the affected entity.
	EntityValue *string `locationName:"entityValue" type:"string"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string `locationName:"eventArn" type:"string"`

	// The most recent time that the entity was updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The most recent status of the entity affected by the event. The possible
	// values are IMPAIRED, UNIMPAIRED, and UNKNOWN.
	StatusCode EntityStatusCode `locationName:"statusCode" type:"string" enum:"true"`

	// A map of entity tags attached to the affected entity.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s AffectedEntity) String() string {
	return awsutil.Prettify(s)
}

// A range of dates and times that is used by the EventFilter and EntityFilter
// objects. If from is set and to is set: match items where the timestamp (startTime,
// endTime, or lastUpdatedTime) is between from and to inclusive. If from is
// set and to is not set: match items where the timestamp value is equal to
// or after from. If from is not set and to is set: match items where the timestamp
// value is equal to or before to.
type DateTimeRange struct {
	_ struct{} `type:"structure"`

	// The starting date and time of a time range.
	From *time.Time `locationName:"from" type:"timestamp"`

	// The ending date and time of a time range.
	To *time.Time `locationName:"to" type:"timestamp"`
}

// String returns the string representation
func (s DateTimeRange) String() string {
	return awsutil.Prettify(s)
}

// The number of entities that are affected by one or more events. Returned
// by the DescribeEntityAggregates operation.
type EntityAggregate struct {
	_ struct{} `type:"structure"`

	// The number entities that match the criteria for the specified events.
	Count *int64 `locationName:"count" type:"integer"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string `locationName:"eventArn" type:"string"`
}

// String returns the string representation
func (s EntityAggregate) String() string {
	return awsutil.Prettify(s)
}

// The values to use to filter results from the DescribeAffectedEntities operation.
type EntityFilter struct {
	_ struct{} `type:"structure"`

	// A list of entity ARNs (unique identifiers).
	EntityArns []string `locationName:"entityArns" min:"1" type:"list"`

	// A list of IDs for affected entities.
	EntityValues []string `locationName:"entityValues" min:"1" type:"list"`

	// A list of event ARNs (unique identifiers). For example: "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	//
	// EventArns is a required field
	EventArns []string `locationName:"eventArns" min:"1" type:"list" required:"true"`

	// A list of the most recent dates and times that the entity was updated.
	LastUpdatedTimes []DateTimeRange `locationName:"lastUpdatedTimes" min:"1" type:"list"`

	// A list of entity status codes (IMPAIRED, UNIMPAIRED, or UNKNOWN).
	StatusCodes []EntityStatusCode `locationName:"statusCodes" min:"1" type:"list"`

	// A map of entity tags attached to the affected entity.
	Tags []map[string]string `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s EntityFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EntityFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EntityFilter"}
	if s.EntityArns != nil && len(s.EntityArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityArns", 1))
	}
	if s.EntityValues != nil && len(s.EntityValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityValues", 1))
	}

	if s.EventArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventArns"))
	}
	if s.EventArns != nil && len(s.EventArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventArns", 1))
	}
	if s.LastUpdatedTimes != nil && len(s.LastUpdatedTimes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LastUpdatedTimes", 1))
	}
	if s.StatusCodes != nil && len(s.StatusCodes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatusCodes", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Summary information about an AWS Health event.
type Event struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	Arn *string `locationName:"arn" type:"string"`

	// The AWS Availability Zone of the event. For example, us-east-1a.
	AvailabilityZone *string `locationName:"availabilityZone" min:"6" type:"string"`

	// The date and time that the event ended.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	EventScopeCode EventScopeCode `locationName:"eventScopeCode" type:"string" enum:"true"`

	// The category of the event. Possible values are issue, scheduledChange, and
	// accountNotification.
	EventTypeCategory EventTypeCategory `locationName:"eventTypeCategory" min:"3" type:"string" enum:"true"`

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION
	// ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	EventTypeCode *string `locationName:"eventTypeCode" min:"3" type:"string"`

	// The most recent date and time that the event was updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The AWS region name of the event.
	Region *string `locationName:"region" min:"2" type:"string"`

	// The AWS service that is affected by the event. For example, EC2, RDS.
	Service *string `locationName:"service" min:"2" type:"string"`

	// The date and time that the event began.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`

	// The most recent status of the event. Possible values are open, closed, and
	// upcoming.
	StatusCode EventStatusCode `locationName:"statusCode" type:"string" enum:"true"`
}

// String returns the string representation
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// The values used to filter results from the DescribeEventDetailsForOrganization
// and DescribeAffectedEntitiesForOrganization operations.
type EventAccountFilter struct {
	_ struct{} `type:"structure"`

	// The 12-digit AWS account numbers that contains the affected entities.
	AwsAccountId *string `locationName:"awsAccountId" type:"string"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	//
	// EventArn is a required field
	EventArn *string `locationName:"eventArn" type:"string" required:"true"`
}

// String returns the string representation
func (s EventAccountFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EventAccountFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EventAccountFilter"}

	if s.EventArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The number of events of each issue type. Returned by the DescribeEventAggregates
// operation.
type EventAggregate struct {
	_ struct{} `type:"structure"`

	// The issue type for the associated count.
	AggregateValue *string `locationName:"aggregateValue" type:"string"`

	// The number of events of the associated issue type.
	Count *int64 `locationName:"count" type:"integer"`
}

// String returns the string representation
func (s EventAggregate) String() string {
	return awsutil.Prettify(s)
}

// The detailed description of the event. Included in the information returned
// by the DescribeEventDetails operation.
type EventDescription struct {
	_ struct{} `type:"structure"`

	// The most recent description of the event.
	LatestDescription *string `locationName:"latestDescription" type:"string"`
}

// String returns the string representation
func (s EventDescription) String() string {
	return awsutil.Prettify(s)
}

// Detailed information about an event. A combination of an Event object, an
// EventDescription object, and additional metadata about the event. Returned
// by the DescribeEventDetails operation.
type EventDetails struct {
	_ struct{} `type:"structure"`

	// Summary information about the event.
	Event *Event `locationName:"event" type:"structure"`

	// The most recent description of the event.
	EventDescription *EventDescription `locationName:"eventDescription" type:"structure"`

	// Additional metadata about the event.
	EventMetadata map[string]string `locationName:"eventMetadata" type:"map"`
}

// String returns the string representation
func (s EventDetails) String() string {
	return awsutil.Prettify(s)
}

// Error information returned when a DescribeEventDetails operation cannot find
// a specified event.
type EventDetailsErrorItem struct {
	_ struct{} `type:"structure"`

	// A message that describes the error.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`

	// The name of the error.
	ErrorName *string `locationName:"errorName" type:"string"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string `locationName:"eventArn" type:"string"`
}

// String returns the string representation
func (s EventDetailsErrorItem) String() string {
	return awsutil.Prettify(s)
}

// The values to use to filter results from the DescribeEvents and DescribeEventAggregates
// operations.
type EventFilter struct {
	_ struct{} `type:"structure"`

	// A list of AWS availability zones.
	AvailabilityZones []string `locationName:"availabilityZones" type:"list"`

	// A list of dates and times that the event ended.
	EndTimes []DateTimeRange `locationName:"endTimes" min:"1" type:"list"`

	// A list of entity ARNs (unique identifiers).
	EntityArns []string `locationName:"entityArns" min:"1" type:"list"`

	// A list of entity identifiers, such as EC2 instance IDs (i-34ab692e) or EBS
	// volumes (vol-426ab23e).
	EntityValues []string `locationName:"entityValues" min:"1" type:"list"`

	// A list of event ARNs (unique identifiers). For example: "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	EventArns []string `locationName:"eventArns" min:"1" type:"list"`

	// A list of event status codes.
	EventStatusCodes []EventStatusCode `locationName:"eventStatusCodes" min:"1" type:"list"`

	// A list of event type category codes (issue, scheduledChange, or accountNotification).
	EventTypeCategories []EventTypeCategory `locationName:"eventTypeCategories" min:"1" type:"list"`

	// A list of unique identifiers for event types. For example, "AWS_EC2_SYSTEM_MAINTENANCE_EVENT","AWS_RDS_MAINTENANCE_SCHEDULED".
	EventTypeCodes []string `locationName:"eventTypeCodes" min:"1" type:"list"`

	// A list of dates and times that the event was last updated.
	LastUpdatedTimes []DateTimeRange `locationName:"lastUpdatedTimes" min:"1" type:"list"`

	// A list of AWS regions.
	Regions []string `locationName:"regions" min:"1" type:"list"`

	// The AWS services associated with the event. For example, EC2, RDS.
	Services []string `locationName:"services" min:"1" type:"list"`

	// A list of dates and times that the event began.
	StartTimes []DateTimeRange `locationName:"startTimes" min:"1" type:"list"`

	// A map of entity tags attached to the affected entity.
	Tags []map[string]string `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s EventFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EventFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EventFilter"}
	if s.EndTimes != nil && len(s.EndTimes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EndTimes", 1))
	}
	if s.EntityArns != nil && len(s.EntityArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityArns", 1))
	}
	if s.EntityValues != nil && len(s.EntityValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityValues", 1))
	}
	if s.EventArns != nil && len(s.EventArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventArns", 1))
	}
	if s.EventStatusCodes != nil && len(s.EventStatusCodes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventStatusCodes", 1))
	}
	if s.EventTypeCategories != nil && len(s.EventTypeCategories) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventTypeCategories", 1))
	}
	if s.EventTypeCodes != nil && len(s.EventTypeCodes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventTypeCodes", 1))
	}
	if s.LastUpdatedTimes != nil && len(s.LastUpdatedTimes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LastUpdatedTimes", 1))
	}
	if s.Regions != nil && len(s.Regions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Regions", 1))
	}
	if s.Services != nil && len(s.Services) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Services", 1))
	}
	if s.StartTimes != nil && len(s.StartTimes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StartTimes", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Metadata about a type of event that is reported by AWS Health. Data consists
// of the category (for example, issue), the service (for example, EC2), and
// the event type code (for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT).
type EventType struct {
	_ struct{} `type:"structure"`

	// A list of event type category codes (issue, scheduledChange, or accountNotification).
	Category EventTypeCategory `locationName:"category" min:"3" type:"string" enum:"true"`

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION
	// ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	Code *string `locationName:"code" min:"3" type:"string"`

	// The AWS service that is affected by the event. For example, EC2, RDS.
	Service *string `locationName:"service" min:"2" type:"string"`
}

// String returns the string representation
func (s EventType) String() string {
	return awsutil.Prettify(s)
}

// The values to use to filter results from the DescribeEventTypes operation.
type EventTypeFilter struct {
	_ struct{} `type:"structure"`

	// A list of event type category codes (issue, scheduledChange, or accountNotification).
	EventTypeCategories []EventTypeCategory `locationName:"eventTypeCategories" min:"1" type:"list"`

	// A list of event type codes.
	EventTypeCodes []string `locationName:"eventTypeCodes" min:"1" type:"list"`

	// The AWS services associated with the event. For example, EC2, RDS.
	Services []string `locationName:"services" min:"1" type:"list"`
}

// String returns the string representation
func (s EventTypeFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EventTypeFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EventTypeFilter"}
	if s.EventTypeCategories != nil && len(s.EventTypeCategories) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventTypeCategories", 1))
	}
	if s.EventTypeCodes != nil && len(s.EventTypeCodes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventTypeCodes", 1))
	}
	if s.Services != nil && len(s.Services) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Services", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Error information returned when a DescribeAffectedEntitiesForOrganization
// operation cannot find or process a specific entity.
type OrganizationAffectedEntitiesErrorItem struct {
	_ struct{} `type:"structure"`

	// The 12-digit AWS account numbers that contains the affected entities.
	AwsAccountId *string `locationName:"awsAccountId" type:"string"`

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION.
	// For example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`

	// The name of the error.
	ErrorName *string `locationName:"errorName" type:"string"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string `locationName:"eventArn" type:"string"`
}

// String returns the string representation
func (s OrganizationAffectedEntitiesErrorItem) String() string {
	return awsutil.Prettify(s)
}

// Summary information about an event, returned by the DescribeEventsForOrganization
// operation.
type OrganizationEvent struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	Arn *string `locationName:"arn" type:"string"`

	// The date and time that the event ended.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	EventScopeCode EventScopeCode `locationName:"eventScopeCode" type:"string" enum:"true"`

	// The category of the event type.
	EventTypeCategory EventTypeCategory `locationName:"eventTypeCategory" min:"3" type:"string" enum:"true"`

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION.
	// For example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	EventTypeCode *string `locationName:"eventTypeCode" min:"3" type:"string"`

	// The most recent date and time that the event was updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The AWS Region name of the event.
	Region *string `locationName:"region" min:"2" type:"string"`

	// The AWS service that is affected by the event. For example, EC2, RDS.
	Service *string `locationName:"service" min:"2" type:"string"`

	// The date and time that the event began.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`

	// The most recent status of the event. Possible values are open, closed, and
	// upcoming.
	StatusCode EventStatusCode `locationName:"statusCode" type:"string" enum:"true"`
}

// String returns the string representation
func (s OrganizationEvent) String() string {
	return awsutil.Prettify(s)
}

// Detailed information about an event. A combination of an Event object, an
// EventDescription object, and additional metadata about the event. Returned
// by the DescribeEventDetailsForOrganization operation.
type OrganizationEventDetails struct {
	_ struct{} `type:"structure"`

	// The 12-digit AWS account numbers that contains the affected entities.
	AwsAccountId *string `locationName:"awsAccountId" type:"string"`

	// Summary information about an AWS Health event.
	Event *Event `locationName:"event" type:"structure"`

	// The detailed description of the event. Included in the information returned
	// by the DescribeEventDetails operation.
	EventDescription *EventDescription `locationName:"eventDescription" type:"structure"`

	// Additional metadata about the event.
	EventMetadata map[string]string `locationName:"eventMetadata" type:"map"`
}

// String returns the string representation
func (s OrganizationEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Error information returned when a DescribeEventDetailsForOrganization operation
// cannot find a specified event.
type OrganizationEventDetailsErrorItem struct {
	_ struct{} `type:"structure"`

	// Error information returned when a DescribeEventDetailsForOrganization operation
	// cannot find a specified event.
	AwsAccountId *string `locationName:"awsAccountId" type:"string"`

	// A message that describes the error.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`

	// The name of the error.
	ErrorName *string `locationName:"errorName" type:"string"`

	// The unique identifier for the event. Format: arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// . Example: Example: arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string `locationName:"eventArn" type:"string"`
}

// String returns the string representation
func (s OrganizationEventDetailsErrorItem) String() string {
	return awsutil.Prettify(s)
}

// The values to filter results from the DescribeEventsForOrganization operation.
type OrganizationEventFilter struct {
	_ struct{} `type:"structure"`

	// A list of 12-digit AWS account numbers that contains the affected entities.
	AwsAccountIds []string `locationName:"awsAccountIds" min:"1" type:"list"`

	// A range of dates and times that is used by the EventFilter and EntityFilter
	// objects. If from is set and to is set: match items where the timestamp (startTime,
	// endTime, or lastUpdatedTime) is between from and to inclusive. If from is
	// set and to is not set: match items where the timestamp value is equal to
	// or after from. If from is not set and to is set: match items where the timestamp
	// value is equal to or before to.
	EndTime *DateTimeRange `locationName:"endTime" type:"structure"`

	// REPLACEME
	EntityArns []string `locationName:"entityArns" min:"1" type:"list"`

	// A list of entity identifiers, such as EC2 instance IDs (i-34ab692e) or EBS
	// volumes (vol-426ab23e).
	EntityValues []string `locationName:"entityValues" min:"1" type:"list"`

	// A list of event status codes.
	EventStatusCodes []EventStatusCode `locationName:"eventStatusCodes" min:"1" type:"list"`

	// REPLACEME
	EventTypeCategories []EventTypeCategory `locationName:"eventTypeCategories" min:"1" type:"list"`

	// A list of unique identifiers for event types. For example, "AWS_EC2_SYSTEM_MAINTENANCE_EVENT","AWS_RDS_MAINTENANCE_SCHEDULED".
	EventTypeCodes []string `locationName:"eventTypeCodes" min:"1" type:"list"`

	// A range of dates and times that is used by the EventFilter and EntityFilter
	// objects. If from is set and to is set: match items where the timestamp (startTime,
	// endTime, or lastUpdatedTime) is between from and to inclusive. If from is
	// set and to is not set: match items where the timestamp value is equal to
	// or after from. If from is not set and to is set: match items where the timestamp
	// value is equal to or before to.
	LastUpdatedTime *DateTimeRange `locationName:"lastUpdatedTime" type:"structure"`

	// A list of AWS Regions.
	Regions []string `locationName:"regions" min:"1" type:"list"`

	// The AWS services associated with the event. For example, EC2, RDS.
	Services []string `locationName:"services" min:"1" type:"list"`

	// A range of dates and times that is used by the EventFilter and EntityFilter
	// objects. If from is set and to is set: match items where the timestamp (startTime,
	// endTime, or lastUpdatedTime) is between from and to inclusive. If from is
	// set and to is not set: match items where the timestamp value is equal to
	// or after from. If from is not set and to is set: match items where the timestamp
	// value is equal to or before to.
	StartTime *DateTimeRange `locationName:"startTime" type:"structure"`
}

// String returns the string representation
func (s OrganizationEventFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OrganizationEventFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OrganizationEventFilter"}
	if s.AwsAccountIds != nil && len(s.AwsAccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountIds", 1))
	}
	if s.EntityArns != nil && len(s.EntityArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityArns", 1))
	}
	if s.EntityValues != nil && len(s.EntityValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityValues", 1))
	}
	if s.EventStatusCodes != nil && len(s.EventStatusCodes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventStatusCodes", 1))
	}
	if s.EventTypeCategories != nil && len(s.EventTypeCategories) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventTypeCategories", 1))
	}
	if s.EventTypeCodes != nil && len(s.EventTypeCodes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventTypeCodes", 1))
	}
	if s.Regions != nil && len(s.Regions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Regions", 1))
	}
	if s.Services != nil && len(s.Services) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Services", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
