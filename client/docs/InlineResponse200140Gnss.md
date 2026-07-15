# InlineResponse200140Gnss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | GNSS status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;not applicable&#39;, &#39;canceled&#39; | [optional] 
**Completed** | Pointer to [**InlineResponse200140GnssCompleted**](InlineResponse200140GnssCompleted.md) |  | [optional] 

## Methods

### NewInlineResponse200140Gnss

`func NewInlineResponse200140Gnss() *InlineResponse200140Gnss`

NewInlineResponse200140Gnss instantiates a new InlineResponse200140Gnss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200140GnssWithDefaults

`func NewInlineResponse200140GnssWithDefaults() *InlineResponse200140Gnss`

NewInlineResponse200140GnssWithDefaults instantiates a new InlineResponse200140Gnss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200140Gnss) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200140Gnss) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200140Gnss) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200140Gnss) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200140Gnss) GetCompleted() InlineResponse200140GnssCompleted`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200140Gnss) GetCompletedOk() (*InlineResponse200140GnssCompleted, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200140Gnss) SetCompleted(v InlineResponse200140GnssCompleted)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200140Gnss) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


