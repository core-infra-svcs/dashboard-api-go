# InlineResponse20049

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

### NewInlineResponse20049

`func NewInlineResponse20049() *InlineResponse20049`

NewInlineResponse20049 instantiates a new InlineResponse20049 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20049WithDefaults

`func NewInlineResponse20049WithDefaults() *InlineResponse20049`

NewInlineResponse20049WithDefaults instantiates a new InlineResponse20049 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *InlineResponse20049) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *InlineResponse20049) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *InlineResponse20049) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *InlineResponse20049) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBundleSize

`func (o *InlineResponse20049) GetBundleSize() int32`

GetBundleSize returns the BundleSize field if non-nil, zero value otherwise.

### GetBundleSizeOk

`func (o *InlineResponse20049) GetBundleSizeOk() (*int32, bool)`

GetBundleSizeOk returns a tuple with the BundleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleSize

`func (o *InlineResponse20049) SetBundleSize(v int32)`

SetBundleSize sets BundleSize field to given value.

### HasBundleSize

`func (o *InlineResponse20049) HasBundleSize() bool`

HasBundleSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse20049) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20049) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20049) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20049) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeviceId

`func (o *InlineResponse20049) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *InlineResponse20049) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *InlineResponse20049) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *InlineResponse20049) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDynamicSize

`func (o *InlineResponse20049) GetDynamicSize() int32`

GetDynamicSize returns the DynamicSize field if non-nil, zero value otherwise.

### GetDynamicSizeOk

`func (o *InlineResponse20049) GetDynamicSizeOk() (*int32, bool)`

GetDynamicSizeOk returns a tuple with the DynamicSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSize

`func (o *InlineResponse20049) SetDynamicSize(v int32)`

SetDynamicSize sets DynamicSize field to given value.

### HasDynamicSize

`func (o *InlineResponse20049) HasDynamicSize() bool`

HasDynamicSize returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20049) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20049) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20049) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20049) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *InlineResponse20049) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InlineResponse20049) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InlineResponse20049) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InlineResponse20049) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInstalledAt

`func (o *InlineResponse20049) GetInstalledAt() string`

GetInstalledAt returns the InstalledAt field if non-nil, zero value otherwise.

### GetInstalledAtOk

`func (o *InlineResponse20049) GetInstalledAtOk() (*string, bool)`

GetInstalledAtOk returns a tuple with the InstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAt

`func (o *InlineResponse20049) SetInstalledAt(v string)`

SetInstalledAt sets InstalledAt field to given value.

### HasInstalledAt

`func (o *InlineResponse20049) HasInstalledAt() bool`

HasInstalledAt returns a boolean if a field has been set.

### GetToInstall

`func (o *InlineResponse20049) GetToInstall() bool`

GetToInstall returns the ToInstall field if non-nil, zero value otherwise.

### GetToInstallOk

`func (o *InlineResponse20049) GetToInstallOk() (*bool, bool)`

GetToInstallOk returns a tuple with the ToInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInstall

`func (o *InlineResponse20049) SetToInstall(v bool)`

SetToInstall sets ToInstall field to given value.

### HasToInstall

`func (o *InlineResponse20049) HasToInstall() bool`

HasToInstall returns a boolean if a field has been set.

### GetIosRedemptionCode

`func (o *InlineResponse20049) GetIosRedemptionCode() bool`

GetIosRedemptionCode returns the IosRedemptionCode field if non-nil, zero value otherwise.

### GetIosRedemptionCodeOk

`func (o *InlineResponse20049) GetIosRedemptionCodeOk() (*bool, bool)`

GetIosRedemptionCodeOk returns a tuple with the IosRedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosRedemptionCode

`func (o *InlineResponse20049) SetIosRedemptionCode(v bool)`

SetIosRedemptionCode sets IosRedemptionCode field to given value.

### HasIosRedemptionCode

`func (o *InlineResponse20049) HasIosRedemptionCode() bool`

HasIosRedemptionCode returns a boolean if a field has been set.

### GetIsManaged

`func (o *InlineResponse20049) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *InlineResponse20049) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *InlineResponse20049) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *InlineResponse20049) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetItunesId

