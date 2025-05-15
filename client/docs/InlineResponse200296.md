# InlineResponse200296

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestOrganizationId** | Pointer to **string** | The ID of the organization to move the SM seats to | [optional] 
**LicenseId** | Pointer to **string** | The ID of the SM license to move the seats from | [optional] 
**SeatCount** | Pointer to **int32** | The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license | [optional] 

## Methods

### NewInlineResponse200296

`func NewInlineResponse200296() *InlineResponse200296`

NewInlineResponse200296 instantiates a new InlineResponse200296 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200296WithDefaults

`func NewInlineResponse200296WithDefaults() *InlineResponse200296`

NewInlineResponse200296WithDefaults instantiates a new InlineResponse200296 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestOrganizationId

`func (o *InlineResponse200296) GetDestOrganizationId() string`

GetDestOrganizationId returns the DestOrganizationId field if non-nil, zero value otherwise.

### GetDestOrganizationIdOk

`func (o *InlineResponse200296) GetDestOrganizationIdOk() (*string, bool)`

GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOrganizationId

`func (o *InlineResponse200296) SetDestOrganizationId(v string)`

SetDestOrganizationId sets DestOrganizationId field to given value.

### HasDestOrganizationId

`func (o *InlineResponse200296) HasDestOrganizationId() bool`

HasDestOrganizationId returns a boolean if a field has been set.

### GetLicenseId

`func (o *InlineResponse200296) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *InlineResponse200296) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *InlineResponse200296) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *InlineResponse200296) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetSeatCount

`func (o *InlineResponse200296) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *InlineResponse200296) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *InlineResponse200296) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.

### HasSeatCount

`func (o *InlineResponse200296) HasSeatCount() bool`

HasSeatCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


