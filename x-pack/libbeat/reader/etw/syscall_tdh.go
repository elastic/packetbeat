// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package etw

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

var (
	tdh                                  = windows.NewLazySystemDLL("tdh.dll")
	tdhEnumerateProviders                = tdh.NewProc("TdhEnumerateProviders")
	tdhEnumerateProviderFieldInformation = tdh.NewProc("TdhEnumerateProviderFieldInformation")
	tdhGetEventInformation               = tdh.NewProc("TdhGetEventInformation")
	tdhGetEventMapInformation            = tdh.NewProc("TdhGetEventMapInformation")
	tdhFormatProperty                    = tdh.NewProc("TdhFormatProperty")
	tdhGetProperty                       = tdh.NewProc("TdhGetProperty")
)

const ANYSIZE_ARRAY = 1 << 25

type ProviderEnumerationInfo struct {
	NumberOfProviders      uint32
	Reserved               uint32
	TraceProviderInfoArray [ANYSIZE_ARRAY]TraceProviderInfo
}

type TraceProviderInfo struct {
	ProviderGuid       GUID
	SchemaSource       uint32
	ProviderNameOffset uint32
}

// Event structs
type EventRecord struct {
	EventHeader       EventHeader
	BufferContext     EtwBufferContext
	ExtendedDataCount uint16
	UserDataLength    uint16
	ExtendedData      EventHeaderExtendedDataItem
	UserData          uintptr // Event data
	UserContext       uintptr
}

const (
	EVENT_HEADER_FLAG_STRING_ONLY   = 0x0004
	EVENT_HEADER_FLAG_32_BIT_HEADER = 0x0020
	EVENT_HEADER_FLAG_64_BIT_HEADER = 0x0040
)

type EventHeader struct {
	Size            uint16
	HeaderType      uint16
	Flags           uint16
	EventProperty   uint16
	ThreadId        uint32
	ProcessId       uint32
	TimeStamp       int64
	ProviderId      GUID
	EventDescriptor EventDescriptor
	Time            int64
	ActivityId      GUID
}

func (e *EventRecord) PointerSize() uint32 {
	if e.EventHeader.Flags&EVENT_HEADER_FLAG_32_BIT_HEADER == EVENT_HEADER_FLAG_32_BIT_HEADER {
		return 4
	}
	return 8
}

type EventDescriptor struct {
	Id      uint16
	Version uint8
	Channel uint8
	Level   uint8
	Opcode  uint8
	Task    uint16
	Keyword uint64
}

type EtwBufferContext struct {
	Union    uint16
	LoggerId uint16
}

type EventHeaderExtendedDataItem struct {
	Reserved1      uint16
	ExtType        uint16
	InternalStruct uint16
	DataSize       uint16
	DataPtr        uint64
}

type TdhContext struct {
	ParameterValue uint32
	ParameterType  int32
	ParameterSize  uint32
}

type TraceEventInfo struct {
	ProviderGUID                GUID
	EventGUID                   GUID
	EventDescriptor             EventDescriptor
	DecodingSource              DecodingSource
	ProviderNameOffset          uint32
	LevelNameOffset             uint32
	ChannelNameOffset           uint32
	KeywordsNameOffset          uint32
	TaskNameOffset              uint32
	OpcodeNameOffset            uint32
	EventMessageOffset          uint32
	ProviderMessageOffset       uint32
	BinaryXMLOffset             uint32
	BinaryXMLSize               uint32
	ActivityIDNameOffset        uint32
	RelatedActivityIDNameOffset uint32
	PropertyCount               uint32
	TopLevelPropertyCount       uint32
	Flags                       TemplateFlags
	EventPropertyInfoArray      [ANYSIZE_ARRAY]EventPropertyInfo
}

type EventPropertyInfo struct {
	Flags      PropertyFlags
	NameOffset uint32
	TypeUnion  struct {
		u1 uint16
		u2 uint16
		u3 uint32
	}
	CountUnion  uint16
	LengthUnion uint16
	ResTagUnion uint32
}

func (i *EventPropertyInfo) Count() uint16 {
	return i.CountUnion
}

func (i *EventPropertyInfo) Length() uint16 {
	return i.LengthUnion
}

func (i *EventPropertyInfo) InType() uint16 {
	return i.TypeUnion.u1
}

func (i *EventPropertyInfo) OutType() uint16 {
	return i.TypeUnion.u2
}

func (i *EventPropertyInfo) StructStartIndex() uint16 {
	return i.InType()
}

