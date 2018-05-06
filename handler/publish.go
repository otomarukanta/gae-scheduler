package handler

import (
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/go-chi/chi"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// GetPublish ...
func GetPublish(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	projectID, ok := os.LookupEnv("PROJECT_ID")
	log.Infof(ctx, "Project ID is %v.", projectID)
	if !ok {
		http.Error(w, http.StatusText(404), 404)
		log.Errorf(ctx, "Project ID not found.")
		return
	}
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		log.Errorf(ctx, "Cannot create pubsub client.")
		return
	}

	topic := chi.URLParam(r, "topic")
	publishedAt := time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60)).String()
	result := client.Topic(topic).Publish(ctx, &pubsub.Message{
		Data: []byte(publishedAt),
	})

	id, err := result.Get(ctx)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		log.Errorf(ctx, "Failed to publish message.")
		return
	}
	log.Infof(ctx, "Published message successfully. ID is  %v", id)
}
