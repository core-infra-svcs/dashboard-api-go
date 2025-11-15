# InlineResponse200325

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200324Items**](InlineResponse200324Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200227Meta**](InlineResponse200227Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200325

`func NewInlineResponse200325() *InlineResponse200325`

NewInlineResponse200325 instantiates a new InlineResponse200325 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200325WithDefaults

`func NewInlineResponse200325WithDefaults() *InlineResponse200325`

NewInlineResponse200325WithDefaults instantiates a new InlineResponse200325 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200325) GetItems() []InlineResponse200324Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200325) GetItemsOk() (*[]InlineResponse200324Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200325) SetItems(v []InlineResponse200324Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200325) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200325) GetMeta() InlineResponse200227Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200325) GetMetaOk() (*InlineResponse200227Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200325) SetMeta(v InlineResponse200227Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200325) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


