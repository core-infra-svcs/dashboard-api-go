# InlineResponse200190Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | 
**DestPort** | Pointer to **string** | Comma-separated list of destination port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**DestCidr** | **string** | Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or &#39;any&#39; | 

## Methods

### NewInlineResponse200190Rules

`func NewInlineResponse200190Rules(policy string, protocol string, destCidr string, ) *InlineResponse200190Rules`

NewInlineResponse200190Rules instantiates a new InlineResponse200190Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200190RulesWithDefaults

`func NewInlineResponse200190RulesWithDefaults() *InlineResponse200190Rules`

NewInlineResponse200190RulesWithDefaults instantiates a new InlineResponse200190Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineResponse200190Rules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse200190Rules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse200190Rules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse200190Rules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200190Rules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200190Rules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200190Rules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *InlineResponse200190Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200190Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200190Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestPort

`func (o *InlineResponse200190Rules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *InlineResponse200190Rules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *InlineResponse200190Rules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *InlineResponse200190Rules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *InlineResponse200190Rules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *InlineResponse200190Rules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *InlineResponse200190Rules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


