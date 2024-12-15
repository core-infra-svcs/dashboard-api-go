# InlineObject39

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20047DefaultDestinations**](InlineResponse20047DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20047Alerts**](InlineResponse20047Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20047Muting**](InlineResponse20047Muting.md) |  | [optional] 

## Methods

### NewInlineObject39

`func NewInlineObject39() *InlineObject39`

NewInlineObject39 instantiates a new InlineObject39 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject39WithDefaults

`func NewInlineObject39WithDefaults() *InlineObject39`

NewInlineObject39WithDefaults instantiates a new InlineObject39 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject39) GetDefaultDestinations() InlineResponse20047DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject39) GetDefaultDestinationsOk() (*InlineResponse20047DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject39) SetDefaultDestinations(v InlineResponse20047DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject39) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject39) GetAlerts() []InlineResponse20047Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject39) GetAlertsOk() (*[]InlineResponse20047Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject39) SetAlerts(v []InlineResponse20047Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject39) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject39) GetMuting() InlineResponse20047Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject39) GetMutingOk() (*InlineResponse20047Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject39) SetMuting(v InlineResponse20047Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject39) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


