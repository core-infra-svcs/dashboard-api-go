# InlineResponse200103Gnss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | GNSS status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;not applicable&#39;, &#39;canceled&#39; | [optional] 
**Completed** | Pointer to [**InlineResponse200103GnssCompleted**](InlineResponse200103GnssCompleted.md) |  | [optional] 

## Methods

### NewInlineResponse200103Gnss

`func NewInlineResponse200103Gnss() *InlineResponse200103Gnss`

NewInlineResponse200103Gnss instantiates a new InlineResponse200103Gnss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200103GnssWithDefaults

`func NewInlineResponse200103GnssWithDefaults() *InlineResponse200103Gnss`

NewInlineResponse200103GnssWithDefaults instantiates a new InlineResponse200103Gnss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200103Gnss) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200103Gnss) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200103Gnss) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200103Gnss) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200103Gnss) GetCompleted() InlineResponse200103GnssCompleted`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200103Gnss) GetCompletedOk() (*InlineResponse200103GnssCompleted, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200103Gnss) SetCompleted(v InlineResponse200103GnssCompleted)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200103Gnss) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


