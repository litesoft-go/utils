package issue

import (
	"strings"

	// standard libs only above!

	"github.com/litesoft-go/utils/strs"
)

type Collector struct {
	sourceTypesByIssue map[string]*issueSourceTypeMap
}

//noinspection GoUnusedExportedFunction
func NewSourceReference(ref string) SourceReference {
	return &baseSR{ref: ref}
}

type SourceReference interface {
	GetReference() string

	AddDetails(label, content string) SourceReference

	AddIndentableDetails(object strs.IndentStringer) SourceReference

	details() []*detail
}

type baseSR struct {
	ref string
}

func (in *baseSR) GetReference() string {
	return in.ref
}

func (in *baseSR) AddDetails(label, content string) SourceReference {
	return &detailedSR{parent: in, detailRef: newDetail(label, content)}
}

func (in *baseSR) AddIndentableDetails(object strs.IndentStringer) SourceReference {
	return &detailedSR{parent: in, detailRef: newIndentableDetails(object)}
}

func (in *baseSR) details() (details []*detail) {
	return
}

type detailedSR struct {
	parent    SourceReference
	detailRef *detail
}

func (in *detailedSR) GetReference() string {
	return in.parent.GetReference()
}

func (in *detailedSR) AddDetails(label, content string) SourceReference {
	return &detailedSR{parent: in, detailRef: newDetail(label, content)}
}

func (in *detailedSR) AddIndentableDetails(object strs.IndentStringer) SourceReference {
	return &detailedSR{parent: in, detailRef: newIndentableDetails(object)}
}

func (in *detailedSR) details() (details []*detail) {
	return
}

func newIndentableDetails(object strs.IndentStringer) *detail {
	return &detail{object: object}
}

func newDetail(label, content string) *detail {
	return &detail{label: label, content: content}
}

type detail struct {
	label   string
	content string
	object  strs.IndentStringer
}

func (in *detail) IndentString(_ string, ib *strs.IndentedBuilder) {
	switch {
	case in == nil:
		return
	case in.object != nil:
		in.object.IndentString("", ib)
	case !strings.Contains(in.content, "\n"):
		ib.AddLine(in.label, ": ", in.content)
	default:
		ib.AddLine(in.label, ":")
		ib.AddIndent()
		ib.AddLine(in.content)
		ib.DropIndent()
	}
}

func (in *Collector) IsEmpty() (empty bool) {
	return (in == nil) || (len(in.sourceTypesByIssue) == 0) // Left to Right!
}

func (in *Collector) Add(problem, problemSourceType string, reference SourceReference) {
	if in.sourceTypesByIssue == nil {
		in.sourceTypesByIssue = make(map[string]*issueSourceTypeMap)
	}
	st := in.sourceTypesByIssue[problem]
	if st == nil {
		st = &issueSourceTypeMap{sourceRefsByType: make(map[string]*issueSourceReferenceList)}
		in.sourceTypesByIssue[problem] = st
	}
	st.add(problemSourceType, reference)
}

type issueSourceTypeMap struct {
	sourceRefsByType map[string]*issueSourceReferenceList
}

func (in *issueSourceTypeMap) add(problemSourceType string, reference SourceReference) {
	wl := in.sourceRefsByType[problemSourceType]
	if wl == nil {
		wl = &issueSourceReferenceList{}
		in.sourceRefsByType[problemSourceType] = wl
	}
	wl.add(reference)
}

func (in *issueSourceTypeMap) IndentString(_ string, ib *strs.IndentedBuilder) {
	if in != nil {
		for kST, vSRs := range in.sourceRefsByType {
			if !vSRs.anyDetails() {
				ib.AddLine(kST, ": ", vSRs.collectReferences())
			} else {
				ib.AddLine(kST, ":")
				ib.AddIndent()
				vSRs.IndentString("", ib)
				ib.DropIndent()
			}
		}
	}
}

type issueSourceReferenceList struct {
	issueSourceReferences []SourceReference
}

func (in *issueSourceReferenceList) add(reference SourceReference) {
	in.issueSourceReferences = append(in.issueSourceReferences, reference)
}

func (in *issueSourceReferenceList) anyDetails() bool {
	for _, sr := range in.issueSourceReferences {
		if len(sr.details()) != 0 {
			return true
		}
	}
	return false
}

func (in *issueSourceReferenceList) collectReferences() (out string) {
	prefix := ""
	for _, sr := range in.issueSourceReferences {
		out += prefix + sr.GetReference()
		prefix = ", "
	}
	return
}

func (in *issueSourceReferenceList) IndentString(_ string, ib *strs.IndentedBuilder) {
	if in != nil {
		for _, sr := range in.issueSourceReferences {
			reference := sr.GetReference()
			details := sr.details()
			if len(details) == 0 {
				ib.AddLine(reference)
			} else {
				ib.AddLine(reference, ":")
				ib.AddIndent()
				for _, detail := range details {
					detail.IndentString("", ib)
				}
				ib.DropIndent()
			}
		}
	}
}

func (in *Collector) String() string {
	ib := strs.NewIndentedBuilder(4)
	in.IndentString("Issues", ib)
	return ib.String()
}

func (in *Collector) IndentString(prefix string, ib *strs.IndentedBuilder) {
	if in.IsEmpty() {
		ib.AddLine("No " + prefix)
	} else {
		ib.AddLine(prefix + ":")
		ib.AddIndent()
		for kIssue, vSTM := range in.sourceTypesByIssue {
			ib.AddLine(kIssue, ":")
			ib.AddIndent()
			vSTM.IndentString("", ib)
			ib.DropIndent()
		}
		ib.DropIndent()
	}
}
