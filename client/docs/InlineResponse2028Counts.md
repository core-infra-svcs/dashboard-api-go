# InlineResponse2028Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**InlineResponse200277CountsJobs**](InlineResponse200277CountsJobs.md) |  | [optional] 
**ByJobOperation** | Pointer to [**[]InlineResponse2028CountsByJobOperation**](InlineResponse2028CountsByJobOperation.md) | Job count breakdown by pipeline job operation. Contains at most one entry per pipeline job operation with matching jobs. New registered operations may add entries over time. | [optional] 

## Methods

### NewInlineResponse2028Counts

`func NewInlineResponse2028Counts() *InlineResponse2028Counts`

NewInlineResponse2028Counts instantiates a new InlineResponse2028Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2028CountsWithDefaults

`func NewInlineResponse2028CountsWithDefaults() *InlineResponse2028Counts`

NewInlineResponse2028CountsWithDefaults instantiates a new InlineResponse2028Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *InlineResponse2028Counts) GetJobs() InlineResponse200277CountsJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *InlineResponse2028Counts) GetJobsOk() (*InlineResponse200277CountsJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *InlineResponse2028Counts) SetJobs(v InlineResponse200277CountsJobs)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *InlineResponse2028Counts) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetByJobOperation

`func (o *InlineResponse2028Counts) GetByJobOperation() []InlineResponse2028CountsByJobOperation`

GetByJobOperation returns the ByJobOperation field if non-nil, zero value otherwise.

### GetByJobOperationOk

`func (o *InlineResponse2028Counts) GetByJobOperationOk() (*[]InlineResponse2028CountsByJobOperation, bool)`

GetByJobOperationOk returns a tuple with the ByJobOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByJobOperation

`func (o *InlineResponse2028Counts) SetByJobOperation(v []InlineResponse2028CountsByJobOperation)`

SetByJobOperation sets ByJobOperation field to given value.

### HasByJobOperation

`func (o *InlineResponse2028Counts) HasByJobOperation() bool`

HasByJobOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


