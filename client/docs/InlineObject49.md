# InlineObject49

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20068DefaultDestinations**](InlineResponse20068DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20068Alerts**](InlineResponse20068Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20068Muting**](InlineResponse20068Muting.md) |  | [optional] 

## Methods

### NewInlineObject49

`func NewInlineObject49() *InlineObject49`

NewInlineObject49 instantiates a new InlineObject49 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject49WithDefaults

`func NewInlineObject49WithDefaults() *InlineObject49`

NewInlineObject49WithDefaults instantiates a new InlineObject49 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject49) GetDefaultDestinations() InlineResponse20068DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject49) GetDefaultDestinationsOk() (*InlineResponse20068DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject49) SetDefaultDestinations(v InlineResponse20068DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject49) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject49) GetAlerts() []InlineResponse20068Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject49) GetAlertsOk() (*[]InlineResponse20068Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject49) SetAlerts(v []InlineResponse20068Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject49) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject49) GetMuting() InlineResponse20068Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject49) GetMutingOk() (*InlineResponse20068Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject49) SetMuting(v InlineResponse20068Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject49) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


