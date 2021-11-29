package cloudtrace

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/cloudtrace/v1"
)

func GetTraceById(traceId, projectId string) (*cloudtrace.Trace, error) {
	ctx := context.Background()
	cloudtraceService, err := cloudtrace.NewService(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Created cloud trace client...")

	traceCall := cloudtraceService.Projects.Traces.Get(projectId, traceId)
	trace, err := traceCall.Do()
	if err != nil {
		return nil, err
	}

	return trace, nil
}

func GetTracesByFilter(projectId, spanName, latency, method, startTime, serviceName string) (*cloudtrace.ListTracesResponse, error) {
	ctx := context.Background()
	cloudtraceService, err := cloudtrace.NewService(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Created cloud trace client...")

	duration, err := time.ParseDuration(fmt.Sprint("-", startTime))
	if err != nil {
		return nil, err
	}

	startTimestamp := time.Now().Add(duration).Format(time.RFC3339)

	filterQuery := constructTraceFilterQuery(spanName, method, latency, serviceName)
	if filterQuery == "" {
		return nil, errors.New("filter query is empty")
	}
	traceList := cloudtraceService.Projects.Traces.List(projectId)

	traceList = traceList.Filter(filterQuery)
	traceList = traceList.StartTime(startTimestamp)
	traces, err := traceList.Do()
	if err != nil {
		return nil, err
	}

	return traces, nil
}

func constructTraceFilterQuery(spanName, method, latency, serviceName string) string {
	var filterQuery string
	if spanName != "" {
		filterQuery += fmt.Sprintf("span:%s ", spanName)
	}
	if method != "" {
		filterQuery += fmt.Sprintf("method:%s ", method)
	}
	if latency != "" {
		filterQuery += fmt.Sprintf("latency:%s ", latency)
	}
	if serviceName != "" {
		filterQuery += fmt.Sprintf("service.name:%s ", serviceName)
	}
	return filterQuery
}
