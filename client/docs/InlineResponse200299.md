# InlineResponse200299

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | The device model | [optional] 
**Count** | Pointer to **int32** | Total number of devices per model | [optional] 
**Usage** | Pointer to [**OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage**](OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage.md) |  | [optional] 

## Methods

### NewInlineResponse200299

`func NewInlineResponse200299() *InlineResponse200299`

NewInlineResponse200299 instantiates a new InlineResponse200299 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200299WithDefaults

`func NewInlineResponse200299WithDefaults() *InlineResponse200299`

NewInlineResponse200299WithDefaults instantiates a new InlineResponse200299 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *InlineResponse200299) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200299) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200299) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200299) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCount

`func (o *InlineResponse200299) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200299) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200299) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse200299) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetUsage

`func (o *InlineResponse200299) GetUsage() OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InlineResponse200299) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InlineResponse200299) SetUsage(v OrganizationsOrganizationIdSummaryTopDevicesModelsByUsageUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InlineResponse200299) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


