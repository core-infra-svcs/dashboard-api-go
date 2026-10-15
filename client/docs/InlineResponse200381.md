# InlineResponse200381

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Device Name | [optional] 
**Serial** | Pointer to **string** | Device Serial Number | [optional] 
**Mac** | Pointer to **string** | MAC Address | [optional] 
**PublicIp** | Pointer to **string** | Public IP Address | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Status** | Pointer to **string** | Device Status | [optional] 
**LastReportedAt** | Pointer to **string** | Device Last Reported Location | [optional] 
**LanIp** | Pointer to **string** | LAN IP Address | [optional] 
**Gateway** | Pointer to **string** | IP Gateway | [optional] 
**IpType** | Pointer to **string** | IP Type | [optional] 
**PrimaryDns** | Pointer to **string** | Primary DNS | [optional] 
**SecondaryDns** | Pointer to **string** | Secondary DNS | [optional] 
**ProductType** | Pointer to **string** | Product Type | [optional] 
**Components** | Pointer to [**OrganizationsOrganizationIdDevicesStatusesComponents**](OrganizationsOrganizationIdDevicesStatusesComponents.md) |  | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 

## Methods

### NewInlineResponse200381

`func NewInlineResponse200381() *InlineResponse200381`

NewInlineResponse200381 instantiates a new InlineResponse200381 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200381WithDefaults

`func NewInlineResponse200381WithDefaults() *InlineResponse200381`

NewInlineResponse200381WithDefaults instantiates a new InlineResponse200381 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200381) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200381) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200381) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200381) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200381) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200381) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200381) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200381) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200381) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200381) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200381) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200381) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse200381) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse200381) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse200381) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse200381) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200381) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200381) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200381) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200381) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200381) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200381) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200381) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200381) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *InlineResponse200381) GetLastReportedAt() string`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200381) GetLastReportedAtOk() (*string, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200381) SetLastReportedAt(v string)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *InlineResponse200381) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetLanIp

`func (o *InlineResponse200381) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse200381) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse200381) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse200381) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetGateway

`func (o *InlineResponse200381) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200381) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200381) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse200381) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpType

`func (o *InlineResponse200381) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *InlineResponse200381) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *InlineResponse200381) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *InlineResponse200381) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *InlineResponse200381) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *InlineResponse200381) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *InlineResponse200381) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *InlineResponse200381) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *InlineResponse200381) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *InlineResponse200381) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *InlineResponse200381) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *InlineResponse200381) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200381) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200381) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200381) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200381) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetComponents

`func (o *InlineResponse200381) GetComponents() OrganizationsOrganizationIdDevicesStatusesComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *InlineResponse200381) GetComponentsOk() (*OrganizationsOrganizationIdDevicesStatusesComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *InlineResponse200381) SetComponents(v OrganizationsOrganizationIdDevicesStatusesComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *InlineResponse200381) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200381) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200381) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200381) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200381) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200381) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200381) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200381) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200381) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


