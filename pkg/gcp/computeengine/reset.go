package computeengine

import (
	"context"
	"log"

	compute "google.golang.org/api/compute/v1"
)

func ResetInstance(projectID, zone, instanceID string) (*compute.Operation, error) {
	ctx := context.Background()

	service := compute.NewInstancesService(&compute.Service{})
	operation, err := service.Reset(projectID, zone, instanceID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	log.Println(operation)

	return operation, nil
}
