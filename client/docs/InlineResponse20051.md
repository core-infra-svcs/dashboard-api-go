# InlineResponse20051

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20051DefaultDestinations**](InlineResponse20051DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20051Alerts**](InlineResponse20051Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20051Muting**](InlineResponse20051Muting.md) |  | [optional] 

## Methods

### NewInlineResponse20051

`func NewInlineResponse20051() *InlineResponse20051`

NewInlineResponse20051 instantiates a new InlineResponse20051 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20051WithDefaults

`func NewInlineResponse20051WithDefaults() *InlineResponse20051`

NewInlineResponse20051WithDefaults instantiates a new InlineResponse20051 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineResponse20051) GetDefaultDestinations() InlineResponse20051DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineResponse20051) GetDefaultDestinationsOk() (*InlineResponse20051DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineResponse20051) SetDefaultDestinations(v InlineResponse20051DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineResponse20051) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineResponse20051) GetAlerts() []InlineResponse20051Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineResponse20051) GetAlertsOk() (*[]InlineResponse20051Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineResponse20051) SetAlerts(v []InlineResponse20051Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineResponse20051) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineResponse20051) GetMuting() InlineResponse20051Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineResponse20051) GetMutingOk() (*InlineResponse20051Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineResponse20051) SetMuting(v InlineResponse20051Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineResponse20051) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


