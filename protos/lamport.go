package protos

import (
	"fmt"
)

type LamportTimestamp struct {
	id        int32
	Timestamp int32
}

func (l *LamportTimestamp) Tick() {
	l.Timestamp += 1
}

func (lamport *LamportTimestamp) RecieveTest(timestamp int32) {
	if lamport.Timestamp < timestamp {
		lamport.Timestamp = timestamp + 1
	} else {
		lamport.Tick()
	}
}

func RecievingCompareToLamport(recieveLamp *LamportTimestamp, sendingLamp int32) int32 {
	// if timestamp of msg sent is greater than timestamp of the recieving end
	// then set recieving timestamp to msg sent timestamp+1 - else increment recieving with one.
	fmt.Printf("recieving time before increment: %d ,", recieveLamp.Timestamp)
	fmt.Printf("sendingLamp time before increment: %d\n", sendingLamp)
	if recieveLamp.Timestamp > sendingLamp {
		sendingLamp = recieveLamp.Timestamp
	} else if sendingLamp > recieveLamp.Timestamp {
		recieveLamp.Timestamp = sendingLamp
	}
	recieveLamp.Timestamp = sendingLamp + 1
	sendingLamp = recieveLamp.Timestamp
	fmt.Printf("reciving time after: %d, ", recieveLamp.Timestamp)
	fmt.Printf("sending time after: %d\n", sendingLamp)
	return sendingLamp
}
