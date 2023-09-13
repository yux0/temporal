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

package replication

import (
	"sync"

	"github.com/dgryski/go-farm"

	replicationspb "go.temporal.io/server/api/replication/v1"

	"go.temporal.io/server/common/collection"
	"go.temporal.io/server/common/definition"
	ctasks "go.temporal.io/server/common/tasks"
)

type (
	SequentialTaskQueue struct {
		id interface{}

		sync.Mutex
		taskQueue collection.Queue[TrackableExecutableTask]
	}
)

func NewSequentialTaskQueue(task TrackableExecutableTask) ctasks.SequentialTaskQueue[TrackableExecutableTask] {
	return &SequentialTaskQueue{
		id: task.QueueID(),

		taskQueue: collection.NewPriorityQueue[TrackableExecutableTask](
			SequentialTaskQueueCompareLess,
		),
	}
}

func (q *SequentialTaskQueue) ID() interface{} {
	return q.id
}

func (q *SequentialTaskQueue) Peek() TrackableExecutableTask {
	q.Lock()
	defer q.Unlock()
	return q.taskQueue.Peek()
}

func (q *SequentialTaskQueue) Add(task TrackableExecutableTask) {
	q.Lock()
	defer q.Unlock()
	q.taskQueue.Add(task)
}

func (q *SequentialTaskQueue) Remove() TrackableExecutableTask {
	q.Lock()
	defer q.Unlock()
	return q.taskQueue.Remove().(TrackableExecutableTask)
}

func (q *SequentialTaskQueue) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()
	return q.taskQueue.IsEmpty()
}

func (q *SequentialTaskQueue) Len() int {
	q.Lock()
	defer q.Unlock()
	return q.taskQueue.Len()
}

func SequentialTaskQueueCompareLess(this TrackableExecutableTask, that TrackableExecutableTask) bool {
	return this.TaskID() < that.TaskID()
}

func TaskHashFn(
	task interface{},
) uint32 {
	workflowKey := TaskWorkflowKey(task)
	if workflowKey == nil {
		return 0
	}
	idBytes := []byte(workflowKey.NamespaceID + "_" + workflowKey.WorkflowID + "_" + workflowKey.RunID)
	return farm.Fingerprint32(idBytes)
}

func TaskWorkflowKey(
	item interface{},
) *definition.WorkflowKey {
	if item == nil {
		return nil
	}
	replicationTask, ok := item.(*replicationspb.ReplicationTask)
	if !ok {
		return nil
	}

	switch attr := replicationTask.Attributes.(type) {
	case *replicationspb.ReplicationTask_HistoryTaskAttributes:
		key := definition.CreateWorkflowKey(attr.HistoryTaskAttributes)
		return &key
	case *replicationspb.ReplicationTask_SyncActivityTaskAttributes:
		key := definition.CreateWorkflowKey(attr.SyncActivityTaskAttributes)
		return &key
	case *replicationspb.ReplicationTask_RawAttributes:
		key := definition.CreateWorkflowKey(attr.RawAttributes)
		return &key
	case *replicationspb.ReplicationTask_SyncWorkflowStateTaskAttributes:
		workflowState := attr.SyncWorkflowStateTaskAttributes.GetWorkflowState()
		workflowKey := definition.NewWorkflowKey(
			workflowState.GetExecutionInfo().GetNamespaceId(),
			workflowState.GetExecutionInfo().GetWorkflowId(),
			workflowState.GetExecutionState().GetRunId(),
		)
		return &workflowKey
	default:
		return nil
	}
}
