// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/enums/v1/task.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TaskSource is the source from which a task was produced.
type TaskSource int32

const (
	TASK_SOURCE_UNSPECIFIED TaskSource = 0
	// Task produced by history service.
	TASK_SOURCE_HISTORY TaskSource = 1
	// Task produced from matching db backlog.
	TASK_SOURCE_DB_BACKLOG TaskSource = 2
)

var TaskSource_name = map[int32]string{
	0: "Unspecified",
	1: "History",
	2: "DbBacklog",
}

var TaskSource_value = map[string]int32{
	"Unspecified": 0,
	"History":     1,
	"DbBacklog":   2,
}

func (TaskSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36a3d3674ca3cfa6, []int{0}
}

type TaskCategory int32

const (
	TASK_CATEGORY_UNSPECIFIED TaskCategory = 0
	// Transfer is the task type for transfer task.
	TASK_CATEGORY_TRANSFER TaskCategory = 1
	// Timer is the task type for timer task.
	TASK_CATEGORY_TIMER TaskCategory = 2
	// Replication is the task type for replication task.
	TASK_CATEGORY_REPLICATION TaskCategory = 3
	// Visibility is the task type for visibility task.
	TASK_CATEGORY_VISIBILITY TaskCategory = 4
)

var TaskCategory_name = map[int32]string{
	0: "Unspecified",
	1: "Transfer",
	2: "Timer",
	3: "Replication",
	4: "Visibility",
}

var TaskCategory_value = map[string]int32{
	"Unspecified": 0,
	"Transfer":    1,
	"Timer":       2,
	"Replication": 3,
	"Visibility":  4,
}

func (TaskCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36a3d3674ca3cfa6, []int{1}
}

type TaskType int32

const (
	TASK_TYPE_UNSPECIFIED                    TaskType = 0
	TASK_TYPE_REPLICATION_HISTORY            TaskType = 1
	TASK_TYPE_REPLICATION_SYNC_ACTIVITY      TaskType = 2
	TASK_TYPE_TRANSFER_WORKFLOW_TASK         TaskType = 3
	TASK_TYPE_TRANSFER_ACTIVITY_TASK         TaskType = 4
	TASK_TYPE_TRANSFER_CLOSE_EXECUTION       TaskType = 5
	TASK_TYPE_TRANSFER_CANCEL_EXECUTION      TaskType = 6
	TASK_TYPE_TRANSFER_START_CHILD_EXECUTION TaskType = 7
	TASK_TYPE_TRANSFER_SIGNAL_EXECUTION      TaskType = 8
	TASK_TYPE_TRANSFER_RESET_WORKFLOW        TaskType = 10
	TASK_TYPE_WORKFLOW_TASK_TIMEOUT          TaskType = 12
	TASK_TYPE_ACTIVITY_TIMEOUT               TaskType = 13
	TASK_TYPE_USER_TIMER                     TaskType = 14
	TASK_TYPE_WORKFLOW_RUN_TIMEOUT           TaskType = 15
	TASK_TYPE_DELETE_HISTORY_EVENT           TaskType = 16
	TASK_TYPE_ACTIVITY_RETRY_TIMER           TaskType = 17
	TASK_TYPE_WORKFLOW_BACKOFF_TIMER         TaskType = 18
	TASK_TYPE_VISIBILITY_START_EXECUTION     TaskType = 19
	TASK_TYPE_VISIBILITY_UPSERT_EXECUTION    TaskType = 20
	TASK_TYPE_VISIBILITY_CLOSE_EXECUTION     TaskType = 21
	TASK_TYPE_VISIBILITY_DELETE_EXECUTION    TaskType = 22
	TASK_TYPE_TIERED_STORAGE                 TaskType = 23
)

var TaskType_name = map[int32]string{
	0:  "Unspecified",
	1:  "ReplicationHistory",
	2:  "ReplicationSyncActivity",
	3:  "TransferWorkflowTask",
	4:  "TransferActivityTask",
	5:  "TransferCloseExecution",
	6:  "TransferCancelExecution",
	7:  "TransferStartChildExecution",
	8:  "TransferSignalExecution",
	10: "TransferResetWorkflow",
	12: "WorkflowTaskTimeout",
	13: "ActivityTimeout",
	14: "UserTimer",
	15: "WorkflowRunTimeout",
	16: "DeleteHistoryEvent",
	17: "ActivityRetryTimer",
	18: "WorkflowBackoffTimer",
	19: "VisibilityStartExecution",
	20: "VisibilityUpsertExecution",
	21: "VisibilityCloseExecution",
	22: "VisibilityDeleteExecution",
	23: "TieredStorage",
}

