# InlineObject44

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20063DefaultDestinations**](InlineResponse20063DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20063Alerts**](InlineResponse20063Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20063Muting**](InlineResponse20063Muting.md) |  | [optional] 

## Methods

### NewInlineObject44

`func NewInlineObject44() *InlineObject44`

NewInlineObject44 instantiates a new InlineObject44 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject44WithDefaults

`func NewInlineObject44WithDefaults() *InlineObject44`

NewInlineObject44WithDefaults instantiates a new InlineObject44 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject44) GetDefaultDestinations() InlineResponse20063DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject44) GetDefaultDestinationsOk() (*InlineResponse20063DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject44) SetDefaultDestinations(v InlineResponse20063DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject44) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject44) GetAlerts() []InlineResponse20063Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject44) GetAlertsOk() (*[]InlineResponse20063Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject44) SetAlerts(v []InlineResponse20063Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject44) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject44) GetMuting() InlineResponse20063Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject44) GetMutingOk() (*InlineResponse20063Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject44) SetMuting(v InlineResponse20063Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject44) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


