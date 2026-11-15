# InlineResponse200396Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse200136Network**](InlineResponse200136Network.md) |  | 
**Available** | [**[]InlineResponse200396Available**](InlineResponse200396Available.md) | The list of roles that can be assigned to a syslog server for this network | 

## Methods

### NewInlineResponse200396Items

`func NewInlineResponse200396Items(network InlineResponse200136Network, available []InlineResponse200396Available, ) *InlineResponse200396Items`

NewInlineResponse200396Items instantiates a new InlineResponse200396Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200396ItemsWithDefaults

`func NewInlineResponse200396ItemsWithDefaults() *InlineResponse200396Items`

NewInlineResponse200396ItemsWithDefaults instantiates a new InlineResponse200396Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200396Items) GetNetwork() InlineResponse200136Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200396Items) GetNetworkOk() (*InlineResponse200136Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200396Items) SetNetwork(v InlineResponse200136Network)`

SetNetwork sets Network field to given value.


### GetAvailable

`func (o *InlineResponse200396Items) GetAvailable() []InlineResponse200396Available`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse200396Items) GetAvailableOk() (*[]InlineResponse200396Available, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse200396Items) SetAvailable(v []InlineResponse200396Available)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


