# InlineResponse200211Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | 
**IpVer** | Pointer to **string** | IP version to which this rule applies (must be one of &#39;both&#39;, &#39;ipv4&#39; or &#39;ipv6&#39;). Defaults to &#39;both&#39; if not specified. | [optional] 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | 
**DestPort** | Pointer to **string** | Comma-separated list of destination port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**DestCidr** | **string** | Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or &#39;any&#39; | 

## Methods

### NewInlineResponse200211Rules

`func NewInlineResponse200211Rules(policy string, protocol string, destCidr string, ) *InlineResponse200211Rules`

NewInlineResponse200211Rules instantiates a new InlineResponse200211Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200211RulesWithDefaults

`func NewInlineResponse200211RulesWithDefaults() *InlineResponse200211Rules`

NewInlineResponse200211RulesWithDefaults instantiates a new InlineResponse200211Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineResponse200211Rules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse200211Rules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse200211Rules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse200211Rules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse200211Rules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200211Rules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200211Rules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetIpVer

`func (o *InlineResponse200211Rules) GetIpVer() string`

GetIpVer returns the IpVer field if non-nil, zero value otherwise.

### GetIpVerOk

`func (o *InlineResponse200211Rules) GetIpVerOk() (*string, bool)`

GetIpVerOk returns a tuple with the IpVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVer

`func (o *InlineResponse200211Rules) SetIpVer(v string)`

SetIpVer sets IpVer field to given value.

### HasIpVer

`func (o *InlineResponse200211Rules) HasIpVer() bool`

HasIpVer returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse200211Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200211Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200211Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestPort

`func (o *InlineResponse200211Rules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *InlineResponse200211Rules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *InlineResponse200211Rules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *InlineResponse200211Rules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *InlineResponse200211Rules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *InlineResponse200211Rules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *InlineResponse200211Rules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