var TaskType_value = map[string]int32{
	"Unspecified":                 0,
	"ReplicationHistory":          1,
	"ReplicationSyncActivity":     2,
	"TransferWorkflowTask":        3,
	"TransferActivityTask":        4,
	"TransferCloseExecution":      5,
	"TransferCancelExecution":     6,
	"TransferStartChildExecution": 7,
	"TransferSignalExecution":     8,
	"TransferResetWorkflow":       10,
	"WorkflowTaskTimeout":         12,
	"ActivityTimeout":             13,
	"UserTimer":                   14,
	"WorkflowRunTimeout":          15,
	"DeleteHistoryEvent":          16,
	"ActivityRetryTimer":          17,
	"WorkflowBackoffTimer":        18,
	"VisibilityStartExecution":    19,
	"VisibilityUpsertExecution":   20,
	"VisibilityCloseExecution":    21,
	"VisibilityDeleteExecution":   22,
	"TieredStorage":               23,
}

func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36a3d3674ca3cfa6, []int{2}
}

func init() {
	proto.RegisterEnum("temporal.server.api.enums.v1.TaskSource", TaskSource_name, TaskSource_value)
	proto.RegisterEnum("temporal.server.api.enums.v1.TaskCategory", TaskCategory_name, TaskCategory_value)
	proto.RegisterEnum("temporal.server.api.enums.v1.TaskType", TaskType_name, TaskType_value)
}

func init() {
	proto.RegisterFile("temporal/server/api/enums/v1/task.proto", fileDescriptor_36a3d3674ca3cfa6)
}

