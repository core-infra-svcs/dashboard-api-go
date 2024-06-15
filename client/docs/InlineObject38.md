# InlineObject38

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20045DefaultDestinations**](InlineResponse20045DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20045Alerts**](InlineResponse20045Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20045Muting**](InlineResponse20045Muting.md) |  | [optional] 

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

`func (o *InlineObject38) GetDefaultDestinations() InlineResponse20045DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject38) GetDefaultDestinationsOk() (*InlineResponse20045DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject38) SetDefaultDestinations(v InlineResponse20045DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject38) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject38) GetAlerts() []InlineResponse20045Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject38) GetAlertsOk() (*[]InlineResponse20045Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject38) SetAlerts(v []InlineResponse20045Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject38) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject38) GetMuting() InlineResponse20045Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject38) GetMutingOk() (*InlineResponse20045Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject38) SetMuting(v InlineResponse20045Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject38) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


