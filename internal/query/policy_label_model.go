package query

import (
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/repository/policy"
)

type LabelPolicyReadModel struct {
	eventstore.ReadModel

	PrimaryColor   string
	SecondaryColor string
	IsActive       bool
}

func (rm *LabelPolicyReadModel) Reduce() error {
	for _, event := range rm.Events {
		switch e := event.(type) {
		case *policy.LabelPolicyAddedEvent:
			rm.PrimaryColor = e.PrimaryColor
			rm.SecondaryColor = e.SecondaryColor
			rm.IsActive = true
		case *policy.LabelPolicyChangedEvent:
			if e.PrimaryColor != nil {
				rm.PrimaryColor = *e.PrimaryColor
			}
			if e.SecondaryColor != nil {
				rm.SecondaryColor = *e.SecondaryColor
			}
		case *policy.LabelPolicyRemovedEvent:
			rm.IsActive = false
		}
	}
	return rm.ReadModel.Reduce()
}