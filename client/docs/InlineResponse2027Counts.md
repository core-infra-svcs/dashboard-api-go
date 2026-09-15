# InlineResponse2027Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200271CountsJobs**](InlineResponse200271CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse2027CountsByJobOperation**](InlineResponse2027CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse2027Counts

`func NewInlineResponse2027Counts() *InlineResponse2027Counts`

NewInlineResponse2027Counts instantiates a new InlineResponse2027Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2027CountsWithDefaults

`func NewInlineResponse2027CountsWithDefaults() *InlineResponse2027Counts`

NewInlineResponse2027CountsWithDefaults instantiates a new InlineResponse2027Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse2027Counts) GetJobs() InlineResponse200271CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse2027Counts) GetJobsOk() (*InlineResponse200271CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse2027Counts) SetJobs(v InlineResponse200271CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse2027Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse2027Counts) GetByJobOperation() []InlineResponse2027CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse2027Counts) GetByJobOperationOk() (*[]InlineResponse2027CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse2027Counts) SetByJobOperation(v []InlineResponse2027CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse2027Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


