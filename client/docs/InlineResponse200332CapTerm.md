# InlineResponse200332CapTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resets** | Pointer to **string** | How often the device resets the aggregated data count. One of: &#39;daily&#39;, &#39;weekly&#39;, or &#39;monthly&#39;. | [optional] 
**Starts** | Pointer to [**InlineResponse200332CapTermStarts**](InlineResponse200332CapTermStarts.md) |  | [optional] 

## Methods

### NewInlineResponse200332CapTerm

`func NewInlineResponse200332CapTerm() *InlineResponse200332CapTerm`

NewInlineResponse200332CapTerm instantiates a new InlineResponse200332CapTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200332CapTermWithDefaults

`func NewInlineResponse200332CapTermWithDefaults() *InlineResponse200332CapTerm`

NewInlineResponse200332CapTermWithDefaults instantiates a new InlineResponse200332CapTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResets

`func (o *InlineResponse200332CapTerm) GetResets() string`

GetResets returns the Resets field if non-nil, zero value otherwise.

### GetResetsOk

`func (o *InlineResponse200332CapTerm) GetResetsOk() (*string, bool)`

GetResetsOk returns a tuple with the Resets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResets

`func (o *InlineResponse200332CapTerm) SetResets(v string)`

SetResets sets Resets field to given value.

### HasResets

`func (o *InlineResponse200332CapTerm) HasResets() bool`

HasResets returns a boolean if a field has been set.

### GetStarts

`func (o *InlineResponse200332CapTerm) GetStarts() InlineResponse200332CapTermStarts`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *InlineResponse200332CapTerm) GetStartsOk() (*InlineResponse200332CapTermStarts, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *InlineResponse200332CapTerm) SetStarts(v InlineResponse200332CapTermStarts)`

SetStarts sets Starts field to given value.

### HasStarts

`func (o *InlineResponse200332CapTerm) HasStarts() bool`

HasStarts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


