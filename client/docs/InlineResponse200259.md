# InlineResponse200259

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the network or config template to copy configuration from | [optional] 
**Name** | Pointer to **string** | The name of the configuration template | [optional] 
**ProductTypes** | Pointer to **[]string** | The product types of the configuration template | [optional] 
**TimeZone** | Pointer to **string** | The timezone of the configuration template. For a list of allowed timezones, please see the &#39;TZ&#39; column in the table in &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones&#39;&gt;this article&lt;/a&gt;. Not applicable if copying from existing network or template | [optional] 

## Methods

### NewInlineResponse200259

`func NewInlineResponse200259() *InlineResponse200259`

NewInlineResponse200259 instantiates a new InlineResponse200259 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200259WithDefaults

`func NewInlineResponse200259WithDefaults() *InlineResponse200259`

NewInlineResponse200259WithDefaults instantiates a new InlineResponse200259 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200259) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200259) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200259) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200259) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200259) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200259) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200259) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200259) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200259) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200259) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200259) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200259) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetTimeZone

`func (o *InlineResponse200259) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *InlineResponse200259) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *InlineResponse200259) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *InlineResponse200259) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


