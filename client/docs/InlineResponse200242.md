# InlineResponse200242

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of the client | [optional] 
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the client | [optional] 
**Records** | Pointer to [**[]InlineResponse200242Records**](InlineResponse200242Records.md) | The clients that appear on any networks within an organization | [optional] 

## Methods

### NewInlineResponse200242

`func NewInlineResponse200242() *InlineResponse200242`

NewInlineResponse200242 instantiates a new InlineResponse200242 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200242WithDefaults

`func NewInlineResponse200242WithDefaults() *InlineResponse200242`

NewInlineResponse200242WithDefaults instantiates a new InlineResponse200242 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InlineResponse200242) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse200242) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse200242) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse200242) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200242) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200242) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200242) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200242) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManufacturer

`func (o *InlineResponse200242) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InlineResponse200242) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InlineResponse200242) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InlineResponse200242) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetRecords

`func (o *InlineResponse200242) GetRecords() []InlineResponse200242Records`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *InlineResponse200242) GetRecordsOk() (*[]InlineResponse200242Records, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *InlineResponse200242) SetRecords(v []InlineResponse200242Records)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *InlineResponse200242) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


