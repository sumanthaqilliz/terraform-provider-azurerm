package experiments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

type ExperimentExecutionDetailsProperties struct {
	CreatedDateUtc    *string                                             `json:"createdDateUtc,omitempty"`
	ExperimentId      *string                                             `json:"experimentId,omitempty"`
	FailureReason     *string                                             `json:"failureReason,omitempty"`
	LastActionDateUtc *string                                             `json:"lastActionDateUtc,omitempty"`
	RunInformation    *ExperimentExecutionDetailsPropertiesRunInformation `json:"runInformation,omitempty"`
	StartDateUtc      *string                                             `json:"startDateUtc,omitempty"`
	Status            *string                                             `json:"status,omitempty"`
	StopDateUtc       *string                                             `json:"stopDateUtc,omitempty"`
}

func (o ExperimentExecutionDetailsProperties) GetCreatedDateUtcAsTime() (*time.Time, error) {
	if o.CreatedDateUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDateUtc, "2006-01-02T15:04:05Z07:00")
}

func (o ExperimentExecutionDetailsProperties) SetCreatedDateUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDateUtc = &formatted
}

func (o ExperimentExecutionDetailsProperties) GetLastActionDateUtcAsTime() (*time.Time, error) {
	if o.LastActionDateUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastActionDateUtc, "2006-01-02T15:04:05Z07:00")
}

func (o ExperimentExecutionDetailsProperties) SetLastActionDateUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastActionDateUtc = &formatted
}

func (o ExperimentExecutionDetailsProperties) GetStartDateUtcAsTime() (*time.Time, error) {
	if o.StartDateUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDateUtc, "2006-01-02T15:04:05Z07:00")
}

func (o ExperimentExecutionDetailsProperties) SetStartDateUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDateUtc = &formatted
}

func (o ExperimentExecutionDetailsProperties) GetStopDateUtcAsTime() (*time.Time, error) {
	if o.StopDateUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StopDateUtc, "2006-01-02T15:04:05Z07:00")
}

func (o ExperimentExecutionDetailsProperties) SetStopDateUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StopDateUtc = &formatted
}
