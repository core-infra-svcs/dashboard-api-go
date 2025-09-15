# InlineResponse200292

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of the device | [optional] 
**Name** | Pointer to **string** | Name assigned to the device | [optional] 
**DeviceStatus** | Pointer to **string** | Status of the device upgrade | [optional] 
**CheckinFinishedAt** | Pointer to **time.Time** | The time the device checkin finished | [optional] 
**CheckinStartedAt** | Pointer to **time.Time** | The time the device checkin started | [optional] 
**DetailedStatus** | Pointer to **string** | The detailed status of the device upgrade | [optional] 
**DownloadFinishedAt** | Pointer to **time.Time** | The time the device upgrade download finished | [optional] 
**DownloadStartedAt** | Pointer to **time.Time** | The time the device upgrade download started | [optional] 
**DownloadStatus** | Pointer to **string** | The status of the device upgrade download | [optional] 
**InstallFinishedAt** | Pointer to **time.Time** | The time the device upgrade install finished | [optional] 
**InstallStartedAt** | Pointer to **time.Time** | The time the device upgrade install started | [optional] 
**InstallStatus** | Pointer to **string** | The status of the device upgrade install | [optional] 
**VerifyFinishedAt** | Pointer to **time.Time** | The time the device upgrade verification finished | [optional] 
**VerifyStartedAt** | Pointer to **time.Time** | The time the device upgrade verification started | [optional] 
**VerifyStatus** | Pointer to **string** | The status of the device upgrade verification | [optional] 
**Upgrade** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade**](OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade.md) |  | [optional] 

## Methods

### NewInlineResponse200292

`func NewInlineResponse200292() *InlineResponse200292`

NewInlineResponse200292 instantiates a new InlineResponse200292 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200292WithDefaults

`func NewInlineResponse200292WithDefaults() *InlineResponse200292`

NewInlineResponse200292WithDefaults instantiates a new InlineResponse200292 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200292) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200292) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200292) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200292) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200292) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200292) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200292) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200292) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *InlineResponse200292) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *InlineResponse200292) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *InlineResponse200292) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *InlineResponse200292) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetCheckinFinishedAt

`func (o *InlineResponse200292) GetCheckinFinishedAt() time.Time`

GetCheckinFinishedAt returns the CheckinFinishedAt field if non-nil, zero value otherwise.

### GetCheckinFinishedAtOk

`func (o *InlineResponse200292) GetCheckinFinishedAtOk() (*time.Time, bool)`

GetCheckinFinishedAtOk returns a tuple with the CheckinFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinFinishedAt

`func (o *InlineResponse200292) SetCheckinFinishedAt(v time.Time)`

SetCheckinFinishedAt sets CheckinFinishedAt field to given value.

### HasCheckinFinishedAt

`func (o *InlineResponse200292) HasCheckinFinishedAt() bool`

HasCheckinFinishedAt returns a boolean if a field has been set.

### GetCheckinStartedAt

`func (o *InlineResponse200292) GetCheckinStartedAt() time.Time`

GetCheckinStartedAt returns the CheckinStartedAt field if non-nil, zero value otherwise.

### GetCheckinStartedAtOk

`func (o *InlineResponse200292) GetCheckinStartedAtOk() (*time.Time, bool)`

GetCheckinStartedAtOk returns a tuple with the CheckinStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinStartedAt

`func (o *InlineResponse200292) SetCheckinStartedAt(v time.Time)`

SetCheckinStartedAt sets CheckinStartedAt field to given value.

### HasCheckinStartedAt

`func (o *InlineResponse200292) HasCheckinStartedAt() bool`

HasCheckinStartedAt returns a boolean if a field has been set.

### GetDetailedStatus

`func (o *InlineResponse200292) GetDetailedStatus() string`

GetDetailedStatus returns the DetailedStatus field if non-nil, zero value otherwise.

### GetDetailedStatusOk

`func (o *InlineResponse200292) GetDetailedStatusOk() (*string, bool)`

GetDetailedStatusOk returns a tuple with the DetailedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedStatus

`func (o *InlineResponse200292) SetDetailedStatus(v string)`

SetDetailedStatus sets DetailedStatus field to given value.

### HasDetailedStatus

`func (o *InlineResponse200292) HasDetailedStatus() bool`

HasDetailedStatus returns a boolean if a field has been set.

### GetDownloadFinishedAt

`func (o *InlineResponse200292) GetDownloadFinishedAt() time.Time`

GetDownloadFinishedAt returns the DownloadFinishedAt field if non-nil, zero value otherwise.

### GetDownloadFinishedAtOk

`func (o *InlineResponse200292) GetDownloadFinishedAtOk() (*time.Time, bool)`

GetDownloadFinishedAtOk returns a tuple with the DownloadFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFinishedAt

`func (o *InlineResponse200292) SetDownloadFinishedAt(v time.Time)`

SetDownloadFinishedAt sets DownloadFinishedAt field to given value.

### HasDownloadFinishedAt

`func (o *InlineResponse200292) HasDownloadFinishedAt() bool`

HasDownloadFinishedAt returns a boolean if a field has been set.

### GetDownloadStartedAt

`func (o *InlineResponse200292) GetDownloadStartedAt() time.Time`

GetDownloadStartedAt returns the DownloadStartedAt field if non-nil, zero value otherwise.

### GetDownloadStartedAtOk

`func (o *InlineResponse200292) GetDownloadStartedAtOk() (*time.Time, bool)`

