package raft

import (
	"fmt"
	"strings"
)

type Entry struct {
	Command interface{}
	Term    int
	Index   int
}

type Log struct {
	Entries []Entry
	index0  int
}

func (l *Log) truncate(i int) {
	l.Entries = l.Entries[:i]
}

func (l *Log) append(entries ...Entry) {
	l.Entries = append(l.Entries, entries...)
}

func (l *Log) len() int {
	return len(l.Entries)
}

func (l *Log) silce(i int) []Entry {
	return l.Entries[i:]
}

func (l *Log) at(i int) *Entry {
	return &l.Entries[i]
}

func (l *Log) lastLog() *Entry {
	return l.at(l.len() - 1)
}

func makeEmptyLog() Log {
	log := Log{
		Entries: make([]Entry, 0),
		index0:  0,
	}
	return log
}

func (e *Entry) String() string {
	return fmt.Sprint(e.Term)
}

func (l *Log) String() string {
	nums := []string{}
	for _, entry := range l.Entries {
		nums = append(nums, fmt.Sprintf("%4d", entry.Term))
	}
	return fmt.Sprint(strings.Join(nums, "|"))
}
