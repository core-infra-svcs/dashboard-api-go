# InlineResponse200424Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**History** | Pointer to [**[]InlineResponse200424History**](InlineResponse200424History.md) | Site level connectivity history | [optional] 
**Devices** | Pointer to [**InlineResponse200424Devices**](InlineResponse200424Devices.md) |  | [optional] 

## Methods

### NewInlineResponse200424Items

`func NewInlineResponse200424Items() *InlineResponse200424Items`

NewInlineResponse200424Items instantiates a new InlineResponse200424Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200424ItemsWithDefaults

`func NewInlineResponse200424ItemsWithDefaults() *InlineResponse200424Items`

NewInlineResponse200424ItemsWithDefaults instantiates a new InlineResponse200424Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *InlineResponse200424Items) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200424Items) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200424Items) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200424Items) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200424Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200424Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200424Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200424Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHistory

`func (o *InlineResponse200424Items) GetHistory() []InlineResponse200424History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *InlineResponse200424Items) GetHistoryOk() (*[]InlineResponse200424History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *InlineResponse200424Items) SetHistory(v []InlineResponse200424History)`

SetHistory sets History field to given value.

### HasHistory

`func (o *InlineResponse200424Items) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200424Items) GetDevices() InlineResponse200424Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200424Items) GetDevicesOk() (*InlineResponse200424Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200424Items) SetDevices(v InlineResponse200424Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200424Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


