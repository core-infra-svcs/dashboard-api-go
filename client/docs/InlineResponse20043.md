# InlineResponse20043

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryId** | Pointer to **string** | Id of the route table summary request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your routing table summary request. | [optional] 
**Status** | Pointer to **string** | Status of the routing table summary request. | [optional] 
**Counts** | Pointer to [**InlineResponse20043Counts**](InlineResponse20043Counts.md) |  | [optional] 
**Errors** | Pointer to **[]string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20043

`func NewInlineResponse20043() *InlineResponse20043`

NewInlineResponse20043 instantiates a new InlineResponse20043 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043WithDefaults

`func NewInlineResponse20043WithDefaults() *InlineResponse20043`

NewInlineResponse20043WithDefaults instantiates a new InlineResponse20043 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryId

`func (o *InlineResponse20043) GetSummaryId() string`

GetSummaryId returns the SummaryId field if non-nil, zero value otherwise.

### GetSummaryIdOk

`func (o *InlineResponse20043) GetSummaryIdOk() (*string, bool)`

GetSummaryIdOk returns a tuple with the SummaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryId

`func (o *InlineResponse20043) SetSummaryId(v string)`

SetSummaryId sets SummaryId field to given value.

### HasSummaryId

`func (o *InlineResponse20043) HasSummaryId() bool`

HasSummaryId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20043) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20043) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20043) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20043) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20043) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20043) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20043) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20043) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse20043) GetCounts() InlineResponse20043Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse20043) GetCountsOk() (*InlineResponse20043Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse20043) SetCounts(v InlineResponse20043Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse20043) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20043) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20043) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20043) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20043) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


