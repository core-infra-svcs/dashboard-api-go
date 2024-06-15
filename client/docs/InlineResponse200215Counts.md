# InlineResponse200215Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of alerts on the organization | 
**BySeverity** | [**[]InlineResponse200215CountsBySeverity**](InlineResponse200215CountsBySeverity.md) | Counts of alerts on organization by severity | 

## Methods

### NewInlineResponse200215Counts

`func NewInlineResponse200215Counts(total int32, bySeverity []InlineResponse200215CountsBySeverity, ) *InlineResponse200215Counts`

NewInlineResponse200215Counts instantiates a new InlineResponse200215Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200215CountsWithDefaults

`func NewInlineResponse200215CountsWithDefaults() *InlineResponse200215Counts`

NewInlineResponse200215CountsWithDefaults instantiates a new InlineResponse200215Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200215Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200215Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200215Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetBySeverity

`func (o *InlineResponse200215Counts) GetBySeverity() []InlineResponse200215CountsBySeverity`

GetBySeverity returns the BySeverity field if non-nil, zero value otherwise.

### GetBySeverityOk

`func (o *InlineResponse200215Counts) GetBySeverityOk() (*[]InlineResponse200215CountsBySeverity, bool)`

GetBySeverityOk returns a tuple with the BySeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySeverity

`func (o *InlineResponse200215Counts) SetBySeverity(v []InlineResponse200215CountsBySeverity)`

SetBySeverity sets BySeverity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


