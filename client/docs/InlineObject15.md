# InlineObject15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | **string** | Required parameter for the SIM slot to update the cellular band mask for | 
**Type** | **string** | Required parameter for the signal type to update the cellular band mask for | 
**Masked** | **[]string** | Required parameter for the band identifiers to mask for the given SIM slot and signal type. For LTE use bands identifiers like &#39;30&#39; and for 5G use band identifiers like &#39;n30&#39;. Maximum 256 bands. | 

## Methods

### NewInlineObject15

`func NewInlineObject15(slot string, type_ string, masked []string, ) *InlineObject15`

NewInlineObject15 instantiates a new InlineObject15 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject15WithDefaults

`func NewInlineObject15WithDefaults() *InlineObject15`

NewInlineObject15WithDefaults instantiates a new InlineObject15 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineObject15) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineObject15) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineObject15) SetSlot(v string)`

SetSlot sets Slot field to given value.


### GetType

`func (o *InlineObject15) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject15) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject15) SetType(v string)`

SetType sets Type field to given value.


### GetMasked

`func (o *InlineObject15) GetMasked() []string`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *InlineObject15) GetMaskedOk() (*[]string, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *InlineObject15) SetMasked(v []string)`

SetMasked sets Masked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


