package protos

import (
	"math"
	"sync"
)

type LamportTimestamp struct {
	id        int32
	Timestamp int32
	mu        sync.Mutex
}

func Tick(l *LamportTimestamp) {
	l.mu.Lock()
	l.Timestamp += 1
	l.mu.Unlock()
}

// message is the given message being sent with the timestamp to the other process
func Recieving(recieveLamp *LamportTimestamp, sendingLamp *LamportTimestamp, message string) {
	// if timestamp of msg sent is greater than timestamp of the recieving end
	// then set recieving timestamp to msg sent timestamp+1 - else increment recieving with one.
	if sendingLamp.Timestamp > recieveLamp.Timestamp {
		recieveLamp.mu.Lock()
		recieveLamp.Timestamp = sendingLamp.Timestamp + 1
		recieveLamp.mu.Unlock()
	} else {
		Tick(recieveLamp)
	}
}

// TEST
func RecievingOneLamportOneInt(recieveLamp *LamportTimestamp, sendingLamp int32) {
	// if timestamp of msg sent is greater than timestamp of the recieving end
	// then set recieving timestamp to msg sent timestamp+1 - else increment recieving with one.
	if sendingLamp > recieveLamp.Timestamp {
		recieveLamp.mu.Lock()
		recieveLamp.Timestamp = sendingLamp + 1
		recieveLamp.mu.Unlock()
	} else {
		Tick(recieveLamp)
	}
}

func RecievingCompareToLamport(recieveLamp *LamportTimestamp, sendingLamp int32) int32 {
	// if timestamp of msg sent is greater than timestamp of the recieving end
	// then set recieving timestamp to msg sent timestamp+1 - else increment recieving with one.
	if recieveLamp.Timestamp > sendingLamp {
		sendingLamp = recieveLamp.Timestamp
	}
	sendingLamp = sendingLamp + 1
	recieveLamp.Timestamp = sendingLamp
	return sendingLamp
}

func RecievingSomething2(recieveLamp *LamportTimestamp, sendingLamp int32) int32 {
	// if timestamp of msg sent is greater than timestamp of the recieving end
	// then set recieving timestamp to msg sent timestamp+1 - else increment recieving with one.
	return int32(math.Max(float64(recieveLamp.Timestamp), float64(sendingLamp)) + 1)
}
