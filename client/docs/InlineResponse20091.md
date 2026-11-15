# InlineResponse20091

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse20091Network**](InlineResponse20091Network.md) |  | 
**Enabled** | **bool** | Whether HTTPS inspection is enabled. | 
**AllowLists** | [**InlineResponse20091AllowLists**](InlineResponse20091AllowLists.md) |  | 

## Methods

### NewInlineResponse20091

`func NewInlineResponse20091(network InlineResponse20091Network, enabled bool, allowLists InlineResponse20091AllowLists, ) *InlineResponse20091`

NewInlineResponse20091 instantiates a new InlineResponse20091 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20091WithDefaults

`func NewInlineResponse20091WithDefaults() *InlineResponse20091`

NewInlineResponse20091WithDefaults instantiates a new InlineResponse20091 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse20091) GetNetwork() InlineResponse20091Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20091) GetNetworkOk() (*InlineResponse20091Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20091) SetNetwork(v InlineResponse20091Network)`

SetNetwork sets Network field to given value.


### GetEnabled

`func (o *InlineResponse20091) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20091) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20091) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowLists

`func (o *InlineResponse20091) GetAllowLists() InlineResponse20091AllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse20091) GetAllowListsOk() (*InlineResponse20091AllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse20091) SetAllowLists(v InlineResponse20091AllowLists)`

SetAllowLists sets AllowLists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


