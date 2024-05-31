package vcoreparams

const (
	// Label is used to select system tests for vCore deployment.
	Label = "vcore"
	// LabelVCoreDeployment is used to select all basic vCore deployment tests.
	LabelVCoreDeployment = "vcoredeployment"
	// LabelVCoreOperators is used to select all vCore initial-deployment deployment and configuration tests
	// excluding odf part.
	LabelVCoreOperators = "vcoreoperators"
	// LabelVCoreOdf is used to select odf configuration tests.
	LabelVCoreOdf = "vcoreodf"
	// LabelUserCases is used to select all vCore user-cases tests.
	LabelUserCases = "vcoreusercases"
	// LabelVCoreRequirements is used to select all vCore requirements tests.
	LabelVCoreRequirements = "vcorerequirements"
	// VCoreLogLevel configures logging level for vCore related tests.
	VCoreLogLevel = 90
	// LabelVCoreDebug is used to select vCore tests under debug.
	LabelVCoreDebug = "vcoredebug"

	// MasterNodeRole master node role.
	MasterNodeRole = "master"

	// WorkerNodeRole master node role.
	WorkerNodeRole = "worker"

	// ExpectedVCorePPNodesCnt expected user-plane-worker nodes count.
	ExpectedVCorePPNodesCnt = 2

	// ExpectedVCoreCpNodesCnt expected control-plane-worker nodes count.
	ExpectedVCoreCpNodesCnt = 2

	// ExpectedOdfNodesCnt expected odf nodes count.
	ExpectedOdfNodesCnt = 3

	// VCorePpMcpName user-plane-worker workers mcp name.
	VCorePpMcpName = "user-plane-worker"

	// VCoreCpMcpName control-plane-worker workers mcp name.
	VCoreCpMcpName = "control-plane-worker"

	// VCoreOdfMcpName odf workers mcp name.
	VCoreOdfMcpName = "odf"

	// OpenshiftMachineAPINamespace openshift machine-api namespace.
	OpenshiftMachineAPINamespace = "openshift-machine-api"

	// MonitoringNetworkPolicyName monitoring networkpolicy name.
	MonitoringNetworkPolicyName = "allow-from-openshift-monitoring-ingress"

	// AllowAllNetworkPolicyName networkpolicy name.
	AllowAllNetworkPolicyName = "allow-all-ingress"

	// SctpModuleName sctp module name.
	SctpModuleName = "load-sctp-module"

	// TemplateFilesFolder path to the template files folder.
	TemplateFilesFolder = "./internal/config-files/"

	// ConfigurationFolderName path to the folder dedicated to the saving all initial-deployment configuration.
	ConfigurationFolderName = "vcore-configfiles"

	// RegistryRepository local registry repository to mirror images to.
	RegistryRepository = "openshift"

	// SccName scc name.
	SccName = "vcore-control-plane-worker-scc"

	// SystemReservedCPU systemreserved cpu value.
	SystemReservedCPU = "500m"

	// SystemReservedMemory systemreserved memory value.
	SystemReservedMemory = "27Gi"
)
