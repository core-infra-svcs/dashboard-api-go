# InlineResponse200361

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of the device | [optional] 
**Name** | Pointer to **string** | Name assigned to the device | [optional] 
**DeviceStatus** | Pointer to **string** | Status of the device upgrade | [optional] 
**CheckinFinishedAt** | Pointer to **NullableTime** | The time the device checkin finished | [optional] 
**CheckinStartedAt** | Pointer to **NullableTime** | The time the device checkin started | [optional] 
**DetailedStatus** | Pointer to **NullableString** | The detailed status of the device upgrade | [optional] 
**DownloadFinishedAt** | Pointer to **NullableTime** | The time the device upgrade download finished | [optional] 
**DownloadStartedAt** | Pointer to **NullableTime** | The time the device upgrade download started | [optional] 
**DownloadStatus** | Pointer to **NullableString** | The status of the device upgrade download | [optional] 
**InstallFinishedAt** | Pointer to **NullableTime** | The time the device upgrade install finished | [optional] 
**InstallStartedAt** | Pointer to **NullableTime** | The time the device upgrade install started | [optional] 
**InstallStatus** | Pointer to **NullableString** | The status of the device upgrade install | [optional] 
**VerifyFinishedAt** | Pointer to **NullableTime** | The time the device upgrade verification finished | [optional] 
**VerifyStartedAt** | Pointer to **NullableTime** | The time the device upgrade verification started | [optional] 
**VerifyStatus** | Pointer to **NullableString** | The status of the device upgrade verification | [optional] 
**Upgrade** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade**](OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade.md) |  | [optional] 

## Methods

### NewInlineResponse200361

`func NewInlineResponse200361() *InlineResponse200361`

NewInlineResponse200361 instantiates a new InlineResponse200361 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200361WithDefaults

`func NewInlineResponse200361WithDefaults() *InlineResponse200361`

NewInlineResponse200361WithDefaults instantiates a new InlineResponse200361 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200361) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200361) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200361) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200361) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200361) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200361) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200361) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200361) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *InlineResponse200361) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *InlineResponse200361) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *InlineResponse200361) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *InlineResponse200361) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetCheckinFinishedAt

`func (o *InlineResponse200361) GetCheckinFinishedAt() time.Time`

GetCheckinFinishedAt returns the CheckinFinishedAt field if non-nil, zero value otherwise.

### GetCheckinFinishedAtOk

`func (o *InlineResponse200361) GetCheckinFinishedAtOk() (*time.Time, bool)`

GetCheckinFinishedAtOk returns a tuple with the CheckinFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinFinishedAt

`func (o *InlineResponse200361) SetCheckinFinishedAt(v time.Time)`

SetCheckinFinishedAt sets CheckinFinishedAt field to given value.

### HasCheckinFinishedAt

`func (o *InlineResponse200361) HasCheckinFinishedAt() bool`

HasCheckinFinishedAt returns a boolean if a field has been set.

### SetCheckinFinishedAtNil

`func (o *InlineResponse200361) SetCheckinFinishedAtNil(b bool)`

 SetCheckinFinishedAtNil sets the value for CheckinFinishedAt to be an explicit nil

### UnsetCheckinFinishedAt
`func (o *InlineResponse200361) UnsetCheckinFinishedAt()`

UnsetCheckinFinishedAt ensures that no value is present for CheckinFinishedAt, not even an explicit nil
### GetCheckinStartedAt

`func (o *InlineResponse200361) GetCheckinStartedAt() time.Time`

GetCheckinStartedAt returns the CheckinStartedAt field if non-nil, zero value otherwise.

### GetCheckinStartedAtOk

`func (o *InlineResponse200361) GetCheckinStartedAtOk() (*time.Time, bool)`

GetCheckinStartedAtOk returns a tuple with the CheckinStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinStartedAt

`func (o *InlineResponse200361) SetCheckinStartedAt(v time.Time)`

SetCheckinStartedAt sets CheckinStartedAt field to given value.

### HasCheckinStartedAt

