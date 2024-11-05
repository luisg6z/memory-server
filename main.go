package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sync"

	game "prueba/proto/game"
	lobby "prueba/proto/lobby"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type lobbyServer struct {
	lobby.UnimplementedLobbyServiceServer
	mu   sync.Mutex
	room *Room // Map of room ID to Room
}

type Room struct {
	players    map[string]*lobby.Player // Map of player IDs to Player
	maxPlayers int32                    // Maximum number of players allowed in the room
}

// JoinRoom allows a player to join an existing room
func (s *lobbyServer) JoinRoom(ctx context.Context, req *lobby.JoinRoomRequest) (*lobby.JoinRoomResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	room := s.room

	// Check if the room is full
	if int32(len(room.players)) >= room.maxPlayers {
		return nil, grpc.Errorf(codes.ResourceExhausted, "Room is full")
	}

	// Add the player to the room
	room.players[req.PlayerId] = &lobby.Player{PlayerId: req.PlayerId}

	// Prepare the response
	otherPlayers := make([]*lobby.Player, 0, len(room.players)-1)
	for id, player := range room.players {
		if id != req.PlayerId {
			otherPlayers = append(otherPlayers, player)
		}
	}

	response := &lobby.JoinRoomResponse{
		OtherPlayers: otherPlayers,
		Me:           room.players[req.PlayerId],
		MaxPlayers:   room.maxPlayers,
	}

	return response, nil
}

type Player struct {
	ID    string
	Score int32
}

type Game struct {
	Players  map[string]*Player // Map of player ID to Player
	Sequence []game.CardType    // Sequence of cards
	Started  bool               // Indicates if the game has started
}

type gameServer struct {
	game.UnimplementedGameServiceServer
	mu           sync.Mutex
	game         *Game
	lobbyService *lobbyServer
}

func (*gameServer) NewGameServer(ls *lobbyServer) *gameServer {
	return &gameServer{lobbyService: ls}
}

// NewGame initializes a new game
func (s *gameServer) StartGame(ctx context.Context, lobby *game.Lobby) (*game.StartGameResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Initialize the game
	s.game = &Game{
		Players:  make(map[string]*Player),
		Sequence: generateSequence(5), // Generate a sequence of 5 cards
		Started:  true,
	}

	// Add players to the game
	for _, playerID := range lobby.PlayersId {
		s.game.Players[playerID] = &Player{ID: playerID, Score: 0}
	}

	// Prepare the response
	response := &game.StartGameResponse{
		Sequence: s.game.Sequence,
	}

	return response, nil
}

// generateSequence generates a random sequence of CardType
func generateSequence(length int) []game.CardType {
	sequence := make([]game.CardType, length)
	for i := 0; i < length; i++ {
		sequence[i] = game.CardType(rand.Intn(17) + 1) // Random CardType from 1 to 17
	}
	return sequence
}

// SubmitAnswer processes a player's answer and calculates the score
func (s *gameServer) SubmitAnswer(ctx context.Context, req *game.SubmitAnswerRequest) (*game.SubmitAnswerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the game has started
	if s.game == nil || !s.game.Started {
		return nil, grpc.Errorf(codes.FailedPrecondition, "Game has not started")
	}

	// Check if the player is part of the game
	player, exists := s.game.Players[req.PlayerId]
	if !exists {
		return nil, grpc.Errorf(codes.NotFound, "Player not found")
	}

	// Check if the answer is correct
	correct := true
	if len(req.Answer) != len(s.game.Sequence) {
		correct = false
	} else {
		for i, card := range req.Answer {
			if card != s.game.Sequence[i] {
				correct = false
				break
			}
		}
	}

	// Calculate score based on correctness and time taken
	score := calculateScore(correct, req.Timestamp)

	// Update player's score
	player.Score += score

	// Prepare the response
	response := &game.SubmitAnswerResponse{
		PlayerId: req.PlayerId,
		Score:    player.Score,
	}

	return response, nil
}

// calculateScore calculates the score based on correctness and time taken
func calculateScore(correct bool, timestamp int64) int32 {
	// For simplicity, let's assume:
	// - If the answer is correct, the player gets 10 points.
	// - If the answer is incorrect, the player gets 0 points.
	// You can modify this logic to include time-based scoring if needed.
	if correct {
		return 10 * int32(timestamp) // Example score for a correct answer
	}
	return 0 // No score for an incorrect answer
}

// GetGameState returns the current state of the game
func (s *gameServer) GetGameState(ctx context.Context, req *emptypb.Empty) (*game.GameState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the game exists
	if s.game == nil {
		return nil, grpc.Errorf(codes.NotFound, "Game not found")
	}

	// Prepare the response
	players := make([]string, 0, len(s.game.Players))
	for playerID := range s.game.Players {
		players = append(players, playerID)
	}

	response := &game.GameState{
		Players:  players,
		Sequence: s.game.Sequence,
	}

	return response, nil
}

// ShowScores returns the scores of all players in the game
func (s *gameServer) ShowScores(ctx context.Context, req *emptypb.Empty) (*game.ShowScoreResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the game exists
	if s.game == nil {
		return nil, grpc.Errorf(codes.NotFound, "Game not found")
	}

	// Prepare the response
	scores := make([]*game.SubmitAnswerResponse, 0, len(s.game.Players))
	for _, player := range s.game.Players {
		scores = append(scores, &game.SubmitAnswerResponse{
			PlayerId: player.ID,
			Score:    player.Score,
		})
	}

	response := &game.ShowScoreResponse{
		Scores: scores,
	}

	return response, nil
}

// Main function to start the gRPC server
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	room := &Room{
		players:    make(map[string]*lobby.Player),
		maxPlayers: 4, // Set maximum players for the room
	}
	lobbyService := &lobbyServer{room: room}
	lobby.RegisterLobbyServiceServer(s, lobbyService)

	gameService := &gameServer{}
	game.RegisterGameServiceServer(s, gameService)

	log.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to create the server: %v", err)
	}
}
