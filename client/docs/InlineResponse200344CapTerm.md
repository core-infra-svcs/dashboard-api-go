# InlineResponse200344CapTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resets** | Pointer to **string** | How often the device resets the aggregated data count. One of: &#39;daily&#39;, &#39;weekly&#39;, or &#39;monthly&#39;. | [optional] 
**Starts** | Pointer to [**InlineResponse200344CapTermStarts**](InlineResponse200344CapTermStarts.md) |  | [optional] 

## Methods

### NewInlineResponse200344CapTerm

`func NewInlineResponse200344CapTerm() *InlineResponse200344CapTerm`

NewInlineResponse200344CapTerm instantiates a new InlineResponse200344CapTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200344CapTermWithDefaults

`func NewInlineResponse200344CapTermWithDefaults() *InlineResponse200344CapTerm`

NewInlineResponse200344CapTermWithDefaults instantiates a new InlineResponse200344CapTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResets

`func (o *InlineResponse200344CapTerm) GetResets() string`

GetResets returns the Resets field if non-nil, zero value otherwise.

### GetResetsOk

`func (o *InlineResponse200344CapTerm) GetResetsOk() (*string, bool)`

GetResetsOk returns a tuple with the Resets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResets

`func (o *InlineResponse200344CapTerm) SetResets(v string)`

SetResets sets Resets field to given value.

### HasResets

`func (o *InlineResponse200344CapTerm) HasResets() bool`

HasResets returns a boolean if a field has been set.

### GetStarts

`func (o *InlineResponse200344CapTerm) GetStarts() InlineResponse200344CapTermStarts`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *InlineResponse200344CapTerm) GetStartsOk() (*InlineResponse200344CapTermStarts, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *InlineResponse200344CapTerm) SetStarts(v InlineResponse200344CapTermStarts)`

SetStarts sets Starts field to given value.

### HasStarts

`func (o *InlineResponse200344CapTerm) HasStarts() bool`

HasStarts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


