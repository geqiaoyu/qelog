package receiver

//import (
//	"context"
//	"github.com/bbdshow/qelog/api/receiverpb"
//	"net"
//
//	"github.com/bbdshow/qelog/infra/httputil"
//	"github.com/bbdshow/qelog/infra/kit"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/peer"
//)
//
//type GRPCService struct {
//	server   *grpc.Server
//	receiver *Service
//}
//
//func NewGRPCService() *GRPCService {
//	srv := &GRPCService{
//		server:   nil,
//		receiver: NewService(),
//	}
//	return srv
//}
//
//func (srv *GRPCService) Run(addr string) error {
//	listen, err := net.Listen("tcp", addr)
//	if err != nil {
//		return err
//	}
//
//	server := grpc.NewServer()
//	srv.server = server
//
//	receiverpb.RegisterReceiverServer(srv.server, srv)
//
//	if err := server.Serve(listen); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (srv *GRPCService) Close() error {
//	if srv.server != nil {
//		srv.server.Stop()
//	}
//	return nil
//}
//
//func (srv *GRPCService) PushPacket(ctx context.Context, in *receiverpb.Packet) (*receiverpb.BaseResp, error) {
//	// 获取 clientIP
//	if err := srv.receiver.InsertPacket(ctx, srv.clientIP(ctx), in); err != nil {
//		e, ok := err.(httputil.Error)
//		if ok {
//			// 数据库操作错误
//			if e.Code == httputil.ErrCodeSystemException {
//				return nil, httputil.ErrSystemException
//			}
//			return &receiverpb.BaseResp{
//				Code:    int32(e.Code),
//				Message: e.Message,
//			}, nil
//		}
//		return nil, err
//	}
//	return &receiverpb.BaseResp{
//		Code:    httputil.CodeSuccess,
//		Message: "success",
//	}, nil
//}
//
//func (srv *GRPCService) clientIP(ctx context.Context) string {
//	ctxPeer, ok := peer.FromContext(ctx)
//	if ok && ctxPeer.Addr != nil {
//		if ipnet, ok := ctxPeer.Addr.(*net.IPNet); ok {
//			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
//				return ipnet.IP.String()
//			}
//		}
//		// 上述解析不成功，则自行拼接
//		return kit.AddrStringToIP(ctxPeer.Addr)
//	}
//	return ""
//}
