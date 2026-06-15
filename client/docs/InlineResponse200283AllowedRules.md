# InlineResponse200283AllowedRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** | A rule identifier of the format meraki:intrusion/snort/GID/&lt;gid&gt;/SID/&lt;sid&gt;. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents | [optional] 
**Message** | Pointer to **NullableString** | Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET | [optional] 

## Methods

### NewInlineResponse200283AllowedRules

`func NewInlineResponse200283AllowedRules() *InlineResponse200283AllowedRules`

NewInlineResponse200283AllowedRules instantiates a new InlineResponse200283AllowedRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200283AllowedRulesWithDefaults

`func NewInlineResponse200283AllowedRulesWithDefaults() *InlineResponse200283AllowedRules`

NewInlineResponse200283AllowedRulesWithDefaults instantiates a new InlineResponse200283AllowedRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *InlineResponse200283AllowedRules) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse200283AllowedRules) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse200283AllowedRules) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse200283AllowedRules) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetMessage

`func (o *InlineResponse200283AllowedRules) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse200283AllowedRules) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse200283AllowedRules) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse200283AllowedRules) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InlineResponse200283AllowedRules) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InlineResponse200283AllowedRules) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


