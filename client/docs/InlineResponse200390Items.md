# InlineResponse200390Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**History** | Pointer to [**[]InlineResponse200390History**](InlineResponse200390History.md) | Site level connectivity history | [optional] 
**Devices** | Pointer to [**InlineResponse200390Devices**](InlineResponse200390Devices.md) |  | [optional] 

## Methods

### NewInlineResponse200390Items

`func NewInlineResponse200390Items() *InlineResponse200390Items`

NewInlineResponse200390Items instantiates a new InlineResponse200390Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200390ItemsWithDefaults

`func NewInlineResponse200390ItemsWithDefaults() *InlineResponse200390Items`

NewInlineResponse200390ItemsWithDefaults instantiates a new InlineResponse200390Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *InlineResponse200390Items) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200390Items) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200390Items) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200390Items) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200390Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200390Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200390Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200390Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHistory

`func (o *InlineResponse200390Items) GetHistory() []InlineResponse200390History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *InlineResponse200390Items) GetHistoryOk() (*[]InlineResponse200390History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *InlineResponse200390Items) SetHistory(v []InlineResponse200390History)`

SetHistory sets History field to given value.

### HasHistory

`func (o *InlineResponse200390Items) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200390Items) GetDevices() InlineResponse200390Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200390Items) GetDevicesOk() (*InlineResponse200390Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200390Items) SetDevices(v InlineResponse200390Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200390Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


