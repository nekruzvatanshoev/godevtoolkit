package computeengine

import (
	"context"
	"log"

	compute "google.golang.org/api/compute/v1"
)

func ResetInstance(projectID, zone, instanceID string) (*compute.Operation, error) {
	ctx := context.Background()

	service, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}
	instanceService := compute.NewInstancesService(service)
	operationCall := instanceService.Reset(projectID, zone, instanceID)
	operation, err := operationCall.Do()
	if err != nil {
		return nil, err
	}
	log.Println(operation)

	return operation, nil
}
