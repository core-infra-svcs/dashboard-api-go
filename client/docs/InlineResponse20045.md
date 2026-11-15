# InlineResponse20045

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryId** | Pointer to **string** | Id of the route table summary request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your routing table summary request. | [optional] 
**Status** | Pointer to **string** | Status of the routing table summary request. | [optional] 
**Counts** | Pointer to [**InlineResponse20045Counts**](InlineResponse20045Counts.md) |  | [optional] 
**Errors** | Pointer to **[]string** | An error message for a failed execution | [optional] 

## Methods

### NewInlineResponse20045

`func NewInlineResponse20045() *InlineResponse20045`

NewInlineResponse20045 instantiates a new InlineResponse20045 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20045WithDefaults

`func NewInlineResponse20045WithDefaults() *InlineResponse20045`

NewInlineResponse20045WithDefaults instantiates a new InlineResponse20045 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryId

`func (o *InlineResponse20045) GetSummaryId() string`

GetSummaryId returns the SummaryId field if non-nil, zero value otherwise.

### GetSummaryIdOk

`func (o *InlineResponse20045) GetSummaryIdOk() (*string, bool)`

GetSummaryIdOk returns a tuple with the SummaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryId

`func (o *InlineResponse20045) SetSummaryId(v string)`

SetSummaryId sets SummaryId field to given value.

### HasSummaryId

`func (o *InlineResponse20045) HasSummaryId() bool`

HasSummaryId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20045) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20045) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20045) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20045) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20045) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20045) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20045) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20045) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse20045) GetCounts() InlineResponse20045Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse20045) GetCountsOk() (*InlineResponse20045Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse20045) SetCounts(v InlineResponse20045Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse20045) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20045) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20045) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20045) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20045) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


