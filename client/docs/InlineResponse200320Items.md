# InlineResponse200320Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | Network identifier | 
**Name** | **string** | Network name | 
**Tunneling** | [**InlineResponse200320Tunneling**](InlineResponse200320Tunneling.md) |  | 
**Site** | Pointer to [**InlineResponse200320Site**](InlineResponse200320Site.md) |  | [optional] 
**Url** | **string** | Network management URL | 
**Counts** | [**InlineResponse200320Counts**](InlineResponse200320Counts.md) |  | 
**Cluster** | [**InlineResponse200320Cluster**](InlineResponse200320Cluster.md) |  | 

## Methods

### NewInlineResponse200320Items

`func NewInlineResponse200320Items(networkId string, name string, tunneling InlineResponse200320Tunneling, url string, counts InlineResponse200320Counts, cluster InlineResponse200320Cluster, ) *InlineResponse200320Items`

NewInlineResponse200320Items instantiates a new InlineResponse200320Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200320ItemsWithDefaults

`func NewInlineResponse200320ItemsWithDefaults() *InlineResponse200320Items`

NewInlineResponse200320ItemsWithDefaults instantiates a new InlineResponse200320Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200320Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200320Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200320Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *InlineResponse200320Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200320Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200320Items) SetName(v string)`

SetName sets Name field to given value.


### GetTunneling

`func (o *InlineResponse200320Items) GetTunneling() InlineResponse200320Tunneling`

GetTunneling returns the Tunneling field if non-nil, zero value otherwise.

### GetTunnelingOk

`func (o *InlineResponse200320Items) GetTunnelingOk() (*InlineResponse200320Tunneling, bool)`

GetTunnelingOk returns a tuple with the Tunneling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneling

`func (o *InlineResponse200320Items) SetTunneling(v InlineResponse200320Tunneling)`

SetTunneling sets Tunneling field to given value.


### GetSite

`func (o *InlineResponse200320Items) GetSite() InlineResponse200320Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InlineResponse200320Items) GetSiteOk() (*InlineResponse200320Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InlineResponse200320Items) SetSite(v InlineResponse200320Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *InlineResponse200320Items) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200320Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200320Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200320Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCounts

`func (o *InlineResponse200320Items) GetCounts() InlineResponse200320Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200320Items) GetCountsOk() (*InlineResponse200320Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200320Items) SetCounts(v InlineResponse200320Counts)`

SetCounts sets Counts field to given value.


### GetCluster

`func (o *InlineResponse200320Items) GetCluster() InlineResponse200320Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200320Items) GetClusterOk() (*InlineResponse200320Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200320Items) SetCluster(v InlineResponse200320Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


