# InlineObject180

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanningEnabled** | Pointer to **bool** | Whether APs will scan for Bluetooth enabled clients. | [optional] 
**AdvertisingEnabled** | Pointer to **bool** | Whether APs will advertise beacons. | [optional] 
**Uuid** | Pointer to **string** | The UUID to be used in the beacon identifier. | [optional] 
**MajorMinorAssignmentMode** | Pointer to **string** | The way major and minor number should be assigned to nodes in the network. (&#39;Unique&#39;, &#39;Non-unique&#39;) | [optional] 
**Major** | Pointer to **int32** | The major number to be used in the beacon identifier. Only valid in &#39;Non-unique&#39; mode. | [optional] 
**Minor** | Pointer to **int32** | The minor number to be used in the beacon identifier. Only valid in &#39;Non-unique&#39; mode. | [optional] 

## Methods

### NewInlineObject180

`func NewInlineObject180() *InlineObject180`

NewInlineObject180 instantiates a new InlineObject180 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject180WithDefaults

`func NewInlineObject180WithDefaults() *InlineObject180`

NewInlineObject180WithDefaults instantiates a new InlineObject180 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanningEnabled

`func (o *InlineObject180) GetScanningEnabled() bool`

GetScanningEnabled returns the ScanningEnabled field if non-nil, zero value otherwise.

### GetScanningEnabledOk

`func (o *InlineObject180) GetScanningEnabledOk() (*bool, bool)`

GetScanningEnabledOk returns a tuple with the ScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanningEnabled

`func (o *InlineObject180) SetScanningEnabled(v bool)`

SetScanningEnabled sets ScanningEnabled field to given value.

### HasScanningEnabled

`func (o *InlineObject180) HasScanningEnabled() bool`

HasScanningEnabled returns a boolean if a field has been set.

### GetAdvertisingEnabled

`func (o *InlineObject180) GetAdvertisingEnabled() bool`

GetAdvertisingEnabled returns the AdvertisingEnabled field if non-nil, zero value otherwise.

### GetAdvertisingEnabledOk

`func (o *InlineObject180) GetAdvertisingEnabledOk() (*bool, bool)`

GetAdvertisingEnabledOk returns a tuple with the AdvertisingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisingEnabled

`func (o *InlineObject180) SetAdvertisingEnabled(v bool)`

SetAdvertisingEnabled sets AdvertisingEnabled field to given value.

### HasAdvertisingEnabled

`func (o *InlineObject180) HasAdvertisingEnabled() bool`

HasAdvertisingEnabled returns a boolean if a field has been set.

### GetUuid

`func (o *InlineObject180) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InlineObject180) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InlineObject180) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InlineObject180) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMajorMinorAssignmentMode

`func (o *InlineObject180) GetMajorMinorAssignmentMode() string`

GetMajorMinorAssignmentMode returns the MajorMinorAssignmentMode field if non-nil, zero value otherwise.

### GetMajorMinorAssignmentModeOk

`func (o *InlineObject180) GetMajorMinorAssignmentModeOk() (*string, bool)`

GetMajorMinorAssignmentModeOk returns a tuple with the MajorMinorAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorMinorAssignmentMode

`func (o *InlineObject180) SetMajorMinorAssignmentMode(v string)`

SetMajorMinorAssignmentMode sets MajorMinorAssignmentMode field to given value.

### HasMajorMinorAssignmentMode

`func (o *InlineObject180) HasMajorMinorAssignmentMode() bool`

HasMajorMinorAssignmentMode returns a boolean if a field has been set.

### GetMajor

`func (o *InlineObject180) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *InlineObject180) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *InlineObject180) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *InlineObject180) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *InlineObject180) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *InlineObject180) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *InlineObject180) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *InlineObject180) HasMinor() bool`

HasMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


