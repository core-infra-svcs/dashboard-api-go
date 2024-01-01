# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimKey** | **string** | The subscription&#39;s claim key | 
**OrganizationId** | **string** | The id of the organization claiming the subscription | 
**Name** | Pointer to **string** | Friendly name to identify the subscription | [optional] 
**Description** | Pointer to **string** | Extra details or notes about the subscription | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject(claimKey string, organizationId string, ) *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimKey

`func (o *InlineObject) GetClaimKey() string`

GetClaimKey returns the ClaimKey field if non-nil, zero value otherwise.

### GetClaimKeyOk

`func (o *InlineObject) GetClaimKeyOk() (*string, bool)`

GetClaimKeyOk returns a tuple with the ClaimKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimKey

`func (o *InlineObject) SetClaimKey(v string)`

SetClaimKey sets ClaimKey field to given value.


### GetOrganizationId

`func (o *InlineObject) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineObject) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineObject) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *InlineObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


