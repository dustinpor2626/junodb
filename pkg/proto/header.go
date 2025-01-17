//
//  Copyright 2023 PayPal Inc.
//
//  Licensed to the Apache Software Foundation (ASF) under one or more
//  contributor license agreements.  See the NOTICE file distributed with
//  this work for additional information regarding copyright ownership.
//  The ASF licenses this file to You under the Apache License, Version 2.0
//  (the "License"); you may not use this file except in compliance with
//  the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package proto

import (
	"fmt"
)

type (
	opMsgFlagT uint8

	messageHeaderT struct {
		magic    uint16
		version  uint8
		typeFlag messageTypeFlagT
		msgSize  uint32
		opaque   uint32
	}
	operationalHeaderT struct {
		opCode          OpCode
		flags           opMsgFlagT
		shardIdOrStatus shardIdOrStatusT
	}

	componentHeaderT struct {
		szComp        uint32
		tagComp       uint8
		szCompPadding uint8
	}

	metaComponentHeaderT struct {
		componentHeaderT
		numFields       uint8
		szHeaderPadding uint8
	}

	payloadComponentHeaderT struct {
		componentHeaderT
		szNamespace uint8
		szKey       uint16
		szValue     uint32
	}
)

func (f opMsgFlagT) IsFlagReplicationSet() bool {
	return (f & 1) != 0
}

func (f *opMsgFlagT) SetReplicationFlag() {
	(*f) |= 1
}

func (f opMsgFlagT) IsFlagDeleteReplicationSet() bool {
	return (f & 0x5) == 0x5
}

func (f *opMsgFlagT) SetDeleteReplicationFlag() {
	(*f) |= 0x5
}

func (f opMsgFlagT) IsFlagMarkDeleteSet() bool {
	return (f & 0x2) != 0
}

func (f *opMsgFlagT) SetMarkDeleteFlag() {
	(*f) |= 0x2
}

func (h *messageHeaderT) reset() {
	h.magic = kMessageMagic
	h.version = kCurrentVersion
	h.typeFlag = 0
	h.msgSize = 0
	h.opaque = 0
}

func (h *messageHeaderT) SetAsResponse() {
	h.typeFlag.setAsResponse()
}

func (h *messageHeaderT) IsSupported() bool {
	if h.magic == kMessageMagic && h.version == kCurrentVersion {
		if h.typeFlag.getMessageType() == kOperationalMessageType {
			return true
		}
	}
	return false
}

func (h *messageHeaderT) GetMsgSize() uint32 {
	return h.msgSize
}

func (h *messageHeaderT) GetOpaque() uint32 {
	return h.opaque
}

func (h *messageHeaderT) SetOpaque(opaque uint32) {
	h.opaque = opaque
}

func (h *messageHeaderT) getMsgType() uint8 {
	return h.typeFlag.getMessageType()
}

func (h *messageHeaderT) PrettyPrint() {
	fmt.Println("\nHeader:")
	fmt.Printf("  Magic\t\t:%#X\n", h.magic)
	fmt.Printf("  Version\t:%d\n", h.version)
	fmt.Printf("  MessageType\t:%d\n", h.getMsgType())
	fmt.Printf("  MessageSize\t:%d\n", h.msgSize)
	fmt.Printf("  OPaque\t:%#X\n", h.opaque)
}