`func (o *InlineResponse20049) GetItunesId() string`

GetItunesId returns the ItunesId field if non-nil, zero value otherwise.

### GetItunesIdOk

`func (o *InlineResponse20049) GetItunesIdOk() (*string, bool)`

GetItunesIdOk returns a tuple with the ItunesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItunesId

`func (o *InlineResponse20049) SetItunesId(v string)`

SetItunesId sets ItunesId field to given value.

### HasItunesId

`func (o *InlineResponse20049) HasItunesId() bool`

HasItunesId returns a boolean if a field has been set.

### GetLicenseKey

`func (o *InlineResponse20049) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *InlineResponse20049) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *InlineResponse20049) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *InlineResponse20049) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20049) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20049) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20049) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20049) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *InlineResponse20049) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InlineResponse20049) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InlineResponse20049) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *InlineResponse20049) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRedemptionCode

`func (o *InlineResponse20049) GetRedemptionCode() int32`

GetRedemptionCode returns the RedemptionCode field if non-nil, zero value otherwise.

### GetRedemptionCodeOk

`func (o *InlineResponse20049) GetRedemptionCodeOk() (*int32, bool)`

GetRedemptionCodeOk returns a tuple with the RedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCode

`func (o *InlineResponse20049) SetRedemptionCode(v int32)`

SetRedemptionCode sets RedemptionCode field to given value.

### HasRedemptionCode

`func (o *InlineResponse20049) HasRedemptionCode() bool`

HasRedemptionCode returns a boolean if a field has been set.

### GetShortVersion

`func (o *InlineResponse20049) GetShortVersion() string`

GetShortVersion returns the ShortVersion field if non-nil, zero value otherwise.

### GetShortVersionOk

`func (o *InlineResponse20049) GetShortVersionOk() (*string, bool)`

GetShortVersionOk returns a tuple with the ShortVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortVersion

`func (o *InlineResponse20049) SetShortVersion(v string)`

SetShortVersion sets ShortVersion field to given value.

### HasShortVersion

`func (o *InlineResponse20049) HasShortVersion() bool`

HasShortVersion returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20049) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20049) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20049) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20049) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToUninstall

`func (o *InlineResponse20049) GetToUninstall() bool`

GetToUninstall returns the ToUninstall field if non-nil, zero value otherwise.

### GetToUninstallOk

`func (o *InlineResponse20049) GetToUninstallOk() (*bool, bool)`

GetToUninstallOk returns a tuple with the ToUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUninstall

`func (o *InlineResponse20049) SetToUninstall(v bool)`

SetToUninstall sets ToUninstall field to given value.

### HasToUninstall

`func (o *InlineResponse20049) HasToUninstall() bool`

HasToUninstall returns a boolean if a field has been set.

### GetUninstalledAt

`func (o *InlineResponse20049) GetUninstalledAt() string`

GetUninstalledAt returns the UninstalledAt field if non-nil, zero value otherwise.

### GetUninstalledAtOk

`func (o *InlineResponse20049) GetUninstalledAtOk() (*string, bool)`

GetUninstalledAtOk returns a tuple with the UninstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalledAt

`func (o *InlineResponse20049) SetUninstalledAt(v string)`

SetUninstalledAt sets UninstalledAt field to given value.

### HasUninstalledAt

`func (o *InlineResponse20049) HasUninstalledAt() bool`

HasUninstalledAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse20049) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse20049) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse20049) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse20049) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVendor

`func (o *InlineResponse20049) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InlineResponse20049) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InlineResponse20049) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InlineResponse20049) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *InlineResponse20049) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineResponse20049) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineResponse20049) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineResponse20049) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


