# OrganizationsOrganizationIdWirelessSsidsProfilesSsid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | SSID name | 
**Advertisement** | Pointer to [**InlineResponse200517SsidAdvertisement**](InlineResponse200517SsidAdvertisement.md) |  | [optional] 
**Security** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurity**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurity.md) |  | [optional] 
**Radius** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius.md) |  | [optional] 
**Addressing** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidAddressing**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidAddressing.md) |  | [optional] 
**Splash** | Pointer to [**InlineResponse200517SsidSplash**](InlineResponse200517SsidSplash.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdWirelessSsidsProfilesSsid

`func NewOrganizationsOrganizationIdWirelessSsidsProfilesSsid(name string, ) *OrganizationsOrganizationIdWirelessSsidsProfilesSsid`

NewOrganizationsOrganizationIdWirelessSsidsProfilesSsid instantiates a new OrganizationsOrganizationIdWirelessSsidsProfilesSsid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidWithDefaults

`func NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidWithDefaults() *OrganizationsOrganizationIdWirelessSsidsProfilesSsid`

NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidWithDefaults instantiates a new OrganizationsOrganizationIdWirelessSsidsProfilesSsid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) SetName(v string)`

SetName sets Name field to given value.


### GetAdvertisement

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetAdvertisement() InlineResponse200517SsidAdvertisement`

GetAdvertisement returns the Advertisement field if non-nil, zero value otherwise.

### GetAdvertisementOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetAdvertisementOk() (*InlineResponse200517SsidAdvertisement, bool)`

GetAdvertisementOk returns a tuple with the Advertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisement

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) SetAdvertisement(v InlineResponse200517SsidAdvertisement)`

SetAdvertisement sets Advertisement field to given value.

### HasAdvertisement

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) HasAdvertisement() bool`

HasAdvertisement returns a boolean if a field has been set.

### GetSecurity

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetSecurity() OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetSecurityOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) SetSecurity(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetRadius

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetRadius() OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetRadiusOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) SetRadius(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetAddressing

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetAddressing() OrganizationsOrganizationIdWirelessSsidsProfilesSsidAddressing`

GetAddressing returns the Addressing field if non-nil, zero value otherwise.

### GetAddressingOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetAddressingOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidAddressing, bool)`

GetAddressingOk returns a tuple with the Addressing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressing

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) SetAddressing(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidAddressing)`

SetAddressing sets Addressing field to given value.

### HasAddressing

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) HasAddressing() bool`

HasAddressing returns a boolean if a field has been set.

### GetSplash

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetSplash() InlineResponse200517SsidSplash`

GetSplash returns the Splash field if non-nil, zero value otherwise.

### GetSplashOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) GetSplashOk() (*InlineResponse200517SsidSplash, bool)`

GetSplashOk returns a tuple with the Splash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplash

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) SetSplash(v InlineResponse200517SsidSplash)`

SetSplash sets Splash field to given value.

### HasSplash

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsid) HasSplash() bool`

HasSplash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


