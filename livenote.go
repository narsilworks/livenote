// LiveNote is lightweight direct logging library
package livenote

import (
	"fmt"
	"runtime"
	"strings"
)

// NoteType
type NoteType string

// NoteType constants
const (
	Info    NoteType = "INF"
	Warn    NoteType = "WRN"
	Error   NoteType = "ERR"
	Fatal   NoteType = "FTL"
	Success NoteType = "SUC"
	App     NoteType = ""
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

// Fmt accepts format and argument to return a string
func Fmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// AddInfo adds an information message
func (r *LiveNote) AddInfo(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Info)
	}
}

// AddWarning adds a warning message
func (r *LiveNote) AddWarning(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Warn)
	}
}

// AddError adds an error message
func (r *LiveNote) AddError(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Error)
	}
}

// AddSuccess adds a success message
func (r *LiveNote) AddSuccess(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Success)
	}
}

// AddAppMsg adds an application message
func (r *LiveNote) AddAppMsg(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, App)
	}
}

// Append adds a note object or more to the current list
func (r *LiveNote) Append(ln ...LiveNoteInfo) {
	r.ln = append(r.ln, ln...)
}

// Clear live notes
func (r *LiveNote) Clear() {
	r.ln = []LiveNoteInfo{}
}

// HasErrors checks if the message array has errors
func (r LiveNote) HasErrors() bool {
	for _, ln := range r.ln {
		if ln.Type == Error {
			return true
		}
	}
	return false
}

// HasWarnings checks if the message array has warnings
func (r LiveNote) HasWarnings() bool {
	for _, ln := range r.ln {
		if ln.Type == Warn {
			return true
		}
	}
	return false
}

// HasInfos checks if the message array has information messages
func (r LiveNote) HasInfos() bool {
	for _, ln := range r.ln {
		if ln.Type == Info {
			return true
		}
	}
	return false
}

// HasSuccess checks if the message array has success messages
func (r LiveNote) HasSucceses() bool {
	for _, ln := range r.ln {
		if ln.Type == Success {
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
	var (
		nfo, wrn, err, suc int
	)

	for _, msg := range *msgs {
		switch msg.Type {
		case Info:
			nfo++
		case Warn:
			wrn++
		case Error:
			err++
		case Success:
			suc++
		}
	}
	if nfo > wrn && nfo > err && nfo > suc {
		return Info
	}
	if wrn > nfo && wrn > err && wrn > suc {
		return Warn
	}
	if err > nfo && err > wrn && err > suc {
		return Error
	}
	if suc > nfo && suc > wrn && suc > err {
		return Success
	}
	return App
}
