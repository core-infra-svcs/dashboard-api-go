# InlineResponse200322Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | Network identifier | 
**Name** | **string** | Network name | 
**Tunneling** | [**InlineResponse200322Tunneling**](InlineResponse200322Tunneling.md) |  | 
**Site** | Pointer to [**InlineResponse200322Site**](InlineResponse200322Site.md) |  | [optional] 
**Url** | **string** | Network management URL | 
**Counts** | [**InlineResponse200322Counts**](InlineResponse200322Counts.md) |  | 
**Cluster** | [**InlineResponse200322Cluster**](InlineResponse200322Cluster.md) |  | 

## Methods

### NewInlineResponse200322Items

`func NewInlineResponse200322Items(networkId string, name string, tunneling InlineResponse200322Tunneling, url string, counts InlineResponse200322Counts, cluster InlineResponse200322Cluster, ) *InlineResponse200322Items`

NewInlineResponse200322Items instantiates a new InlineResponse200322Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200322ItemsWithDefaults

`func NewInlineResponse200322ItemsWithDefaults() *InlineResponse200322Items`

NewInlineResponse200322ItemsWithDefaults instantiates a new InlineResponse200322Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200322Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200322Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200322Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *InlineResponse200322Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200322Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200322Items) SetName(v string)`

SetName sets Name field to given value.


### GetTunneling

`func (o *InlineResponse200322Items) GetTunneling() InlineResponse200322Tunneling`

GetTunneling returns the Tunneling field if non-nil, zero value otherwise.

### GetTunnelingOk

`func (o *InlineResponse200322Items) GetTunnelingOk() (*InlineResponse200322Tunneling, bool)`

GetTunnelingOk returns a tuple with the Tunneling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneling

`func (o *InlineResponse200322Items) SetTunneling(v InlineResponse200322Tunneling)`

SetTunneling sets Tunneling field to given value.


### GetSite

`func (o *InlineResponse200322Items) GetSite() InlineResponse200322Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InlineResponse200322Items) GetSiteOk() (*InlineResponse200322Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InlineResponse200322Items) SetSite(v InlineResponse200322Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *InlineResponse200322Items) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200322Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200322Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200322Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCounts

`func (o *InlineResponse200322Items) GetCounts() InlineResponse200322Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200322Items) GetCountsOk() (*InlineResponse200322Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200322Items) SetCounts(v InlineResponse200322Counts)`

SetCounts sets Counts field to given value.


### GetCluster

`func (o *InlineResponse200322Items) GetCluster() InlineResponse200322Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200322Items) GetClusterOk() (*InlineResponse200322Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200322Items) SetCluster(v InlineResponse200322Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


