package protos

type LamportTimestamp struct {
	id        int32
	timestamp int32
}

func tick(l *LamportTimestamp) {
	l.timestamp += 1
}

// message is the given message being sent with the timestamp to the other process
func recieving(recieveLamp *LamportTimestamp, sendingLamp *LamportTimestamp, message string) {
	// if timestamp of msg sent is greater than timestamp of the recieving end
	// then set recieving timestamp to msg sent timestamp+1 - else increment recieving with one.
	if sendingLamp.timestamp > recieveLamp.timestamp {
		recieveLamp.timestamp = sendingLamp.timestamp + 1
	} else {
		tick(recieveLamp)
	}
}