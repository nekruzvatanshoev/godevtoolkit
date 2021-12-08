package computeengine

import (
	"context"
	"log"

	compute "google.golang.org/api/compute/v1"
)

func GetInstance(projectID, zone, instanceID string) (*compute.Instance, error) {
	ctx := context.Background()

	service := compute.NewInstancesService(&compute.Service{})
	instance, err := service.Get(projectID, zone, instanceID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	log.Println(instance)

	return instance, nil
}
