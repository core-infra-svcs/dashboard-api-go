# InlineResponse200267Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200267CountsJobs**](InlineResponse200267CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse200267CountsByJobOperation**](InlineResponse200267CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse200267Counts

`func NewInlineResponse200267Counts() *InlineResponse200267Counts`

NewInlineResponse200267Counts instantiates a new InlineResponse200267Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200267CountsWithDefaults

`func NewInlineResponse200267CountsWithDefaults() *InlineResponse200267Counts`

NewInlineResponse200267CountsWithDefaults instantiates a new InlineResponse200267Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse200267Counts) GetJobs() InlineResponse200267CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse200267Counts) GetJobsOk() (*InlineResponse200267CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse200267Counts) SetJobs(v InlineResponse200267CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse200267Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse200267Counts) GetByJobOperation() []InlineResponse200267CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse200267Counts) GetByJobOperationOk() (*[]InlineResponse200267CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse200267Counts) SetByJobOperation(v []InlineResponse200267CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse200267Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


