package disposal

import (
	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
)

type disposalUpdate struct {
	Filter map[string]interface{} `json:"filter"`
	Update map[string]interface{} `json:"update"`
}

type updateResult struct {
	MatchedCount  int64 `json:"matchedCount,omitempty"`
	ModifiedCount int64 `json:"modifiedCount,omitempty"`
}

// etcd *clientv3.Client,
// 	collection *mongo.Collection,
// 	event *model.Event,
// ) *model.Document

// Update handles "update" events.
func Update(collection *mongo.Collection, event *model.Event) *model.Document {
	// switch event.ServiceAction {
	// case "createSale":
	// 	return createSale(etcd, collection, event)
	// default:
	// 	return UpdateDisposal(collection, event)
	// }

	return UpdateDisposal(collection, event)
}
