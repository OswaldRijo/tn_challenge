package service_conf

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"truenorth/packages/utils"
)

func getTranceValues(ctx context.Context) (string, string, string) {
	spanIDStr := "none"
	requestIDStr := "none"
	traceIDStr := "none"
	spanID := ctx.Value("spanId")
	if spanID != nil {
		spanIDStr = spanID.(string)
	}
	requestID := ctx.Value("requestId")
	if requestID != nil {
		requestIDStr = requestID.(string)
	}
	traceID := ctx.Value("traceId")
	if traceID != nil {
		traceIDStr = traceID.(string)
	}

	return traceIDStr, spanIDStr, requestIDStr
}

func clientInterceptor(s *utils.ServerStartUp) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string, req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		u, _ := uuid.NewUUID()
		traceId, spanId, _ := getTranceValues(ctx)
		md := metadata.New(map[string]string{
			"requestId": u.String(),
			"traceId":   traceId,
			"spanId":    spanId,
		})
		clientCtx := metadata.NewOutgoingContext(ctx, md)
		log.Infow(clientCtx, "Client request started",
			"Method requested:", method,
			"Started at", start.UTC().String())
		// Calls the handler
		err := invoker(clientCtx, method,
			req,
			reply,
			cc,
			opts...,
		)
		log.Infow(clientCtx, "Request ended ",
			"Method: ", method,
			"Request: ", req,
			"Response: ", reply,
			"Duration: ", time.Since(start),
			"Error: ", err)
		return err
	}
}

func withClientUnaryInterceptor(s *utils.ServerStartUp) grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor(s))
}

func GetClientConn(serverPath string) *grpc.ClientConn {
	s := GetLocalServerInstance()
	opts := []grpc.DialOption{
		withClientUnaryInterceptor(s),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(serverPath, opts...)
	if err != nil {
		panic("connection failed")
	}
	return conn
}
