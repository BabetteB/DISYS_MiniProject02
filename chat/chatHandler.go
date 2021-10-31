package chat

import (
	logger "github.com/BabetteB/DISYS_MiniProject02/logFile"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

var (
	broadcastMessage = ""
	broadcaster      = ""
	isNewMessage     bool
	clients          = make(map[int32]string)
)

type Server struct {
	timestamp lamport_timestamp
}

// used in the client as well - maybe move to seperate file? + its functions?
type lamport_timestamp struct {
	id        int32
	timestamp int32
}

func (s *Server) setServer() {
	s.timestamp.id = 0
}

func (s *Server) Publish(ctx context.Context, in *ClientMessage) (*StatusMessage, error) {
	logger.InfoLogger.Printf("Received message from client: %v", in.Msg)
	response := StatusMessage{
		Operation: "Operation: Publish",
		Status:    Status_SUCCESS}
	broadcastMessage = in.Msg
	broadcaster = in.UserName
	isNewMessage = true
	s.timestamp.timestamp += 1
	logger.InfoLogger.Println("Response successfull")
	return &response, nil
}

func (s *Server) Broadcast(ctx context.Context, in *google_protobuf.Empty) (*ChatRoomMessages, error) {
	receivers := 0
	s.timestamp.timestamp += 1
	logger.InfoLogger.Println("Requesting brodcast")
	if isNewMessage {
		receivers++
		if receivers == len(clients) {
			isNewMessage = false
			receivers = 0
		}
    logger.InfoLogger.Println("Brodcast successfull")
		return &ChatRoomMessages{
			Msg:              broadcastMessage,
			LamportTimestamp: s.timestamp.timestamp,
			Username:         broadcaster,
		}, nil
	} else {
    logger.WarningLogger.Println("Bodcasting empty message and empty user")
		return &ChatRoomMessages{
			Msg:              "",
			LamportTimestamp: s.timestamp.timestamp,
			Username:         "",
		}, nil
	}
}

func (s *Server) Connect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	logger.InfoLogger.Println("Requesting connection")
	var newId int32 = int32(len(clients) + 1)
	clients[newId] = in.Name
	response := StatusMessage{
		Operation: "Operation: Connect",
		Status:    Status_SUCCESS,
		NewId:     &newId,
	}
	logger.InfoLogger.Printf("Connection status: %s", Status_SUCCESS)
	return &response, nil
}

func (s *Server) Disconnect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	logger.InfoLogger.Println("Requesting disconnection")
	response := StatusMessage{
		Operation: "Operation: Disconnect",
		Status:    Status_SUCCESS,
	}
	logger.InfoLogger.Printf("Disconnection status: %s", Status_SUCCESS)
	return &response, nil
}
