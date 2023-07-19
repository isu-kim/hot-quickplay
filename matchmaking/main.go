package main

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	pb "simple-mm/proto/simple-mm"
	"simple-mm/utils"
	"sync"
	"time"
)

const (
	hostAddr  = ":50051"
	jaegerURL = "http://127.0.0.1:14268/api/traces"
	minMaxMMR = 15
)

type matchMakingServer struct {
	pb.UnimplementedMatchmakingServiceServer
	tr      trace.Tracer
	mu      sync.Mutex
	mqUsers []string
	mqMMR   []int64
}

// mmUser will store a match making user in the queue.
type mmUser struct {
	username string
	mmr      int64
}

// removeString removes a string entry from slice
// https://stackoverflow.com/questions/37334119/
func removeString(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// removeInt removes a string entry from slice
// https://stackoverflow.com/questions/37334119/
func removeInt(slice []int64, s int) []int64 {
	return append(slice[:s], slice[s+1:]...)
}

// checkQueue will check the matchmaking queue
func (s *matchMakingServer) checkQueue(ctx context.Context, user mmUser) (bool, []string, []int64) {
	// Crate a span for checking queue
	ctx, span := s.tr.Start(ctx, "checkQueue", trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(attribute.Key("username").String(user.username))
	span.SetAttributes(attribute.Key("MMR").Int64(user.mmr))
	defer span.End()

	// Lock mutex for checking the match making queue
	s.mu.Lock()
	span.AddEvent("checking matchmaking queue",
		trace.WithAttributes(
			attribute.Key("queue-length").Int64(int64(len(s.mqUsers))),
			attribute.Key("queue-users").StringSlice(s.mqUsers),
		),
	)

	// Find all candidates, this is O(n), stupid
	candidateIndexes := make([]int, 0)
	candidateUsernames := make([]string, 0)
	candidateMMRs := make([]int64, 0)

	for i, val := range s.mqMMR {
		// Check MMR difference
		dif := int(math.Abs(float64(user.mmr - val)))
		if dif < minMaxMMR {
			candidateIndexes = append(candidateIndexes, i)
			candidateUsernames = append(candidateUsernames, s.mqUsers[i])
			candidateMMRs = append(candidateMMRs, s.mqMMR[i])
		}

		// Early match making break
		if len(candidateIndexes) == 10 {
			break
		}
	}

	//  Add span event
	span.AddEvent("checking matchmaking queue finished",
		trace.WithAttributes(
			attribute.Key("candidate-count").Int64(int64(len(candidateIndexes))),
			attribute.Key("candidate-users").StringSlice(candidateUsernames),
			attribute.Key("candidate-mmrs").Int64Slice(candidateMMRs),
		),
	)

	// Check if there were 10 candidates
	if len(candidateIndexes) != 10 {
		// Matchmaking failed
		span.AddEvent("matchmaking failed")
		s.mu.Unlock()
		return false, []string{}, []int64{}
	} else {
		span.AddEvent("matchmaking finished")

		// Remove users from match making queue for candidates.
		for _, i := range candidateIndexes {
			s.mqUsers = removeString(s.mqUsers, i)
			s.mqMMR = removeInt(s.mqMMR, i)
		}
		s.mu.Unlock()
		return true, candidateUsernames, candidateMMRs
	}
}

// enqueueUser adds user into the match making queue
func (s *matchMakingServer) enqueueUser(ctx context.Context, user mmUser) {
	// Lock mutex for the match making queue
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mqUsers = append(s.mqUsers, user.username)
	s.mqMMR = append(s.mqMMR, user.mmr)

	// Add span for enqueue
	ctx, span := s.tr.Start(ctx, "enqueueUser", trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(attribute.Key("username").String(user.username))
	span.SetAttributes(attribute.Key("MMR").Int64(user.mmr))
	defer span.End()
}

// AddMatchmaking is for handling gRPC when adding real match making user
func (s *matchMakingServer) AddMatchmaking(ctx context.Context, req *pb.MatchmakingRequest) (*pb.MatchmakingResponse, error) {
	// You can access the username and password from req.Username and req.Password respectively
	ctx, span := s.tr.Start(ctx, "AddMatchmaking", trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(attribute.Key("username").String(req.Username))
	span.SetAttributes(attribute.Key("MMR").Int64(req.Mmr))
	defer span.End()

	// enqueue user to the match making queue
	mu := mmUser{username: req.Username, mmr: req.Mmr}
	s.enqueueUser(ctx, mu)

	// Loop until match making was complete
	for {
		ok, users, mmrs := s.checkQueue(ctx, mu)
		if ok {
			// Return gRPC response
			response := &pb.MatchmakingResponse{
				Usernames: users,
				Mmr:       mmrs,
			}
			return response, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// AddFakeUser is for handling gRPC when adding fake users
func (s *matchMakingServer) AddFakeUser(ctx context.Context, req *pb.MatchmakingRequest) (*pb.MatchmakingResponse, error) {
	// You can access the username and password from req.Username and req.Password respectively
	ctx, span := s.tr.Start(ctx, "AddFakeUser", trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(attribute.Key("username").String(req.Username))
	span.SetAttributes(attribute.Key("MMR").Int64(req.Mmr))
	defer span.End()

	// enqueue user to the match making queue
	mu := mmUser{username: req.Username, mmr: req.Mmr}
	s.enqueueUser(ctx, mu)
	return &pb.MatchmakingResponse{}, nil
}

func main() {
	// Initialize tracer
	tp, err := utils.TracerProvider("matchmaking", "production", 2, jaegerURL)
	if err != nil {
		log.Fatalf("could not initialize tracer provider: %v", err)
		return
	}
	log.Printf("tracer provider initialized")

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	ctx, cancel := context.WithCancel(context.Background())

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	// Start listening
	lis, err := net.Listen("tcp", hostAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create gRPC server, but with arguments for interceptors
	server := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor(otelgrpc.WithTracerProvider(tp))),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor(otelgrpc.WithTracerProvider(tp))),
	)

	// Start gRPC server
	pb.RegisterMatchmakingServiceServer(server, &matchMakingServer{tr: otel.Tracer(""),
		mqUsers: make([]string, 0), mqMMR: make([]int64, 0)})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
