# InlineResponse200298

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | License ID | [optional] 
**LicenseType** | Pointer to **string** | License type | [optional] 
**LicenseKey** | Pointer to **string** | License key | [optional] 
**OrderNumber** | Pointer to **string** | Order number | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device the license is assigned to | [optional] 
**NetworkId** | Pointer to **string** | ID of the network the license is assigned to | [optional] 
**State** | Pointer to **string** | The state of the license. All queued licenses have a status of &#x60;recentlyQueued&#x60;. | [optional] 
**SeatCount** | Pointer to **int32** | The number of seats of the license. Only applicable to SM licenses. | [optional] 
**TotalDurationInDays** | Pointer to **int32** | The duration of the license plus all permanently queued licenses associated with it | [optional] 
**DurationInDays** | Pointer to **int32** | The duration of the individual license | [optional] 
**PermanentlyQueuedLicenses** | Pointer to [**[]OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses**](OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses.md) | DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial&#x3D; to retrieved queued licenses for a given device. | [optional] 
**ClaimDate** | Pointer to **string** | The date the license was claimed into the organization | [optional] 
**ActivationDate** | Pointer to **string** | The date the license started burning | [optional] 
**ExpirationDate** | Pointer to **string** | The date the license will expire | [optional] 
**HeadLicenseId** | Pointer to **string** | The id of the head license this license is queued behind. If there is no head license, it returns nil. | [optional] 

## Methods

### NewInlineResponse200298

`func NewInlineResponse200298() *InlineResponse200298`

NewInlineResponse200298 instantiates a new InlineResponse200298 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200298WithDefaults

`func NewInlineResponse200298WithDefaults() *InlineResponse200298`

NewInlineResponse200298WithDefaults instantiates a new InlineResponse200298 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200298) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200298) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200298) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200298) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseType

`func (o *InlineResponse200298) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *InlineResponse200298) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *InlineResponse200298) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *InlineResponse200298) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseKey

`func (o *InlineResponse200298) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *InlineResponse200298) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *InlineResponse200298) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *InlineResponse200298) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetOrderNumber

`func (o *InlineResponse200298) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *InlineResponse200298) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *InlineResponse200298) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *InlineResponse200298) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse200298) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse200298) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse200298) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse200298) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200298) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200298) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200298) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200298) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetState

`func (o *InlineResponse200298) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse200298) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse200298) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InlineResponse200298) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSeatCount

`func (o *InlineResponse200298) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *InlineResponse200298) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *InlineResponse200298) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.

### HasSeatCount

`func (o *InlineResponse200298) HasSeatCount() bool`

HasSeatCount returns a boolean if a field has been set.

### GetTotalDurationInDays

`func (o *InlineResponse200298) GetTotalDurationInDays() int32`

GetTotalDurationInDays returns the TotalDurationInDays field if non-nil, zero value otherwise.

### GetTotalDurationInDaysOk

`func (o *InlineResponse200298) GetTotalDurationInDaysOk() (*int32, bool)`

GetTotalDurationInDaysOk returns a tuple with the TotalDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationInDays

`func (o *InlineResponse200298) SetTotalDurationInDays(v int32)`

SetTotalDurationInDays sets TotalDurationInDays field to given value.

### HasTotalDurationInDays

`func (o *InlineResponse200298) HasTotalDurationInDays() bool`

HasTotalDurationInDays returns a boolean if a field has been set.

### GetDurationInDays

`func (o *InlineResponse200298) GetDurationInDays() int32`

GetDurationInDays returns the DurationInDays field if non-nil, zero value otherwise.

### GetDurationInDaysOk

`func (o *InlineResponse200298) GetDurationInDaysOk() (*int32, bool)`

GetDurationInDaysOk returns a tuple with the DurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInDays

`func (o *InlineResponse200298) SetDurationInDays(v int32)`

SetDurationInDays sets DurationInDays field to given value.

### HasDurationInDays

`func (o *InlineResponse200298) HasDurationInDays() bool`

HasDurationInDays returns a boolean if a field has been set.

### GetPermanentlyQueuedLicenses

`func (o *InlineResponse200298) GetPermanentlyQueuedLicenses() []OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses`

GetPermanentlyQueuedLicenses returns the PermanentlyQueuedLicenses field if non-nil, zero value otherwise.

### GetPermanentlyQueuedLicensesOk

`func (o *InlineResponse200298) GetPermanentlyQueuedLicensesOk() (*[]OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses, bool)`

GetPermanentlyQueuedLicensesOk returns a tuple with the PermanentlyQueuedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentlyQueuedLicenses

`func (o *InlineResponse200298) SetPermanentlyQueuedLicenses(v []OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses)`

SetPermanentlyQueuedLicenses sets PermanentlyQueuedLicenses field to given value.

### HasPermanentlyQueuedLicenses

`func (o *InlineResponse200298) HasPermanentlyQueuedLicenses() bool`

HasPermanentlyQueuedLicenses returns a boolean if a field has been set.

### GetClaimDate

`func (o *InlineResponse200298) GetClaimDate() string`

GetClaimDate returns the ClaimDate field if non-nil, zero value otherwise.

### GetClaimDateOk

`func (o *InlineResponse200298) GetClaimDateOk() (*string, bool)`

GetClaimDateOk returns a tuple with the ClaimDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimDate

`func (o *InlineResponse200298) SetClaimDate(v string)`

SetClaimDate sets ClaimDate field to given value.

### HasClaimDate

`func (o *InlineResponse200298) HasClaimDate() bool`

HasClaimDate returns a boolean if a field has been set.

### GetActivationDate

`func (o *InlineResponse200298) GetActivationDate() string`

GetActivationDate returns the ActivationDate field if non-nil, zero value otherwise.

### GetActivationDateOk

`func (o *InlineResponse200298) GetActivationDateOk() (*string, bool)`

GetActivationDateOk returns a tuple with the ActivationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationDate

`func (o *InlineResponse200298) SetActivationDate(v string)`

SetActivationDate sets ActivationDate field to given value.

### HasActivationDate

`func (o *InlineResponse200298) HasActivationDate() bool`

HasActivationDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InlineResponse200298) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InlineResponse200298) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InlineResponse200298) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InlineResponse200298) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHeadLicenseId

`func (o *InlineResponse200298) GetHeadLicenseId() string`

GetHeadLicenseId returns the HeadLicenseId field if non-nil, zero value otherwise.

### GetHeadLicenseIdOk

`func (o *InlineResponse200298) GetHeadLicenseIdOk() (*string, bool)`

GetHeadLicenseIdOk returns a tuple with the HeadLicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadLicenseId

`func (o *InlineResponse200298) SetHeadLicenseId(v string)`

SetHeadLicenseId sets HeadLicenseId field to given value.

### HasHeadLicenseId

`func (o *InlineResponse200298) HasHeadLicenseId() bool`

HasHeadLicenseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


