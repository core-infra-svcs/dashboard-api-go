# InlineResponse200384Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse200133Network**](InlineResponse200133Network.md) |  | 
**Available** | [**[]InlineResponse200384Available**](InlineResponse200384Available.md) | The list of roles that can be assigned to a syslog server for this network | 

## Methods

### NewInlineResponse200384Items

`func NewInlineResponse200384Items(network InlineResponse200133Network, available []InlineResponse200384Available, ) *InlineResponse200384Items`

NewInlineResponse200384Items instantiates a new InlineResponse200384Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200384ItemsWithDefaults

`func NewInlineResponse200384ItemsWithDefaults() *InlineResponse200384Items`

NewInlineResponse200384ItemsWithDefaults instantiates a new InlineResponse200384Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200384Items) GetNetwork() InlineResponse200133Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200384Items) GetNetworkOk() (*InlineResponse200133Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200384Items) SetNetwork(v InlineResponse200133Network)`

SetNetwork sets Network field to given value.


### GetAvailable

`func (o *InlineResponse200384Items) GetAvailable() []InlineResponse200384Available`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse200384Items) GetAvailableOk() (*[]InlineResponse200384Available, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse200384Items) SetAvailable(v []InlineResponse200384Available)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


