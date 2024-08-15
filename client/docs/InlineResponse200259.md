# InlineResponse200259

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | License status (Co-termination licensing only) | [optional] 
**ExpirationDate** | Pointer to **string** | License expiration date (Co-termination licensing only) | [optional] 
**LicensedDeviceCounts** | Pointer to **map[string]int32** | License counts (Co-termination licensing only) | [optional] 
**LicenseCount** | Pointer to **int32** | Total number of licenses (Per-device licensing only) | [optional] 
**States** | Pointer to [**InlineResponse200259States**](InlineResponse200259States.md) |  | [optional] 
**LicenseTypes** | Pointer to [**[]InlineResponse200259LicenseTypes**](InlineResponse200259LicenseTypes.md) | Data by license type (Per-device licensing only) | [optional] 
**SystemsManager** | Pointer to [**InlineResponse200259SystemsManager**](InlineResponse200259SystemsManager.md) |  | [optional] 

## Methods

### NewInlineResponse200259

`func NewInlineResponse200259() *InlineResponse200259`

NewInlineResponse200259 instantiates a new InlineResponse200259 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200259WithDefaults

`func NewInlineResponse200259WithDefaults() *InlineResponse200259`

NewInlineResponse200259WithDefaults instantiates a new InlineResponse200259 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200259) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200259) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200259) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200259) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InlineResponse200259) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InlineResponse200259) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InlineResponse200259) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InlineResponse200259) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicensedDeviceCounts

`func (o *InlineResponse200259) GetLicensedDeviceCounts() map[string]int32`

GetLicensedDeviceCounts returns the LicensedDeviceCounts field if non-nil, zero value otherwise.

### GetLicensedDeviceCountsOk

`func (o *InlineResponse200259) GetLicensedDeviceCountsOk() (*map[string]int32, bool)`

GetLicensedDeviceCountsOk returns a tuple with the LicensedDeviceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedDeviceCounts

`func (o *InlineResponse200259) SetLicensedDeviceCounts(v map[string]int32)`

SetLicensedDeviceCounts sets LicensedDeviceCounts field to given value.

### HasLicensedDeviceCounts

`func (o *InlineResponse200259) HasLicensedDeviceCounts() bool`

HasLicensedDeviceCounts returns a boolean if a field has been set.

### GetLicenseCount

`func (o *InlineResponse200259) GetLicenseCount() int32`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *InlineResponse200259) GetLicenseCountOk() (*int32, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *InlineResponse200259) SetLicenseCount(v int32)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *InlineResponse200259) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetStates

`func (o *InlineResponse200259) GetStates() InlineResponse200259States`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *InlineResponse200259) GetStatesOk() (*InlineResponse200259States, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *InlineResponse200259) SetStates(v InlineResponse200259States)`

SetStates sets States field to given value.

### HasStates

`func (o *InlineResponse200259) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetLicenseTypes

`func (o *InlineResponse200259) GetLicenseTypes() []InlineResponse200259LicenseTypes`

GetLicenseTypes returns the LicenseTypes field if non-nil, zero value otherwise.

### GetLicenseTypesOk

`func (o *InlineResponse200259) GetLicenseTypesOk() (*[]InlineResponse200259LicenseTypes, bool)`

GetLicenseTypesOk returns a tuple with the LicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypes

`func (o *InlineResponse200259) SetLicenseTypes(v []InlineResponse200259LicenseTypes)`

SetLicenseTypes sets LicenseTypes field to given value.

### HasLicenseTypes

`func (o *InlineResponse200259) HasLicenseTypes() bool`

HasLicenseTypes returns a boolean if a field has been set.

### GetSystemsManager

`func (o *InlineResponse200259) GetSystemsManager() InlineResponse200259SystemsManager`

GetSystemsManager returns the SystemsManager field if non-nil, zero value otherwise.

### GetSystemsManagerOk

`func (o *InlineResponse200259) GetSystemsManagerOk() (*InlineResponse200259SystemsManager, bool)`

GetSystemsManagerOk returns a tuple with the SystemsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsManager

`func (o *InlineResponse200259) SetSystemsManager(v InlineResponse200259SystemsManager)`

SetSystemsManager sets SystemsManager field to given value.

### HasSystemsManager

`func (o *InlineResponse200259) HasSystemsManager() bool`

HasSystemsManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


