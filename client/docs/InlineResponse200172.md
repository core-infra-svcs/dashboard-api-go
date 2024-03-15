# InlineResponse200172

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Model** | Pointer to **string** | Model type of the device | [optional] 
**NetworkId** | Pointer to **string** | Network Id of the device | [optional] 
**OrderNumber** | Pointer to **string** | Order number of the device | [optional] 
**ClaimedAt** | Pointer to **time.Time** | Claimed time of the device | [optional] 
**LicenseExpirationDate** | Pointer to **time.Time** | License expiration date of the device | [optional] 
**Tags** | Pointer to **[]string** | Device tags | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 
**CountryCode** | Pointer to **string** | Country/region code from device, network, or store order | [optional] 
**Details** | Pointer to [**[]NetworksNetworkIdFloorPlansDetails**](NetworksNetworkIdFloorPlansDetails.md) | Additional device information | [optional] 

## Methods

### NewInlineResponse200172

`func NewInlineResponse200172() *InlineResponse200172`

NewInlineResponse200172 instantiates a new InlineResponse200172 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200172WithDefaults

`func NewInlineResponse200172WithDefaults() *InlineResponse200172`

NewInlineResponse200172WithDefaults instantiates a new InlineResponse200172 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse200172) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200172) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200172) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200172) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200172) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200172) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200172) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200172) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200172) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200172) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200172) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200172) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200172) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200172) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200172) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200172) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200172) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200172) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200172) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200172) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *InlineResponse200172) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *InlineResponse200172) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *InlineResponse200172) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *InlineResponse200172) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetClaimedAt

`func (o *InlineResponse200172) GetClaimedAt() time.Time`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *InlineResponse200172) GetClaimedAtOk() (*time.Time, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *InlineResponse200172) SetClaimedAt(v time.Time)`

SetClaimedAt sets ClaimedAt field to given value.

### HasClaimedAt

`func (o *InlineResponse200172) HasClaimedAt() bool`

HasClaimedAt returns a boolean if a field has been set.

### GetLicenseExpirationDate

`func (o *InlineResponse200172) GetLicenseExpirationDate() time.Time`

GetLicenseExpirationDate returns the LicenseExpirationDate field if non-nil, zero value otherwise.

### GetLicenseExpirationDateOk

`func (o *InlineResponse200172) GetLicenseExpirationDateOk() (*time.Time, bool)`

GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationDate

`func (o *InlineResponse200172) SetLicenseExpirationDate(v time.Time)`

SetLicenseExpirationDate sets LicenseExpirationDate field to given value.

### HasLicenseExpirationDate

`func (o *InlineResponse200172) HasLicenseExpirationDate() bool`

HasLicenseExpirationDate returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200172) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200172) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200172) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200172) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200172) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200172) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200172) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200172) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetCountryCode

`func (o *InlineResponse200172) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InlineResponse200172) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InlineResponse200172) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *InlineResponse200172) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200172) GetDetails() []NetworksNetworkIdFloorPlansDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200172) GetDetailsOk() (*[]NetworksNetworkIdFloorPlansDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200172) SetDetails(v []NetworksNetworkIdFloorPlansDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200172) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


