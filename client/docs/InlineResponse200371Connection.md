# InlineResponse200371Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cell** | Pointer to [**InlineResponse200371ConnectionCell**](InlineResponse200371ConnectionCell.md) |  | [optional] 
**Tac** | Pointer to **NullableString** | Tracking Area Code in hexadecimal with a 0x prefix, representing the coverage area of one or more cell towers. Null when the device has not reported tower details. | [optional] 

## Methods

### NewInlineResponse200371Connection

`func NewInlineResponse200371Connection() *InlineResponse200371Connection`

NewInlineResponse200371Connection instantiates a new InlineResponse200371Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200371ConnectionWithDefaults

`func NewInlineResponse200371ConnectionWithDefaults() *InlineResponse200371Connection`

NewInlineResponse200371ConnectionWithDefaults instantiates a new InlineResponse200371Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCell

`func (o *InlineResponse200371Connection) GetCell() InlineResponse200371ConnectionCell`

GetCell returns the Cell field if non-nil, zero value otherwise.

### GetCellOk

`func (o *InlineResponse200371Connection) GetCellOk() (*InlineResponse200371ConnectionCell, bool)`

GetCellOk returns a tuple with the Cell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCell

`func (o *InlineResponse200371Connection) SetCell(v InlineResponse200371ConnectionCell)`

SetCell sets Cell field to given value.

### HasCell

`func (o *InlineResponse200371Connection) HasCell() bool`

HasCell returns a boolean if a field has been set.

### GetTac

`func (o *InlineResponse200371Connection) GetTac() string`

GetTac returns the Tac field if non-nil, zero value otherwise.

### GetTacOk

`func (o *InlineResponse200371Connection) GetTacOk() (*string, bool)`

GetTacOk returns a tuple with the Tac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTac

`func (o *InlineResponse200371Connection) SetTac(v string)`

SetTac sets Tac field to given value.

### HasTac

`func (o *InlineResponse200371Connection) HasTac() bool`

HasTac returns a boolean if a field has been set.

### SetTacNil

`func (o *InlineResponse200371Connection) SetTacNil(b bool)`

 SetTacNil sets the value for Tac to be an explicit nil

### UnsetTac
`func (o *InlineResponse200371Connection) UnsetTac()`

UnsetTac ensures that no value is present for Tac, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


