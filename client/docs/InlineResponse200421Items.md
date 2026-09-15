# InlineResponse200421Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID, may not be unique in result set | [optional] 
**Type** | Pointer to **string** | The site&#39;s type (one of: &#39;Meraki spoke&#39;, &#39;Meraki hub&#39;, &#39;Meraki template&#39;) | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**Region** | Pointer to [**InlineResponse200421Region**](InlineResponse200421Region.md) |  | [optional] 
**Device** | Pointer to [**InlineResponse200421Device**](InlineResponse200421Device.md) |  | [optional] 
**Address** | Pointer to [**InlineResponse200421Address**](InlineResponse200421Address.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse200421Vpn**](InlineResponse200421Vpn.md) |  | [optional] 
**Routing** | Pointer to [**InlineResponse200421Routing**](InlineResponse200421Routing.md) |  | [optional] 

## Methods

### NewInlineResponse200421Items

`func NewInlineResponse200421Items() *InlineResponse200421Items`

NewInlineResponse200421Items instantiates a new InlineResponse200421Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200421ItemsWithDefaults

`func NewInlineResponse200421ItemsWithDefaults() *InlineResponse200421Items`

NewInlineResponse200421ItemsWithDefaults instantiates a new InlineResponse200421Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200421Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200421Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200421Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200421Items) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200421Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200421Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200421Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200421Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200421Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200421Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200421Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200421Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *InlineResponse200421Items) GetRegion() InlineResponse200421Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineResponse200421Items) GetRegionOk() (*InlineResponse200421Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineResponse200421Items) SetRegion(v InlineResponse200421Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineResponse200421Items) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200421Items) GetDevice() InlineResponse200421Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200421Items) GetDeviceOk() (*InlineResponse200421Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200421Items) SetDevice(v InlineResponse200421Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200421Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200421Items) GetAddress() InlineResponse200421Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200421Items) GetAddressOk() (*InlineResponse200421Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200421Items) SetAddress(v InlineResponse200421Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200421Items) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse200421Items) GetVpn() InlineResponse200421Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse200421Items) GetVpnOk() (*InlineResponse200421Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse200421Items) SetVpn(v InlineResponse200421Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse200421Items) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetRouting

`func (o *InlineResponse200421Items) GetRouting() InlineResponse200421Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *InlineResponse200421Items) GetRoutingOk() (*InlineResponse200421Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *InlineResponse200421Items) SetRouting(v InlineResponse200421Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *InlineResponse200421Items) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


