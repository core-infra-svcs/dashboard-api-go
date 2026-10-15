# InlineResponse200358

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Time, in ISO8601 format, when the configuration change was made. | [optional] 
**AdminName** | Pointer to **string** | The name of the admin who made the configuration change. | [optional] 
**AdminEmail** | Pointer to **string** | The email address of the admin who made the configuration change. This attribute may be null. | [optional] 
**AdminId** | Pointer to **string** | The ID of the admin who made the configuration change. This attribute may be null. | [optional] 
**NetworkName** | Pointer to **string** | The name of the network that the configuration change was applied to. This attribute may be null. | [optional] 
**NetworkId** | Pointer to **string** | The ID of the network that the configuration change was applied to. This attribute may be null. | [optional] 
**NetworkUrl** | Pointer to **string** | The url of the network that the configuration change was applied to. This attribute may be null. | [optional] 
**SsidName** | Pointer to **string** | The name of the ssid that the configuration change was applied to, if applicable. This attribute may be null. | [optional] 
**SsidNumber** | Pointer to **int32** | The ssid number that the configuration change was applied to, if applicable. This attribute may be null. | [optional] 
**Page** | Pointer to **string** | The name of the Meraki Dashboard page on which the configuration change was made. | [optional] 
**Label** | Pointer to **string** | Description of the configuration change. | [optional] 
**OldValue** | Pointer to **string** | The value of the configuration, before the change was applied. | [optional] 
**NewValue** | Pointer to **string** | The value of the configuration, after the change was applied. | [optional] 
**Client** | Pointer to [**OrganizationsOrganizationIdConfigurationChangesClient**](OrganizationsOrganizationIdConfigurationChangesClient.md) |  | [optional] 

## Methods

### NewInlineResponse200358

`func NewInlineResponse200358() *InlineResponse200358`

NewInlineResponse200358 instantiates a new InlineResponse200358 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200358WithDefaults

`func NewInlineResponse200358WithDefaults() *InlineResponse200358`

NewInlineResponse200358WithDefaults instantiates a new InlineResponse200358 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200358) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200358) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200358) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200358) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetAdminName

`func (o *InlineResponse200358) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *InlineResponse200358) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *InlineResponse200358) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *InlineResponse200358) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminEmail

`func (o *InlineResponse200358) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *InlineResponse200358) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *InlineResponse200358) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *InlineResponse200358) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetAdminId

`func (o *InlineResponse200358) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *InlineResponse200358) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *InlineResponse200358) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *InlineResponse200358) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetNetworkName

`func (o *InlineResponse200358) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200358) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200358) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *InlineResponse200358) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200358) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200358) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200358) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200358) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkUrl

`func (o *InlineResponse200358) GetNetworkUrl() string`

GetNetworkUrl returns the NetworkUrl field if non-nil, zero value otherwise.

### GetNetworkUrlOk

`func (o *InlineResponse200358) GetNetworkUrlOk() (*string, bool)`

GetNetworkUrlOk returns a tuple with the NetworkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUrl

`func (o *InlineResponse200358) SetNetworkUrl(v string)`

SetNetworkUrl sets NetworkUrl field to given value.

### HasNetworkUrl

`func (o *InlineResponse200358) HasNetworkUrl() bool`

HasNetworkUrl returns a boolean if a field has been set.

### GetSsidName

`func (o *InlineResponse200358) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *InlineResponse200358) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *InlineResponse200358) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *InlineResponse200358) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.

### GetSsidNumber

`func (o *InlineResponse200358) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200358) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200358) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200358) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetPage

`func (o *InlineResponse200358) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InlineResponse200358) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InlineResponse200358) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *InlineResponse200358) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLabel

`func (o *InlineResponse200358) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InlineResponse200358) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InlineResponse200358) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InlineResponse200358) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOldValue

`func (o *InlineResponse200358) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *InlineResponse200358) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *InlineResponse200358) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *InlineResponse200358) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetNewValue

`func (o *InlineResponse200358) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *InlineResponse200358) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *InlineResponse200358) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *InlineResponse200358) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetClient

`func (o *InlineResponse200358) GetClient() OrganizationsOrganizationIdConfigurationChangesClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *InlineResponse200358) GetClientOk() (*OrganizationsOrganizationIdConfigurationChangesClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *InlineResponse200358) SetClient(v OrganizationsOrganizationIdConfigurationChangesClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *InlineResponse200358) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


