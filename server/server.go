package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	logger "github.com/BabetteB/DISYS_MiniProject02/logFile"

	"github.com/BabetteB/DISYS_MiniProject02/protos"
	"google.golang.org/grpc"
)

type message struct {
	ClientUniqueCode  int32
	ClientName        string
	Msg               string
	MessageUniqueCode int32
}

type raw struct {
	MQue []message
	mu   sync.Mutex
}

type Server struct {
	protos.UnimplementedChittyChatServiceServer
	subscribers sync.Map
}

type sub struct {
	stream   protos.ChittyChatService_BroadcastServer
	finished chan<- bool
}

var messageHandle = raw{}

var (
	//mockServerTimeStamp string = "YYYY/YY/YY YY:YY:YY"
)

func (s *Server) Broadcast(request *protos.Subscription, stream protos.ChittyChatService_BroadcastServer) error {
	log.Printf("Received subscribe request from ID: %d", request.ClientId)
	fin := make(chan bool)

	s.subscribers.Store(request.ClientId, sub{stream: stream, finished: fin})

	ctx := stream.Context()
	go s.sendToClients(stream)


	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %d", request.ClientId)
			return nil
		case <-ctx.Done():
			log.Printf("Client ID %d has disconnected", request.ClientId)
			return nil
		}
	}
}


func (s *Server)sendToClients(srv protos.ChittyChatService_BroadcastServer) {

	//implement a loop
	for {

		//loop through messages in MQue
		for {

			time.Sleep(500 * time.Millisecond)

			messageHandle.mu.Lock()

			if len(messageHandle.MQue) == 0 {
				messageHandle.mu.Unlock()
				break
			}

			senderUniqueCode := messageHandle.MQue[0].ClientUniqueCode
			senderName4Client := messageHandle.MQue[0].ClientName
			// LamportTimestamp :=
			messageFromServer := messageHandle.MQue[0].Msg

			messageHandle.mu.Unlock()

			var unsubscribe []int32

			s.subscribers.Range(func(k, v interface{}) bool {
				id, ok := k.(int32)
				if !ok {
					log.Printf("Failed to cast subscriber key: %T", k)
					return false
				}
				sub, ok := v.(sub)
				if !ok {
					log.Printf("Failed to cast subscriber value: %T", v)
					return false
				}
				// Send data over the gRPC stream to the client
				if err := sub.stream.Send(&protos.ChatRoomMessages{
					Msg:              messageFromServer,
					LamportTimestamp: 321321,
					Username:         senderName4Client,
					ClientId:         senderUniqueCode,
				}); err != nil {
					log.Printf("Failed to send data to client: %v", err)
					select {
					case sub.finished <- true:
						log.Printf("Unsubscribed client: %d", id)
					default:
						// Default case is to avoid blocking in case client has already unsubscribed
					}
					// In case of error the client would re-subscribe so close the subscriber stream
					unsubscribe = append(unsubscribe, id)
				}
				return true
			})
	
			// Unsubscribe erroneous client streams
			for _, id := range unsubscribe {
				s.subscribers.Delete(id)
			} 

			messageHandle.mu.Lock()

			if len(messageHandle.MQue) > 1 {
				messageHandle.MQue = messageHandle.MQue[1:] // delete the message at index 0 after sending to receiver
			} else {
				messageHandle.MQue = []message{}
			}

			messageHandle.mu.Unlock()
		}
		time.Sleep(100 * time.Millisecond)
	}
}



func (s *Server) Publish(srv protos.ChittyChatService_PublishServer) error {

	errch := make(chan error)

	// receive messages - init a go routine
	go receiveFromStream(srv, errch)
	go sendToStream(srv, errch)

	return <-errch
}

//receive messages
func receiveFromStream(srv protos.ChittyChatService_PublishServer, errch_ chan error) {

	//implement a loop
	for {
		mssg, err := srv.Recv()
		if err != nil {
			errch_ <- err
		} else {
			id := mssg.ClientId

			messageHandle.mu.Lock()

			messageHandle.MQue = append(messageHandle.MQue, message{
				ClientUniqueCode:  id,
				ClientName:        mssg.UserName,
				Msg:               mssg.Msg,
				MessageUniqueCode: int32(rand.Intn(1e6)), // Maybe delete
			})


			//log.Printf("%v", messageHandle.MQue[len(messageHandle.MQue)-1])

			messageHandle.mu.Unlock()
		}
	}
}

func sendToStream(srv protos.ChittyChatService_PublishServer, errch_ chan error) {
	//implement a loop
	for {

		//loop through messages in MQue
		for {

			time.Sleep(500 * time.Millisecond)

			messageHandle.mu.Lock()

			if len(messageHandle.MQue) == 0 {
				messageHandle.mu.Unlock()
				break
			}

			//senderUniqueCode := messageHandle.MQue[0].ClientUniqueCode
			// senderName4Client := messageHandle.MQue[0].ClientName
			// message4Client := messageHandle.MQue[0].Msg

			messageHandle.mu.Unlock()

			//send message to designated client (do not send to the same client)
			//if senderUniqueCode != clientUniqueCode_ {

			err := srv.Send(&protos.StatusMessage{
				Operation: "Publish()",
				Status:    protos.Status_SUCCESS,
			})

			if err != nil {
				errch_ <- err
			}
			//}

		}
		time.Sleep(100 * time.Millisecond)
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
		println("Unsubscribed client: %d", request.ClientId)
	default:
		// Default case is to avoid blocking in case client has already unsubscribed
	}
	s.subscribers.Delete(request.ClientId)
	return &protos.StatusMessage{
		Operation: "Connect()",
		Status:    protos.Status_SUCCESS,
	}, nil
}

func main() {
	logger.LogFileInit()

	Output("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		logger.ErrorLogger.Fatalf("FATAL: Connection unable to establish. Failed to listen: %v", err)
	}
	logger.InfoLogger.Println("Connection established through TCP, listening at port 3000")

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
	logger.InfoLogger.Println("Exit successfull. Server closing...")

	os.Exit(3)
}

func Output(input string) {
	fmt.Println(input)
}
