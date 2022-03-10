package raft

import (
	"math/rand"
	"time"
)

func (rf *Raft) resetElectionTimer() {
	t := time.Now()
	electionTimeout := time.Duration(150+rand.Intn(150)) * time.Microsecond
	rf.electionTime = t.Add(electionTimeout)
}
