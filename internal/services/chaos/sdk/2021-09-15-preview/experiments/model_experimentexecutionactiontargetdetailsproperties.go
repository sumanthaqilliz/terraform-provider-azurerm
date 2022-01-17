package experiments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

type ExperimentExecutionActionTargetDetailsProperties struct {
	CompletedDateUtc *string                                      `json:"completedDateUtc,omitempty"`
	Error            *ExperimentExecutionActionTargetDetailsError `json:"error,omitempty"`
	FailedDateUtc    *string                                      `json:"failedDateUtc,omitempty"`
	Status           *string                                      `json:"status,omitempty"`
	Target           *string                                      `json:"target,omitempty"`
}

func (o ExperimentExecutionActionTargetDetailsProperties) GetCompletedDateUtcAsTime() (*time.Time, error) {
	if o.CompletedDateUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CompletedDateUtc, "2006-01-02T15:04:05Z07:00")
}

func (o ExperimentExecutionActionTargetDetailsProperties) SetCompletedDateUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CompletedDateUtc = &formatted
}

func (o ExperimentExecutionActionTargetDetailsProperties) GetFailedDateUtcAsTime() (*time.Time, error) {
	if o.FailedDateUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FailedDateUtc, "2006-01-02T15:04:05Z07:00")
}

func (o ExperimentExecutionActionTargetDetailsProperties) SetFailedDateUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FailedDateUtc = &formatted
}
