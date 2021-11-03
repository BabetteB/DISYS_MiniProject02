package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	logger "github.com/BabetteB/DISYS_MiniProject02/logFile"

	"github.com/BabetteB/DISYS_MiniProject02/protos"
	"google.golang.org/grpc"
)

/*
MessageCode:
1 - client joined server
2 - Chat
3 - client left server
4 - Server closing
*/

// burde nok have timestamp inkluderet... ??????????????????????
type message struct {
	ClientUniqueCode int32
	ClientName       string
	Msg              string
	MessageCode      int32
	Lamport          int32
}

type raw struct {
	MessageQue []message
	mu         sync.Mutex
}

type Server struct {
	protos.UnimplementedChittyChatServiceServer
	subscribers sync.Map
	lamport     protos.LamportTimestamp
}

type sub struct {
	stream   protos.ChittyChatService_BroadcastServer
	finished chan<- bool
}

var messageHandle = raw{}

func (s *Server) Broadcast(request *protos.Subscription, stream protos.ChittyChatService_BroadcastServer) error {
	protos.Tick(&s.lamport)
	logger.InfoLogger.Printf("Lamp.t.: %d, Received subscribe request from ID: %d", s.lamport.Timestamp, request.ClientId)
	fin := make(chan bool)

	s.subscribers.Store(request.ClientId, sub{stream: stream, finished: fin})

	// Connecting
	addToMessageQueue(request.ClientId, s.lamport.Timestamp, 1, request.UserName, "")
	Output(fmt.Sprintf("ID: %v Name: %v, joined chat at timestamp %d", request.ClientId, request.UserName, s.lamport.Timestamp))

	ctx := stream.Context()
	go s.sendToClients(stream)

	for {
		select {
		case <-fin:
			protos.Tick(&s.lamport)
			logger.InfoLogger.Printf("Lamp.t.: %d, Closing stream for client ID: %d", s.lamport.Timestamp, request.ClientId)
			return nil
		case <-ctx.Done():
			protos.Tick(&s.lamport)
			logger.InfoLogger.Printf("Lamp.t.: %d, Client ID %d has disconnected", s.lamport.Timestamp, request.ClientId)
			return nil
		}
	}
}

