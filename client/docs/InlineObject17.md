# InlineObject17

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | **string** | Required parameter for the SIM slot to update the cellular band mask for | 
**Type** | **string** | Required parameter for the signal type to update the cellular band mask for | 
**Masked** | **[]string** | Required parameter for the band identifiers to mask for the given SIM slot and signal type. For LTE use bands identifiers like &#39;30&#39;, for 5G use band identifiers like &#39;n30&#39;, or use &#39;all&#39; to mask all bands for that signal type. Maximum 256 bands. | 

## Methods

### NewInlineObject17

`func NewInlineObject17(slot string, type_ string, masked []string, ) *InlineObject17`

NewInlineObject17 instantiates a new InlineObject17 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject17WithDefaults

`func NewInlineObject17WithDefaults() *InlineObject17`

NewInlineObject17WithDefaults instantiates a new InlineObject17 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineObject17) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineObject17) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineObject17) SetSlot(v string)`

SetSlot sets Slot field to given value.


### GetType

`func (o *InlineObject17) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject17) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject17) SetType(v string)`

SetType sets Type field to given value.


### GetMasked

`func (o *InlineObject17) GetMasked() []string`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *InlineObject17) GetMaskedOk() (*[]string, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *InlineObject17) SetMasked(v []string)`

SetMasked sets Masked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


