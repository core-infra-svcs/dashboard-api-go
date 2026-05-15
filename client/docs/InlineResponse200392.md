# InlineResponse200392

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Unique identifier for the enrolled site | [optional] 
**Network** | Pointer to [**InlineResponse200389Network**](InlineResponse200389Network.md) |  | [optional] 
**Type** | Pointer to **string** | The site&#39;s type | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**Region** | Pointer to [**InlineResponse200389Region**](InlineResponse200389Region.md) |  | [optional] 
**Model** | Pointer to **string** | Model of the network&#39;s primary MX device | [optional] 
**Address** | Pointer to [**InlineResponse200387Address**](InlineResponse200387Address.md) |  | [optional] 
**Vpn** | Pointer to [**InlineResponse200387Vpn**](InlineResponse200387Vpn.md) |  | [optional] 
**Routing** | Pointer to [**InlineResponse200387Routing**](InlineResponse200387Routing.md) |  | [optional] 

## Methods

### NewInlineResponse200392

`func NewInlineResponse200392() *InlineResponse200392`

NewInlineResponse200392 instantiates a new InlineResponse200392 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200392WithDefaults

`func NewInlineResponse200392WithDefaults() *InlineResponse200392`

NewInlineResponse200392WithDefaults instantiates a new InlineResponse200392 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *InlineResponse200392) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200392) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200392) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200392) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200392) GetNetwork() InlineResponse200389Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200392) GetNetworkOk() (*InlineResponse200389Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200392) SetNetwork(v InlineResponse200389Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200392) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200392) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200392) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200392) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200392) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200392) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200392) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200392) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200392) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *InlineResponse200392) GetRegion() InlineResponse200389Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InlineResponse200392) GetRegionOk() (*InlineResponse200389Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InlineResponse200392) SetRegion(v InlineResponse200389Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InlineResponse200392) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200392) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200392) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200392) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200392) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200392) GetAddress() InlineResponse200387Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200392) GetAddressOk() (*InlineResponse200387Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200392) SetAddress(v InlineResponse200387Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200392) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVpn

`func (o *InlineResponse200392) GetVpn() InlineResponse200387Vpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *InlineResponse200392) GetVpnOk() (*InlineResponse200387Vpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *InlineResponse200392) SetVpn(v InlineResponse200387Vpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *InlineResponse200392) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetRouting

`func (o *InlineResponse200392) GetRouting() InlineResponse200387Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *InlineResponse200392) GetRoutingOk() (*InlineResponse200387Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *InlineResponse200392) SetRouting(v InlineResponse200387Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *InlineResponse200392) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


