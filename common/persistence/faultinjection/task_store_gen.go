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

// Code generated by cmd/tools/genfaultinjection. DO NOT EDIT.

package faultinjection

import (
	"context"

	"go.temporal.io/server/common/persistence"
)

type (
	faultInjectionTaskStore struct {
		baseStore persistence.TaskStore
		generator faultGenerator
	}
)

func newFaultInjectionTaskStore(
	baseStore persistence.TaskStore,
	generator faultGenerator,
) *faultInjectionTaskStore {
	return &faultInjectionTaskStore{
		baseStore: baseStore,
		generator: generator,
	}
}

func (c *faultInjectionTaskStore) Close() {
	c.baseStore.Close()
}

func (c *faultInjectionTaskStore) CompleteTasksLessThan(
	ctx context.Context,
	request *persistence.CompleteTasksLessThanRequest,
) (int, error) {
	return inject1(c.generator.generate("CompleteTasksLessThan"), func() (int, error) {
		return c.baseStore.CompleteTasksLessThan(ctx, request)
	})
}

func (c *faultInjectionTaskStore) CountTaskQueuesByBuildId(
	ctx context.Context,
	request *persistence.CountTaskQueuesByBuildIdRequest,
) (int, error) {
	return inject1(c.generator.generate("CountTaskQueuesByBuildId"), func() (int, error) {
		return c.baseStore.CountTaskQueuesByBuildId(ctx, request)
	})
}

func (c *faultInjectionTaskStore) CreateTaskQueue(
	ctx context.Context,
	request *persistence.InternalCreateTaskQueueRequest,
) error {
	return inject0(c.generator.generate("CreateTaskQueue"), func() error {
		return c.baseStore.CreateTaskQueue(ctx, request)
	})
}

func (c *faultInjectionTaskStore) CreateTasks(
	ctx context.Context,
	request *persistence.InternalCreateTasksRequest,
) (*persistence.CreateTasksResponse, error) {
	return inject1(c.generator.generate("CreateTasks"), func() (*persistence.CreateTasksResponse, error) {
		return c.baseStore.CreateTasks(ctx, request)
	})
}

func (c *faultInjectionTaskStore) DeleteTaskQueue(
	ctx context.Context,
	request *persistence.DeleteTaskQueueRequest,
) error {
	return inject0(c.generator.generate("DeleteTaskQueue"), func() error {
		return c.baseStore.DeleteTaskQueue(ctx, request)
	})
}

func (c *faultInjectionTaskStore) GetName() string {
	return c.baseStore.GetName()
}

func (c *faultInjectionTaskStore) GetTaskQueue(
	ctx context.Context,
	request *persistence.InternalGetTaskQueueRequest,
) (*persistence.InternalGetTaskQueueResponse, error) {
	return inject1(c.generator.generate("GetTaskQueue"), func() (*persistence.InternalGetTaskQueueResponse, error) {
		return c.baseStore.GetTaskQueue(ctx, request)
	})
}

func (c *faultInjectionTaskStore) GetTaskQueueUserData(
	ctx context.Context,
	request *persistence.GetTaskQueueUserDataRequest,
) (*persistence.InternalGetTaskQueueUserDataResponse, error) {
	return inject1(c.generator.generate("GetTaskQueueUserData"), func() (*persistence.InternalGetTaskQueueUserDataResponse, error) {
		return c.baseStore.GetTaskQueueUserData(ctx, request)
	})
}

func (c *faultInjectionTaskStore) GetTaskQueuesByBuildId(
	ctx context.Context,
	request *persistence.GetTaskQueuesByBuildIdRequest,
) ([]string, error) {
	return inject1(c.generator.generate("GetTaskQueuesByBuildId"), func() ([]string, error) {
		return c.baseStore.GetTaskQueuesByBuildId(ctx, request)
	})
}

func (c *faultInjectionTaskStore) GetTasks(
	ctx context.Context,
	request *persistence.GetTasksRequest,
) (*persistence.InternalGetTasksResponse, error) {
	return inject1(c.generator.generate("GetTasks"), func() (*persistence.InternalGetTasksResponse, error) {
		return c.baseStore.GetTasks(ctx, request)
	})
}

func (c *faultInjectionTaskStore) ListTaskQueue(
	ctx context.Context,
	request *persistence.ListTaskQueueRequest,
) (*persistence.InternalListTaskQueueResponse, error) {
	return inject1(c.generator.generate("ListTaskQueue"), func() (*persistence.InternalListTaskQueueResponse, error) {
		return c.baseStore.ListTaskQueue(ctx, request)
	})
}

func (c *faultInjectionTaskStore) ListTaskQueueUserDataEntries(
	ctx context.Context,
	request *persistence.ListTaskQueueUserDataEntriesRequest,
) (*persistence.InternalListTaskQueueUserDataEntriesResponse, error) {
	return inject1(c.generator.generate("ListTaskQueueUserDataEntries"), func() (*persistence.InternalListTaskQueueUserDataEntriesResponse, error) {
		return c.baseStore.ListTaskQueueUserDataEntries(ctx, request)
	})
}

func (c *faultInjectionTaskStore) UpdateTaskQueue(
	ctx context.Context,
	request *persistence.InternalUpdateTaskQueueRequest,
) (*persistence.UpdateTaskQueueResponse, error) {
	return inject1(c.generator.generate("UpdateTaskQueue"), func() (*persistence.UpdateTaskQueueResponse, error) {
		return c.baseStore.UpdateTaskQueue(ctx, request)
	})
}

func (c *faultInjectionTaskStore) UpdateTaskQueueUserData(
	ctx context.Context,
	request *persistence.InternalUpdateTaskQueueUserDataRequest,
) error {
	return inject0(c.generator.generate("UpdateTaskQueueUserData"), func() error {
		return c.baseStore.UpdateTaskQueueUserData(ctx, request)
	})
}
