package disposal

import (
	"encoding/json"
	"log"

	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"
)

func UpdateDisposal(coll *mongo.Collection, event *model.Event) *model.Document {
	disposeUpdate := &disposalUpdate{}

	err := json.Unmarshal(event.Data, disposeUpdate)
	if err != nil {
		err = errors.Wrap(err, "Update: Error while unmarshalling Event-data")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	if len(disposeUpdate.Filter) == 0 {
		err = errors.New("blank filter provided")
		err = errors.Wrap(err, "Update")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}
	if len(disposeUpdate.Update) == 0 {
		err = errors.New("blank update provided")
		err = errors.Wrap(err, "Update")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}
	if disposeUpdate.Update["itemID"] == (uuuid.UUID{}).String() {
		err = errors.New("found blank itemID in update")
		err = errors.Wrap(err, "Update")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	updateStats, err := coll.UpdateMany(disposeUpdate.Filter, disposeUpdate.Update)
	if err != nil {
		err = errors.Wrap(err, "Update: Error in UpdateMany")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     DatabaseError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	result := &updateResult{
		MatchedCount:  updateStats.MatchedCount,
		ModifiedCount: updateStats.ModifiedCount,
	}
	resultMarshal, err := json.Marshal(result)
	if err != nil {
		err = errors.Wrap(err, "Update: Error marshalling Inventory Update-result")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	return &model.Document{
		AggregateID:   event.AggregateID,
		CorrelationID: event.CorrelationID,
		Result:        resultMarshal,
		EventAction:   event.EventAction,
		ServiceAction: event.ServiceAction,
		UUID:          event.UUID,
	}
}