func (i *EventPropertyInfo) NumOfStructMembers() uint16 {
	return i.OutType()
}

func (i *EventPropertyInfo) MapNameOffset() uint32 {
	return i.TypeUnion.u3
}

const (
	TdhIntypeBinary = 14
	TdhOuttypeIpv6  = 24
)

type DecodingSource int32
type TemplateFlags int32

type PropertyFlags int32

const (
	PropertyStruct           = PropertyFlags(0x1)
	PropertyParamLength      = PropertyFlags(0x2)
	PropertyParamCount       = PropertyFlags(0x4)
	PropertyWBEMXmlFragment  = PropertyFlags(0x8)
	PropertyParamFixedLength = PropertyFlags(0x10)
)

type EventMapInfo struct {
	NameOffset    uint32
	Flag          MapFlags
	EntryCount    uint32
	Union         uint32
	MapEntryArray [ANYSIZE_ARRAY]EventMapEntry
}

type MapFlags int32

type EventMapEntry struct {
	OutputOffset uint32
	Union        uint32
}

type PropertyDataDescriptor struct {
	PropertyName unsafe.Pointer
	ArrayIndex   uint32
	Reserved     uint32
}

func _TdhEnumerateProviders(
	pBuffer *ProviderEnumerationInfo,
	pBufferSize *uint32) error {
	r0, _, _ := tdhEnumerateProviders.Call(
		uintptr(unsafe.Pointer(pBuffer)),
		uintptr(unsafe.Pointer(pBufferSize)))
	if r0 == 0 {
		return nil
	}
	return syscall.Errno(r0)
}

func _TdhGetEventInformation(pEvent *EventRecord,
	tdhContextCount uint32,
	pTdhContext *TdhContext,
	pBuffer *TraceEventInfo,
	pBufferSize *uint32) error {
	r0, _, _ := tdhGetEventInformation.Call(
		uintptr(unsafe.Pointer(pEvent)),
		uintptr(tdhContextCount),
		uintptr(unsafe.Pointer(pTdhContext)),
		uintptr(unsafe.Pointer(pBuffer)),
		uintptr(unsafe.Pointer(pBufferSize)))
	if r0 == 0 {
		return nil
	}
	return syscall.Errno(r0)
}

func _TdhFormatProperty(
	eventInfo *TraceEventInfo,
	mapInfo *EventMapInfo,
	pointerSize uint32,
	propertyInType uint16,
	propertyOutType uint16,
	propertyLength uint16,
	userDataLength uint16,
	userData *byte,
	bufferSize *uint32,
	buffer *uint16,
	userDataConsumed *uint16) error {
	r0, _, _ := tdhFormatProperty.Call(
		uintptr(unsafe.Pointer(eventInfo)),
		uintptr(unsafe.Pointer(mapInfo)),
		uintptr(pointerSize),
		uintptr(propertyInType),
		uintptr(propertyOutType),
		uintptr(propertyLength),
		uintptr(userDataLength),
		uintptr(unsafe.Pointer(userData)),
		uintptr(unsafe.Pointer(bufferSize)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(userDataConsumed)))
	if r0 == 0 {
		return nil
	}
	return syscall.Errno(r0)
}

func _TdhGetProperty(pEvent *EventRecord,
	tdhContextCount uint32,
	pTdhContext *TdhContext,
	propertyDataCount uint32,
	pPropertyData *PropertyDataDescriptor,
	bufferSize uint32,
	pBuffer *byte) error {
	r0, _, _ := tdhGetProperty.Call(
		uintptr(unsafe.Pointer(pEvent)),
		uintptr(tdhContextCount),
		uintptr(unsafe.Pointer(pTdhContext)),
		uintptr(propertyDataCount),
		uintptr(unsafe.Pointer(pPropertyData)),
		uintptr(bufferSize),
		uintptr(unsafe.Pointer(pBuffer)))
	if r0 == 0 {
		return nil
	}
	return syscall.Errno(r0)
}

func _TdhGetEventMapInformation(pEvent *EventRecord,
	pMapName *uint16,
	pBuffer *EventMapInfo,
	pBufferSize *uint32) error {
	r0, _, _ := tdhGetEventMapInformation.Call(
		uintptr(unsafe.Pointer(pEvent)),
		uintptr(unsafe.Pointer(pMapName)),
		uintptr(unsafe.Pointer(pBuffer)),
		uintptr(unsafe.Pointer(pBufferSize)))
	if r0 == 0 {
		return nil
	}
	return syscall.Errno(r0)
}
