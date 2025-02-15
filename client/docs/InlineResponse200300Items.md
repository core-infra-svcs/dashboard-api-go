# InlineResponse200300Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | The Id of the Network | [optional] 
**Policies** | Pointer to [**[]InlineResponse200300Policies**](InlineResponse200300Policies.md) | Array of Sentry Group Policies for the Network | [optional] 

## Methods

### NewInlineResponse200300Items

`func NewInlineResponse200300Items() *InlineResponse200300Items`

NewInlineResponse200300Items instantiates a new InlineResponse200300Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200300ItemsWithDefaults

`func NewInlineResponse200300ItemsWithDefaults() *InlineResponse200300Items`

NewInlineResponse200300ItemsWithDefaults instantiates a new InlineResponse200300Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200300Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200300Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200300Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200300Items) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200300Items) GetPolicies() []InlineResponse200300Policies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200300Items) GetPoliciesOk() (*[]InlineResponse200300Policies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200300Items) SetPolicies(v []InlineResponse200300Policies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200300Items) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


