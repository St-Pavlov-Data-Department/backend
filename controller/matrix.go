package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/datamodel"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/St-Pavlov-Data-Department/backend/requests"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"github.com/gin-gonic/gin"
)

func (r *PavlovController) matrixHandler(c *gin.Context) {
	logger := log.CurrentModuleLogger()
	logger.Info("matrixHandler")

	request := &requests.MatrixRequest{}
	/*
		if err := c.ShouldBind(request); err != nil {
			logger.WithError(err).
				Errorf("gin context bind parameter error")
			RespJSONWithError(c, nil, err)
			return
		}
	*/

	request.Stages = utils.StrToInt64Arr(c.Query("stages"))
	request.Items = utils.StrToInt64Arr(c.Query("items"))
	request.Server = c.Query("server")

	response, err := r.getMatrix(request)

	RespJSONWithError(c, response, err)
}

// --------

type MatrixArray []*MatrixPoint

type MatrixPoint struct {
	StageId        int64 `json:"stage_id"`
	ItemId         int64 `json:"item_id"`
	StartTimeMilli int64 `json:"start_time_milli"`
	EndTimeMilli   int64 `json:"end_time_milli"`
	Quantity       int64 `json:"quantity"`
	ReplayCount    int64 `json:"replay_count"`
}

func (r *PavlovController) getMatrix(req *requests.MatrixRequest) (
	MatrixArray, error,
) {
	logger := log.CurrentModuleLogger()
	logger.WithField("MatrixRequest", req).Info("")

	/*
		// query for metrix data
		matrixList := datamodel.MatrixList{}
		if err := matrixList.LoadByRequest(r.db, req); err != nil {
			logger.WithError(err).
				Errorf("MatrixList LoadByRequest error")
			return nil, err
		}
	*/

	// in the early stage of development, just calculate among all the reports first.
	reportReq := &requests.QueryReportRequest{
		Server: req.Server,
		Stages: req.Stages,
		Items:  req.Items,
	}

	reportList := datamodel.LootReportList{}
	if err := reportList.LoadByRequest(r.db, reportReq); err != nil {
		logger.WithError(err).
			Errorf("load reportList error")
		return nil, err
	}

	// map[stage_id] replay_count
	stageReplay := map[int64]int64{}
	// map[stage_id] map[item_id] item_count
	stageItems := map[int64]map[int64]int64{}

	// merge data
	stageSet := utils.NewSet[int64](req.Stages...)
	for _, r := range reportList {
		if stageSet.Contains(r.StageID) {
			stageReplay[r.StageID] += r.ReplayLevel
			for _, item := range r.Loot {
				if _, ok := stageItems[r.StageID]; !ok {
					stageItems[r.StageID] = map[int64]int64{}
				}
				stageItems[r.StageID][item.ItemID] += item.Quantity
			}
		}
	}

	var response MatrixArray = make([]*MatrixPoint, 0)
	for stageId, replayCount := range stageReplay {
		for itemId, itemCount := range stageItems[stageId] {
			response = append(response,
				&MatrixPoint{
					StageId:        stageId,
					ItemId:         itemId,
					StartTimeMilli: 0, // TODO: collect time rage from reports
					EndTimeMilli:   0,
					Quantity:       itemCount,
					ReplayCount:    replayCount,
				},
			)
		}
	}

	/*
		var response MatrixArray = make([]*MatrixPoint, len(matrixList), len(matrixList))
		for i, m := range matrixList {
			response[i] = &MatrixPoint{
				StageId:        m.StageId,
				ItemId:         m.ItemId,
				StartTimeMilli: m.StartTimeMilli,
				EndTimeMilli:   m.EndTimeMilli,
				Quantity:       m.Quantity,
				ReplayCount:    m.ReplayCount,
			}
		}
	*/

	return response, nil
}
