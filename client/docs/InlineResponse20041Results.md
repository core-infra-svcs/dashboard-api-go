# InlineResponse20041Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instant** | Pointer to **float32** | Instantaneous device power usage measured when this job result was generated. | [optional] 
**Peak** | Pointer to **float32** | Highest device power usage observed since the last reboot. | [optional] 
**Budget** | Pointer to **float32** | Total power reserved for the device, defined by hardware and system specifications. | [optional] 

## Methods

### NewInlineResponse20041Results

`func NewInlineResponse20041Results() *InlineResponse20041Results`

NewInlineResponse20041Results instantiates a new InlineResponse20041Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20041ResultsWithDefaults

`func NewInlineResponse20041ResultsWithDefaults() *InlineResponse20041Results`

NewInlineResponse20041ResultsWithDefaults instantiates a new InlineResponse20041Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstant

`func (o *InlineResponse20041Results) GetInstant() float32`

GetInstant returns the Instant field if non-nil, zero value otherwise.

### GetInstantOk

`func (o *InlineResponse20041Results) GetInstantOk() (*float32, bool)`

GetInstantOk returns a tuple with the Instant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstant

`func (o *InlineResponse20041Results) SetInstant(v float32)`

SetInstant sets Instant field to given value.

### HasInstant

`func (o *InlineResponse20041Results) HasInstant() bool`

HasInstant returns a boolean if a field has been set.

### GetPeak

`func (o *InlineResponse20041Results) GetPeak() float32`

GetPeak returns the Peak field if non-nil, zero value otherwise.

### GetPeakOk

`func (o *InlineResponse20041Results) GetPeakOk() (*float32, bool)`

GetPeakOk returns a tuple with the Peak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeak

`func (o *InlineResponse20041Results) SetPeak(v float32)`

SetPeak sets Peak field to given value.

### HasPeak

`func (o *InlineResponse20041Results) HasPeak() bool`

HasPeak returns a boolean if a field has been set.

### GetBudget

`func (o *InlineResponse20041Results) GetBudget() float32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *InlineResponse20041Results) GetBudgetOk() (*float32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *InlineResponse20041Results) SetBudget(v float32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *InlineResponse20041Results) HasBudget() bool`

HasBudget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


