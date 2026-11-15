# InlineResponse200342Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | Network identifier | 
**Name** | **string** | Network name | 
**Tunneling** | [**InlineResponse200342Tunneling**](InlineResponse200342Tunneling.md) |  | 
**Site** | Pointer to [**InlineResponse200342Site**](InlineResponse200342Site.md) |  | [optional] 
**Url** | **string** | Network management URL | 
**Counts** | [**InlineResponse200342Counts**](InlineResponse200342Counts.md) |  | 
**Cluster** | [**InlineResponse200342Cluster**](InlineResponse200342Cluster.md) |  | 

## Methods

### NewInlineResponse200342Items

`func NewInlineResponse200342Items(networkId string, name string, tunneling InlineResponse200342Tunneling, url string, counts InlineResponse200342Counts, cluster InlineResponse200342Cluster, ) *InlineResponse200342Items`

NewInlineResponse200342Items instantiates a new InlineResponse200342Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200342ItemsWithDefaults

`func NewInlineResponse200342ItemsWithDefaults() *InlineResponse200342Items`

NewInlineResponse200342ItemsWithDefaults instantiates a new InlineResponse200342Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200342Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200342Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200342Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *InlineResponse200342Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200342Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200342Items) SetName(v string)`

SetName sets Name field to given value.


### GetTunneling

`func (o *InlineResponse200342Items) GetTunneling() InlineResponse200342Tunneling`

GetTunneling returns the Tunneling field if non-nil, zero value otherwise.

### GetTunnelingOk

`func (o *InlineResponse200342Items) GetTunnelingOk() (*InlineResponse200342Tunneling, bool)`

GetTunnelingOk returns a tuple with the Tunneling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneling

`func (o *InlineResponse200342Items) SetTunneling(v InlineResponse200342Tunneling)`

SetTunneling sets Tunneling field to given value.


### GetSite

`func (o *InlineResponse200342Items) GetSite() InlineResponse200342Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InlineResponse200342Items) GetSiteOk() (*InlineResponse200342Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InlineResponse200342Items) SetSite(v InlineResponse200342Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *InlineResponse200342Items) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200342Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200342Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200342Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCounts

`func (o *InlineResponse200342Items) GetCounts() InlineResponse200342Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200342Items) GetCountsOk() (*InlineResponse200342Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200342Items) SetCounts(v InlineResponse200342Counts)`

SetCounts sets Counts field to given value.


### GetCluster

`func (o *InlineResponse200342Items) GetCluster() InlineResponse200342Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200342Items) GetClusterOk() (*InlineResponse200342Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200342Items) SetCluster(v InlineResponse200342Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


