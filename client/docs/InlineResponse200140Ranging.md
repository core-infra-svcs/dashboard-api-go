# InlineResponse200140Ranging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Ranging status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;no neighbors&#39; | [optional] 
**Completed** | Pointer to [**InlineResponse200140RangingCompleted**](InlineResponse200140RangingCompleted.md) |  | [optional] 

## Methods

### NewInlineResponse200140Ranging

`func NewInlineResponse200140Ranging() *InlineResponse200140Ranging`

NewInlineResponse200140Ranging instantiates a new InlineResponse200140Ranging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200140RangingWithDefaults

`func NewInlineResponse200140RangingWithDefaults() *InlineResponse200140Ranging`

NewInlineResponse200140RangingWithDefaults instantiates a new InlineResponse200140Ranging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200140Ranging) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200140Ranging) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200140Ranging) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200140Ranging) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200140Ranging) GetCompleted() InlineResponse200140RangingCompleted`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200140Ranging) GetCompletedOk() (*InlineResponse200140RangingCompleted, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200140Ranging) SetCompleted(v InlineResponse200140RangingCompleted)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200140Ranging) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