`func (o *InlineResponse200361) HasCheckinStartedAt() bool`

HasCheckinStartedAt returns a boolean if a field has been set.

### SetCheckinStartedAtNil

`func (o *InlineResponse200361) SetCheckinStartedAtNil(b bool)`

 SetCheckinStartedAtNil sets the value for CheckinStartedAt to be an explicit nil

### UnsetCheckinStartedAt
`func (o *InlineResponse200361) UnsetCheckinStartedAt()`

UnsetCheckinStartedAt ensures that no value is present for CheckinStartedAt, not even an explicit nil
### GetDetailedStatus

`func (o *InlineResponse200361) GetDetailedStatus() string`

GetDetailedStatus returns the DetailedStatus field if non-nil, zero value otherwise.

### GetDetailedStatusOk

`func (o *InlineResponse200361) GetDetailedStatusOk() (*string, bool)`

GetDetailedStatusOk returns a tuple with the DetailedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedStatus

`func (o *InlineResponse200361) SetDetailedStatus(v string)`

SetDetailedStatus sets DetailedStatus field to given value.

### HasDetailedStatus

`func (o *InlineResponse200361) HasDetailedStatus() bool`

HasDetailedStatus returns a boolean if a field has been set.

### SetDetailedStatusNil

`func (o *InlineResponse200361) SetDetailedStatusNil(b bool)`

 SetDetailedStatusNil sets the value for DetailedStatus to be an explicit nil

### UnsetDetailedStatus
`func (o *InlineResponse200361) UnsetDetailedStatus()`

UnsetDetailedStatus ensures that no value is present for DetailedStatus, not even an explicit nil
### GetDownloadFinishedAt

`func (o *InlineResponse200361) GetDownloadFinishedAt() time.Time`

GetDownloadFinishedAt returns the DownloadFinishedAt field if non-nil, zero value otherwise.

### GetDownloadFinishedAtOk

`func (o *InlineResponse200361) GetDownloadFinishedAtOk() (*time.Time, bool)`

GetDownloadFinishedAtOk returns a tuple with the DownloadFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFinishedAt

`func (o *InlineResponse200361) SetDownloadFinishedAt(v time.Time)`

SetDownloadFinishedAt sets DownloadFinishedAt field to given value.

### HasDownloadFinishedAt

`func (o *InlineResponse200361) HasDownloadFinishedAt() bool`

HasDownloadFinishedAt returns a boolean if a field has been set.

### SetDownloadFinishedAtNil

`func (o *InlineResponse200361) SetDownloadFinishedAtNil(b bool)`

 SetDownloadFinishedAtNil sets the value for DownloadFinishedAt to be an explicit nil

### UnsetDownloadFinishedAt
`func (o *InlineResponse200361) UnsetDownloadFinishedAt()`

UnsetDownloadFinishedAt ensures that no value is present for DownloadFinishedAt, not even an explicit nil
### GetDownloadStartedAt

`func (o *InlineResponse200361) GetDownloadStartedAt() time.Time`

GetDownloadStartedAt returns the DownloadStartedAt field if non-nil, zero value otherwise.

### GetDownloadStartedAtOk

`func (o *InlineResponse200361) GetDownloadStartedAtOk() (*time.Time, bool)`

GetDownloadStartedAtOk returns a tuple with the DownloadStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStartedAt

`func (o *InlineResponse200361) SetDownloadStartedAt(v time.Time)`

SetDownloadStartedAt sets DownloadStartedAt field to given value.

### HasDownloadStartedAt

`func (o *InlineResponse200361) HasDownloadStartedAt() bool`

HasDownloadStartedAt returns a boolean if a field has been set.

### SetDownloadStartedAtNil

`func (o *InlineResponse200361) SetDownloadStartedAtNil(b bool)`

 SetDownloadStartedAtNil sets the value for DownloadStartedAt to be an explicit nil

### UnsetDownloadStartedAt
`func (o *InlineResponse200361) UnsetDownloadStartedAt()`

