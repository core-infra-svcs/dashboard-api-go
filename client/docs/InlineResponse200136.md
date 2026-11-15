# InlineResponse200136

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse200136Network**](InlineResponse200136Network.md) |  | 
**Servers** | [**[]InlineResponse200136Servers**](InlineResponse200136Servers.md) | The list of syslog server configurations that are assigned to this network; suggested maximum array size is 10 | 

## Methods

### NewInlineResponse200136

`func NewInlineResponse200136(network InlineResponse200136Network, servers []InlineResponse200136Servers, ) *InlineResponse200136`

NewInlineResponse200136 instantiates a new InlineResponse200136 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200136WithDefaults

`func NewInlineResponse200136WithDefaults() *InlineResponse200136`

NewInlineResponse200136WithDefaults instantiates a new InlineResponse200136 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200136) GetNetwork() InlineResponse200136Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200136) GetNetworkOk() (*InlineResponse200136Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200136) SetNetwork(v InlineResponse200136Network)`

SetNetwork sets Network field to given value.


### GetServers

`func (o *InlineResponse200136) GetServers() []InlineResponse200136Servers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineResponse200136) GetServersOk() (*[]InlineResponse200136Servers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineResponse200136) SetServers(v []InlineResponse200136Servers)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


