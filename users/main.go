package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"net/http"
	pb "simple-mm/proto/simple-mm"
	"simple-mm/utils"
	"time"
)

const (
	jaegerURL = "http://127.0.0.1:14268/api/traces"
)

type User struct {
	Username string `json:"username"`
	MMR      int    `json:"mmr"`
}

func handleMatchmaking(w http.ResponseWriter, r *http.Request) {
	// Init OTEL tracer
	tr := otel.Tracer("frontend")
	ctx, span := tr.Start(context.Background(), "handleMatchmaking")
	defer span.End()

	// Check request was valid
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON to struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Log and add attribute to span
	log.Printf("%s entered matchmaking! MMR: %d\n", user.Username, user.MMR)
	span.SetAttributes(attribute.Key("username").String(user.Username))
	span.SetAttributes(attribute.Key("MMR").Int(user.MMR))

	// Start gRPC server dial
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))),
		grpc.WithStreamInterceptor(
			otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))))
	if err != nil {
		fmt.Println("could not dial grpc", err)
		return
	}

	// Execute gRPC
	cl := pb.NewMatchmakingServiceClient(conn)
	users, err := cl.AddMatchmaking(ctx, &pb.MatchmakingRequest{Username: user.Username, Mmr: int64(user.MMR)})
	if err != nil {
		log.Printf("could not make matchmaking: %v", err)
		return
	}

	// Convert usernames to JSON
	jsonData, err := json.Marshal(users.Usernames)
	if err != nil {
		log.Println("Failed to marshal usernames to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("Failed to write JSON response:", err)
	}

	return
}

func main() {
	// Init OTEL trace provider
	tp, err := utils.TracerProvider("fake-user", "production", 1, jaegerURL)
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

	// Add random users
	// Start gRPC server dial
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))),
		grpc.WithStreamInterceptor(
			otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))))
	if err != nil {
		fmt.Println("could not dial grpc", err)
		return
	}

	// Generate seed for random username
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Execute gRPC, add random users
	for {
		username := func() string {
			b := make([]rune, 10)
			for i := range b {
				b[i] = letterRunes[rand.Intn(len(letterRunes))]
			}
			return string(b)
		}()
		mmr := rand.Intn(100-0) + 0
		log.Printf("adding fake user %s(%d)\n", username, mmr)

		cl := pb.NewMatchmakingServiceClient(conn)
		_, err := cl.AddFakeUser(ctx, &pb.MatchmakingRequest{Username: username, Mmr: int64(mmr)})
		if err != nil {
			log.Printf("could not add fake user: %v(%d)\n", username, mmr)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