UnsetDownloadStartedAt ensures that no value is present for DownloadStartedAt, not even an explicit nil
### GetDownloadStatus

`func (o *InlineResponse200361) GetDownloadStatus() string`

GetDownloadStatus returns the DownloadStatus field if non-nil, zero value otherwise.

### GetDownloadStatusOk

`func (o *InlineResponse200361) GetDownloadStatusOk() (*string, bool)`

GetDownloadStatusOk returns a tuple with the DownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStatus

`func (o *InlineResponse200361) SetDownloadStatus(v string)`

SetDownloadStatus sets DownloadStatus field to given value.

### HasDownloadStatus

`func (o *InlineResponse200361) HasDownloadStatus() bool`

HasDownloadStatus returns a boolean if a field has been set.

### SetDownloadStatusNil

`func (o *InlineResponse200361) SetDownloadStatusNil(b bool)`

 SetDownloadStatusNil sets the value for DownloadStatus to be an explicit nil

### UnsetDownloadStatus
`func (o *InlineResponse200361) UnsetDownloadStatus()`

UnsetDownloadStatus ensures that no value is present for DownloadStatus, not even an explicit nil
### GetInstallFinishedAt

`func (o *InlineResponse200361) GetInstallFinishedAt() time.Time`

GetInstallFinishedAt returns the InstallFinishedAt field if non-nil, zero value otherwise.

### GetInstallFinishedAtOk

`func (o *InlineResponse200361) GetInstallFinishedAtOk() (*time.Time, bool)`

GetInstallFinishedAtOk returns a tuple with the InstallFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallFinishedAt

`func (o *InlineResponse200361) SetInstallFinishedAt(v time.Time)`

SetInstallFinishedAt sets InstallFinishedAt field to given value.

### HasInstallFinishedAt

`func (o *InlineResponse200361) HasInstallFinishedAt() bool`

HasInstallFinishedAt returns a boolean if a field has been set.

### SetInstallFinishedAtNil

`func (o *InlineResponse200361) SetInstallFinishedAtNil(b bool)`

 SetInstallFinishedAtNil sets the value for InstallFinishedAt to be an explicit nil

### UnsetInstallFinishedAt
`func (o *InlineResponse200361) UnsetInstallFinishedAt()`

UnsetInstallFinishedAt ensures that no value is present for InstallFinishedAt, not even an explicit nil
### GetInstallStartedAt

`func (o *InlineResponse200361) GetInstallStartedAt() time.Time`

GetInstallStartedAt returns the InstallStartedAt field if non-nil, zero value otherwise.

### GetInstallStartedAtOk

`func (o *InlineResponse200361) GetInstallStartedAtOk() (*time.Time, bool)`

GetInstallStartedAtOk returns a tuple with the InstallStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStartedAt

`func (o *InlineResponse200361) SetInstallStartedAt(v time.Time)`

SetInstallStartedAt sets InstallStartedAt field to given value.

### HasInstallStartedAt

`func (o *InlineResponse200361) HasInstallStartedAt() bool`

HasInstallStartedAt returns a boolean if a field has been set.

### SetInstallStartedAtNil

`func (o *InlineResponse200361) SetInstallStartedAtNil(b bool)`

 SetInstallStartedAtNil sets the value for InstallStartedAt to be an explicit nil

### UnsetInstallStartedAt
`func (o *InlineResponse200361) UnsetInstallStartedAt()`

UnsetInstallStartedAt ensures that no value is present for InstallStartedAt, not even an explicit nil
### GetInstallStatus

`func (o *InlineResponse200361) GetInstallStatus() string`

GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.

### GetInstallStatusOk

`func (o *InlineResponse200361) GetInstallStatusOk() (*string, bool)`

GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStatus

`func (o *InlineResponse200361) SetInstallStatus(v string)`

SetInstallStatus sets InstallStatus field to given value.

### HasInstallStatus

`func (o *InlineResponse200361) HasInstallStatus() bool`

HasInstallStatus returns a boolean if a field has been set.

### SetInstallStatusNil

