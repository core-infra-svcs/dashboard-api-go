# InlineResponse200261SplitTunnelRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol for this split tunnel rule. | [optional] 
**DestCidr** | Pointer to **string** | Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or &#39;any&#39;. | [optional] 
**DestPort** | Pointer to **string** | Destination port for this split tunnel rule, (integer in the range 1-65535), or &#39;any&#39;. | [optional] 
**Policy** | Pointer to **string** | Traffic policy specified for this split tunnel rule, &#39;allow&#39; or &#39;deny&#39;. | [optional] 
**Comment** | Pointer to **string** | Description for this split tunnel rule (optional). | [optional] 

## Methods

### NewInlineResponse200261SplitTunnelRules

`func NewInlineResponse200261SplitTunnelRules() *InlineResponse200261SplitTunnelRules`

NewInlineResponse200261SplitTunnelRules instantiates a new InlineResponse200261SplitTunnelRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200261SplitTunnelRulesWithDefaults

`func NewInlineResponse200261SplitTunnelRulesWithDefaults() *InlineResponse200261SplitTunnelRules`

NewInlineResponse200261SplitTunnelRulesWithDefaults instantiates a new InlineResponse200261SplitTunnelRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse200261SplitTunnelRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200261SplitTunnelRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200261SplitTunnelRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200261SplitTunnelRules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestCidr

`func (o *InlineResponse200261SplitTunnelRules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *InlineResponse200261SplitTunnelRules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *InlineResponse200261SplitTunnelRules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.

### HasDestCidr

`func (o *InlineResponse200261SplitTunnelRules) HasDestCidr() bool`

HasDestCidr returns a boolean if a field has been set.

### GetDestPort

`func (o *InlineResponse200261SplitTunnelRules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *InlineResponse200261SplitTunnelRules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *InlineResponse200261SplitTunnelRules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *InlineResponse200261SplitTunnelRules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200261SplitTunnelRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200261SplitTunnelRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200261SplitTunnelRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200261SplitTunnelRules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetComment

`func (o *InlineResponse200261SplitTunnelRules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse200261SplitTunnelRules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse200261SplitTunnelRules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse200261SplitTunnelRules) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


