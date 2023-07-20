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

	request.Episodes = utils.StrToInt64Arr(c.Query("episodes"))
	request.Items = utils.StrToInt64Arr(c.Query("items"))
	request.Server = c.Query("server")

	response, err := r.getMatrix(request)

	RespJSONWithError(c, response, err)
}

// --------

type MatrixArray []*MatrixPoint

type MatrixPoint struct {
	EpisodeID      int64 `json:"episode_id"`
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
		Server:   req.Server,
		Episodes: req.Episodes,
		Items:    req.Items,
	}

	reportList := datamodel.LootReportList{}
	if err := reportList.LoadByRequest(r.db, reportReq); err != nil {
		logger.WithError(err).
			Errorf("load reportList error")
		return nil, err
	}

	// map[episode_id] replay_count
	episodeReplay := map[int64]int64{}
	// map[episode_id] map[item_id] item_count
	episodeItems := map[int64]map[int64]int64{}

	// merge data
	for _, r := range reportList {
		episodeReplay[r.EpisodeID] += r.ReplayLevel
		for _, item := range r.Loot {
			if _, ok := episodeItems[r.EpisodeID]; !ok {
				episodeItems[r.EpisodeID] = map[int64]int64{}
			}
			episodeItems[r.EpisodeID][item.ItemID] += item.Quantity
		}
	}

	var response MatrixArray = make([]*MatrixPoint, 0)
	for episodeID, replayCount := range episodeReplay {
		for itemId, itemCount := range episodeItems[episodeID] {
			response = append(response,
				&MatrixPoint{
					EpisodeID:      episodeID,
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
				EpisodeID:        m.EpisodeID,
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
