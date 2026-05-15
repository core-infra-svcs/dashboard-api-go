# InlineResponse200246SplitTunnelRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol for this split tunnel rule. | [optional] 
**DestCidr** | Pointer to **string** | Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or &#39;any&#39;. | [optional] 
**DestPort** | Pointer to **string** | Destination port for this split tunnel rule, (integer in the range 1-65535), or &#39;any&#39;. | [optional] 
**Policy** | Pointer to **string** | Traffic policy specified for this split tunnel rule, &#39;allow&#39; or &#39;deny&#39;. | [optional] 
**Comment** | Pointer to **string** | Description for this split tunnel rule (optional). | [optional] 

## Methods

### NewInlineResponse200246SplitTunnelRules

`func NewInlineResponse200246SplitTunnelRules() *InlineResponse200246SplitTunnelRules`

NewInlineResponse200246SplitTunnelRules instantiates a new InlineResponse200246SplitTunnelRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200246SplitTunnelRulesWithDefaults

`func NewInlineResponse200246SplitTunnelRulesWithDefaults() *InlineResponse200246SplitTunnelRules`

NewInlineResponse200246SplitTunnelRulesWithDefaults instantiates a new InlineResponse200246SplitTunnelRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse200246SplitTunnelRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200246SplitTunnelRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200246SplitTunnelRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200246SplitTunnelRules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestCidr

`func (o *InlineResponse200246SplitTunnelRules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *InlineResponse200246SplitTunnelRules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *InlineResponse200246SplitTunnelRules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.

### HasDestCidr

`func (o *InlineResponse200246SplitTunnelRules) HasDestCidr() bool`

HasDestCidr returns a boolean if a field has been set.

### GetDestPort

`func (o *InlineResponse200246SplitTunnelRules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *InlineResponse200246SplitTunnelRules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *InlineResponse200246SplitTunnelRules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *InlineResponse200246SplitTunnelRules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200246SplitTunnelRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200246SplitTunnelRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200246SplitTunnelRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse200246SplitTunnelRules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetComment

`func (o *InlineResponse200246SplitTunnelRules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse200246SplitTunnelRules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse200246SplitTunnelRules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse200246SplitTunnelRules) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


