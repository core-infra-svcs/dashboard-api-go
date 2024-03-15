# InlineResponse207

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | The ID of the job that was used to create all of the device swaps. | [optional] 
**Swaps** | Pointer to [**[]InlineResponse207Swaps**](InlineResponse207Swaps.md) | An array of recent swap requests and their statuses. | [optional] 

## Methods

### NewInlineResponse207

`func NewInlineResponse207() *InlineResponse207`

NewInlineResponse207 instantiates a new InlineResponse207 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse207WithDefaults

`func NewInlineResponse207WithDefaults() *InlineResponse207`

NewInlineResponse207WithDefaults instantiates a new InlineResponse207 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *InlineResponse207) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *InlineResponse207) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *InlineResponse207) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *InlineResponse207) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSwaps

`func (o *InlineResponse207) GetSwaps() []InlineResponse207Swaps`

GetSwaps returns the Swaps field if non-nil, zero value otherwise.

### GetSwapsOk

`func (o *InlineResponse207) GetSwapsOk() (*[]InlineResponse207Swaps, bool)`

GetSwapsOk returns a tuple with the Swaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaps

`func (o *InlineResponse207) SetSwaps(v []InlineResponse207Swaps)`

SetSwaps sets Swaps field to given value.

### HasSwaps

`func (o *InlineResponse207) HasSwaps() bool`

HasSwaps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


