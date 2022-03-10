package raft

type Entry struct {
	Command interface{}
	Term    int
	Index   int
}

type Log struct {
	Entries []Entry
	index0  int
}

func (l *Log) append(entries ...Entry) {
	l.Entries = append(l.Entries, entries...)
}

func makeEmptyLog() Log {
	log := Log{
		Entries: make([]Entry, 0),
		index0:  0,
	}
	return log
}
