# OrganizationsOrganizationIdDevicesCellularDataProfilesCap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **int32** | The total bandwidth available for the specified term in megabytes. | 
**Threshold** | Pointer to **float32** | The data usage threshold at which the rule should be triggered. This is expressed as a float between 0.01 and 1.0 (where 0.01 is equal to 1% of total and 1.0 is equal to 100%). By default this will be 1.0. | [optional] 
**Term** | [**OrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm**](OrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm.md) |  | 

## Methods

### NewOrganizationsOrganizationIdDevicesCellularDataProfilesCap

`func NewOrganizationsOrganizationIdDevicesCellularDataProfilesCap(value int32, term OrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm, ) *OrganizationsOrganizationIdDevicesCellularDataProfilesCap`

NewOrganizationsOrganizationIdDevicesCellularDataProfilesCap instantiates a new OrganizationsOrganizationIdDevicesCellularDataProfilesCap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdDevicesCellularDataProfilesCapWithDefaults

`func NewOrganizationsOrganizationIdDevicesCellularDataProfilesCapWithDefaults() *OrganizationsOrganizationIdDevicesCellularDataProfilesCap`

NewOrganizationsOrganizationIdDevicesCellularDataProfilesCapWithDefaults instantiates a new OrganizationsOrganizationIdDevicesCellularDataProfilesCap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) SetValue(v int32)`

SetValue sets Value field to given value.


### GetThreshold

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTerm

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) GetTerm() OrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) GetTermOk() (*OrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesCap) SetTerm(v OrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm)`

SetTerm sets Term field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


