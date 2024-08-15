# InlineResponse20047Destinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address to test connectivity with | [optional] 
**Description** | Pointer to **string** | Description of the testing destination. Optional, defaults to an empty string | [optional] 
**Default** | Pointer to **bool** | Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed | [optional] 

## Methods

### NewInlineResponse20047Destinations

`func NewInlineResponse20047Destinations() *InlineResponse20047Destinations`

NewInlineResponse20047Destinations instantiates a new InlineResponse20047Destinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20047DestinationsWithDefaults

`func NewInlineResponse20047DestinationsWithDefaults() *InlineResponse20047Destinations`

NewInlineResponse20047DestinationsWithDefaults instantiates a new InlineResponse20047Destinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse20047Destinations) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20047Destinations) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20047Destinations) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20047Destinations) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20047Destinations) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20047Destinations) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20047Destinations) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20047Destinations) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefault

`func (o *InlineResponse20047Destinations) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *InlineResponse20047Destinations) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *InlineResponse20047Destinations) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *InlineResponse20047Destinations) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