var fileDescriptor_36a3d3674ca3cfa6 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xcd, 0x4e, 0xdb, 0x4e,
	0x14, 0xc5, 0x6d, 0xc8, 0x9f, 0x7f, 0xb8, 0xd0, 0x76, 0x3a, 0x7c, 0x53, 0x98, 0x96, 0x00, 0x85,
	0xa2, 0x2a, 0x11, 0xea, 0xb2, 0x2b, 0x67, 0x32, 0x09, 0x23, 0x5c, 0x3b, 0x9a, 0x99, 0x40, 0xd3,
	0x05, 0x56, 0x5a, 0x59, 0x08, 0x51, 0xea, 0xc8, 0x09, 0x48, 0xec, 0xfa, 0x08, 0x7d, 0x83, 0x6e,
	0xfb, 0x28, 0x5d, 0xd2, 0x1d, 0xcb, 0x62, 0x36, 0x5d, 0xf2, 0x08, 0x95, 0x4d, 0xe2, 0x0f, 0xea,
	0xec, 0x22, 0x9d, 0x5f, 0xce, 0xbd, 0xf7, 0xdc, 0xeb, 0x81, 0xad, 0xbe, 0x7b, 0xd6, 0xf5, 0xfc,
	0xce, 0xe7, 0x4a, 0xcf, 0xf5, 0x2f, 0x5c, 0xbf, 0xd2, 0xe9, 0x9e, 0x54, 0xdc, 0x2f, 0xe7, 0x67,
	0xbd, 0xca, 0xc5, 0x6e, 0xa5, 0xdf, 0xe9, 0x9d, 0x96, 0xbb, 0xbe, 0xd7, 0xf7, 0xf0, 0xca, 0x10,
	0x2c, 0xdf, 0x83, 0xe5, 0x4e, 0xf7, 0xa4, 0x1c, 0x81, 0xe5, 0x8b, 0xdd, 0x9d, 0x23, 0x00, 0xd5,
	0xe9, 0x9d, 0x4a, 0xef, 0xdc, 0xff, 0xe4, 0xe2, 0x67, 0xb0, 0xa0, 0x0c, 0xb9, 0xef, 0x48, 0xbb,
	0x25, 0x28, 0x73, 0x5a, 0x96, 0x6c, 0x32, 0xca, 0xeb, 0x9c, 0xd5, 0x90, 0x86, 0x17, 0x60, 0x26,
	0x2d, 0xee, 0x71, 0xa9, 0x6c, 0xd1, 0x46, 0x3a, 0x5e, 0x86, 0xf9, 0xb4, 0x50, 0xab, 0x3a, 0x55,
	0x83, 0xee, 0x9b, 0x76, 0x03, 0x8d, 0xed, 0x7c, 0xd7, 0x61, 0x3a, 0x2c, 0x40, 0x3b, 0x7d, 0xf7,
	0xd8, 0xf3, 0x2f, 0xf1, 0x2a, 0x2c, 0x45, 0x30, 0x35, 0x14, 0x6b, 0xd8, 0xa2, 0xfd, 0xa0, 0xc8,
	0xd0, 0x2b, 0x96, 0x95, 0x30, 0x2c, 0x59, 0x67, 0x02, 0xe9, 0x71, 0x03, 0x89, 0xc6, 0xdf, 0x31,
	0x81, 0xc6, 0xfe, 0xf5, 0x14, 0xac, 0x69, 0x72, 0x6a, 0x28, 0x6e, 0x5b, 0x68, 0x1c, 0xaf, 0xc0,
	0x62, 0x56, 0x3e, 0xe0, 0x92, 0x57, 0xb9, 0xc9, 0x55, 0x1b, 0x15, 0x76, 0x7e, 0x4d, 0x40, 0x31,
	0xec, 0x50, 0x5d, 0x76, 0x5d, 0xbc, 0x04, 0x73, 0x11, 0xaa, 0xda, 0xcd, 0x87, 0xe3, 0xaf, 0xc1,
	0x6a, 0x22, 0xa5, 0x0a, 0xa4, 0x82, 0xd8, 0x82, 0xf5, 0x7c, 0x44, 0xb6, 0x2d, 0xea, 0x18, 0x54,
	0xf1, 0x83, 0xb0, 0xe6, 0x18, 0xde, 0x80, 0x17, 0x09, 0x38, 0x9c, 0xd0, 0x39, 0xb4, 0xc5, 0x7e,
	0xdd, 0xb4, 0x0f, 0x9d, 0x50, 0x43, 0xe3, 0x23, 0xa8, 0xa1, 0xcd, 0x3d, 0x55, 0xc0, 0x2f, 0xa1,
	0x94, 0x43, 0x51, 0xd3, 0x96, 0xcc, 0x61, 0xef, 0x19, 0x6d, 0x45, 0x29, 0xfc, 0x97, 0x6d, 0x2e,
	0xe1, 0x0c, 0x8b, 0x32, 0x33, 0x05, 0x4e, 0xe0, 0xd7, 0xb0, 0x9d, 0x03, 0x4a, 0x65, 0x08, 0xe5,
	0xd0, 0x3d, 0x6e, 0xd6, 0x52, 0xf4, 0xff, 0x23, 0x6c, 0x25, 0x6f, 0x58, 0x46, 0xda, 0xb6, 0x88,
	0x37, 0x61, 0x2d, 0x07, 0x14, 0x4c, 0x32, 0x15, 0x4f, 0x8e, 0x00, 0xaf, 0xc3, 0xf3, 0x04, 0xcb,
	0x24, 0x12, 0xad, 0xdb, 0x6e, 0x29, 0x34, 0x8d, 0x09, 0x2c, 0x27, 0x50, 0x12, 0xc8, 0x40, 0x7f,
	0x84, 0x17, 0x61, 0x36, 0xb5, 0x46, 0xc9, 0xc4, 0xe0, 0x54, 0x1e, 0xe3, 0x12, 0x90, 0x1c, 0x7b,
	0xd1, 0xb2, 0xe2, 0x7f, 0x3f, 0xc9, 0x32, 0x35, 0x66, 0x32, 0x15, 0x5f, 0xbb, 0xc3, 0x0e, 0x98,
	0xa5, 0x10, 0xca, 0x32, 0x71, 0x07, 0x82, 0xa9, 0xf8, 0x2c, 0x9f, 0x66, 0xf7, 0x17, 0xd7, 0x0a,
	0xbf, 0x0d, 0xbb, 0x5e, 0x1f, 0x50, 0x18, 0x6f, 0xc3, 0x46, 0x42, 0x25, 0x97, 0x39, 0x08, 0x3c,
	0x49, 0x70, 0x06, 0xbf, 0x82, 0xcd, 0x5c, 0xb2, 0xd5, 0x94, 0x2c, 0x83, 0xce, 0x8e, 0x34, 0x7d,
	0x78, 0x16, 0x73, 0x23, 0x4d, 0x07, 0x73, 0x27, 0xe8, 0x7c, 0xfc, 0x1d, 0xdd, 0x6f, 0x90, 0x33,
	0xc1, 0x6a, 0x4e, 0x98, 0x8a, 0xd1, 0x60, 0x68, 0xa1, 0x54, 0x28, 0x4e, 0xa2, 0xc9, 0x52, 0xa1,
	0x38, 0x85, 0xa6, 0xaa, 0x47, 0x57, 0x37, 0x44, 0xbb, 0xbe, 0x21, 0xda, 0xdd, 0x0d, 0xd1, 0xbf,
	0x06, 0x44, 0xff, 0x11, 0x10, 0xfd, 0x67, 0x40, 0xf4, 0xab, 0x80, 0xe8, 0xbf, 0x03, 0xa2, 0xff,
	0x09, 0x88, 0x76, 0x17, 0x10, 0xfd, 0xdb, 0x2d, 0xd1, 0xae, 0x6e, 0x89, 0x76, 0x7d, 0x4b, 0xb4,
	0x0f, 0xdb, 0xc7, 0x5e, 0x39, 0x7e, 0xac, 0x4e, 0xbc, 0xbc, 0x87, 0xed, 0x6d, 0xf4, 0xe3, 0xe3,
	0x44, 0xf4, 0xb4, 0xbd, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x82, 0xb9, 0x79, 0xa8, 0x05, 0x05,
	0x00, 0x00,
}

func (x TaskSource) String() string {
	s, ok := TaskSource_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TaskCategory) String() string {
	s, ok := TaskCategory_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TaskType) String() string {
	s, ok := TaskType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
