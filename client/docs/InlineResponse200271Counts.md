# InlineResponse200271Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200271CountsJobs**](InlineResponse200271CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse200271CountsByJobOperation**](InlineResponse200271CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse200271Counts

`func NewInlineResponse200271Counts() *InlineResponse200271Counts`

NewInlineResponse200271Counts instantiates a new InlineResponse200271Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200271CountsWithDefaults

`func NewInlineResponse200271CountsWithDefaults() *InlineResponse200271Counts`

NewInlineResponse200271CountsWithDefaults instantiates a new InlineResponse200271Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse200271Counts) GetJobs() InlineResponse200271CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse200271Counts) GetJobsOk() (*InlineResponse200271CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse200271Counts) SetJobs(v InlineResponse200271CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse200271Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse200271Counts) GetByJobOperation() []InlineResponse200271CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse200271Counts) GetByJobOperationOk() (*[]InlineResponse200271CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse200271Counts) SetByJobOperation(v []InlineResponse200271CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse200271Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


