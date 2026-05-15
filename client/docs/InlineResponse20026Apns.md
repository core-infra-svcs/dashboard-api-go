# InlineResponse20026Apns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | APN name. | 
**AllowedIpTypes** | **[]string** | IP versions to support (permitted values include &#39;ipv4&#39;, &#39;ipv6&#39;). | 
**Authentication** | Pointer to [**InlineResponse20026Authentication**](InlineResponse20026Authentication.md) |  | [optional] 

## Methods

### NewInlineResponse20026Apns

`func NewInlineResponse20026Apns(name string, allowedIpTypes []string, ) *InlineResponse20026Apns`

NewInlineResponse20026Apns instantiates a new InlineResponse20026Apns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20026ApnsWithDefaults

`func NewInlineResponse20026ApnsWithDefaults() *InlineResponse20026Apns`

NewInlineResponse20026ApnsWithDefaults instantiates a new InlineResponse20026Apns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20026Apns) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20026Apns) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20026Apns) SetName(v string)`

SetName sets Name field to given value.


### GetAllowedIpTypes

`func (o *InlineResponse20026Apns) GetAllowedIpTypes() []string`

GetAllowedIpTypes returns the AllowedIpTypes field if non-nil, zero value otherwise.

### GetAllowedIpTypesOk

`func (o *InlineResponse20026Apns) GetAllowedIpTypesOk() (*[]string, bool)`

GetAllowedIpTypesOk returns a tuple with the AllowedIpTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpTypes

`func (o *InlineResponse20026Apns) SetAllowedIpTypes(v []string)`

SetAllowedIpTypes sets AllowedIpTypes field to given value.


### GetAuthentication

`func (o *InlineResponse20026Apns) GetAuthentication() InlineResponse20026Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse20026Apns) GetAuthenticationOk() (*InlineResponse20026Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse20026Apns) SetAuthentication(v InlineResponse20026Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse20026Apns) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


