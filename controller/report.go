package controller

import (
	"fmt"
	"github.com/St-Pavlov-Data-Department/backend/datamodel"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/St-Pavlov-Data-Department/backend/requests"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

func (r *PavlovController) reportHandler(c *gin.Context) {
	logger := log.CurrentModuleLogger()
	logger.Info("reportHandler")

	request := &requests.UploadReportRequest{}
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

func (r *PavlovController) report(req *requests.UploadReportRequest) (resp *ReportResponse, err error) {
	logger := log.CurrentModuleLogger()
	logger.WithField("UploadReportRequest", req).Info("")

	// validate report info
	if err := r.validateReportStage(req); err != nil {
		logger.WithError(err).Error("validateReportStage error")
		return nil, err
	}

	for _, item := range req.Loot {
		if err := r.validateReportItem(req, item); err != nil {
			logger.WithError(err).Error("validateReportItem error")
			return nil, err
		}
	}

	response := &ReportResponse{}

	// wrap the creation of loot items and the report itself
	// into a single transaction
	if err := utils.WithTransaction(r.db, func(tx *gorm.DB) error {
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

func (r *PavlovController) validateReportStage(report *requests.UploadReportRequest) error {
	if !r.gameResource.Episode.Contains(report.StageID) {
		return fmt.Errorf("stage_id: %d not exist", report.StageID)
	}

	// TODO: for time-limited stages, validate the stage active time range

	return nil
}

func (r *PavlovController) validateReportItem(report *requests.UploadReportRequest, item *requests.LootItem) error {
	gameItem, ok := r.gameResource.Item.Get(item.ItemID)
	if !ok {
		// item_id not exist
		return fmt.Errorf("item_id: %d not exist", item.ItemID)
	}

	itemSourceStages := utils.NewSet(strings.Split(gameItem.Sources, "|")...)

	if !itemSourceStages.Contains(fmt.Sprintf("%d", report.StageID)) {
		// item could not come from this stage
		return fmt.Errorf("item_id: %d could not come from this stage_id: %s", item.ItemID, report.StageID)
	}

	return nil
}
