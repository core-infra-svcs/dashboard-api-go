# InlineResponse20052

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20052DefaultDestinations**](InlineResponse20052DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20052Alerts**](InlineResponse20052Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20052Muting**](InlineResponse20052Muting.md) |  | [optional] 

## Methods

### NewInlineResponse20052

`func NewInlineResponse20052() *InlineResponse20052`

NewInlineResponse20052 instantiates a new InlineResponse20052 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20052WithDefaults

`func NewInlineResponse20052WithDefaults() *InlineResponse20052`

NewInlineResponse20052WithDefaults instantiates a new InlineResponse20052 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineResponse20052) GetDefaultDestinations() InlineResponse20052DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineResponse20052) GetDefaultDestinationsOk() (*InlineResponse20052DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineResponse20052) SetDefaultDestinations(v InlineResponse20052DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineResponse20052) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineResponse20052) GetAlerts() []InlineResponse20052Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineResponse20052) GetAlertsOk() (*[]InlineResponse20052Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineResponse20052) SetAlerts(v []InlineResponse20052Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineResponse20052) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineResponse20052) GetMuting() InlineResponse20052Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineResponse20052) GetMutingOk() (*InlineResponse20052Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineResponse20052) SetMuting(v InlineResponse20052Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineResponse20052) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


