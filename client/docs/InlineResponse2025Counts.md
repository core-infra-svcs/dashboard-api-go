# InlineResponse2025Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200271CountsJobs**](InlineResponse200271CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse2025CountsByJobOperation**](InlineResponse2025CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse2025Counts

`func NewInlineResponse2025Counts() *InlineResponse2025Counts`

NewInlineResponse2025Counts instantiates a new InlineResponse2025Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2025CountsWithDefaults

`func NewInlineResponse2025CountsWithDefaults() *InlineResponse2025Counts`

NewInlineResponse2025CountsWithDefaults instantiates a new InlineResponse2025Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse2025Counts) GetJobs() InlineResponse200271CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse2025Counts) GetJobsOk() (*InlineResponse200271CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse2025Counts) SetJobs(v InlineResponse200271CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse2025Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse2025Counts) GetByJobOperation() []InlineResponse2025CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse2025Counts) GetByJobOperationOk() (*[]InlineResponse2025CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse2025Counts) SetByJobOperation(v []InlineResponse2025CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse2025Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


