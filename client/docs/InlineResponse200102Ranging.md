# InlineResponse200102Ranging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Ranging status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;no neighbors&#39; | [optional] 
**Completed** | Pointer to [**InlineResponse200102RangingCompleted**](InlineResponse200102RangingCompleted.md) |  | [optional] 

## Methods

### NewInlineResponse200102Ranging

`func NewInlineResponse200102Ranging() *InlineResponse200102Ranging`

NewInlineResponse200102Ranging instantiates a new InlineResponse200102Ranging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102RangingWithDefaults

`func NewInlineResponse200102RangingWithDefaults() *InlineResponse200102Ranging`

NewInlineResponse200102RangingWithDefaults instantiates a new InlineResponse200102Ranging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200102Ranging) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200102Ranging) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200102Ranging) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200102Ranging) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200102Ranging) GetCompleted() InlineResponse200102RangingCompleted`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200102Ranging) GetCompletedOk() (*InlineResponse200102RangingCompleted, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200102Ranging) SetCompleted(v InlineResponse200102RangingCompleted)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200102Ranging) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


