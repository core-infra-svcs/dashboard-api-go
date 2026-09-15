# InlineResponse200133

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse200133Network**](InlineResponse200133Network.md) |  | 
**Servers** | [**[]InlineResponse200133Servers**](InlineResponse200133Servers.md) | The list of syslog server configurations that are assigned to this network; suggested maximum array size is 10 | 

## Methods

### NewInlineResponse200133

`func NewInlineResponse200133(network InlineResponse200133Network, servers []InlineResponse200133Servers, ) *InlineResponse200133`

NewInlineResponse200133 instantiates a new InlineResponse200133 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200133WithDefaults

`func NewInlineResponse200133WithDefaults() *InlineResponse200133`

NewInlineResponse200133WithDefaults instantiates a new InlineResponse200133 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200133) GetNetwork() InlineResponse200133Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200133) GetNetworkOk() (*InlineResponse200133Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200133) SetNetwork(v InlineResponse200133Network)`

SetNetwork sets Network field to given value.


### GetServers

`func (o *InlineResponse200133) GetServers() []InlineResponse200133Servers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineResponse200133) GetServersOk() (*[]InlineResponse200133Servers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineResponse200133) SetServers(v []InlineResponse200133Servers)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


