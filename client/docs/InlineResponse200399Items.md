# InlineResponse200399Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID, may not be unique in result set | [optional] 
**Type** | Pointer to **string** | The site&#39;s type (one of: &#39;Meraki spoke&#39;, &#39;Meraki hub&#39;, &#39;Meraki template&#39;) | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**Region** | Pointer to [**InlineResponse200399Region**](InlineResponse200399Region.md) |  | [optional] 
**Device** | Pointer to [**InlineResponse200399Device**](InlineResponse200399Device.md) |  | [optional] 
**Address** | Pointer to [**InlineResponse200399Address**](InlineResponse200399Address.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse200399Vpn**](InlineResponse200399Vpn.md) |  | [optional] 
**Routing** | Pointer to [**InlineResponse200399Routing**](InlineResponse200399Routing.md) |  | [optional] 

## Methods

### NewInlineResponse200399Items

`func NewInlineResponse200399Items() *InlineResponse200399Items`

NewInlineResponse200399Items instantiates a new InlineResponse200399Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200399ItemsWithDefaults

`func NewInlineResponse200399ItemsWithDefaults() *InlineResponse200399Items`

NewInlineResponse200399ItemsWithDefaults instantiates a new InlineResponse200399Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200399Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200399Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200399Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200399Items) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200399Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200399Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200399Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200399Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200399Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200399Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200399Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200399Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *InlineResponse200399Items) GetRegion() InlineResponse200399Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineResponse200399Items) GetRegionOk() (*InlineResponse200399Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineResponse200399Items) SetRegion(v InlineResponse200399Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineResponse200399Items) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200399Items) GetDevice() InlineResponse200399Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200399Items) GetDeviceOk() (*InlineResponse200399Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200399Items) SetDevice(v InlineResponse200399Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200399Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200399Items) GetAddress() InlineResponse200399Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200399Items) GetAddressOk() (*InlineResponse200399Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200399Items) SetAddress(v InlineResponse200399Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200399Items) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse200399Items) GetVpn() InlineResponse200399Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse200399Items) GetVpnOk() (*InlineResponse200399Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse200399Items) SetVpn(v InlineResponse200399Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse200399Items) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetRouting

`func (o *InlineResponse200399Items) GetRouting() InlineResponse200399Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *InlineResponse200399Items) GetRoutingOk() (*InlineResponse200399Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *InlineResponse200399Items) SetRouting(v InlineResponse200399Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *InlineResponse200399Items) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


