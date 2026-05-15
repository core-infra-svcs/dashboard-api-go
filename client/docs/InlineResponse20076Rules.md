# InlineResponse20076Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**PublicIp** | Pointer to **string** | The IP address that will be used to access the internal resource from the WAN | [optional] 
**LanIp** | Pointer to **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | [optional] 
**Uplink** | Pointer to **string** | The physical WAN interface on which the traffic will arrive, formatted as &#39;internetN&#39; where N is an integer representing a valid uplink for the network&#39;s appliance | [optional] 
**AllowedInbound** | Pointer to [**[]InlineResponse20076AllowedInbound**](InlineResponse20076AllowedInbound.md) | The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource | [optional] 

## Methods

### NewInlineResponse20076Rules

`func NewInlineResponse20076Rules() *InlineResponse20076Rules`

NewInlineResponse20076Rules instantiates a new InlineResponse20076Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20076RulesWithDefaults

`func NewInlineResponse20076RulesWithDefaults() *InlineResponse20076Rules`

NewInlineResponse20076RulesWithDefaults instantiates a new InlineResponse20076Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20076Rules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20076Rules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20076Rules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20076Rules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse20076Rules) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse20076Rules) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse20076Rules) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse20076Rules) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetLanIp

`func (o *InlineResponse20076Rules) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse20076Rules) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse20076Rules) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse20076Rules) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse20076Rules) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20076Rules) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20076Rules) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20076Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetAllowedInbound

`func (o *InlineResponse20076Rules) GetAllowedInbound() []InlineResponse20076AllowedInbound`

GetAllowedInbound returns the AllowedInbound field if non-nil, zero value otherwise.

### GetAllowedInboundOk

`func (o *InlineResponse20076Rules) GetAllowedInboundOk() (*[]InlineResponse20076AllowedInbound, bool)`

GetAllowedInboundOk returns a tuple with the AllowedInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInbound

`func (o *InlineResponse20076Rules) SetAllowedInbound(v []InlineResponse20076AllowedInbound)`

SetAllowedInbound sets AllowedInbound field to given value.

### HasAllowedInbound

`func (o *InlineResponse20076Rules) HasAllowedInbound() bool`

HasAllowedInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


