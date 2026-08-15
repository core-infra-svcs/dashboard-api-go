# InlineResponse200350Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total cellular data used in bytes across all SIMs | [optional] 
**BySim** | Pointer to [**[]InlineResponse200350UsageBySim**](InlineResponse200350UsageBySim.md) | Cellular data used in bytes by SIM. Maximum 3 SIM entries. | [optional] 

## Methods

### NewInlineResponse200350Usage

`func NewInlineResponse200350Usage() *InlineResponse200350Usage`

NewInlineResponse200350Usage instantiates a new InlineResponse200350Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200350UsageWithDefaults

`func NewInlineResponse200350UsageWithDefaults() *InlineResponse200350Usage`

NewInlineResponse200350UsageWithDefaults instantiates a new InlineResponse200350Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200350Usage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200350Usage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200350Usage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200350Usage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBySim

`func (o *InlineResponse200350Usage) GetBySim() []InlineResponse200350UsageBySim`

GetBySim returns the BySim field if non-nil, zero value otherwise.

### GetBySimOk

`func (o *InlineResponse200350Usage) GetBySimOk() (*[]InlineResponse200350UsageBySim, bool)`

GetBySimOk returns a tuple with the BySim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySim

`func (o *InlineResponse200350Usage) SetBySim(v []InlineResponse200350UsageBySim)`

SetBySim sets BySim field to given value.

### HasBySim

`func (o *InlineResponse200350Usage) HasBySim() bool`

HasBySim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


