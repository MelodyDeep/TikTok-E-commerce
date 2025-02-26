// Code generated by Kitex v0.12.1. DO NOT EDIT.

package orderservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	order_service "order/rpc/order/kitex_gen/demo/order_service"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateOrder": kitex.NewMethodInfo(
		createOrderHandler,
		newCreateOrderArgs,
		newCreateOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"QueryOrderById": kitex.NewMethodInfo(
		queryOrderByIdHandler,
		newQueryOrderByIdArgs,
		newQueryOrderByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"QueryOrdersByUserId": kitex.NewMethodInfo(
		queryOrdersByUserIdHandler,
		newQueryOrdersByUserIdArgs,
		newQueryOrdersByUserIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateOrder": kitex.NewMethodInfo(
		updateOrderHandler,
		newUpdateOrderArgs,
		newUpdateOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateOrderStatus": kitex.NewMethodInfo(
		updateOrderStatusHandler,
		newUpdateOrderStatusArgs,
		newUpdateOrderStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateOrderAddresseeInfo": kitex.NewMethodInfo(
		updateOrderAddresseeInfoHandler,
		newUpdateOrderAddresseeInfoArgs,
		newUpdateOrderAddresseeInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order_service.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "order_service",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func createOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order_service.CreateOrderRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order_service.OrderService).CreateOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateOrderArgs:
		success, err := handler.(order_service.OrderService).CreateOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateOrderArgs() interface{} {
	return &CreateOrderArgs{}
}

func newCreateOrderResult() interface{} {
	return &CreateOrderResult{}
}

type CreateOrderArgs struct {
	Req *order_service.CreateOrderRequest
}

func (p *CreateOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order_service.CreateOrderRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateOrderArgs) Unmarshal(in []byte) error {
	msg := new(order_service.CreateOrderRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateOrderArgs_Req_DEFAULT *order_service.CreateOrderRequest

func (p *CreateOrderArgs) GetReq() *order_service.CreateOrderRequest {
	if !p.IsSetReq() {
		return CreateOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateOrderResult struct {
	Success *order_service.BaseResponse
}

var CreateOrderResult_Success_DEFAULT *order_service.BaseResponse

func (p *CreateOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order_service.BaseResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateOrderResult) Unmarshal(in []byte) error {
	msg := new(order_service.BaseResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateOrderResult) GetSuccess() *order_service.BaseResponse {
	if !p.IsSetSuccess() {
		return CreateOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order_service.BaseResponse)
}

func (p *CreateOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateOrderResult) GetResult() interface{} {
	return p.Success
}

func queryOrderByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order_service.QueryOrderByIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order_service.OrderService).QueryOrderById(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *QueryOrderByIdArgs:
		success, err := handler.(order_service.OrderService).QueryOrderById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryOrderByIdResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newQueryOrderByIdArgs() interface{} {
	return &QueryOrderByIdArgs{}
}

func newQueryOrderByIdResult() interface{} {
	return &QueryOrderByIdResult{}
}

type QueryOrderByIdArgs struct {
	Req *order_service.QueryOrderByIdRequest
}

func (p *QueryOrderByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order_service.QueryOrderByIdRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryOrderByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryOrderByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryOrderByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *QueryOrderByIdArgs) Unmarshal(in []byte) error {
	msg := new(order_service.QueryOrderByIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryOrderByIdArgs_Req_DEFAULT *order_service.QueryOrderByIdRequest

func (p *QueryOrderByIdArgs) GetReq() *order_service.QueryOrderByIdRequest {
	if !p.IsSetReq() {
		return QueryOrderByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryOrderByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryOrderByIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryOrderByIdResult struct {
	Success *order_service.QueryOrderResponse
}

var QueryOrderByIdResult_Success_DEFAULT *order_service.QueryOrderResponse

func (p *QueryOrderByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order_service.QueryOrderResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryOrderByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryOrderByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryOrderByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *QueryOrderByIdResult) Unmarshal(in []byte) error {
	msg := new(order_service.QueryOrderResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryOrderByIdResult) GetSuccess() *order_service.QueryOrderResponse {
	if !p.IsSetSuccess() {
		return QueryOrderByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryOrderByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*order_service.QueryOrderResponse)
}

func (p *QueryOrderByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryOrderByIdResult) GetResult() interface{} {
	return p.Success
}

func queryOrdersByUserIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order_service.QueryOrdersByUserIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order_service.OrderService).QueryOrdersByUserId(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *QueryOrdersByUserIdArgs:
		success, err := handler.(order_service.OrderService).QueryOrdersByUserId(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryOrdersByUserIdResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newQueryOrdersByUserIdArgs() interface{} {
	return &QueryOrdersByUserIdArgs{}
}

func newQueryOrdersByUserIdResult() interface{} {
	return &QueryOrdersByUserIdResult{}
}

type QueryOrdersByUserIdArgs struct {
	Req *order_service.QueryOrdersByUserIdRequest
}

func (p *QueryOrdersByUserIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order_service.QueryOrdersByUserIdRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryOrdersByUserIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryOrdersByUserIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryOrdersByUserIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *QueryOrdersByUserIdArgs) Unmarshal(in []byte) error {
	msg := new(order_service.QueryOrdersByUserIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryOrdersByUserIdArgs_Req_DEFAULT *order_service.QueryOrdersByUserIdRequest

func (p *QueryOrdersByUserIdArgs) GetReq() *order_service.QueryOrdersByUserIdRequest {
	if !p.IsSetReq() {
		return QueryOrdersByUserIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryOrdersByUserIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryOrdersByUserIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryOrdersByUserIdResult struct {
	Success *order_service.QueryOrdersResponse
}

var QueryOrdersByUserIdResult_Success_DEFAULT *order_service.QueryOrdersResponse

func (p *QueryOrdersByUserIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order_service.QueryOrdersResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryOrdersByUserIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryOrdersByUserIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryOrdersByUserIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *QueryOrdersByUserIdResult) Unmarshal(in []byte) error {
	msg := new(order_service.QueryOrdersResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryOrdersByUserIdResult) GetSuccess() *order_service.QueryOrdersResponse {
	if !p.IsSetSuccess() {
		return QueryOrdersByUserIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryOrdersByUserIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*order_service.QueryOrdersResponse)
}

func (p *QueryOrdersByUserIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryOrdersByUserIdResult) GetResult() interface{} {
	return p.Success
}

func updateOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order_service.UpdateOrderRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order_service.OrderService).UpdateOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateOrderArgs:
		success, err := handler.(order_service.OrderService).UpdateOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateOrderArgs() interface{} {
	return &UpdateOrderArgs{}
}

func newUpdateOrderResult() interface{} {
	return &UpdateOrderResult{}
}

type UpdateOrderArgs struct {
	Req *order_service.UpdateOrderRequest
}

func (p *UpdateOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order_service.UpdateOrderRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateOrderArgs) Unmarshal(in []byte) error {
	msg := new(order_service.UpdateOrderRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateOrderArgs_Req_DEFAULT *order_service.UpdateOrderRequest

func (p *UpdateOrderArgs) GetReq() *order_service.UpdateOrderRequest {
	if !p.IsSetReq() {
		return UpdateOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateOrderResult struct {
	Success *order_service.BaseResponse
}

var UpdateOrderResult_Success_DEFAULT *order_service.BaseResponse

func (p *UpdateOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order_service.BaseResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateOrderResult) Unmarshal(in []byte) error {
	msg := new(order_service.BaseResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateOrderResult) GetSuccess() *order_service.BaseResponse {
	if !p.IsSetSuccess() {
		return UpdateOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order_service.BaseResponse)
}

func (p *UpdateOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateOrderResult) GetResult() interface{} {
	return p.Success
}

func updateOrderStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order_service.UpdateOrderStatusRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order_service.OrderService).UpdateOrderStatus(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateOrderStatusArgs:
		success, err := handler.(order_service.OrderService).UpdateOrderStatus(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateOrderStatusResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateOrderStatusArgs() interface{} {
	return &UpdateOrderStatusArgs{}
}

func newUpdateOrderStatusResult() interface{} {
	return &UpdateOrderStatusResult{}
}

type UpdateOrderStatusArgs struct {
	Req *order_service.UpdateOrderStatusRequest
}

func (p *UpdateOrderStatusArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order_service.UpdateOrderStatusRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateOrderStatusArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateOrderStatusArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateOrderStatusArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateOrderStatusArgs) Unmarshal(in []byte) error {
	msg := new(order_service.UpdateOrderStatusRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateOrderStatusArgs_Req_DEFAULT *order_service.UpdateOrderStatusRequest

func (p *UpdateOrderStatusArgs) GetReq() *order_service.UpdateOrderStatusRequest {
	if !p.IsSetReq() {
		return UpdateOrderStatusArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateOrderStatusArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateOrderStatusArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateOrderStatusResult struct {
	Success *order_service.BaseResponse
}

var UpdateOrderStatusResult_Success_DEFAULT *order_service.BaseResponse

func (p *UpdateOrderStatusResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order_service.BaseResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateOrderStatusResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateOrderStatusResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateOrderStatusResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateOrderStatusResult) Unmarshal(in []byte) error {
	msg := new(order_service.BaseResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateOrderStatusResult) GetSuccess() *order_service.BaseResponse {
	if !p.IsSetSuccess() {
		return UpdateOrderStatusResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateOrderStatusResult) SetSuccess(x interface{}) {
	p.Success = x.(*order_service.BaseResponse)
}

func (p *UpdateOrderStatusResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateOrderStatusResult) GetResult() interface{} {
	return p.Success
}

func updateOrderAddresseeInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order_service.UpdateOrderAddresseeInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order_service.OrderService).UpdateOrderAddresseeInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateOrderAddresseeInfoArgs:
		success, err := handler.(order_service.OrderService).UpdateOrderAddresseeInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateOrderAddresseeInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateOrderAddresseeInfoArgs() interface{} {
	return &UpdateOrderAddresseeInfoArgs{}
}

func newUpdateOrderAddresseeInfoResult() interface{} {
	return &UpdateOrderAddresseeInfoResult{}
}

type UpdateOrderAddresseeInfoArgs struct {
	Req *order_service.UpdateOrderAddresseeInfoRequest
}

func (p *UpdateOrderAddresseeInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order_service.UpdateOrderAddresseeInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateOrderAddresseeInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateOrderAddresseeInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateOrderAddresseeInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateOrderAddresseeInfoArgs) Unmarshal(in []byte) error {
	msg := new(order_service.UpdateOrderAddresseeInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateOrderAddresseeInfoArgs_Req_DEFAULT *order_service.UpdateOrderAddresseeInfoRequest

func (p *UpdateOrderAddresseeInfoArgs) GetReq() *order_service.UpdateOrderAddresseeInfoRequest {
	if !p.IsSetReq() {
		return UpdateOrderAddresseeInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateOrderAddresseeInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateOrderAddresseeInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateOrderAddresseeInfoResult struct {
	Success *order_service.BaseResponse
}

var UpdateOrderAddresseeInfoResult_Success_DEFAULT *order_service.BaseResponse

func (p *UpdateOrderAddresseeInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order_service.BaseResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateOrderAddresseeInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateOrderAddresseeInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateOrderAddresseeInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateOrderAddresseeInfoResult) Unmarshal(in []byte) error {
	msg := new(order_service.BaseResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateOrderAddresseeInfoResult) GetSuccess() *order_service.BaseResponse {
	if !p.IsSetSuccess() {
		return UpdateOrderAddresseeInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateOrderAddresseeInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*order_service.BaseResponse)
}

func (p *UpdateOrderAddresseeInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateOrderAddresseeInfoResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateOrder(ctx context.Context, Req *order_service.CreateOrderRequest) (r *order_service.BaseResponse, err error) {
	var _args CreateOrderArgs
	_args.Req = Req
	var _result CreateOrderResult
	if err = p.c.Call(ctx, "CreateOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryOrderById(ctx context.Context, Req *order_service.QueryOrderByIdRequest) (r *order_service.QueryOrderResponse, err error) {
	var _args QueryOrderByIdArgs
	_args.Req = Req
	var _result QueryOrderByIdResult
	if err = p.c.Call(ctx, "QueryOrderById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryOrdersByUserId(ctx context.Context, Req *order_service.QueryOrdersByUserIdRequest) (r *order_service.QueryOrdersResponse, err error) {
	var _args QueryOrdersByUserIdArgs
	_args.Req = Req
	var _result QueryOrdersByUserIdResult
	if err = p.c.Call(ctx, "QueryOrdersByUserId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateOrder(ctx context.Context, Req *order_service.UpdateOrderRequest) (r *order_service.BaseResponse, err error) {
	var _args UpdateOrderArgs
	_args.Req = Req
	var _result UpdateOrderResult
	if err = p.c.Call(ctx, "UpdateOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateOrderStatus(ctx context.Context, Req *order_service.UpdateOrderStatusRequest) (r *order_service.BaseResponse, err error) {
	var _args UpdateOrderStatusArgs
	_args.Req = Req
	var _result UpdateOrderStatusResult
	if err = p.c.Call(ctx, "UpdateOrderStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateOrderAddresseeInfo(ctx context.Context, Req *order_service.UpdateOrderAddresseeInfoRequest) (r *order_service.BaseResponse, err error) {
	var _args UpdateOrderAddresseeInfoArgs
	_args.Req = Req
	var _result UpdateOrderAddresseeInfoResult
	if err = p.c.Call(ctx, "UpdateOrderAddresseeInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
