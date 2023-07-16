package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/datamodel"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/St-Pavlov-Data-Department/backend/requests"
	"github.com/gin-gonic/gin"
)

func (r *PavlovController) matrixHandler(c *gin.Context) {
	logger := log.CurrentModuleLogger()
	logger.Info("matrixHandler")

	request := &requests.MatrixRequest{}
	if err := c.ShouldBind(request); err != nil {
		logger.WithError(err).
			Errorf("gin context bind parameter error")
		RespJSONWithError(c, nil, err)
		return
	}

	response, err := r.getMatrix(request)

	RespJSONWithError(c, response, err)
}

// --------

type MatrixArray []*MatrixPoint

type MatrixPoint struct {
	StageId        string `json:"stage_id"`
	ItemId         string `json:"item_id"`
	StartTimeMilli int64  `json:"start_time_milli"`
	EndTimeMilli   int64  `json:"end_time_milli"`
	Quantity       int    `json:"quantity"`
	ReplayCount    int    `json:"replay_count"`
}

func (r *PavlovController) getMatrix(req *requests.MatrixRequest) (
	MatrixArray, error,
) {
	logger := log.CurrentModuleLogger()
	logger.WithField("MatrixRequest", req).Info("")

	// query for metrix data
	matrixList := datamodel.MatrixList{}
	if err := matrixList.LoadByRequest(r.db, req); err != nil {
		logger.WithError(err).
			Errorf("MatrixList LoadByRequest error")
		return nil, err
	}

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

	return response, nil
}
