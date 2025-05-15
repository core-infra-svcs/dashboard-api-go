# InlineResponse200139

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The Meraki managed application Id for this record on a particular device. | [optional] 
**BundleSize** | Pointer to **int32** | The size of the software bundle. | [optional] 
**CreatedAt** | Pointer to **string** | When the Meraki record for the software was created. | [optional] 
**DeviceId** | Pointer to **string** | The Meraki managed device Id. | [optional] 
**DynamicSize** | Pointer to **int32** | The size of the data stored in the application. | [optional] 
**Id** | Pointer to **string** | The Meraki software Id. | [optional] 
**Identifier** | Pointer to **string** | Software bundle identifier. | [optional] 
**InstalledAt** | Pointer to **string** | When the Software was installed on the device. | [optional] 
**ToInstall** | Pointer to **bool** | A boolean indicating this software record should be installed on the associated device. | [optional] 
**IosRedemptionCode** | Pointer to **bool** | A boolean indicating whether or not an iOS redemption code was used. | [optional] 
**IsManaged** | Pointer to **bool** | A boolean indicating whether or not the software is managed by Meraki. | [optional] 
**ItunesId** | Pointer to **string** | The itunes numerical identifier. | [optional] 
**LicenseKey** | Pointer to **string** | The license key associated with this software installation. | [optional] 
**Name** | Pointer to **string** | The name of the software. | [optional] 
**Path** | Pointer to **string** | The path on the device where the software record is located. | [optional] 
**RedemptionCode** | Pointer to **int32** | The redemption code used for this software. | [optional] 
**ShortVersion** | Pointer to **string** | Short version notation for the software. | [optional] 
**Status** | Pointer to **string** | The management status of the software. | [optional] 
**ToUninstall** | Pointer to **bool** | A boolean indicating this software record should be uninstalled on the associated device. | [optional] 
**UninstalledAt** | Pointer to **string** | When the record was uninstalled from the device. | [optional] 
**UpdatedAt** | Pointer to **string** | When the record was last updated by Meraki. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the software. | [optional] 
**Version** | Pointer to **string** | Full version notation for the software. | [optional] 

## Methods

### NewInlineResponse200139

`func NewInlineResponse200139() *InlineResponse200139`

NewInlineResponse200139 instantiates a new InlineResponse200139 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200139WithDefaults

`func NewInlineResponse200139WithDefaults() *InlineResponse200139`

NewInlineResponse200139WithDefaults instantiates a new InlineResponse200139 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *InlineResponse200139) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *InlineResponse200139) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *InlineResponse200139) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *InlineResponse200139) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBundleSize

`func (o *InlineResponse200139) GetBundleSize() int32`

GetBundleSize returns the BundleSize field if non-nil, zero value otherwise.

### GetBundleSizeOk

`func (o *InlineResponse200139) GetBundleSizeOk() (*int32, bool)`

GetBundleSizeOk returns a tuple with the BundleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleSize

`func (o *InlineResponse200139) SetBundleSize(v int32)`

SetBundleSize sets BundleSize field to given value.

### HasBundleSize

`func (o *InlineResponse200139) HasBundleSize() bool`

HasBundleSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200139) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200139) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200139) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200139) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeviceId

`func (o *InlineResponse200139) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *InlineResponse200139) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *InlineResponse200139) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *InlineResponse200139) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDynamicSize

`func (o *InlineResponse200139) GetDynamicSize() int32`

GetDynamicSize returns the DynamicSize field if non-nil, zero value otherwise.

### GetDynamicSizeOk

`func (o *InlineResponse200139) GetDynamicSizeOk() (*int32, bool)`

GetDynamicSizeOk returns a tuple with the DynamicSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSize

`func (o *InlineResponse200139) SetDynamicSize(v int32)`

SetDynamicSize sets DynamicSize field to given value.

### HasDynamicSize

`func (o *InlineResponse200139) HasDynamicSize() bool`

HasDynamicSize returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200139) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200139) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200139) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200139) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *InlineResponse200139) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InlineResponse200139) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InlineResponse200139) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InlineResponse200139) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInstalledAt

`func (o *InlineResponse200139) GetInstalledAt() string`

GetInstalledAt returns the InstalledAt field if non-nil, zero value otherwise.

### GetInstalledAtOk

`func (o *InlineResponse200139) GetInstalledAtOk() (*string, bool)`

GetInstalledAtOk returns a tuple with the InstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAt

`func (o *InlineResponse200139) SetInstalledAt(v string)`

SetInstalledAt sets InstalledAt field to given value.

### HasInstalledAt

`func (o *InlineResponse200139) HasInstalledAt() bool`

HasInstalledAt returns a boolean if a field has been set.

### GetToInstall

`func (o *InlineResponse200139) GetToInstall() bool`

GetToInstall returns the ToInstall field if non-nil, zero value otherwise.

### GetToInstallOk

`func (o *InlineResponse200139) GetToInstallOk() (*bool, bool)`

GetToInstallOk returns a tuple with the ToInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInstall

`func (o *InlineResponse200139) SetToInstall(v bool)`

SetToInstall sets ToInstall field to given value.

### HasToInstall

`func (o *InlineResponse200139) HasToInstall() bool`

HasToInstall returns a boolean if a field has been set.

### GetIosRedemptionCode

`func (o *InlineResponse200139) GetIosRedemptionCode() bool`

GetIosRedemptionCode returns the IosRedemptionCode field if non-nil, zero value otherwise.

### GetIosRedemptionCodeOk

`func (o *InlineResponse200139) GetIosRedemptionCodeOk() (*bool, bool)`

GetIosRedemptionCodeOk returns a tuple with the IosRedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosRedemptionCode

`func (o *InlineResponse200139) SetIosRedemptionCode(v bool)`

SetIosRedemptionCode sets IosRedemptionCode field to given value.

### HasIosRedemptionCode

`func (o *InlineResponse200139) HasIosRedemptionCode() bool`

HasIosRedemptionCode returns a boolean if a field has been set.

### GetIsManaged

`func (o *InlineResponse200139) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *InlineResponse200139) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *InlineResponse200139) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *InlineResponse200139) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetItunesId

