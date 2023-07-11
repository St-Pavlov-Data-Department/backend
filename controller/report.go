package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/controller/requests"
	"github.com/St-Pavlov-Data-Department/backend/datamodel"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r *PavlovController) reportHandler(c *gin.Context) {
	logger := log.CurrentModuleLogger()
	logger.Info("reportHandler")

	request := &requests.ReportRequest{}
	if err := c.ShouldBind(request); err != nil {
		logger.WithError(err).
			Errorf("gin context bind parameter error")
		RespJSONWithError(c, nil, err)
		return
	}
	request.ClientIP = c.ClientIP()

	response, err := r.report(request)

	RespJSONWithError(c, response, err)
}

// --------

type ReportResponse struct {
	ReportUUID string `json:"report_uuid"`
}

func (r *PavlovController) report(req *requests.ReportRequest) (resp *ReportResponse, err error) {
	logger := log.CurrentModuleLogger()
	logger.WithField("ReportRequest", req).Info("")

	response := &ReportResponse{}

	// wrap the creation of loot items and the report itself
	// into a single transaction
	if err := WithTransaction(r.db, func(tx *gorm.DB) error {
		lootReport := &datamodel.LootReport{
			Server:   req.Server,
			Source:   req.Source,
			Version:  req.Version,
			ClientIP: req.ClientIP,

			StageID:     req.StageID,
			ReplayLevel: req.ReplayLevel,
		}
		if err := lootReport.Save(tx); err != nil {
			return err
		}

		for _, reportItem := range req.Loot {
			lootItem := datamodel.LootItem{
				ReportUUID: lootReport.UUID,

				ItemID:   reportItem.ItemID,
				LootType: reportItem.LootType,
				Quantity: reportItem.Quantity,
			}
			if err := lootItem.Save(tx); err != nil {
				return err
			}
		}

		response.ReportUUID = lootReport.UUID

		return nil
	}); err != nil {
		logger.WithError(err).
			Error("lootItem save error")
		return nil, err
	} else {
		logger.Info("lootReport saved")
	}

	return response, nil
}
