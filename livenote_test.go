package livenote

import "testing"

// Note: To run in the terminal, please type: C:\Go\bin\go.exe test -timeout 30s -run ^TestLiveNote$ github.com/narsilworks/livenote -v -count=1
func TestLiveNote(t *testing.T) {

	mm := &LiveNote{
		Prefix: "Example",
	}

	// Add through methods
	mm.AddInfo("This is an information message!")
	mm.AddInfo("This is an information message too!")
	mm.AddInfo("This is an information message too too!")

	mm.AddWarning("This is a warning!")
	mm.AddWarning("This is a warning too!")
	mm.AddWarning("This is a warning too too!")

	mm.AddError("This is an error!")
	mm.AddError("This is an error too!")
	mm.AddError("This is an error too too!")

	mm.AddAppMsg("This is an application message!")
	mm.AddAppMsg("This is an application message too too!")
	mm.AddAppMsg("This is an application message too too!")

	t.Log(`Dominant Message`, mm.Prevailing())
	t.Log(`Has Error Messages`, mm.HasErrors())
	t.Log(`Has Warning Messages`, mm.HasWarnings())
	t.Log(`Has Info Messages`, mm.HasInfos())

	t.Log(mm.ToString())
}
