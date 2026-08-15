# InlineResponse200460Clients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impacted** | Pointer to [**InlineResponse200460ClientsImpacted**](InlineResponse200460ClientsImpacted.md) |  | [optional] 
**Total** | Pointer to **int32** | Total number of unique clients observed on the SSID. | [optional] 

## Methods

### NewInlineResponse200460Clients

`func NewInlineResponse200460Clients() *InlineResponse200460Clients`

NewInlineResponse200460Clients instantiates a new InlineResponse200460Clients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200460ClientsWithDefaults

`func NewInlineResponse200460ClientsWithDefaults() *InlineResponse200460Clients`

NewInlineResponse200460ClientsWithDefaults instantiates a new InlineResponse200460Clients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpacted

`func (o *InlineResponse200460Clients) GetImpacted() InlineResponse200460ClientsImpacted`

GetImpacted returns the Impacted field if non-nil, zero value otherwise.

### GetImpactedOk

`func (o *InlineResponse200460Clients) GetImpactedOk() (*InlineResponse200460ClientsImpacted, bool)`

GetImpactedOk returns a tuple with the Impacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacted

`func (o *InlineResponse200460Clients) SetImpacted(v InlineResponse200460ClientsImpacted)`

SetImpacted sets Impacted field to given value.

### HasImpacted

`func (o *InlineResponse200460Clients) HasImpacted() bool`

HasImpacted returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200460Clients) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200460Clients) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200460Clients) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200460Clients) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


