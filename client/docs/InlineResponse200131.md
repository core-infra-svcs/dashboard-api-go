# InlineResponse200131

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse200131Network**](InlineResponse200131Network.md) |  | 
**Servers** | [**[]InlineResponse200131Servers**](InlineResponse200131Servers.md) | The list of syslog server configurations that are assigned to this network; suggested maximum array size is 10 | 

## Methods

### NewInlineResponse200131

`func NewInlineResponse200131(network InlineResponse200131Network, servers []InlineResponse200131Servers, ) *InlineResponse200131`

NewInlineResponse200131 instantiates a new InlineResponse200131 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200131WithDefaults

`func NewInlineResponse200131WithDefaults() *InlineResponse200131`

NewInlineResponse200131WithDefaults instantiates a new InlineResponse200131 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200131) GetNetwork() InlineResponse200131Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200131) GetNetworkOk() (*InlineResponse200131Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200131) SetNetwork(v InlineResponse200131Network)`

SetNetwork sets Network field to given value.


### GetServers

`func (o *InlineResponse200131) GetServers() []InlineResponse200131Servers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineResponse200131) GetServersOk() (*[]InlineResponse200131Servers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineResponse200131) SetServers(v []InlineResponse200131Servers)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


