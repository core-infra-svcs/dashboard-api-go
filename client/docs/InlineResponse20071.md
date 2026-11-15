# InlineResponse20071

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20071DefaultDestinations**](InlineResponse20071DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20071Alerts**](InlineResponse20071Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20071Muting**](InlineResponse20071Muting.md) |  | [optional] 

## Methods

### NewInlineResponse20071

`func NewInlineResponse20071() *InlineResponse20071`

NewInlineResponse20071 instantiates a new InlineResponse20071 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20071WithDefaults

`func NewInlineResponse20071WithDefaults() *InlineResponse20071`

NewInlineResponse20071WithDefaults instantiates a new InlineResponse20071 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineResponse20071) GetDefaultDestinations() InlineResponse20071DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineResponse20071) GetDefaultDestinationsOk() (*InlineResponse20071DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineResponse20071) SetDefaultDestinations(v InlineResponse20071DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineResponse20071) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineResponse20071) GetAlerts() []InlineResponse20071Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineResponse20071) GetAlertsOk() (*[]InlineResponse20071Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineResponse20071) SetAlerts(v []InlineResponse20071Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineResponse20071) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineResponse20071) GetMuting() InlineResponse20071Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineResponse20071) GetMutingOk() (*InlineResponse20071Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineResponse20071) SetMuting(v InlineResponse20071Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineResponse20071) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


