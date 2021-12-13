package computeengine

import (
	"context"
	"log"

	compute "google.golang.org/api/compute/v1"
)

func ListInstances(projectID, zone string) (*compute.InstanceList, error) {
	ctx := context.Background()

	service, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}
	instanceService := compute.NewInstancesService(service)
	instanceCall := instanceService.List(projectID, zone)
	instances, err := instanceCall.Do()
	if err != nil {
		return nil, err
	}

	log.Println(instances)

	return instances, nil
}
