# InlineResponse200147

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | device ID | [optional] 
**SsidName** | Pointer to **string** | SSID name | [optional] 
**Name** | Pointer to **string** | device name | [optional] 
**Scope** | Pointer to **string** | scope | [optional] 
**Tags** | Pointer to **[]string** | device tags | [optional] 
**TimeboundType** | Pointer to **string** | type of access period, either a static range or a dynamic period | [optional] 
**SendExpirationEmails** | Pointer to **bool** | Send Email Notifications | [optional] 
**NotifyTimeBeforeAccessEnds** | Pointer to **int32** | Time before access expiration reminder email sends | [optional] 
**AdditionalEmailText** | Pointer to **string** | Optional email text | [optional] 
**AccessStartAt** | Pointer to **time.Time** | time that access starts | [optional] 
**AccessEndAt** | Pointer to **time.Time** | time that access ends | [optional] 

## Methods

### NewInlineResponse200147

`func NewInlineResponse200147() *InlineResponse200147`

NewInlineResponse200147 instantiates a new InlineResponse200147 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200147WithDefaults

`func NewInlineResponse200147WithDefaults() *InlineResponse200147`

NewInlineResponse200147WithDefaults instantiates a new InlineResponse200147 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200147) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200147) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200147) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200147) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSsidName

`func (o *InlineResponse200147) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *InlineResponse200147) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *InlineResponse200147) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *InlineResponse200147) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200147) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200147) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200147) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200147) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *InlineResponse200147) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineResponse200147) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineResponse200147) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineResponse200147) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200147) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200147) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200147) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200147) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeboundType

`func (o *InlineResponse200147) GetTimeboundType() string`

GetTimeboundType returns the TimeboundType field if non-nil, zero value otherwise.

### GetTimeboundTypeOk

`func (o *InlineResponse200147) GetTimeboundTypeOk() (*string, bool)`

GetTimeboundTypeOk returns a tuple with the TimeboundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeboundType

`func (o *InlineResponse200147) SetTimeboundType(v string)`

SetTimeboundType sets TimeboundType field to given value.

### HasTimeboundType

`func (o *InlineResponse200147) HasTimeboundType() bool`

HasTimeboundType returns a boolean if a field has been set.

### GetSendExpirationEmails

`func (o *InlineResponse200147) GetSendExpirationEmails() bool`

GetSendExpirationEmails returns the SendExpirationEmails field if non-nil, zero value otherwise.

### GetSendExpirationEmailsOk

`func (o *InlineResponse200147) GetSendExpirationEmailsOk() (*bool, bool)`

GetSendExpirationEmailsOk returns a tuple with the SendExpirationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendExpirationEmails

`func (o *InlineResponse200147) SetSendExpirationEmails(v bool)`

SetSendExpirationEmails sets SendExpirationEmails field to given value.

### HasSendExpirationEmails

`func (o *InlineResponse200147) HasSendExpirationEmails() bool`

HasSendExpirationEmails returns a boolean if a field has been set.

### GetNotifyTimeBeforeAccessEnds

`func (o *InlineResponse200147) GetNotifyTimeBeforeAccessEnds() int32`

GetNotifyTimeBeforeAccessEnds returns the NotifyTimeBeforeAccessEnds field if non-nil, zero value otherwise.

### GetNotifyTimeBeforeAccessEndsOk

`func (o *InlineResponse200147) GetNotifyTimeBeforeAccessEndsOk() (*int32, bool)`

GetNotifyTimeBeforeAccessEndsOk returns a tuple with the NotifyTimeBeforeAccessEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTimeBeforeAccessEnds

`func (o *InlineResponse200147) SetNotifyTimeBeforeAccessEnds(v int32)`

SetNotifyTimeBeforeAccessEnds sets NotifyTimeBeforeAccessEnds field to given value.

### HasNotifyTimeBeforeAccessEnds

`func (o *InlineResponse200147) HasNotifyTimeBeforeAccessEnds() bool`

HasNotifyTimeBeforeAccessEnds returns a boolean if a field has been set.

### GetAdditionalEmailText

`func (o *InlineResponse200147) GetAdditionalEmailText() string`

GetAdditionalEmailText returns the AdditionalEmailText field if non-nil, zero value otherwise.

### GetAdditionalEmailTextOk

`func (o *InlineResponse200147) GetAdditionalEmailTextOk() (*string, bool)`

GetAdditionalEmailTextOk returns a tuple with the AdditionalEmailText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEmailText

`func (o *InlineResponse200147) SetAdditionalEmailText(v string)`

SetAdditionalEmailText sets AdditionalEmailText field to given value.

### HasAdditionalEmailText

`func (o *InlineResponse200147) HasAdditionalEmailText() bool`

HasAdditionalEmailText returns a boolean if a field has been set.

### GetAccessStartAt

`func (o *InlineResponse200147) GetAccessStartAt() time.Time`

GetAccessStartAt returns the AccessStartAt field if non-nil, zero value otherwise.

### GetAccessStartAtOk

`func (o *InlineResponse200147) GetAccessStartAtOk() (*time.Time, bool)`

GetAccessStartAtOk returns a tuple with the AccessStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStartAt

`func (o *InlineResponse200147) SetAccessStartAt(v time.Time)`

SetAccessStartAt sets AccessStartAt field to given value.

### HasAccessStartAt

`func (o *InlineResponse200147) HasAccessStartAt() bool`

HasAccessStartAt returns a boolean if a field has been set.

### GetAccessEndAt

`func (o *InlineResponse200147) GetAccessEndAt() time.Time`

GetAccessEndAt returns the AccessEndAt field if non-nil, zero value otherwise.

### GetAccessEndAtOk

`func (o *InlineResponse200147) GetAccessEndAtOk() (*time.Time, bool)`

GetAccessEndAtOk returns a tuple with the AccessEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessEndAt

`func (o *InlineResponse200147) SetAccessEndAt(v time.Time)`

SetAccessEndAt sets AccessEndAt field to given value.

### HasAccessEndAt

`func (o *InlineResponse200147) HasAccessEndAt() bool`

HasAccessEndAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


