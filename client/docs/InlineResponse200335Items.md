# InlineResponse200335Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | Network identifier | 
**Name** | **string** | Network name | 
**Tunneling** | [**InlineResponse200335Tunneling**](InlineResponse200335Tunneling.md) |  | 
**Site** | Pointer to [**InlineResponse200335Site**](InlineResponse200335Site.md) |  | [optional] 
**Url** | **string** | Network management URL | 
**Counts** | [**InlineResponse200335Counts**](InlineResponse200335Counts.md) |  | 
**Cluster** | [**InlineResponse200335Cluster**](InlineResponse200335Cluster.md) |  | 

## Methods

### NewInlineResponse200335Items

`func NewInlineResponse200335Items(networkId string, name string, tunneling InlineResponse200335Tunneling, url string, counts InlineResponse200335Counts, cluster InlineResponse200335Cluster, ) *InlineResponse200335Items`

NewInlineResponse200335Items instantiates a new InlineResponse200335Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200335ItemsWithDefaults

`func NewInlineResponse200335ItemsWithDefaults() *InlineResponse200335Items`

NewInlineResponse200335ItemsWithDefaults instantiates a new InlineResponse200335Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200335Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200335Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200335Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *InlineResponse200335Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200335Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200335Items) SetName(v string)`

SetName sets Name field to given value.


### GetTunneling

`func (o *InlineResponse200335Items) GetTunneling() InlineResponse200335Tunneling`

GetTunneling returns the Tunneling field if non-nil, zero value otherwise.

### GetTunnelingOk

`func (o *InlineResponse200335Items) GetTunnelingOk() (*InlineResponse200335Tunneling, bool)`

GetTunnelingOk returns a tuple with the Tunneling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneling

`func (o *InlineResponse200335Items) SetTunneling(v InlineResponse200335Tunneling)`

SetTunneling sets Tunneling field to given value.


### GetSite

`func (o *InlineResponse200335Items) GetSite() InlineResponse200335Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InlineResponse200335Items) GetSiteOk() (*InlineResponse200335Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InlineResponse200335Items) SetSite(v InlineResponse200335Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *InlineResponse200335Items) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200335Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200335Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200335Items) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCounts

`func (o *InlineResponse200335Items) GetCounts() InlineResponse200335Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200335Items) GetCountsOk() (*InlineResponse200335Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200335Items) SetCounts(v InlineResponse200335Counts)`

SetCounts sets Counts field to given value.


### GetCluster

`func (o *InlineResponse200335Items) GetCluster() InlineResponse200335Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InlineResponse200335Items) GetClusterOk() (*InlineResponse200335Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InlineResponse200335Items) SetCluster(v InlineResponse200335Cluster)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


