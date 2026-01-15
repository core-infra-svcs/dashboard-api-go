# InlineResponse200160ArpInspection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable Dynamic ARP Inspection on the network. Default value is false. | [optional] 
**UnsupportedModels** | Pointer to **[]string** | List of switch models that does not support dynamic ARP inspection | [optional] 

## Methods

### NewInlineResponse200160ArpInspection

`func NewInlineResponse200160ArpInspection() *InlineResponse200160ArpInspection`

NewInlineResponse200160ArpInspection instantiates a new InlineResponse200160ArpInspection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200160ArpInspectionWithDefaults

`func NewInlineResponse200160ArpInspectionWithDefaults() *InlineResponse200160ArpInspection`

NewInlineResponse200160ArpInspectionWithDefaults instantiates a new InlineResponse200160ArpInspection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200160ArpInspection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200160ArpInspection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200160ArpInspection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200160ArpInspection) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUnsupportedModels

`func (o *InlineResponse200160ArpInspection) GetUnsupportedModels() []string`

GetUnsupportedModels returns the UnsupportedModels field if non-nil, zero value otherwise.

### GetUnsupportedModelsOk

`func (o *InlineResponse200160ArpInspection) GetUnsupportedModelsOk() (*[]string, bool)`

GetUnsupportedModelsOk returns a tuple with the UnsupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedModels

`func (o *InlineResponse200160ArpInspection) SetUnsupportedModels(v []string)`

SetUnsupportedModels sets UnsupportedModels field to given value.

### HasUnsupportedModels

`func (o *InlineResponse200160ArpInspection) HasUnsupportedModels() bool`

HasUnsupportedModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


