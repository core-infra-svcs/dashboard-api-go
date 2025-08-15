# InlineObject279

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestOrganizationId** | **string** | The ID of the organization to move the SM seats to | 
**LicenseId** | **string** | The ID of the SM license to move the seats from | 
**SeatCount** | **int32** | The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license | 

## Methods

### NewInlineObject279

`func NewInlineObject279(destOrganizationId string, licenseId string, seatCount int32, ) *InlineObject279`

NewInlineObject279 instantiates a new InlineObject279 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject279WithDefaults

`func NewInlineObject279WithDefaults() *InlineObject279`

NewInlineObject279WithDefaults instantiates a new InlineObject279 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestOrganizationId

`func (o *InlineObject279) GetDestOrganizationId() string`

GetDestOrganizationId returns the DestOrganizationId field if non-nil, zero value otherwise.

### GetDestOrganizationIdOk

`func (o *InlineObject279) GetDestOrganizationIdOk() (*string, bool)`

GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOrganizationId

`func (o *InlineObject279) SetDestOrganizationId(v string)`

SetDestOrganizationId sets DestOrganizationId field to given value.


### GetLicenseId

`func (o *InlineObject279) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *InlineObject279) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *InlineObject279) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetSeatCount

`func (o *InlineObject279) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *InlineObject279) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *InlineObject279) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


