# InlineResponse200311Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | The Id of the Network | [optional] 
**Policies** | Pointer to [**[]InlineResponse200311Policies**](InlineResponse200311Policies.md) | Array of Sentry Group Policies for the Network | [optional] 

## Methods

### NewInlineResponse200311Items

`func NewInlineResponse200311Items() *InlineResponse200311Items`

NewInlineResponse200311Items instantiates a new InlineResponse200311Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200311ItemsWithDefaults

`func NewInlineResponse200311ItemsWithDefaults() *InlineResponse200311Items`

NewInlineResponse200311ItemsWithDefaults instantiates a new InlineResponse200311Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200311Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200311Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200311Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200311Items) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200311Items) GetPolicies() []InlineResponse200311Policies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200311Items) GetPoliciesOk() (*[]InlineResponse200311Policies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200311Items) SetPolicies(v []InlineResponse200311Policies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200311Items) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


