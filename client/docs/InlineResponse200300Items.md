# InlineResponse200300Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **NullableString** | Identifier for the policy (present for default base policies). | [optional] 
**Counts** | Pointer to [**InlineResponse200300Counts**](InlineResponse200300Counts.md) |  | [optional] 

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

### GetPolicyId

`func (o *InlineResponse200300Items) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *InlineResponse200300Items) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *InlineResponse200300Items) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *InlineResponse200300Items) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *InlineResponse200300Items) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *InlineResponse200300Items) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetCounts

`func (o *InlineResponse200300Items) GetCounts() InlineResponse200300Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200300Items) GetCountsOk() (*InlineResponse200300Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200300Items) SetCounts(v InlineResponse200300Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200300Items) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


