# InlineResponse200180

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200179Items**](InlineResponse200179Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200177Meta**](InlineResponse200177Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200180

`func NewInlineResponse200180() *InlineResponse200180`

NewInlineResponse200180 instantiates a new InlineResponse200180 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200180WithDefaults

`func NewInlineResponse200180WithDefaults() *InlineResponse200180`

NewInlineResponse200180WithDefaults instantiates a new InlineResponse200180 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200180) GetItems() []InlineResponse200179Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200180) GetItemsOk() (*[]InlineResponse200179Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200180) SetItems(v []InlineResponse200179Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200180) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200180) GetMeta() InlineResponse200177Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200180) GetMetaOk() (*InlineResponse200177Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200180) SetMeta(v InlineResponse200177Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200180) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


