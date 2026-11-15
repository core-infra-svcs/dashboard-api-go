# InlineResponse200379Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total cellular data used in bytes across all SIMs | [optional] 
**BySim** | Pointer to [**[]InlineResponse200379UsageBySim**](InlineResponse200379UsageBySim.md) | Cellular data used in bytes by SIM. Maximum 3 SIM entries. | [optional] 

## Methods

### NewInlineResponse200379Usage

`func NewInlineResponse200379Usage() *InlineResponse200379Usage`

NewInlineResponse200379Usage instantiates a new InlineResponse200379Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200379UsageWithDefaults

`func NewInlineResponse200379UsageWithDefaults() *InlineResponse200379Usage`

NewInlineResponse200379UsageWithDefaults instantiates a new InlineResponse200379Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200379Usage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200379Usage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200379Usage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200379Usage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBySim

`func (o *InlineResponse200379Usage) GetBySim() []InlineResponse200379UsageBySim`

GetBySim returns the BySim field if non-nil, zero value otherwise.

### GetBySimOk

`func (o *InlineResponse200379Usage) GetBySimOk() (*[]InlineResponse200379UsageBySim, bool)`

GetBySimOk returns a tuple with the BySim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySim

`func (o *InlineResponse200379Usage) SetBySim(v []InlineResponse200379UsageBySim)`

SetBySim sets BySim field to given value.

### HasBySim

`func (o *InlineResponse200379Usage) HasBySim() bool`

HasBySim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


