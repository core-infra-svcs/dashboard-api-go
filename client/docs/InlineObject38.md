# InlineObject38

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20046DefaultDestinations**](InlineResponse20046DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20046Alerts**](InlineResponse20046Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20046Muting**](InlineResponse20046Muting.md) |  | [optional] 

## Methods

### NewInlineObject38

`func NewInlineObject38() *InlineObject38`

NewInlineObject38 instantiates a new InlineObject38 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject38WithDefaults

`func NewInlineObject38WithDefaults() *InlineObject38`

NewInlineObject38WithDefaults instantiates a new InlineObject38 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject38) GetDefaultDestinations() InlineResponse20046DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject38) GetDefaultDestinationsOk() (*InlineResponse20046DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject38) SetDefaultDestinations(v InlineResponse20046DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject38) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject38) GetAlerts() []InlineResponse20046Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject38) GetAlertsOk() (*[]InlineResponse20046Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject38) SetAlerts(v []InlineResponse20046Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject38) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject38) GetMuting() InlineResponse20046Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject38) GetMutingOk() (*InlineResponse20046Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject38) SetMuting(v InlineResponse20046Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject38) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


