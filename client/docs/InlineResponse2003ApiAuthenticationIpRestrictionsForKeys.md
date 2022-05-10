# InlineResponse2003ApiAuthenticationIpRestrictionsForKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges. | [optional] 
**Ranges** | Pointer to **[]string** | List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets. | [optional] 

## Methods

### NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeys

`func NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeys() *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys`

NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeys instantiates a new InlineResponse2003ApiAuthenticationIpRestrictionsForKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeysWithDefaults

`func NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeysWithDefaults() *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys`

NewInlineResponse2003ApiAuthenticationIpRestrictionsForKeysWithDefaults instantiates a new InlineResponse2003ApiAuthenticationIpRestrictionsForKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetRanges() []string`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) GetRangesOk() (*[]string, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) SetRanges(v []string)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *InlineResponse2003ApiAuthenticationIpRestrictionsForKeys) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


