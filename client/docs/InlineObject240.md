# InlineObject240

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseId** | **string** | The ID of the SM license to assign seats from | 
**NetworkId** | **string** | The ID of the SM network to assign the seats to | 
**SeatCount** | **int32** | The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license | 

## Methods

### NewInlineObject240

`func NewInlineObject240(licenseId string, networkId string, seatCount int32, ) *InlineObject240`

NewInlineObject240 instantiates a new InlineObject240 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject240WithDefaults

`func NewInlineObject240WithDefaults() *InlineObject240`

NewInlineObject240WithDefaults instantiates a new InlineObject240 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseId

`func (o *InlineObject240) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *InlineObject240) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *InlineObject240) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetNetworkId

`func (o *InlineObject240) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineObject240) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineObject240) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetSeatCount

`func (o *InlineObject240) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *InlineObject240) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *InlineObject240) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


