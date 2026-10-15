# InlineResponse200274Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200274CountsJobs**](InlineResponse200274CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse200274CountsByJobOperation**](InlineResponse200274CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse200274Counts

`func NewInlineResponse200274Counts() *InlineResponse200274Counts`

NewInlineResponse200274Counts instantiates a new InlineResponse200274Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200274CountsWithDefaults

`func NewInlineResponse200274CountsWithDefaults() *InlineResponse200274Counts`

NewInlineResponse200274CountsWithDefaults instantiates a new InlineResponse200274Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse200274Counts) GetJobs() InlineResponse200274CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse200274Counts) GetJobsOk() (*InlineResponse200274CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse200274Counts) SetJobs(v InlineResponse200274CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse200274Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse200274Counts) GetByJobOperation() []InlineResponse200274CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse200274Counts) GetByJobOperationOk() (*[]InlineResponse200274CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse200274Counts) SetByJobOperation(v []InlineResponse200274CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse200274Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