func (s *Server) sendToClients(srv protos.ChittyChatService_BroadcastServer) {
	logger.InfoLogger.Println("Request send to clients")
	//implement a loop
	for {

		//loop through messages in MessageQue
		for {

			time.Sleep(500 * time.Millisecond)

			messageHandle.mu.Lock()

			if len(messageHandle.MessageQue) == 0 {
				messageHandle.mu.Unlock()
				break
			}

			senderUniqueCode := messageHandle.MessageQue[0].ClientUniqueCode
			senderName := messageHandle.MessageQue[0].ClientName
			LamportTimestamp := messageHandle.MessageQue[0].Lamport
			messageFromServer := messageHandle.MessageQue[0].Msg
			messageCode := messageHandle.MessageQue[0].MessageCode

			messageHandle.mu.Unlock()

			var unsubscribe []int32

			s.subscribers.Range(func(k, v interface{}) bool {
				id, ok := k.(int32)
				if !ok {
					logger.WarningLogger.Panicf("Failed to cast subscriber key: %T", k)
					return false
				}
				sub, ok := v.(sub)
				if !ok {
					logger.WarningLogger.Panicf("Failed to cast subscriber value: %T", v)
					return false
				}
				// Send data over the gRPC stream to the client
				if err := sub.stream.Send(&protos.ChatRoomMessages{
					Msg:              messageFromServer,
					LamportTimestamp: LamportTimestamp,
					Username:         senderName,
					ClientId:         senderUniqueCode,
					Code:             messageCode,
				}); err != nil {
					logger.ErrorLogger.Output(2, (fmt.Sprintf("Failed to send data to client: %v", err)))
					select {
					case sub.finished <- true:
						logger.InfoLogger.Printf("Unsubscribed successfully client: %d", id)
					default:
						// Default case is to avoid blocking in case client has already unsubscribed
					}
					// In case of error the client would re-subscribe so close the subscriber stream
					unsubscribe = append(unsubscribe, id)
				}
				return true
			})
			logger.InfoLogger.Println("Brodcasting message success.")

			// Unsubscribe erroneous client streams
			for _, id := range unsubscribe {
				logger.InfoLogger.Printf("Killed client: %v", id)
				s.subscribers.Delete(id)
			}

			messageHandle.mu.Lock()

			if len(messageHandle.MessageQue) > 1 {
				messageHandle.MessageQue = messageHandle.MessageQue[1:] // delete the message at index 0 after sending to receiver
			} else {
				messageHandle.MessageQue = []message{}
			}

			messageHandle.mu.Unlock()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Server) Publish(srv protos.ChittyChatService_PublishServer) error {
	logger.InfoLogger.Println("Requests publish")
	errch := make(chan error)

	// receive messages - init a go routine
	go receiveFromStream(srv, errch)
	go sendToStream(srv, errch)
	return <-errch
}

func receiveFromStream(srv protos.ChittyChatService_PublishServer, errch_ chan error) {

	//implement a loop
	for {
		mssg, err := srv.Recv()
		if err != nil {
			logger.WarningLogger.Printf("Error occured when recieving message: %v", err)
			errch_ <- err
		} else {
			addToMessageQueue(mssg.ClientId, mssg.LamportTimestamp, 2, mssg.UserName, mssg.Msg)
		}
	}
}

func addToMessageQueue(id, lamport, code int32, username, msg string) {
	messageHandle.mu.Lock()

	messageHandle.MessageQue = append(messageHandle.MessageQue, message{
		ClientUniqueCode: id,
		ClientName:       username,
		Msg:              msg,
		MessageCode:      code, // Maybe delete
		Lamport:          lamport,
	})

	logger.InfoLogger.Printf("Message successfully recieved and queued: %v", id)

	messageHandle.mu.Unlock()
}

func sendToStream(srv protos.ChittyChatService_PublishServer, errch_ chan error) {
	//implement a loop
	for {
		time.Sleep(500 * time.Millisecond)

		err := srv.Send(&protos.StatusMessage{
			Operation: "Publish()",
			Status:    protos.Status_SUCCESS,
		})

		if err != nil {
			logger.WarningLogger.Panicf("An error occured when sending message: %v", err)
			errch_ <- err
		}
	}
}

func (s *Server) Disconnect(ctx context.Context, request *protos.Subscription) (*protos.StatusMessage, error) {
	v, ok := s.subscribers.Load(request.ClientId)
	if !ok {
		return nil, fmt.Errorf("failed to load subscriber key: %d", request.ClientId)
	}
	sub, ok := v.(sub)
	if !ok {
		return nil, fmt.Errorf("failed to cast subscriber value: %T", v)
	}
	select {
		case sub.finished <- true:
			println("Client %d disconnected", request.ClientId)
		default:
			// Default case is to avoid blocking in case client has already unsubscribed
	}
	s.subscribers.Delete(request.ClientId)

	// Noget her skal ændres	
	return &protos.StatusMessage{
		Operation: "Disconnected()",
		Status:    protos.Status_SUCCESS,
	}, nil
}

func main() {
	logger.LogFileInit()

	Output("Server started")

	port := 3000

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.ErrorLogger.Fatalf("FATAL: Connection unable to establish. Failed to listen: %v", err)
	}
	logger.InfoLogger.Printf("Connection established through TCP, listening at port %v", port)

	s := &Server{}

	grpcServer := grpc.NewServer()

	protos.RegisterChittyChatServiceServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.ErrorLogger.Fatalf("FATAL: Server connection failed: %s", err)
		}
	}()

	Output("Server started. Press any key to stop")

	var o string
	fmt.Scanln(&o)
	Output("Server exiting... ")

	addToMessageQueue(0, s.lamport.Timestamp, 4, "", "")
	Output("Exit successfull")

	logger.InfoLogger.Println("Exit successfull. Server closing...")

	os.Exit(3)
}

func Output(input string) {
	fmt.Println(input)
}
