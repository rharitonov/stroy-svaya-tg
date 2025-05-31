package model

import "time"

type PileDrivingRecordLine struct {
	ProjectId    int       `json:"project_id"`
	PileNumber   string    `json:"pile_number"`
	PileFieldId  int       `json:"pile_field_id"`
	StartDate    time.Time `json:"start_date"`
	FactPileHead float32   `json:"fact_pile_head"`
	RecordedBy   string    `json:"recorded_by"`
}
