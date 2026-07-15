# InlineResponse200354Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse200131Network**](InlineResponse200131Network.md) |  | 
**Available** | [**[]InlineResponse200354Available**](InlineResponse200354Available.md) | The list of roles that can be assigned to a syslog server for this network | 

## Methods

### NewInlineResponse200354Items

`func NewInlineResponse200354Items(network InlineResponse200131Network, available []InlineResponse200354Available, ) *InlineResponse200354Items`

NewInlineResponse200354Items instantiates a new InlineResponse200354Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200354ItemsWithDefaults

`func NewInlineResponse200354ItemsWithDefaults() *InlineResponse200354Items`

NewInlineResponse200354ItemsWithDefaults instantiates a new InlineResponse200354Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200354Items) GetNetwork() InlineResponse200131Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200354Items) GetNetworkOk() (*InlineResponse200131Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200354Items) SetNetwork(v InlineResponse200131Network)`

SetNetwork sets Network field to given value.


### GetAvailable

`func (o *InlineResponse200354Items) GetAvailable() []InlineResponse200354Available`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse200354Items) GetAvailableOk() (*[]InlineResponse200354Available, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse200354Items) SetAvailable(v []InlineResponse200354Available)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


