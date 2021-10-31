package chat

import (
	logger  "github.com/BabetteB/DISYS_MiniProject02/logFile"
	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

const (
	mockTimestamp = "2021-10-29 00:00:00";
) 

var (
	broadcastMessage = "";
	broadcaster = "";
	isNewMessage bool;
	clients = make(map[int32] string);
)


type Server struct {
	logger.InfoLogger("Server initialized")
}

func (s *Server) Publish(ctx context.Context, in *ClientMessage) (*StatusMessage, error) {
	log.InfoLogger("Received message from client: %s", in.Body)

	response := StatusMessage{ 
		Operation: "Operation: Publish",
		Status: Status_SUCCESS};
	broadcastMessage = in.Msg
	broadcaster = in.UserName 
	isNewMessage = true
	// println(broadcastMessage) // Test
	log.InfoLogger("Response successfull")
	return &response, nil
}

func (s *Server) Broadcast(ctx context.Context, in *google_protobuf.Empty) (*ChatRoomMessages, error) {
	logger.InfoLogger()
	receivers := 0;
	if isNewMessage {
		receivers++
		if(receivers == len(clients)) {
			isNewMessage = false
			receivers = 0
		}
		return &ChatRoomMessages{
			Msg: broadcastMessage,
			Timestamp: mockTimestamp,
			Username: broadcaster,
			}, nil
	} else {
		return &ChatRoomMessages{
			Msg: "",
			Timestamp: mockTimestamp,
			Username: "",
			}, nil
	}	
}

func (s *Server) Connect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	// log.Printf("")
	var newId int32 = int32(len(clients)+1)
	clients[newId] = in.Name
	response := StatusMessage{
		Operation: "Operation: Connect",
		Status: Status_SUCCESS,
		NewId: &newId,
	}
	return &response, nil
}

func (s *Server) Disconnect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	// log.Printf("")
	response := StatusMessage{
		Operation: "Operation: Disconnect",
		Status: Status_SUCCESS,
	}
	return &response, nil
}

