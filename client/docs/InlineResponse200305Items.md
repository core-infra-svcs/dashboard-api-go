# InlineResponse200305Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | The Id of the Network | [optional] 
**Policies** | Pointer to [**[]InlineResponse200305Policies**](InlineResponse200305Policies.md) | Array of Sentry Group Policies for the Network | [optional] 

## Methods

### NewInlineResponse200305Items

`func NewInlineResponse200305Items() *InlineResponse200305Items`

NewInlineResponse200305Items instantiates a new InlineResponse200305Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200305ItemsWithDefaults

`func NewInlineResponse200305ItemsWithDefaults() *InlineResponse200305Items`

NewInlineResponse200305ItemsWithDefaults instantiates a new InlineResponse200305Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200305Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200305Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200305Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200305Items) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200305Items) GetPolicies() []InlineResponse200305Policies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200305Items) GetPoliciesOk() (*[]InlineResponse200305Policies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200305Items) SetPolicies(v []InlineResponse200305Policies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200305Items) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


