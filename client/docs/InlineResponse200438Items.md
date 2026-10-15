# InlineResponse200438Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID, may not be unique in result set | [optional] 
**Type** | Pointer to **string** | The site&#39;s type (one of: &#39;Meraki spoke&#39;, &#39;Meraki hub&#39;, &#39;Meraki template&#39;) | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**Region** | Pointer to [**InlineResponse200438Region**](InlineResponse200438Region.md) |  | [optional] 
**Device** | Pointer to [**InlineResponse200438Device**](InlineResponse200438Device.md) |  | [optional] 
**Address** | Pointer to [**InlineResponse200438Address**](InlineResponse200438Address.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse200438Vpn**](InlineResponse200438Vpn.md) |  | [optional] 
**Routing** | Pointer to [**InlineResponse200438Routing**](InlineResponse200438Routing.md) |  | [optional] 

## Methods

### NewInlineResponse200438Items

`func NewInlineResponse200438Items() *InlineResponse200438Items`

NewInlineResponse200438Items instantiates a new InlineResponse200438Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200438ItemsWithDefaults

`func NewInlineResponse200438ItemsWithDefaults() *InlineResponse200438Items`

NewInlineResponse200438ItemsWithDefaults instantiates a new InlineResponse200438Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200438Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200438Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200438Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200438Items) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200438Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200438Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200438Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200438Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200438Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200438Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200438Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200438Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *InlineResponse200438Items) GetRegion() InlineResponse200438Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineResponse200438Items) GetRegionOk() (*InlineResponse200438Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineResponse200438Items) SetRegion(v InlineResponse200438Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineResponse200438Items) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200438Items) GetDevice() InlineResponse200438Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200438Items) GetDeviceOk() (*InlineResponse200438Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200438Items) SetDevice(v InlineResponse200438Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200438Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200438Items) GetAddress() InlineResponse200438Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200438Items) GetAddressOk() (*InlineResponse200438Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200438Items) SetAddress(v InlineResponse200438Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200438Items) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse200438Items) GetVpn() InlineResponse200438Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse200438Items) GetVpnOk() (*InlineResponse200438Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse200438Items) SetVpn(v InlineResponse200438Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse200438Items) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetRouting

`func (o *InlineResponse200438Items) GetRouting() InlineResponse200438Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *InlineResponse200438Items) GetRoutingOk() (*InlineResponse200438Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *InlineResponse200438Items) SetRouting(v InlineResponse200438Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *InlineResponse200438Items) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


