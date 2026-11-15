# InlineResponse20084Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**PublicIp** | Pointer to **string** | The IP address that will be used to access the internal resource from the WAN | [optional] 
**LanIp** | Pointer to **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | [optional] 
**Uplink** | Pointer to **string** | The physical WAN interface on which the traffic will arrive, formatted as &#39;internetN&#39; where N is an integer representing a valid uplink for the network&#39;s appliance | [optional] 
**AllowedInbound** | Pointer to [**[]InlineResponse20084AllowedInbound**](InlineResponse20084AllowedInbound.md) | The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource | [optional] 

## Methods

### NewInlineResponse20084Rules

`func NewInlineResponse20084Rules() *InlineResponse20084Rules`

NewInlineResponse20084Rules instantiates a new InlineResponse20084Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20084RulesWithDefaults

`func NewInlineResponse20084RulesWithDefaults() *InlineResponse20084Rules`

NewInlineResponse20084RulesWithDefaults instantiates a new InlineResponse20084Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20084Rules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20084Rules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20084Rules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20084Rules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse20084Rules) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse20084Rules) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse20084Rules) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse20084Rules) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetLanIp

`func (o *InlineResponse20084Rules) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse20084Rules) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse20084Rules) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse20084Rules) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse20084Rules) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20084Rules) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20084Rules) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20084Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetAllowedInbound

`func (o *InlineResponse20084Rules) GetAllowedInbound() []InlineResponse20084AllowedInbound`

GetAllowedInbound returns the AllowedInbound field if non-nil, zero value otherwise.

### GetAllowedInboundOk

`func (o *InlineResponse20084Rules) GetAllowedInboundOk() (*[]InlineResponse20084AllowedInbound, bool)`

GetAllowedInboundOk returns a tuple with the AllowedInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInbound

`func (o *InlineResponse20084Rules) SetAllowedInbound(v []InlineResponse20084AllowedInbound)`

SetAllowedInbound sets AllowedInbound field to given value.

### HasAllowedInbound

`func (o *InlineResponse20084Rules) HasAllowedInbound() bool`

HasAllowedInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


