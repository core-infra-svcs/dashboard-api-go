# InlineResponse200452Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Unique identifier for the enrolled site | [optional] 
**Network** | Pointer to [**InlineResponse200452Network**](InlineResponse200452Network.md) |  | [optional] 
**Type** | Pointer to **string** | The site&#39;s type | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**Region** | Pointer to [**InlineResponse200452Region**](InlineResponse200452Region.md) |  | [optional] 
**Model** | Pointer to **string** | Model of the network&#39;s primary MX device | [optional] 
**Address** | Pointer to [**InlineResponse200450Address**](InlineResponse200450Address.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse200450Vpn**](InlineResponse200450Vpn.md) |  | [optional] 
**Routing** | Pointer to [**InlineResponse200450Routing**](InlineResponse200450Routing.md) |  | [optional] 
**Devices** | Pointer to [**InlineResponse200452Devices**](InlineResponse200452Devices.md) |  | [optional] 
**Subnets** | Pointer to [**[]InlineResponse200452Subnets**](InlineResponse200452Subnets.md) | List of local networks in the site | [optional] 
**Url** | Pointer to **string** | Dashboard URL for the site&#39;s network | [optional] 

## Methods

### NewInlineResponse200452Items

`func NewInlineResponse200452Items() *InlineResponse200452Items`

NewInlineResponse200452Items instantiates a new InlineResponse200452Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200452ItemsWithDefaults

`func NewInlineResponse200452ItemsWithDefaults() *InlineResponse200452Items`

NewInlineResponse200452ItemsWithDefaults instantiates a new InlineResponse200452Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *InlineResponse200452Items) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200452Items) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200452Items) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200452Items) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200452Items) GetNetwork() InlineResponse200452Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200452Items) GetNetworkOk() (*InlineResponse200452Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200452Items) SetNetwork(v InlineResponse200452Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200452Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200452Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200452Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200452Items) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200452Items) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200452Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200452Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200452Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200452Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *InlineResponse200452Items) GetRegion() InlineResponse200452Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineResponse200452Items) GetRegionOk() (*InlineResponse200452Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineResponse200452Items) SetRegion(v InlineResponse200452Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineResponse200452Items) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200452Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200452Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200452Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200452Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200452Items) GetAddress() InlineResponse200450Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200452Items) GetAddressOk() (*InlineResponse200450Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200452Items) SetAddress(v InlineResponse200450Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200452Items) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse200452Items) GetVpn() InlineResponse200450Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse200452Items) GetVpnOk() (*InlineResponse200450Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse200452Items) SetVpn(v InlineResponse200450Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse200452Items) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetRouting

`func (o *InlineResponse200452Items) GetRouting() InlineResponse200450Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *InlineResponse200452Items) GetRoutingOk() (*InlineResponse200450Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *InlineResponse200452Items) SetRouting(v InlineResponse200450Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *InlineResponse200452Items) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200452Items) GetDevices() InlineResponse200452Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200452Items) GetDevicesOk() (*InlineResponse200452Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200452Items) SetDevices(v InlineResponse200452Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200452Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse200452Items) GetSubnets() []InlineResponse200452Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse200452Items) GetSubnetsOk() (*[]InlineResponse200452Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse200452Items) SetSubnets(v []InlineResponse200452Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse200452Items) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200452Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200452Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200452Items) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200452Items) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


