# InlineResponse200269Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200269CountsJobs**](InlineResponse200269CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse200269CountsByJobOperation**](InlineResponse200269CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse200269Counts

`func NewInlineResponse200269Counts() *InlineResponse200269Counts`

NewInlineResponse200269Counts instantiates a new InlineResponse200269Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200269CountsWithDefaults

`func NewInlineResponse200269CountsWithDefaults() *InlineResponse200269Counts`

NewInlineResponse200269CountsWithDefaults instantiates a new InlineResponse200269Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse200269Counts) GetJobs() InlineResponse200269CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse200269Counts) GetJobsOk() (*InlineResponse200269CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse200269Counts) SetJobs(v InlineResponse200269CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse200269Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse200269Counts) GetByJobOperation() []InlineResponse200269CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse200269Counts) GetByJobOperationOk() (*[]InlineResponse200269CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse200269Counts) SetByJobOperation(v []InlineResponse200269CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse200269Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


