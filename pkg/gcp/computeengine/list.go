package computeengine

import (
	"context"
	"log"

	compute "google.golang.org/api/compute/v1"
)

func ListInstances(projectID, zone string) (*compute.InstanceList, error) {
	ctx := context.Background()

	service := compute.NewInstancesService(&compute.Service{})
	instances, err := service.List(projectID, zone).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	log.Println(instances)

	return instances, nil
}
