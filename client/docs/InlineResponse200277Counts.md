# InlineResponse200277Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200277CountsJobs**](InlineResponse200277CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse200277CountsByJobOperation**](InlineResponse200277CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse200277Counts

`func NewInlineResponse200277Counts() *InlineResponse200277Counts`

NewInlineResponse200277Counts instantiates a new InlineResponse200277Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200277CountsWithDefaults

`func NewInlineResponse200277CountsWithDefaults() *InlineResponse200277Counts`

NewInlineResponse200277CountsWithDefaults instantiates a new InlineResponse200277Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse200277Counts) GetJobs() InlineResponse200277CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse200277Counts) GetJobsOk() (*InlineResponse200277CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse200277Counts) SetJobs(v InlineResponse200277CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse200277Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse200277Counts) GetByJobOperation() []InlineResponse200277CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse200277Counts) GetByJobOperationOk() (*[]InlineResponse200277CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse200277Counts) SetByJobOperation(v []InlineResponse200277CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse200277Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


