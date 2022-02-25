package utils

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/meilisearch/meilisearch-go"
)

type IndexResponse struct {
	*meilisearch.Index
}
type TaskResponse struct {
	*meilisearch.Task
}

func (rd *IndexResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshaled and sent across the wire
	return nil
}

func (rd *TaskResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshaled and sent across the wire
	return nil
}

func NewIndexResponse(index *meilisearch.Index) *IndexResponse {
	resp := &IndexResponse{Index: index}

	return resp
}

func NewTaskResponse(task *meilisearch.Task) *TaskResponse {
	resp := &TaskResponse{Task: task}

	return resp
}

func IndexListResponse(Indexes []*meilisearch.Index) []render.Renderer {
	list := []render.Renderer{}
	for _, Index := range Indexes {
		list = append(list, NewIndexResponse(Index))
	}
	return list
}