`func (o *InlineResponse200361) SetInstallStatusNil(b bool)`

 SetInstallStatusNil sets the value for InstallStatus to be an explicit nil

### UnsetInstallStatus
`func (o *InlineResponse200361) UnsetInstallStatus()`

UnsetInstallStatus ensures that no value is present for InstallStatus, not even an explicit nil
### GetVerifyFinishedAt

`func (o *InlineResponse200361) GetVerifyFinishedAt() time.Time`

GetVerifyFinishedAt returns the VerifyFinishedAt field if non-nil, zero value otherwise.

### GetVerifyFinishedAtOk

`func (o *InlineResponse200361) GetVerifyFinishedAtOk() (*time.Time, bool)`

GetVerifyFinishedAtOk returns a tuple with the VerifyFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyFinishedAt

`func (o *InlineResponse200361) SetVerifyFinishedAt(v time.Time)`

SetVerifyFinishedAt sets VerifyFinishedAt field to given value.

### HasVerifyFinishedAt

`func (o *InlineResponse200361) HasVerifyFinishedAt() bool`

HasVerifyFinishedAt returns a boolean if a field has been set.

### SetVerifyFinishedAtNil

`func (o *InlineResponse200361) SetVerifyFinishedAtNil(b bool)`

 SetVerifyFinishedAtNil sets the value for VerifyFinishedAt to be an explicit nil

### UnsetVerifyFinishedAt
`func (o *InlineResponse200361) UnsetVerifyFinishedAt()`

UnsetVerifyFinishedAt ensures that no value is present for VerifyFinishedAt, not even an explicit nil
### GetVerifyStartedAt

`func (o *InlineResponse200361) GetVerifyStartedAt() time.Time`

GetVerifyStartedAt returns the VerifyStartedAt field if non-nil, zero value otherwise.

### GetVerifyStartedAtOk

`func (o *InlineResponse200361) GetVerifyStartedAtOk() (*time.Time, bool)`

GetVerifyStartedAtOk returns a tuple with the VerifyStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyStartedAt

`func (o *InlineResponse200361) SetVerifyStartedAt(v time.Time)`

SetVerifyStartedAt sets VerifyStartedAt field to given value.

### HasVerifyStartedAt

`func (o *InlineResponse200361) HasVerifyStartedAt() bool`

HasVerifyStartedAt returns a boolean if a field has been set.

### SetVerifyStartedAtNil

`func (o *InlineResponse200361) SetVerifyStartedAtNil(b bool)`

 SetVerifyStartedAtNil sets the value for VerifyStartedAt to be an explicit nil

### UnsetVerifyStartedAt
`func (o *InlineResponse200361) UnsetVerifyStartedAt()`

UnsetVerifyStartedAt ensures that no value is present for VerifyStartedAt, not even an explicit nil
### GetVerifyStatus

`func (o *InlineResponse200361) GetVerifyStatus() string`

GetVerifyStatus returns the VerifyStatus field if non-nil, zero value otherwise.

### GetVerifyStatusOk

`func (o *InlineResponse200361) GetVerifyStatusOk() (*string, bool)`

GetVerifyStatusOk returns a tuple with the VerifyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyStatus

`func (o *InlineResponse200361) SetVerifyStatus(v string)`

SetVerifyStatus sets VerifyStatus field to given value.

### HasVerifyStatus

`func (o *InlineResponse200361) HasVerifyStatus() bool`

HasVerifyStatus returns a boolean if a field has been set.

### SetVerifyStatusNil

`func (o *InlineResponse200361) SetVerifyStatusNil(b bool)`

 SetVerifyStatusNil sets the value for VerifyStatus to be an explicit nil

### UnsetVerifyStatus
`func (o *InlineResponse200361) UnsetVerifyStatus()`

UnsetVerifyStatus ensures that no value is present for VerifyStatus, not even an explicit nil
### GetUpgrade

`func (o *InlineResponse200361) GetUpgrade() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *InlineResponse200361) GetUpgradeOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *InlineResponse200361) SetUpgrade(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *InlineResponse200361) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


