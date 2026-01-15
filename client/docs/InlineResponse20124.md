# InlineResponse20124

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlId** | **string** | ID of the CRL | 
**Data** | [**[]InlineResponse20124Data**](InlineResponse20124Data.md) | CRL parsed data | 
**IsDelta** | **bool** | Whether it&#39;s a delta CRL or not | 
**CaId** | **string** | ID of the CRL issuer | 
**CreatedAt** | Pointer to **string** | Date of CRL creation | [optional] 
**LastUpdatedAt** | Pointer to **string** | Date of CRL update | [optional] 

## Methods

### NewInlineResponse20124

`func NewInlineResponse20124(crlId string, data []InlineResponse20124Data, isDelta bool, caId string, ) *InlineResponse20124`

NewInlineResponse20124 instantiates a new InlineResponse20124 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20124WithDefaults

`func NewInlineResponse20124WithDefaults() *InlineResponse20124`

NewInlineResponse20124WithDefaults instantiates a new InlineResponse20124 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrlId

`func (o *InlineResponse20124) GetCrlId() string`

GetCrlId returns the CrlId field if non-nil, zero value otherwise.

### GetCrlIdOk

`func (o *InlineResponse20124) GetCrlIdOk() (*string, bool)`

GetCrlIdOk returns a tuple with the CrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlId

`func (o *InlineResponse20124) SetCrlId(v string)`

SetCrlId sets CrlId field to given value.


### GetData

`func (o *InlineResponse20124) GetData() []InlineResponse20124Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20124) GetDataOk() (*[]InlineResponse20124Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20124) SetData(v []InlineResponse20124Data)`

SetData sets Data field to given value.


### GetIsDelta

`func (o *InlineResponse20124) GetIsDelta() bool`

GetIsDelta returns the IsDelta field if non-nil, zero value otherwise.

### GetIsDeltaOk

`func (o *InlineResponse20124) GetIsDeltaOk() (*bool, bool)`

GetIsDeltaOk returns a tuple with the IsDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelta

`func (o *InlineResponse20124) SetIsDelta(v bool)`

SetIsDelta sets IsDelta field to given value.


### GetCaId

`func (o *InlineResponse20124) GetCaId() string`

GetCaId returns the CaId field if non-nil, zero value otherwise.

### GetCaIdOk

`func (o *InlineResponse20124) GetCaIdOk() (*string, bool)`

GetCaIdOk returns a tuple with the CaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaId

`func (o *InlineResponse20124) SetCaId(v string)`

SetCaId sets CaId field to given value.


### GetCreatedAt

`func (o *InlineResponse20124) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20124) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20124) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20124) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse20124) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse20124) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse20124) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse20124) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


