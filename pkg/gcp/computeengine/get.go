package computeengine

import (
	"context"
	"log"

	compute "google.golang.org/api/compute/v1"
)

func GetInstance(projectID, zone, instanceID string) (*compute.Instance, error) {
	ctx := context.Background()

	service, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}
	instanceService := compute.NewInstancesService(service)
	instanceCall := instanceService.Get(projectID, zone, instanceID)
	instance, err := instanceCall.Do()
	if err != nil {
		return nil, err
	}
	log.Println(instance)

	return instance, nil
}