`func (o *InlineResponse200139) GetItunesId() string`

GetItunesId returns the ItunesId field if non-nil, zero value otherwise.

### GetItunesIdOk

`func (o *InlineResponse200139) GetItunesIdOk() (*string, bool)`

GetItunesIdOk returns a tuple with the ItunesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItunesId

`func (o *InlineResponse200139) SetItunesId(v string)`

SetItunesId sets ItunesId field to given value.

### HasItunesId

`func (o *InlineResponse200139) HasItunesId() bool`

HasItunesId returns a boolean if a field has been set.

### GetLicenseKey

`func (o *InlineResponse200139) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *InlineResponse200139) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *InlineResponse200139) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *InlineResponse200139) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200139) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200139) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200139) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200139) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *InlineResponse200139) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InlineResponse200139) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InlineResponse200139) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *InlineResponse200139) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRedemptionCode

`func (o *InlineResponse200139) GetRedemptionCode() int32`

GetRedemptionCode returns the RedemptionCode field if non-nil, zero value otherwise.

### GetRedemptionCodeOk

`func (o *InlineResponse200139) GetRedemptionCodeOk() (*int32, bool)`

GetRedemptionCodeOk returns a tuple with the RedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCode

`func (o *InlineResponse200139) SetRedemptionCode(v int32)`

SetRedemptionCode sets RedemptionCode field to given value.

### HasRedemptionCode

`func (o *InlineResponse200139) HasRedemptionCode() bool`

HasRedemptionCode returns a boolean if a field has been set.

### GetShortVersion

`func (o *InlineResponse200139) GetShortVersion() string`

GetShortVersion returns the ShortVersion field if non-nil, zero value otherwise.

### GetShortVersionOk

`func (o *InlineResponse200139) GetShortVersionOk() (*string, bool)`

GetShortVersionOk returns a tuple with the ShortVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortVersion

`func (o *InlineResponse200139) SetShortVersion(v string)`

SetShortVersion sets ShortVersion field to given value.

### HasShortVersion

`func (o *InlineResponse200139) HasShortVersion() bool`

HasShortVersion returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200139) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200139) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200139) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200139) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToUninstall

`func (o *InlineResponse200139) GetToUninstall() bool`

GetToUninstall returns the ToUninstall field if non-nil, zero value otherwise.

### GetToUninstallOk

`func (o *InlineResponse200139) GetToUninstallOk() (*bool, bool)`

GetToUninstallOk returns a tuple with the ToUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUninstall

`func (o *InlineResponse200139) SetToUninstall(v bool)`

SetToUninstall sets ToUninstall field to given value.

### HasToUninstall

`func (o *InlineResponse200139) HasToUninstall() bool`

HasToUninstall returns a boolean if a field has been set.

### GetUninstalledAt

`func (o *InlineResponse200139) GetUninstalledAt() string`

GetUninstalledAt returns the UninstalledAt field if non-nil, zero value otherwise.

### GetUninstalledAtOk

`func (o *InlineResponse200139) GetUninstalledAtOk() (*string, bool)`

GetUninstalledAtOk returns a tuple with the UninstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalledAt

`func (o *InlineResponse200139) SetUninstalledAt(v string)`

SetUninstalledAt sets UninstalledAt field to given value.

### HasUninstalledAt

`func (o *InlineResponse200139) HasUninstalledAt() bool`

HasUninstalledAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200139) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200139) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200139) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200139) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVendor

`func (o *InlineResponse200139) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InlineResponse200139) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InlineResponse200139) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InlineResponse200139) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *InlineResponse200139) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineResponse200139) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineResponse200139) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineResponse200139) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


