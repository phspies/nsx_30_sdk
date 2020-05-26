# PacketCaptureRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | Define the transport node to capture data. | [optional] [default to null]
**Direction** | **string** | Define the capture direction. Support three types INPUT/OUTPUT/DUAL. | [optional] [default to null]
**Capduration** | **int32** | Define the packet capture duration time. After the capture duration time, the capture process will stop working. | [optional] [default to null]
**Capamount** | **int32** | Define the packet capture amount size. | [optional] [default to null]
**Capsource** | **string** | This type is used to differenite the incoming request from CLI/UI. | [default to null]
**NodeIp** | **string** | Define the transport node to capture data. | [optional] [default to null]
**Capvalue** | **string** | Define the capture value of given capture point. | [optional] [default to null]
**Filtertype** | **string** | Define the capture filter type. Support PRE/POST mode. | [optional] [default to null]
**Cappoint** | **string** | Define the point to capture data. | [default to null]
**Capfilesize** | **int32** | Define the packet capture file size limit. | [optional] [default to null]
**Options** | [***PacketCaptureOptionList**](PacketCaptureOptionList.md) |  | [optional] [default to null]
**Streamport** | **int32** | Set the stream port to receive the capture packet. The STREAM mode is based on GRE-in-UDP Encapsulation(RFC8086). Packets are sent to UDP port 4754. | [optional] [default to null]
**Caprate** | **int32** | Define the rate of packet capture process. | [optional] [default to null]
**Capcore** | **int32** | The CPU core id on Edge node. | [optional] [default to null]
**Capsnaplen** | **int32** | Limit the number of bytes captured from each packet. | [optional] [default to null]
**Streamaddress** | **string** | Set the stream address to receive the capture packet. | [optional] [default to null]
**Capmode** | **string** | Define the capture streaming mode. The STREAM mode will send the data to given stream address and port. And the STANDALONE mode will save the capture file in local folder. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

