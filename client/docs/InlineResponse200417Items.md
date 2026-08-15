# InlineResponse200417Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID, may not be unique in result set | [optional] 
**Type** | Pointer to **string** | The site&#39;s type (one of: &#39;Meraki spoke&#39;, &#39;Meraki hub&#39;, &#39;Meraki template&#39;) | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**Region** | Pointer to [**InlineResponse200417Region**](InlineResponse200417Region.md) |  | [optional] 
**Device** | Pointer to [**InlineResponse200417Device**](InlineResponse200417Device.md) |  | [optional] 
**Address** | Pointer to [**InlineResponse200417Address**](InlineResponse200417Address.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse200417Vpn**](InlineResponse200417Vpn.md) |  | [optional] 
**Routing** | Pointer to [**InlineResponse200417Routing**](InlineResponse200417Routing.md) |  | [optional] 

## Methods

### NewInlineResponse200417Items

`func NewInlineResponse200417Items() *InlineResponse200417Items`

NewInlineResponse200417Items instantiates a new InlineResponse200417Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200417ItemsWithDefaults

`func NewInlineResponse200417ItemsWithDefaults() *InlineResponse200417Items`

NewInlineResponse200417ItemsWithDefaults instantiates a new InlineResponse200417Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200417Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200417Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200417Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200417Items) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200417Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200417Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200417Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200417Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200417Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200417Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200417Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200417Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *InlineResponse200417Items) GetRegion() InlineResponse200417Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineResponse200417Items) GetRegionOk() (*InlineResponse200417Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineResponse200417Items) SetRegion(v InlineResponse200417Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineResponse200417Items) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200417Items) GetDevice() InlineResponse200417Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200417Items) GetDeviceOk() (*InlineResponse200417Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200417Items) SetDevice(v InlineResponse200417Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200417Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200417Items) GetAddress() InlineResponse200417Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200417Items) GetAddressOk() (*InlineResponse200417Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200417Items) SetAddress(v InlineResponse200417Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200417Items) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse200417Items) GetVpn() InlineResponse200417Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse200417Items) GetVpnOk() (*InlineResponse200417Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse200417Items) SetVpn(v InlineResponse200417Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse200417Items) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetRouting

`func (o *InlineResponse200417Items) GetRouting() InlineResponse200417Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *InlineResponse200417Items) GetRoutingOk() (*InlineResponse200417Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *InlineResponse200417Items) SetRouting(v InlineResponse200417Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *InlineResponse200417Items) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


