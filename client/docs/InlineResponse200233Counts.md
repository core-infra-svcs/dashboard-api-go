# InlineResponse200233Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of alerts on the organization | 
**BySeverity** | [**[]InlineResponse200233CountsBySeverity**](InlineResponse200233CountsBySeverity.md) | Counts of alerts on organization by severity | 

## Methods

### NewInlineResponse200233Counts

`func NewInlineResponse200233Counts(total int32, bySeverity []InlineResponse200233CountsBySeverity, ) *InlineResponse200233Counts`

NewInlineResponse200233Counts instantiates a new InlineResponse200233Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200233CountsWithDefaults

`func NewInlineResponse200233CountsWithDefaults() *InlineResponse200233Counts`

NewInlineResponse200233CountsWithDefaults instantiates a new InlineResponse200233Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200233Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200233Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200233Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetBySeverity

`func (o *InlineResponse200233Counts) GetBySeverity() []InlineResponse200233CountsBySeverity`

GetBySeverity returns the BySeverity field if non-nil, zero value otherwise.

### GetBySeverityOk

`func (o *InlineResponse200233Counts) GetBySeverityOk() (*[]InlineResponse200233CountsBySeverity, bool)`

GetBySeverityOk returns a tuple with the BySeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySeverity

`func (o *InlineResponse200233Counts) SetBySeverity(v []InlineResponse200233CountsBySeverity)`

SetBySeverity sets BySeverity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


