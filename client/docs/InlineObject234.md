# InlineObject234

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestOrganizationId** | **string** | The ID of the organization to move the SM seats to | 
**LicenseId** | **string** | The ID of the SM license to move the seats from | 
**SeatCount** | **int32** | The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license | 

## Methods

### NewInlineObject234

`func NewInlineObject234(destOrganizationId string, licenseId string, seatCount int32, ) *InlineObject234`

NewInlineObject234 instantiates a new InlineObject234 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject234WithDefaults

`func NewInlineObject234WithDefaults() *InlineObject234`

NewInlineObject234WithDefaults instantiates a new InlineObject234 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestOrganizationId

`func (o *InlineObject234) GetDestOrganizationId() string`

GetDestOrganizationId returns the DestOrganizationId field if non-nil, zero value otherwise.

### GetDestOrganizationIdOk

`func (o *InlineObject234) GetDestOrganizationIdOk() (*string, bool)`

GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOrganizationId

`func (o *InlineObject234) SetDestOrganizationId(v string)`

SetDestOrganizationId sets DestOrganizationId field to given value.


### GetLicenseId

`func (o *InlineObject234) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *InlineObject234) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *InlineObject234) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetSeatCount

`func (o *InlineObject234) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *InlineObject234) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *InlineObject234) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


