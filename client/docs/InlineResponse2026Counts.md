# InlineResponse2026Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200269CountsJobs**](InlineResponse200269CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse2026CountsByJobOperation**](InlineResponse2026CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse2026Counts

`func NewInlineResponse2026Counts() *InlineResponse2026Counts`

NewInlineResponse2026Counts instantiates a new InlineResponse2026Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2026CountsWithDefaults

`func NewInlineResponse2026CountsWithDefaults() *InlineResponse2026Counts`

NewInlineResponse2026CountsWithDefaults instantiates a new InlineResponse2026Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse2026Counts) GetJobs() InlineResponse200269CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse2026Counts) GetJobsOk() (*InlineResponse200269CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse2026Counts) SetJobs(v InlineResponse200269CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse2026Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse2026Counts) GetByJobOperation() []InlineResponse2026CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse2026Counts) GetByJobOperationOk() (*[]InlineResponse2026CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse2026Counts) SetByJobOperation(v []InlineResponse2026CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse2026Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


