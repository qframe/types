package qtypes_messages

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestNewBase(t *testing.T) {
	before := time.Now()
	b := NewBase("src1")
	after := time.Now()
	assert.Equal(t, version, b.BaseVersion)
	assert.Equal(t, "src1", b.SourcePath[0])
	assert.True(t, before.UnixNano() < b.Time.UnixNano())
	assert.True(t, after.UnixNano() > b.Time.UnixNano())
}


func TestNewTimedBase(t *testing.T) {
	now := time.Now()
	b := NewTimedBase("src1", now)
	assert.Equal(t, now, b.Time)
}

func TestBase_GetTimeUnix(t *testing.T) {
	now := time.Now()
	b := NewTimedBase("src1", now)
	assert.Equal(t, now.Unix(), b.GetTimeUnix())
}

func TestBase_GetTimeUnixNano(t *testing.T) {
	now := time.Now()
	b := NewTimedBase( "src1", now)
	assert.Equal(t, now.UnixNano(), b.GetTimeUnixNano())
}


func TestBase_AppendSrc(t *testing.T) {
	b := NewBase("src1")
	b.AppendSource("src2")
	assert.Equal(t, "src1", b.SourcePath[0])
	assert.Equal(t, "src2", b.SourcePath[1])
}

func TestBase_IsLastSource(t *testing.T) {
	b := NewBase("src1")
	assert.True(t, b.IsLastSource("src1"), "Last source should be 'src1'")
	b.AppendSource("src2")
	assert.True(t, b.IsLastSource("src2"), "Last source should be 'src2'")
}

func TestBase_InputsMatch(t *testing.T) {
	b := NewBase("src1")
	assert.True(t, b.InputsMatch([]string{"src2", "src1"}), "Should match input list 'src2', 'src1'")
	assert.False(t, b.InputsMatch([]string{"src2"}), "Should not match input list 'src2'")
}

func TestSha1HashString(t *testing.T) {
	s := "sha1 this string"
	assert.Equal(t, "cf23df2207d99a74fbe169e3eba035e633b65d94", Sha1HashString(s))
}

func TestBase_GenDefaultID(t *testing.T) {
	ts := time.Unix(1499156134, 0)
	b := NewTimedBase("src1", ts)
	exp := "27188913c2c2ce1a6cfc5c3a56d072b0297a202f"
	got := b.GenDefaultID()
	assert.Equal(t, exp, got)
}

func TestBase_GetMessageDigest(t *testing.T) {
	b := NewBase("src")
	b.ID = "27188913c2c2ce1a6cfc5c3a56d072b0297a202f"
	exp := "27188913c2c2c"
	got := b.GetMessageDigest()
	assert.Equal(t, exp, got)
}

func TestBase_GetTimeRFC(t *testing.T) {
	ts := time.Unix(1499156134, 123124)
	b := NewTimedBase("src1", ts)
	exp := "2017-07-04T10:15:34.000123+02:00"
	got := b.GetTimeRFC()
	assert.Equal(t, exp, got)
}

func TestNewBaseFromBase(t *testing.T) {
	ts := time.Unix(1499156134, 123124)
	b1 := NewTimedBase("src1", ts)
	b1.Tags["key1"] = "val1"
	b2 := NewBaseFromBase("src2", b1)
	assert.Equal(t, b1.BaseVersion, b2.BaseVersion)
	assert.Equal(t, b1.ID, b2.ID)
	assert.Equal(t, b1.Time, b2.Time)
	assert.Equal(t, b1.SourceID, b2.SourceID)
	assert.Equal(t, append(b1.SourcePath,"src2"), b2.SourcePath)
	assert.Equal(t, b1.SourceSuccess, b2.SourceSuccess)
	assert.Equal(t, b1.Tags, b2.Tags)
	assert.Equal(t, b1.Data, b2.Data)
}