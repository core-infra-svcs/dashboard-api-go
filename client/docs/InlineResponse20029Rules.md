# InlineResponse20029Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**LanIp** | Pointer to **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | [optional] 
**PublicPort** | Pointer to **string** | A port or port ranges that will be forwarded to the host on the LAN | [optional] 
**LocalPort** | Pointer to **string** | A port or port ranges that will receive the forwarded traffic from the WAN | [optional] 
**AllowedIps** | Pointer to **[]string** | An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges. | [optional] 
**Protocol** | Pointer to **string** | TCP or UDP | [optional] 
**Access** | Pointer to **string** | &#x60;any&#x60; or &#x60;restricted&#x60;. Specify the right to make inbound connections on the specified ports or port ranges. If &#x60;restricted&#x60;, a list of allowed IPs is mandatory. | [optional] 

## Methods

### NewInlineResponse20029Rules

`func NewInlineResponse20029Rules() *InlineResponse20029Rules`

NewInlineResponse20029Rules instantiates a new InlineResponse20029Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029RulesWithDefaults

`func NewInlineResponse20029RulesWithDefaults() *InlineResponse20029Rules`

NewInlineResponse20029RulesWithDefaults instantiates a new InlineResponse20029Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20029Rules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20029Rules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20029Rules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20029Rules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanIp

`func (o *InlineResponse20029Rules) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse20029Rules) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse20029Rules) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse20029Rules) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetPublicPort

`func (o *InlineResponse20029Rules) GetPublicPort() string`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *InlineResponse20029Rules) GetPublicPortOk() (*string, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *InlineResponse20029Rules) SetPublicPort(v string)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *InlineResponse20029Rules) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetLocalPort

`func (o *InlineResponse20029Rules) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *InlineResponse20029Rules) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *InlineResponse20029Rules) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *InlineResponse20029Rules) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetAllowedIps

`func (o *InlineResponse20029Rules) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *InlineResponse20029Rules) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *InlineResponse20029Rules) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *InlineResponse20029Rules) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20029Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20029Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20029Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20029Rules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAccess

`func (o *InlineResponse20029Rules) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineResponse20029Rules) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineResponse20029Rules) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineResponse20029Rules) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


