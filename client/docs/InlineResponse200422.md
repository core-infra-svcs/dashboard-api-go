# InlineResponse200422

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Unique identifier for the enrolled site | [optional] 
**Network** | Pointer to [**InlineResponse200419Network**](InlineResponse200419Network.md) |  | [optional] 
**Type** | Pointer to **string** | The site&#39;s type | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**Region** | Pointer to [**InlineResponse200419Region**](InlineResponse200419Region.md) |  | [optional] 
**Model** | Pointer to **string** | Model of the network&#39;s primary MX device | [optional] 
**Address** | Pointer to [**InlineResponse200417Address**](InlineResponse200417Address.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse200417Vpn**](InlineResponse200417Vpn.md) |  | [optional] 
**Routing** | Pointer to [**InlineResponse200417Routing**](InlineResponse200417Routing.md) |  | [optional] 

## Methods

### NewInlineResponse200422

`func NewInlineResponse200422() *InlineResponse200422`

NewInlineResponse200422 instantiates a new InlineResponse200422 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200422WithDefaults

`func NewInlineResponse200422WithDefaults() *InlineResponse200422`

NewInlineResponse200422WithDefaults instantiates a new InlineResponse200422 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *InlineResponse200422) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200422) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200422) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200422) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200422) GetNetwork() InlineResponse200419Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200422) GetNetworkOk() (*InlineResponse200419Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200422) SetNetwork(v InlineResponse200419Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200422) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200422) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200422) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200422) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200422) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200422) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200422) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200422) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200422) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *InlineResponse200422) GetRegion() InlineResponse200419Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineResponse200422) GetRegionOk() (*InlineResponse200419Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineResponse200422) SetRegion(v InlineResponse200419Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineResponse200422) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200422) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200422) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200422) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200422) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200422) GetAddress() InlineResponse200417Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200422) GetAddressOk() (*InlineResponse200417Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200422) SetAddress(v InlineResponse200417Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200422) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse200422) GetVpn() InlineResponse200417Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse200422) GetVpnOk() (*InlineResponse200417Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse200422) SetVpn(v InlineResponse200417Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse200422) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetRouting

`func (o *InlineResponse200422) GetRouting() InlineResponse200417Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *InlineResponse200422) GetRoutingOk() (*InlineResponse200417Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *InlineResponse200422) SetRouting(v InlineResponse200417Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *InlineResponse200422) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


