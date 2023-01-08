package raft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {
	log := NewPersistentLog("raft/test")

	assert.Equal(t, "raft/test/log", log.Path())
	assert.Nil(t, log.File())
	assert.Zero(t, log.Size())
	assert.Zero(t, log.LastIndex())
	assert.Zero(t, log.LastTerm())
}

func TestOpenNew(t *testing.T) {
	path := t.TempDir()
	log := NewPersistentLog(path)
	t.Cleanup(func() { log.Close() })
	log.Open()

	assert.NotNil(t, log.File())
	assert.Equal(t, log.File().Name(), log.Path())
	assert.Zero(t, log.Size())
}

func TestIsOpen(t *testing.T) {
	path := t.TempDir()
	log := NewPersistentLog(path)
	t.Cleanup(func() { log.Close() })

	assert.False(t, log.IsOpen())
	log.Open()
	assert.True(t, log.IsOpen())
}

func TestAppendEntries(t *testing.T) {
	log := NewTestLog(t)

	var entry1, entry2 *LogEntry

	var entry1Index uint64 = 1
	var entry1Term uint64 = 1
	entry1Data := []byte("entry1")
	entry1 = NewLogEntry(entry1Index, entry1Term, entry1Data)

	var entry2Index uint64 = 2
	var entry2Term uint64 = 2
	entry2Data := []byte("entry2")
	entry2 = NewLogEntry(entry2Index, entry2Term, entry2Data)

	log.AppendEntries(entry1, entry2)

	validateLogSize(t, log.Size(), 2)

	entry1 = log.GetEntry(entry1Index)
	validateLogEntry(t, entry1, entry1Index, entry1Term, entry1Data)

	entry2 = log.GetEntry(entry2Index)
	validateLogEntry(t, entry2, entry2Index, entry2Term, entry2Data)

	assert.Equal(t, log.LastTerm(), entry2Term)
	assert.Equal(t, log.LastIndex(), entry2Index)
}

func TestTruncate(t *testing.T) {
	log := NewTestLog(t)

	var entry1, entry2 *LogEntry

	var entry1Index uint64 = 1
	var entry1Term uint64 = 1
	entry1Data := []byte("entry1")
	entry1 = NewLogEntry(entry1Index, entry1Term, entry1Data)

	var entry2Index uint64 = 2
	var entry2Term uint64 = 2
	entry2Data := []byte("entry2")
	entry2 = NewLogEntry(entry2Index, entry2Term, entry2Data)

	log.AppendEntries(entry1, entry2)

	log.Truncate(entry2Index)

	checkLog := func() {
		validateLogSize(t, log.Size(), 1)
		entry1 = log.GetEntry(entry1Index)
		validateLogEntry(t, entry1, entry1Index, entry1Term, entry1Data)
	}

	checkLog()

	log.Close()
	log.Open()

	checkLog()

	assert.Equal(t, log.LastTerm(), entry1Term)
	assert.Equal(t, log.LastIndex(), entry1Index)
}

func TestAppendEntriesTruncate(t *testing.T) {
	log := NewTestLog(t)

	var entry1, entry2, entry3 *LogEntry

	var entry1Index uint64 = 1
	var entry1Term uint64 = 1
	entry1Data := []byte("entry1")
	entry1 = NewLogEntry(entry1Index, entry1Term, entry1Data)

	var entry2Index uint64 = 2
	var entry2Term uint64 = 2
	entry2Data := []byte("entry2")
	entry2 = NewLogEntry(entry2Index, entry2Term, entry2Data)

	var entry3Index uint64 = 3
	var entry3Term uint64 = 3
	entry3Data := []byte("entry3")
	entry3 = NewLogEntry(entry3Index, entry3Term, entry3Data)

	log.AppendEntries(entry1, entry2, entry3)

	var entry4Index uint64 = 2
	var entry4Term uint64 = 3
	entry4Data := []byte("entry4")
	entry4 := NewLogEntry(entry4Index, entry4Term, entry4Data)

	var entry5Index uint64 = 3
	var entry5Term uint64 = 3
	entry5Data := []byte("entry5")
	entry5 := NewLogEntry(entry5Index, entry5Term, entry5Data)

	log.AppendEntries(entry4, entry5)

	checkLog := func() {
		validateLogSize(t, log.Size(), 3)

		entry1 = log.GetEntry(entry1Index)
		validateLogEntry(t, entry1, entry1Index, entry1Term, entry1Data)

		entry2 = log.GetEntry(entry2Index)
		validateLogEntry(t, entry2, entry4Index, entry4Term, entry4Data)

		entry3 = log.GetEntry(entry3Index)
		validateLogEntry(t, entry3, entry5Index, entry5Term, entry5Data)
	}

	checkLog()

	log.Close()
	log.Open()

	checkLog()

	assert.Equal(t, log.LastTerm(), entry5Term)
	assert.Equal(t, log.LastIndex(), entry5Index)
}
