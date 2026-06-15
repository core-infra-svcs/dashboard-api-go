# InlineResponse200266Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200266CountsJobs**](InlineResponse200266CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse200266CountsByJobOperation**](InlineResponse200266CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse200266Counts

`func NewInlineResponse200266Counts() *InlineResponse200266Counts`

NewInlineResponse200266Counts instantiates a new InlineResponse200266Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200266CountsWithDefaults

`func NewInlineResponse200266CountsWithDefaults() *InlineResponse200266Counts`

NewInlineResponse200266CountsWithDefaults instantiates a new InlineResponse200266Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse200266Counts) GetJobs() InlineResponse200266CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse200266Counts) GetJobsOk() (*InlineResponse200266CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse200266Counts) SetJobs(v InlineResponse200266CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse200266Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse200266Counts) GetByJobOperation() []InlineResponse200266CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse200266Counts) GetByJobOperationOk() (*[]InlineResponse200266CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse200266Counts) SetByJobOperation(v []InlineResponse200266CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse200266Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


