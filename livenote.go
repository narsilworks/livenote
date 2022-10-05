// LiveNote is lightweight direct logging library
package livenote

import (
	"runtime"
	"strings"
)

// NoteType
type NoteType string

// NoteType constants
const (
	Info  NoteType = "INF"
	Warn  NoteType = "WRN"
	Error NoteType = "ERR"
	App   NoteType = ""
)

const DelimMsgType string = `: `

type LiveNote struct {
	Prefix  string // Prefix
	ln      []LiveNoteInfo
	osIsWin bool
}

type LiveNoteInfo struct {
	Type    NoteType
	Prefix  string
	Message string
}

func NewLiveNote(prefix string) *LiveNote {
	return &LiveNote{
		Prefix:  prefix,
		ln:      make([]LiveNoteInfo, 0),
		osIsWin: runtime.GOOS == "windows",
	}
}

// AddInfo adds an information message
func (r *LiveNote) AddInfo(Message ...string) {
	for _, m := range Message {
		addMessage(&r.ln, r.Prefix, m, Info)
	}
}

// AddWarning adds a warning message
func (r *LiveNote) AddWarning(Message ...string) {
	for _, m := range Message {
		addMessage(&r.ln, r.Prefix, m, Warn)
	}
}

// AddError adds an error message
func (r *LiveNote) AddError(Message ...string) {
	for _, m := range Message {
		addMessage(&r.ln, r.Prefix, m, Error)
	}
}

// AddAppMsg adds an application message
func (r *LiveNote) AddAppMsg(Message ...string) {
	for _, m := range Message {
		addMessage(&r.ln, r.Prefix, m, App)
	}
}

// Append adds a note object to the current list
func (r *LiveNote) Append(ln LiveNoteInfo) {
	r.ln = append(r.ln, ln)
}

// HasErrors - Checks if the message array has errors
func (r LiveNote) HasErrors() bool {

	for _, ln := range r.ln {
		if ln.Type == Error {
			return true
		}
	}
	return false
}

// HasWarnings - Checks if the message array has warnings
func (r LiveNote) HasWarnings() bool {

	for _, ln := range r.ln {
		if ln.Type == Warn {
			return true
		}
	}

	return false
}

// HasInfos - Checks if the message array has information messages
func (r LiveNote) HasInfos() bool {
	for _, ln := range r.ln {
		if ln.Type == Info {
			return true
		}
	}
	return false
}

// Prevailing checks for a dominant message
func (r *LiveNote) Prevailing() NoteType {
	return getDominantNoteType(&r.ln)
}

// Notes will list all notes
func (r *LiveNote) Notes() []LiveNoteInfo {
	return r.ln
}

// ToString return the messages as a carriage/return delimited string
func (r *LiveNote) ToString() string {

	lf := "\n"
	if r.osIsWin {
		lf = "\r\n"
	}

	sb := strings.Builder{}
	for _, v := range r.ln {
		sb.Write([]byte(v.ToString() + lf))
	}

	return sb.String()
}

// ToString return the messages as a carriage/return delimited string
func (lni *LiveNoteInfo) ToString() string {

	td := ""
	td += string(lni.Type)

	if lni.Prefix != "" {
		td += "[" + lni.Prefix + "]"
	}
	td += DelimMsgType
	td += lni.Message

	return td
}

// add new message to the message array
func addMessage(nt *[]LiveNoteInfo, prefix, msg string, typ NoteType) {

	msg = strings.TrimSpace(msg)
	*nt = append(*nt, LiveNoteInfo{
		Prefix:  prefix,
		Message: msg,
		Type:    typ,
	})
}

// get dominant message
func getDominantNoteType(msgs *[]LiveNoteInfo) NoteType {

	nfo := 0
	wrn := 0
	err := 0

	for _, msg := range *msgs {
		switch msg.Type {
		case Info:
			nfo++
		case Warn:
			wrn++
		case Error:
			err++
		}
	}

	if nfo > wrn && nfo > err {
		return Info
	}

	if wrn > nfo && wrn > err {
		return Warn
	}

	if err > nfo && err > wrn {
		return Error
	}

	return App
}
