package errorgroups

import (
	"strconv"
	"strings"

	"github.com/highlight-run/highlight/backend/model"
	privateModel "github.com/highlight-run/highlight/backend/private-graph/graph/model"
)

func GetFingerprints(projectID int, errorTraces []privateModel.ErrorTrace) []*model.ErrorFingerprint {
	fingerprints := []*model.ErrorFingerprint{}

	for idx, frame := range errorTraces {
		codeVal := joinStringPtrs(frame.LinesBefore, frame.LineContent, frame.LinesAfter)
		if codeVal != "" {
			code := model.ErrorFingerprint{
				ProjectID: projectID,
				Type:      model.Fingerprint.StackFrameCode,
				Value:     codeVal,
				Index:     idx,
			}
			fingerprints = append(fingerprints, &code)
		}

		metaVal := joinStringPtrs(frame.FileName, frame.FunctionName) +
			joinIntPtrs(frame.LineNumber, frame.ColumnNumber)
		if metaVal != "" {
			meta := model.ErrorFingerprint{
				ProjectID: projectID,
				Type:      model.Fingerprint.StackFrameMetadata,
				Value:     metaVal,
				Index:     idx,
			}
			fingerprints = append(fingerprints, &meta)
		}
	}
	return fingerprints
}

func joinStringPtrs(ptrs ...*string) string {
	var sb strings.Builder
	for _, ptr := range ptrs {
		if ptr != nil {
			sb.WriteString(*ptr)
			sb.WriteString(";")
		}
	}
	return sb.String()
}

func joinIntPtrs(ptrs ...*int) string {
	var sb strings.Builder
	for _, ptr := range ptrs {
		if ptr != nil {
			sb.WriteString(strconv.Itoa(*ptr))
			sb.WriteString(";")
		}
	}
	return sb.String()
}
