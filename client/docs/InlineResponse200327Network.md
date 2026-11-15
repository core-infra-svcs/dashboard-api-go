# InlineResponse200327Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network ID. | [optional] 
**Name** | Pointer to **string** | Network name. | [optional] 
**Url** | Pointer to **NullableString** | Network Dashboard URL. Null when network details are unavailable. | [optional] 
**Type** | Pointer to **string** | Network type: Wireless, Wired, or Unknown. | [optional] 

## Methods

### NewInlineResponse200327Network

`func NewInlineResponse200327Network() *InlineResponse200327Network`

NewInlineResponse200327Network instantiates a new InlineResponse200327Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327NetworkWithDefaults

`func NewInlineResponse200327NetworkWithDefaults() *InlineResponse200327Network`

NewInlineResponse200327NetworkWithDefaults instantiates a new InlineResponse200327Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200327Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200327Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200327Network) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200327Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200327Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200327Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200327Network) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200327Network) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200327Network) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200327Network) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200327Network) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200327Network) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *InlineResponse200327Network) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *InlineResponse200327Network) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetType

`func (o *InlineResponse200327Network) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200327Network) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200327Network) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200327Network) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


