# InlineResponse200291

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | License status (Co-termination licensing only) | [optional] 
**ExpirationDate** | Pointer to **string** | License expiration date (Co-termination licensing only) | [optional] 
**LicensedDeviceCounts** | Pointer to **map[string]int32** | License counts (Co-termination licensing only) | [optional] 
**LicenseCount** | Pointer to **int32** | Total number of licenses (Per-device licensing only) | [optional] 
**States** | Pointer to [**InlineResponse200291States**](InlineResponse200291States.md) |  | [optional] 
**LicenseTypes** | Pointer to [**[]InlineResponse200291LicenseTypes**](InlineResponse200291LicenseTypes.md) | Data by license type (Per-device licensing only) | [optional] 
**SystemsManager** | Pointer to [**InlineResponse200291SystemsManager**](InlineResponse200291SystemsManager.md) |  | [optional] 

## Methods

### NewInlineResponse200291

`func NewInlineResponse200291() *InlineResponse200291`

NewInlineResponse200291 instantiates a new InlineResponse200291 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200291WithDefaults

`func NewInlineResponse200291WithDefaults() *InlineResponse200291`

NewInlineResponse200291WithDefaults instantiates a new InlineResponse200291 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200291) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200291) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200291) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200291) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InlineResponse200291) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InlineResponse200291) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InlineResponse200291) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InlineResponse200291) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicensedDeviceCounts

`func (o *InlineResponse200291) GetLicensedDeviceCounts() map[string]int32`

GetLicensedDeviceCounts returns the LicensedDeviceCounts field if non-nil, zero value otherwise.

### GetLicensedDeviceCountsOk

`func (o *InlineResponse200291) GetLicensedDeviceCountsOk() (*map[string]int32, bool)`

GetLicensedDeviceCountsOk returns a tuple with the LicensedDeviceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedDeviceCounts

`func (o *InlineResponse200291) SetLicensedDeviceCounts(v map[string]int32)`

SetLicensedDeviceCounts sets LicensedDeviceCounts field to given value.

### HasLicensedDeviceCounts

`func (o *InlineResponse200291) HasLicensedDeviceCounts() bool`

HasLicensedDeviceCounts returns a boolean if a field has been set.

### GetLicenseCount

`func (o *InlineResponse200291) GetLicenseCount() int32`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *InlineResponse200291) GetLicenseCountOk() (*int32, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *InlineResponse200291) SetLicenseCount(v int32)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *InlineResponse200291) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetStates

`func (o *InlineResponse200291) GetStates() InlineResponse200291States`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *InlineResponse200291) GetStatesOk() (*InlineResponse200291States, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *InlineResponse200291) SetStates(v InlineResponse200291States)`

SetStates sets States field to given value.

### HasStates

`func (o *InlineResponse200291) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetLicenseTypes

`func (o *InlineResponse200291) GetLicenseTypes() []InlineResponse200291LicenseTypes`

GetLicenseTypes returns the LicenseTypes field if non-nil, zero value otherwise.

### GetLicenseTypesOk

`func (o *InlineResponse200291) GetLicenseTypesOk() (*[]InlineResponse200291LicenseTypes, bool)`

GetLicenseTypesOk returns a tuple with the LicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypes

`func (o *InlineResponse200291) SetLicenseTypes(v []InlineResponse200291LicenseTypes)`

SetLicenseTypes sets LicenseTypes field to given value.

### HasLicenseTypes

`func (o *InlineResponse200291) HasLicenseTypes() bool`

HasLicenseTypes returns a boolean if a field has been set.

### GetSystemsManager

`func (o *InlineResponse200291) GetSystemsManager() InlineResponse200291SystemsManager`

GetSystemsManager returns the SystemsManager field if non-nil, zero value otherwise.

### GetSystemsManagerOk

`func (o *InlineResponse200291) GetSystemsManagerOk() (*InlineResponse200291SystemsManager, bool)`

GetSystemsManagerOk returns a tuple with the SystemsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsManager

`func (o *InlineResponse200291) SetSystemsManager(v InlineResponse200291SystemsManager)`

SetSystemsManager sets SystemsManager field to given value.

### HasSystemsManager

`func (o *InlineResponse200291) HasSystemsManager() bool`

HasSystemsManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


