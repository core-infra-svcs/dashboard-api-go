# InlineObject203

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseIdToRenew** | **string** | The ID of the SM license to renew. This license must already be assigned to an SM network | 
**UnusedLicenseId** | **string** | The SM license to use to renew the seats on &#39;licenseIdToRenew&#39;. This license must have at least as many seats available as there are seats on &#39;licenseIdToRenew&#39; | 

## Methods

### NewInlineObject203

`func NewInlineObject203(licenseIdToRenew string, unusedLicenseId string, ) *InlineObject203`

NewInlineObject203 instantiates a new InlineObject203 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject203WithDefaults

`func NewInlineObject203WithDefaults() *InlineObject203`

NewInlineObject203WithDefaults instantiates a new InlineObject203 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseIdToRenew

`func (o *InlineObject203) GetLicenseIdToRenew() string`

GetLicenseIdToRenew returns the LicenseIdToRenew field if non-nil, zero value otherwise.

### GetLicenseIdToRenewOk

`func (o *InlineObject203) GetLicenseIdToRenewOk() (*string, bool)`

GetLicenseIdToRenewOk returns a tuple with the LicenseIdToRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseIdToRenew

`func (o *InlineObject203) SetLicenseIdToRenew(v string)`

SetLicenseIdToRenew sets LicenseIdToRenew field to given value.


### GetUnusedLicenseId

`func (o *InlineObject203) GetUnusedLicenseId() string`

GetUnusedLicenseId returns the UnusedLicenseId field if non-nil, zero value otherwise.

### GetUnusedLicenseIdOk

`func (o *InlineObject203) GetUnusedLicenseIdOk() (*string, bool)`

GetUnusedLicenseIdOk returns a tuple with the UnusedLicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedLicenseId

`func (o *InlineObject203) SetUnusedLicenseId(v string)`

SetUnusedLicenseId sets UnusedLicenseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


