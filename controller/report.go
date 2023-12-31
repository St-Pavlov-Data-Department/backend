package controller

import (
	"fmt"
	"github.com/St-Pavlov-Data-Department/backend/datamodel"
	"github.com/St-Pavlov-Data-Department/backend/gameresource"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/St-Pavlov-Data-Department/backend/requests"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// validate report request
	if err := r.validateReportReq(req); err != nil {
		logger.WithError(err).Error("validateReportReq error")
		return nil, err
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

			EpisodeID:   req.EpisodeID,
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

func (r *PavlovController) validateReportReq(report *requests.UploadReportRequest) error {
	// validate episode_id
	episodeInfo, ok := r.gameResource.Episode.Get(report.EpisodeID)
	if !ok {
		return fmt.Errorf("episode_id: %d not exist", report.EpisodeID)
	}

	// TODO: for time-limited episodes, validate the episode active time range

	episodeRewards := utils.NewSet(
		utils.Apply(episodeInfo.RewardList,
			func(ref *gameresource.Reference) int64 {
				// find item id from reference
				switch ref.Table {
				case gameresource.ITEM:
					item, ok := r.gameResource.Item.Get(ref.ID)
					if !ok {
						return 0
					}
					return item.Id

				default:
					return 0
				}
			},
		)...,
	)

	for _, lootItem := range report.Loot {
		// validate loot_item
		if !r.gameResource.Item.Contains(lootItem.ItemID) {
			// item_id not exist
			return fmt.Errorf("item_id: %d not exist", lootItem.ItemID)
		}

		// validate loot_item possible to be found in this episode
		if !episodeRewards.Contains(lootItem.ItemID) {
			return fmt.Errorf("item_id: %d could not come from this episode_id: %d", lootItem.ItemID, report.EpisodeID)
		}
	}

	return nil
}
