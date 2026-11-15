# InlineResponse200499Clients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impacted** | Pointer to [**InlineResponse200499ClientsImpacted**](InlineResponse200499ClientsImpacted.md) |  | [optional] 
**Total** | Pointer to **int32** | Total number of unique clients observed on the SSID. | [optional] 

## Methods

### NewInlineResponse200499Clients

`func NewInlineResponse200499Clients() *InlineResponse200499Clients`

NewInlineResponse200499Clients instantiates a new InlineResponse200499Clients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200499ClientsWithDefaults

`func NewInlineResponse200499ClientsWithDefaults() *InlineResponse200499Clients`

NewInlineResponse200499ClientsWithDefaults instantiates a new InlineResponse200499Clients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpacted

`func (o *InlineResponse200499Clients) GetImpacted() InlineResponse200499ClientsImpacted`

GetImpacted returns the Impacted field if non-nil, zero value otherwise.

### GetImpactedOk

`func (o *InlineResponse200499Clients) GetImpactedOk() (*InlineResponse200499ClientsImpacted, bool)`

GetImpactedOk returns a tuple with the Impacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacted

`func (o *InlineResponse200499Clients) SetImpacted(v InlineResponse200499ClientsImpacted)`

SetImpacted sets Impacted field to given value.

### HasImpacted

`func (o *InlineResponse200499Clients) HasImpacted() bool`

HasImpacted returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200499Clients) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200499Clients) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200499Clients) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200499Clients) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


