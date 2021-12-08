package cmd

const (
	AppName        = "godevtoolkit"
	AppDescription = "godevtoolkit is a tool kit for Golang developers"
	AppVersion     = "0.0.1"
)

const (
	ResourceTrace   = "trace"
	ResourceCompute = "compute"
)

const (
	ActionDescribe = "describe"
	ActionGet      = "get"
	ActionList     = "list"
	ActionReset    = "reset"
)

const (
	TraceCmdName          = ResourceTrace
	TraceCmdUsage         = ""
	TraceCmdDescription   = ""
	ActionDescribeCmdName = ActionDescribe
	ActionListCmdName     = ActionList
	ActionGetCmdName      = ActionGet
	ActionResetCmdName    = ActionReset

	ComputeCmdName        = ResourceCompute
	ComputeCmdUsage       = ""
	ComputeCmdDescription = ""
)

const (
	FlagNameGCPProjectId     = "gcp-project-id"
	FlagNameTraceSpanName    = "span-name"
	FlagNameTraceLatency     = "latency"
	FlagNameTraceMethod      = "method"
	FlagNameTraceStartTime   = "start-time"
	FlagNameTraceServiceName = "service-name"
	FlagNameComputeZone      = "gcp-compute-zone"
)
