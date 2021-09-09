// Code generated by sdkgen. DO NOT EDIT.

//nolint
package instancegroup

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	instancegroup "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1/instancegroup"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// InstanceGroupServiceClient is a instancegroup.InstanceGroupServiceClient with
// lazy GRPC connection initialization.
type InstanceGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Create(ctx context.Context, in *instancegroup.CreateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Create(ctx, in, opts...)
}

// CreateFromYaml implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) CreateFromYaml(ctx context.Context, in *instancegroup.CreateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).CreateFromYaml(ctx, in, opts...)
}

// Delete implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Delete(ctx context.Context, in *instancegroup.DeleteInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteInstances implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) DeleteInstances(ctx context.Context, in *instancegroup.DeleteInstancesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).DeleteInstances(ctx, in, opts...)
}

// Get implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Get(ctx context.Context, in *instancegroup.GetInstanceGroupRequest, opts ...grpc.CallOption) (*instancegroup.InstanceGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) List(ctx context.Context, in *instancegroup.ListInstanceGroupsRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).List(ctx, in, opts...)
}

type InstanceGroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupsRequest

	items []*instancegroup.InstanceGroup
}

func (c *InstanceGroupServiceClient) InstanceGroupIterator(ctx context.Context, req *instancegroup.ListInstanceGroupsRequest, opts ...grpc.CallOption) *InstanceGroupIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceGroupIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceGroupIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.InstanceGroups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupIterator) Take(size int64) ([]*instancegroup.InstanceGroup, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*instancegroup.InstanceGroup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *InstanceGroupIterator) TakeAll() ([]*instancegroup.InstanceGroup, error) {
	return it.Take(0)
}

func (it *InstanceGroupIterator) Value() *instancegroup.InstanceGroup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupIterator) Error() error {
	return it.err
}

// ListAccessBindings implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type InstanceGroupAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceGroupServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *InstanceGroupServiceClient) InstanceGroupAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *InstanceGroupAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceGroupAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceGroupAccessBindingsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *InstanceGroupAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *InstanceGroupAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupAccessBindingsIterator) Error() error {
	return it.err
}

// ListInstances implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ListInstances(ctx context.Context, in *instancegroup.ListInstanceGroupInstancesRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupInstancesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ListInstances(ctx, in, opts...)
}

type InstanceGroupInstancesIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupInstancesRequest

	items []*instancegroup.ManagedInstance
}

func (c *InstanceGroupServiceClient) InstanceGroupInstancesIterator(ctx context.Context, req *instancegroup.ListInstanceGroupInstancesRequest, opts ...grpc.CallOption) *InstanceGroupInstancesIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceGroupInstancesIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceGroupInstancesIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListInstances(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Instances
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupInstancesIterator) Take(size int64) ([]*instancegroup.ManagedInstance, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*instancegroup.ManagedInstance

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *InstanceGroupInstancesIterator) TakeAll() ([]*instancegroup.ManagedInstance, error) {
	return it.Take(0)
}

func (it *InstanceGroupInstancesIterator) Value() *instancegroup.ManagedInstance {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupInstancesIterator) Error() error {
	return it.err
}

// ListLogRecords implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ListLogRecords(ctx context.Context, in *instancegroup.ListInstanceGroupLogRecordsRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupLogRecordsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ListLogRecords(ctx, in, opts...)
}

type InstanceGroupLogRecordsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupLogRecordsRequest

	items []*instancegroup.LogRecord
}

func (c *InstanceGroupServiceClient) InstanceGroupLogRecordsIterator(ctx context.Context, req *instancegroup.ListInstanceGroupLogRecordsRequest, opts ...grpc.CallOption) *InstanceGroupLogRecordsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceGroupLogRecordsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceGroupLogRecordsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListLogRecords(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.LogRecords
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupLogRecordsIterator) Take(size int64) ([]*instancegroup.LogRecord, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*instancegroup.LogRecord

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *InstanceGroupLogRecordsIterator) TakeAll() ([]*instancegroup.LogRecord, error) {
	return it.Take(0)
}

func (it *InstanceGroupLogRecordsIterator) Value() *instancegroup.LogRecord {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupLogRecordsIterator) Error() error {
	return it.err
}

// ListOperations implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ListOperations(ctx context.Context, in *instancegroup.ListInstanceGroupOperationsRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type InstanceGroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupOperationsRequest

	items []*operation.Operation
}

func (c *InstanceGroupServiceClient) InstanceGroupOperationsIterator(ctx context.Context, req *instancegroup.ListInstanceGroupOperationsRequest, opts ...grpc.CallOption) *InstanceGroupOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &InstanceGroupOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *InstanceGroupOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *InstanceGroupOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *InstanceGroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupOperationsIterator) Error() error {
	return it.err
}

// PauseProcesses implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) PauseProcesses(ctx context.Context, in *instancegroup.PauseInstanceGroupProcessesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).PauseProcesses(ctx, in, opts...)
}

// ResumeProcesses implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ResumeProcesses(ctx context.Context, in *instancegroup.ResumeInstanceGroupProcessesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ResumeProcesses(ctx, in, opts...)
}

// SetAccessBindings implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Start implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Start(ctx context.Context, in *instancegroup.StartInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Start(ctx, in, opts...)
}

// Stop implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Stop(ctx context.Context, in *instancegroup.StopInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Stop(ctx, in, opts...)
}

// StopInstances implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) StopInstances(ctx context.Context, in *instancegroup.StopInstancesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).StopInstances(ctx, in, opts...)
}

// Update implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Update(ctx context.Context, in *instancegroup.UpdateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}

// UpdateFromYaml implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) UpdateFromYaml(ctx context.Context, in *instancegroup.UpdateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).UpdateFromYaml(ctx, in, opts...)
}