GetDownloadStartedAtOk returns a tuple with the DownloadStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStartedAt

`func (o *InlineResponse200292) SetDownloadStartedAt(v time.Time)`

SetDownloadStartedAt sets DownloadStartedAt field to given value.

### HasDownloadStartedAt

`func (o *InlineResponse200292) HasDownloadStartedAt() bool`

HasDownloadStartedAt returns a boolean if a field has been set.

### GetDownloadStatus

`func (o *InlineResponse200292) GetDownloadStatus() string`

GetDownloadStatus returns the DownloadStatus field if non-nil, zero value otherwise.

### GetDownloadStatusOk

`func (o *InlineResponse200292) GetDownloadStatusOk() (*string, bool)`

GetDownloadStatusOk returns a tuple with the DownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStatus

`func (o *InlineResponse200292) SetDownloadStatus(v string)`

SetDownloadStatus sets DownloadStatus field to given value.

### HasDownloadStatus

`func (o *InlineResponse200292) HasDownloadStatus() bool`

HasDownloadStatus returns a boolean if a field has been set.

### GetInstallFinishedAt

`func (o *InlineResponse200292) GetInstallFinishedAt() time.Time`

GetInstallFinishedAt returns the InstallFinishedAt field if non-nil, zero value otherwise.

### GetInstallFinishedAtOk

`func (o *InlineResponse200292) GetInstallFinishedAtOk() (*time.Time, bool)`

GetInstallFinishedAtOk returns a tuple with the InstallFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallFinishedAt

`func (o *InlineResponse200292) SetInstallFinishedAt(v time.Time)`

SetInstallFinishedAt sets InstallFinishedAt field to given value.

### HasInstallFinishedAt

`func (o *InlineResponse200292) HasInstallFinishedAt() bool`

HasInstallFinishedAt returns a boolean if a field has been set.

### GetInstallStartedAt

`func (o *InlineResponse200292) GetInstallStartedAt() time.Time`

GetInstallStartedAt returns the InstallStartedAt field if non-nil, zero value otherwise.

### GetInstallStartedAtOk

`func (o *InlineResponse200292) GetInstallStartedAtOk() (*time.Time, bool)`

GetInstallStartedAtOk returns a tuple with the InstallStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStartedAt

`func (o *InlineResponse200292) SetInstallStartedAt(v time.Time)`

SetInstallStartedAt sets InstallStartedAt field to given value.

### HasInstallStartedAt

`func (o *InlineResponse200292) HasInstallStartedAt() bool`

HasInstallStartedAt returns a boolean if a field has been set.

### GetInstallStatus

`func (o *InlineResponse200292) GetInstallStatus() string`

GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.

### GetInstallStatusOk

`func (o *InlineResponse200292) GetInstallStatusOk() (*string, bool)`

GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStatus

`func (o *InlineResponse200292) SetInstallStatus(v string)`

SetInstallStatus sets InstallStatus field to given value.

### HasInstallStatus

`func (o *InlineResponse200292) HasInstallStatus() bool`

HasInstallStatus returns a boolean if a field has been set.

### GetVerifyFinishedAt

`func (o *InlineResponse200292) GetVerifyFinishedAt() time.Time`

GetVerifyFinishedAt returns the VerifyFinishedAt field if non-nil, zero value otherwise.

### GetVerifyFinishedAtOk

`func (o *InlineResponse200292) GetVerifyFinishedAtOk() (*time.Time, bool)`

GetVerifyFinishedAtOk returns a tuple with the VerifyFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyFinishedAt

`func (o *InlineResponse200292) SetVerifyFinishedAt(v time.Time)`

SetVerifyFinishedAt sets VerifyFinishedAt field to given value.

### HasVerifyFinishedAt

`func (o *InlineResponse200292) HasVerifyFinishedAt() bool`

HasVerifyFinishedAt returns a boolean if a field has been set.

### GetVerifyStartedAt

`func (o *InlineResponse200292) GetVerifyStartedAt() time.Time`

GetVerifyStartedAt returns the VerifyStartedAt field if non-nil, zero value otherwise.

### GetVerifyStartedAtOk

`func (o *InlineResponse200292) GetVerifyStartedAtOk() (*time.Time, bool)`

GetVerifyStartedAtOk returns a tuple with the VerifyStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyStartedAt

`func (o *InlineResponse200292) SetVerifyStartedAt(v time.Time)`

SetVerifyStartedAt sets VerifyStartedAt field to given value.

### HasVerifyStartedAt

`func (o *InlineResponse200292) HasVerifyStartedAt() bool`

HasVerifyStartedAt returns a boolean if a field has been set.

### GetVerifyStatus

`func (o *InlineResponse200292) GetVerifyStatus() string`

GetVerifyStatus returns the VerifyStatus field if non-nil, zero value otherwise.

### GetVerifyStatusOk

`func (o *InlineResponse200292) GetVerifyStatusOk() (*string, bool)`

GetVerifyStatusOk returns a tuple with the VerifyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyStatus

`func (o *InlineResponse200292) SetVerifyStatus(v string)`

SetVerifyStatus sets VerifyStatus field to given value.

### HasVerifyStatus

`func (o *InlineResponse200292) HasVerifyStatus() bool`

HasVerifyStatus returns a boolean if a field has been set.

### GetUpgrade

`func (o *InlineResponse200292) GetUpgrade() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *InlineResponse200292) GetUpgradeOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *InlineResponse200292) SetUpgrade(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *InlineResponse200292) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


