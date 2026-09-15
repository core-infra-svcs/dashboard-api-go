# InlineResponse20082Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**PublicIp** | Pointer to **string** | The IP address that will be used to access the internal resource from the WAN | [optional] 
**LanIp** | Pointer to **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | [optional] 
**Uplink** | Pointer to **string** | The physical WAN interface on which the traffic will arrive, formatted as &#39;internetN&#39; where N is an integer representing a valid uplink for the network&#39;s appliance | [optional] 
**AllowedInbound** | Pointer to [**[]InlineResponse20082AllowedInbound**](InlineResponse20082AllowedInbound.md) | The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource | [optional] 

## Methods

### NewInlineResponse20082Rules

`func NewInlineResponse20082Rules() *InlineResponse20082Rules`

NewInlineResponse20082Rules instantiates a new InlineResponse20082Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20082RulesWithDefaults

`func NewInlineResponse20082RulesWithDefaults() *InlineResponse20082Rules`

NewInlineResponse20082RulesWithDefaults instantiates a new InlineResponse20082Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20082Rules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20082Rules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20082Rules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20082Rules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse20082Rules) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse20082Rules) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse20082Rules) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse20082Rules) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetLanIp

`func (o *InlineResponse20082Rules) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse20082Rules) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse20082Rules) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse20082Rules) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse20082Rules) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20082Rules) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20082Rules) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20082Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetAllowedInbound

`func (o *InlineResponse20082Rules) GetAllowedInbound() []InlineResponse20082AllowedInbound`

GetAllowedInbound returns the AllowedInbound field if non-nil, zero value otherwise.

### GetAllowedInboundOk

`func (o *InlineResponse20082Rules) GetAllowedInboundOk() (*[]InlineResponse20082AllowedInbound, bool)`

GetAllowedInboundOk returns a tuple with the AllowedInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInbound

`func (o *InlineResponse20082Rules) SetAllowedInbound(v []InlineResponse20082AllowedInbound)`

SetAllowedInbound sets AllowedInbound field to given value.

### HasAllowedInbound

`func (o *InlineResponse20082Rules) HasAllowedInbound() bool`

HasAllowedInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


