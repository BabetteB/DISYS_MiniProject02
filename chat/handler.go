package chat

import (
	//"log"
	"math/rand"
	"sync"
	"time"
	// logger "github.com/BabetteB/DISYS_MiniProject02/logFile"

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
	UnimplementedChittyChatServiceServer
}

var messageHandle = raw{}

//define ChatService
func (s *Server) Publish(srv ChittyChatService_PublishServer) error {

	clientUniqueCode := int32(rand.Intn(1e6))
	errch := make(chan error)

	// receive messages - init a go routine
	go receiveFromStream(srv, clientUniqueCode, errch)

	// send messages - init a go routine
	go sendToStream(srv, clientUniqueCode, errch)

	return <-errch

}

//receive messages
func receiveFromStream(srv ChittyChatService_PublishServer, clientUniqueCode_ int32, errch_ chan error) {

	//implement a loop
	for {
		mssg, err := srv.Recv()
		if err != nil {
			errch_ <- err
		} else {

			messageHandle.mu.Lock()

			messageHandle.MQue = append(messageHandle.MQue, message{
				ClientUniqueCode:  mssg.ClientId,
				ClientName:        mssg.UserName,
				Msg:               mssg.Msg,
				MessageUniqueCode: int32(rand.Intn(1e6)),
			}) 
			println(mssg.Msg) 

			//log.Printf("%v", messageHandle.MQue[len(messageHandle.MQue)-1])

			messageHandle.mu.Unlock()
		}
	}
}

//send message
func sendToStream(srv ChittyChatService_PublishServer, clientUniqueCode_ int32, errch_ chan error) {

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
			// senderName4Client := messageHandle.MQue[0].ClientName
			// message4Client := messageHandle.MQue[0].Msg

			messageHandle.mu.Unlock()

			//send message to designated client (do not send to the same client)
			if senderUniqueCode != clientUniqueCode_ {

				err := srv.Send(&StatusMessage{
					Operation: "publish",
					Status:    Status_SUCCESS,
				})

				if err != nil {
					errch_ <- err
				}

				messageHandle.mu.Lock()

				if len(messageHandle.MQue) > 1 {
					messageHandle.MQue = messageHandle.MQue[1:] // delete the message at index 0 after sending to receiver
				} else {
					messageHandle.MQue = []message{}
				}

				messageHandle.mu.Unlock()

			}

		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Server) Broadcast(srv ChittyChatService_BroadcastServer) error {
	return nil
}

