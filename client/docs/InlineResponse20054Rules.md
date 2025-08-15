# InlineResponse20054Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | Pointer to **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | [optional] 
**Protocol** | Pointer to **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | [optional] 
**SrcPort** | Pointer to **string** | Comma-separated list of source port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**SrcCidr** | Pointer to **string** | Comma-separated list of source IP address(es) (in IP or CIDR notation), or &#39;any&#39; (note: FQDN not supported for source addresses) | [optional] 
**DestPort** | Pointer to **string** | Comma-separated list of destination port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**DestCidr** | Pointer to **string** | Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or &#39;any&#39; | [optional] 
**SyslogEnabled** | Pointer to **bool** | Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional) | [optional] 

## Methods

### NewInlineResponse20054Rules

`func NewInlineResponse20054Rules() *InlineResponse20054Rules`

NewInlineResponse20054Rules instantiates a new InlineResponse20054Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20054RulesWithDefaults

`func NewInlineResponse20054RulesWithDefaults() *InlineResponse20054Rules`

NewInlineResponse20054RulesWithDefaults instantiates a new InlineResponse20054Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineResponse20054Rules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineResponse20054Rules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineResponse20054Rules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineResponse20054Rules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse20054Rules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse20054Rules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse20054Rules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse20054Rules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20054Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20054Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20054Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20054Rules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *InlineResponse20054Rules) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *InlineResponse20054Rules) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *InlineResponse20054Rules) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *InlineResponse20054Rules) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcCidr

`func (o *InlineResponse20054Rules) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *InlineResponse20054Rules) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *InlineResponse20054Rules) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.

### HasSrcCidr

`func (o *InlineResponse20054Rules) HasSrcCidr() bool`

HasSrcCidr returns a boolean if a field has been set.

### GetDestPort

`func (o *InlineResponse20054Rules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *InlineResponse20054Rules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *InlineResponse20054Rules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *InlineResponse20054Rules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *InlineResponse20054Rules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *InlineResponse20054Rules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *InlineResponse20054Rules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.

### HasDestCidr

`func (o *InlineResponse20054Rules) HasDestCidr() bool`

HasDestCidr returns a boolean if a field has been set.

### GetSyslogEnabled

`func (o *InlineResponse20054Rules) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *InlineResponse20054Rules) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *InlineResponse20054Rules) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *InlineResponse20054Rules) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


