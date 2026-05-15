# InlineResponse20077Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanIp** | Pointer to **string** | IP address of the device subject to port forwarding | [optional] 
**AllowedIps** | Pointer to **[]string** | An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any) | [optional] 
**Name** | Pointer to **string** | Name of the rule | [optional] 
**Protocol** | Pointer to **string** | Protocol the rule applies to | [optional] 
**PublicPort** | Pointer to **string** | The port or port range forwarded to the host on the LAN | [optional] 
**LocalPort** | Pointer to **string** | The port or port range that receives forwarded traffic from the WAN | [optional] 
**Uplink** | Pointer to **string** | The physical WAN interface on which the traffic arrives; allowed values vary by appliance model and configuration | [optional] 

## Methods

### NewInlineResponse20077Rules

`func NewInlineResponse20077Rules() *InlineResponse20077Rules`

NewInlineResponse20077Rules instantiates a new InlineResponse20077Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20077RulesWithDefaults

`func NewInlineResponse20077RulesWithDefaults() *InlineResponse20077Rules`

NewInlineResponse20077RulesWithDefaults instantiates a new InlineResponse20077Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanIp

`func (o *InlineResponse20077Rules) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse20077Rules) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse20077Rules) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse20077Rules) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetAllowedIps

`func (o *InlineResponse20077Rules) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *InlineResponse20077Rules) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *InlineResponse20077Rules) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *InlineResponse20077Rules) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20077Rules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20077Rules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20077Rules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20077Rules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20077Rules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20077Rules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20077Rules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20077Rules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublicPort

`func (o *InlineResponse20077Rules) GetPublicPort() string`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *InlineResponse20077Rules) GetPublicPortOk() (*string, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *InlineResponse20077Rules) SetPublicPort(v string)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *InlineResponse20077Rules) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetLocalPort

`func (o *InlineResponse20077Rules) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *InlineResponse20077Rules) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *InlineResponse20077Rules) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *InlineResponse20077Rules) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse20077Rules) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20077Rules) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20077Rules) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20077Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


