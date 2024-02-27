package controller

import (
	api "github.com/mpedrozoduran/go-orchestrator/internal/api/entities"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence"
	db "github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence/entities"
)

type AuditTrailController struct {
	auditTrailRepository persistence.Repository[db.AuditTrail]
}

func NewAuditTrailController(repo persistence.Repository[db.AuditTrail]) AuditTrailController {
	return AuditTrailController{
		auditTrailRepository: repo,
	}
}

func (a AuditTrailController) GetAuditTrails() ([]api.AuditTrail, error) {
	items, err := a.auditTrailRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var res []api.AuditTrail
	for _, item := range items {
		trail := api.AuditTrail{
			ID:       item.KeyId,
			Request:  item.Request,
			Response: item.Response,
			Created:  item.Created,
		}
		res = append(res, trail)
	}
	return res, nil
}
