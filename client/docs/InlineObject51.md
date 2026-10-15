# InlineObject51

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20069DefaultDestinations**](InlineResponse20069DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20069Alerts**](InlineResponse20069Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20069Muting**](InlineResponse20069Muting.md) |  | [optional] 

## Methods

### NewInlineObject51

`func NewInlineObject51() *InlineObject51`

NewInlineObject51 instantiates a new InlineObject51 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject51WithDefaults

`func NewInlineObject51WithDefaults() *InlineObject51`

NewInlineObject51WithDefaults instantiates a new InlineObject51 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject51) GetDefaultDestinations() InlineResponse20069DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject51) GetDefaultDestinationsOk() (*InlineResponse20069DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject51) SetDefaultDestinations(v InlineResponse20069DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject51) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject51) GetAlerts() []InlineResponse20069Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject51) GetAlertsOk() (*[]InlineResponse20069Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject51) SetAlerts(v []InlineResponse20069Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject51) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject51) GetMuting() InlineResponse20069Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject51) GetMutingOk() (*InlineResponse20069Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject51) SetMuting(v InlineResponse20069Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject51) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


