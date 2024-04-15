# InlineObject37

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20044DefaultDestinations**](InlineResponse20044DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20044Alerts**](InlineResponse20044Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20044Muting**](InlineResponse20044Muting.md) |  | [optional] 

## Methods

### NewInlineObject37

`func NewInlineObject37() *InlineObject37`

NewInlineObject37 instantiates a new InlineObject37 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject37WithDefaults

`func NewInlineObject37WithDefaults() *InlineObject37`

NewInlineObject37WithDefaults instantiates a new InlineObject37 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject37) GetDefaultDestinations() InlineResponse20044DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject37) GetDefaultDestinationsOk() (*InlineResponse20044DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject37) SetDefaultDestinations(v InlineResponse20044DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject37) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject37) GetAlerts() []InlineResponse20044Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject37) GetAlertsOk() (*[]InlineResponse20044Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject37) SetAlerts(v []InlineResponse20044Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject37) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject37) GetMuting() InlineResponse20044Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject37) GetMutingOk() (*InlineResponse20044Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject37) SetMuting(v InlineResponse20044Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject37) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


