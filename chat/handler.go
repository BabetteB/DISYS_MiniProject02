package chat

import (
	//"log"
	"fmt"
	"math/rand"
	"sync"
	"time"
	// logger "github.com/BabetteB/DISYS_MiniProject02/logFile"
)

var (
	//mockServerTimeStamp string = "YYYY/YY/YY YY:YY:YY"
	currentID      int32 = 0
	checkingStatus bool  = true // true: prints statusmessages
	numberOfClients int 
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

func (s *Server) Publish(srv ChittyChatService_PublishServer) error {

	errch := make(chan error)

	// receive messages - init a go routine
	go receiveFromStream(srv, errch)
	go sendToStream(srv, currentID, errch)

	return <-errch

}

//receive messages
func receiveFromStream(srv ChittyChatService_PublishServer, errch_ chan error) {

	//implement a loop
	for {
		mssg, err := srv.Recv()
		if err != nil {
			errch_ <- err
		} else {
			id := mssg.ClientId

			if id == 0 {
				id = int32(rand.Intn(1e6))
				++numberOfClients
			}

			messageHandle.mu.Lock()

			messageHandle.MQue = append(messageHandle.MQue, message{
				ClientUniqueCode:  id,
				ClientName:        mssg.UserName,
				Msg:               mssg.Msg,
				MessageUniqueCode: int32(rand.Intn(1e6)), // Maybe delete
			})

			currentID = id

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
					Operation: "Publish()",
					Status:    Status_SUCCESS,
					NewId:     &senderUniqueCode,
				})

				if err != nil {
					errch_ <- err
				}
			}

		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Server) Broadcast(srv ChittyChatService_BroadcastServer) error {
	errch := make(chan error)

	// receive messages - init a go routine
	//go receiveStatusStream(csi, clientUniqueCode, errch)

	// send messages - init a go routine
	go sendToClients(srv, errch)
	go recvFromClient(srv, errch)

	return <-errch

}

func recvFromClient(srv ChittyChatService_BroadcastServer, errch_ chan error) {
	//implement a loop
	for {
		if (!checkingStatus) {
			break
		}
		mssg, err := srv.Recv()
		if err != nil {
			errch_ <- err
		} else {
			if checkingStatus {
				println(fmt.Sprintf("%s : %s \n", mssg.Operation, mssg.Status))
				checkingStatus = false
			}
		}
	}
}

func sendToClients(srv ChittyChatService_BroadcastServer, errch_ chan error) {

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

			//send message to designated client (do not send to the same client)

			err := srv.Send(&ChatRoomMessages{
				Msg:              messageFromServer,
				LamportTimestamp: 321321,
				Username:         senderName4Client,
				ClientId:         senderUniqueCode,
			})

			if err != nil {
				errch_ <- err
			}

			messageHandle.mu.Lock()

			if len(messageHandle.MQue) > 1 {
				messageHandle.MQue = messageHandle.MQue[1:] // delete the message at index 0 after sending to receiver
			} else {
				messageHandle.MQue = []message{}
				checkingStatus = true
			}

			messageHandle.mu.Unlock()

		}

		time.Sleep(100 * time.Millisecond)
	}
}
