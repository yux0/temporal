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

func GetRemoteShardIDs(localShardId int32, localShardCount int32, remoteShardCount int32) []int32 {
	var pollingShards []int32
	if remoteShardCount <= localShardCount {
		if localShardId <= remoteShardCount {
			pollingShards = append(pollingShards, localShardId)
		}
		return pollingShards
	} else {
		// remoteShardCount > localShardCount
		// 30 3
		// 147,10,13,16,19,22,25,28    258,11,14,17,20,23,26,29,    3,6,9,12,15,18,21,24,27,30
		// X % 30 = 7 X % 3 = 1
		// 101 % 30 = 11 101 % 3 = 2
		// 24  % 30 = 24 24 % 3 = 0 // 3
		for i := localShardId; i <= remoteShardCount; i += localShardCount {
			pollingShards = append(pollingShards, i)
		}
	}
	return pollingShards
}
