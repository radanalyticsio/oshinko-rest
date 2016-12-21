package integration

import (
	"strings"
	"testing"

	kapierror "k8s.io/kubernetes/pkg/api/errors"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	buildapi "github.com/openshift/origin/pkg/build/api"
	"github.com/openshift/origin/pkg/client"
	policy "github.com/openshift/origin/pkg/cmd/admin/policy"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	imageapi "github.com/openshift/origin/pkg/image/api"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestPolicyBasedRestrictionOfBuildCreateAndCloneByStrategy(t *testing.T) {
	defer testutil.DumpEtcdOnFailure(t)
	clusterAdminClient, projectAdminClient, projectEditorClient := setupBuildStrategyTest(t, false)

	clients := map[string]*client.Client{"admin": projectAdminClient, "editor": projectEditorClient}
	builds := map[string]*buildapi.Build{}

	// Create builds to setup test
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			var err error
			if builds[string(strategy)+clientType], err = createBuild(t, client.Builds(testutil.Namespace()), strategy); err != nil {
				t.Errorf("unexpected error for strategy %s and client %s: %v", strategy, clientType, err)
			}
		}
	}

	// by default amdins and editors can clone builds
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := cloneBuild(t, client.Builds(testutil.Namespace()), builds[string(strategy)+clientType]); err != nil {
				t.Errorf("unexpected clone error for strategy %s and client %s: %v", strategy, clientType, err)
			}
		}
	}

	removeBuildStrategyRoleResources(t, clusterAdminClient, projectAdminClient, projectEditorClient)

	// make sure builds are rejected
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := createBuild(t, client.Builds(testutil.Namespace()), strategy); !kapierror.IsForbidden(err) {
				t.Errorf("expected forbidden for strategy %s and client %s: got %v", strategy, clientType, err)
			}
		}
	}

	// make sure build updates are rejected
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := updateBuild(t, client.Builds(testutil.Namespace()), builds[string(strategy)+clientType]); !kapierror.IsForbidden(err) {
				t.Errorf("expected forbidden for strategy %s and client %s: got %v", strategy, clientType, err)
			}
		}
	}

	// make sure clone is rejected
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := cloneBuild(t, client.Builds(testutil.Namespace()), builds[string(strategy)+clientType]); !kapierror.IsForbidden(err) {
				t.Errorf("expected forbidden for strategy %s and client %s: got %v", strategy, clientType, err)
			}
		}
	}
}

func TestPolicyBasedRestrictionOfBuildConfigCreateAndInstantiateByStrategy(t *testing.T) {
	defer testutil.DumpEtcdOnFailure(t)
	clusterAdminClient, projectAdminClient, projectEditorClient := setupBuildStrategyTest(t, true)

	clients := map[string]*client.Client{"admin": projectAdminClient, "editor": projectEditorClient}
	buildConfigs := map[string]*buildapi.BuildConfig{}

	// by default admins and editors can create all type of buildconfigs
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			var err error
			if buildConfigs[string(strategy)+clientType], err = createBuildConfig(t, client.BuildConfigs(testutil.Namespace()), strategy); err != nil {
				t.Errorf("unexpected error for strategy %s and client %s: %v", strategy, clientType, err)
			}
		}
	}

	// by default admins and editors can instantiate build configs
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := instantiateBuildConfig(t, client.BuildConfigs(testutil.Namespace()), buildConfigs[string(strategy)+clientType]); err != nil {
				t.Errorf("unexpected instantiate error for strategy %s and client %s: %v", strategy, clientType, err)
			}
		}
	}

	removeBuildStrategyRoleResources(t, clusterAdminClient, projectAdminClient, projectEditorClient)

	// make sure buildconfigs are rejected
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := createBuildConfig(t, client.BuildConfigs(testutil.Namespace()), strategy); !kapierror.IsForbidden(err) {
				t.Errorf("expected forbidden for strategy %s and client %s: got %v", strategy, clientType, err)
			}
		}
	}

	// make sure buildconfig updates are rejected
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := updateBuildConfig(t, client.BuildConfigs(testutil.Namespace()), buildConfigs[string(strategy)+clientType]); !kapierror.IsForbidden(err) {
				t.Errorf("expected forbidden for strategy %s and client %s: got %v", strategy, clientType, err)
			}
		}
	}

	// make sure instantiate is rejected
	for _, strategy := range buildStrategyTypes() {
		for clientType, client := range clients {
			if _, err := instantiateBuildConfig(t, client.BuildConfigs(testutil.Namespace()), buildConfigs[string(strategy)+clientType]); !kapierror.IsForbidden(err) {
				t.Errorf("expected forbidden for strategy %s and client %s: got %v", strategy, clientType, err)
			}
		}
	}
}

func buildStrategyTypes() []string {
	return []string{"source", "docker", "custom", "jenkinspipeline"}
}

func setupBuildStrategyTest(t *testing.T, includeControllers bool) (clusterAdminClient, projectAdminClient, projectEditorClient *client.Client) {
	testutil.RequireEtcd(t)
	namespace := testutil.Namespace()
	var clusterAdminKubeConfig string
	var err error

	if includeControllers {
		_, clusterAdminKubeConfig, err = testserver.StartTestMaster()
	} else {
		_, clusterAdminKubeConfig, err = testserver.StartTestMasterAPI()
	}
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	clusterAdminClient, err = testutil.GetClusterAdminClient(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	clusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projectAdminClient, err = testserver.CreateNewProject(clusterAdminClient, *clusterAdminClientConfig, namespace, "harold")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	projectEditorClient, _, _, err = testutil.GetClientForUser(*clusterAdminClientConfig, "joe")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	addJoe := &policy.RoleModificationOptions{
		RoleNamespace:       "",
		RoleName:            bootstrappolicy.EditRoleName,
		RoleBindingAccessor: policy.NewLocalRoleBindingAccessor(namespace, projectAdminClient),
		Users:               []string{"joe"},
	}
	if err := addJoe.AddRole(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := testutil.WaitForPolicyUpdate(projectEditorClient, namespace, "create", buildapi.Resource(authorizationapi.DockerBuildResource), true); err != nil {
		t.Fatalf(err.Error())
	}

	// Create builder image stream and tag
	imageStream := &imageapi.ImageStream{}
	imageStream.Name = "builderimage"
	_, err = clusterAdminClient.ImageStreams(testutil.Namespace()).Create(imageStream)
	if err != nil {
		t.Fatalf("Couldn't create ImageStream: %v", err)
	}
	// Create image stream mapping
	imageStreamMapping := &imageapi.ImageStreamMapping{}
	imageStreamMapping.Name = "builderimage"
	imageStreamMapping.Tag = "latest"
	imageStreamMapping.Image.Name = "image-id"
	imageStreamMapping.Image.DockerImageReference = "test/builderimage:latest"
	err = clusterAdminClient.ImageStreamMappings(testutil.Namespace()).Create(imageStreamMapping)
	if err != nil {
		t.Fatalf("Couldn't create ImageStreamMapping: %v", err)
	}

	template, err := testutil.GetTemplateFixture("../../examples/jenkins/jenkins-ephemeral-template.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	template.Name = "jenkins"
	template.Namespace = "openshift"

	_, err = clusterAdminClient.Templates("openshift").Create(template)
	if err != nil {
		t.Fatalf("Couldn't create jenkins template: %v", err)
	}

	return
}

func removeBuildStrategyRoleResources(t *testing.T, clusterAdminClient, projectAdminClient, projectEditorClient *client.Client) {
	// remove resources from role so that certain build strategies are forbidden
	for _, role := range []string{bootstrappolicy.BuildStrategyCustomRoleName, bootstrappolicy.BuildStrategyDockerRoleName, bootstrappolicy.BuildStrategySourceRoleName, bootstrappolicy.BuildStrategyJenkinsPipelineRoleName} {
		remove := &policy.RoleModificationOptions{
			RoleNamespace:       "",
			RoleName:            role,
			RoleBindingAccessor: policy.NewClusterRoleBindingAccessor(clusterAdminClient),
			Groups:              []string{"system:authenticated"},
		}
		if err := remove.RemoveRole(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if err := testutil.WaitForPolicyUpdate(projectEditorClient, testutil.Namespace(), "create", buildapi.Resource(authorizationapi.DockerBuildResource), false); err != nil {
		t.Fatal(err)
	}
	if err := testutil.WaitForPolicyUpdate(projectEditorClient, testutil.Namespace(), "create", buildapi.Resource(authorizationapi.SourceBuildResource), false); err != nil {
		t.Fatal(err)
	}
	if err := testutil.WaitForPolicyUpdate(projectEditorClient, testutil.Namespace(), "create", buildapi.Resource(authorizationapi.CustomBuildResource), false); err != nil {
		t.Fatal(err)
	}
	if err := testutil.WaitForPolicyUpdate(projectEditorClient, testutil.Namespace(), "create", buildapi.Resource(authorizationapi.JenkinsPipelineBuildResource), false); err != nil {
		t.Fatal(err)
	}
}

func strategyForType(t *testing.T, strategy string) buildapi.BuildStrategy {
	buildStrategy := buildapi.BuildStrategy{}
	switch strategy {
	case "docker":
		buildStrategy.DockerStrategy = &buildapi.DockerBuildStrategy{}
	case "custom":
		buildStrategy.CustomStrategy = &buildapi.CustomBuildStrategy{}
		buildStrategy.CustomStrategy.From.Name = "builderimage:latest"
	case "source":
		buildStrategy.SourceStrategy = &buildapi.SourceBuildStrategy{}
		buildStrategy.SourceStrategy.From.Name = "builderimage:latest"
	case "jenkinspipeline":
		buildStrategy.JenkinsPipelineStrategy = &buildapi.JenkinsPipelineBuildStrategy{}
	default:
		t.Fatalf("unknown strategy: %#v", strategy)
	}
	return buildStrategy
}

func createBuild(t *testing.T, buildInterface client.BuildInterface, strategy string) (*buildapi.Build, error) {
	build := &buildapi.Build{}
	build.ObjectMeta.Labels = map[string]string{
		buildapi.BuildConfigLabel:    "mock-build-config",
		buildapi.BuildRunPolicyLabel: string(buildapi.BuildRunPolicyParallel),
	}
	build.GenerateName = strings.ToLower(string(strategy)) + "-build-"
	build.Spec.Strategy = strategyForType(t, strategy)
	build.Spec.Source.Git = &buildapi.GitBuildSource{URI: "example.org"}

	return buildInterface.Create(build)
}

func updateBuild(t *testing.T, buildInterface client.BuildInterface, build *buildapi.Build) (*buildapi.Build, error) {
	build.Labels = map[string]string{"updated": "true"}
	return buildInterface.Update(build)
}

func createBuildConfig(t *testing.T, buildConfigInterface client.BuildConfigInterface, strategy string) (*buildapi.BuildConfig, error) {
	buildConfig := &buildapi.BuildConfig{}
	buildConfig.Spec.RunPolicy = buildapi.BuildRunPolicyParallel
	buildConfig.GenerateName = strings.ToLower(string(strategy)) + "-buildconfig-"
	buildConfig.Spec.Strategy = strategyForType(t, strategy)
	buildConfig.Spec.Source.Git = &buildapi.GitBuildSource{URI: "example.org"}

	return buildConfigInterface.Create(buildConfig)
}

func cloneBuild(t *testing.T, buildInterface client.BuildInterface, build *buildapi.Build) (*buildapi.Build, error) {
	req := &buildapi.BuildRequest{}
	req.Name = build.Name
	return buildInterface.Clone(req)
}

func instantiateBuildConfig(t *testing.T, buildConfigInterface client.BuildConfigInterface, buildConfig *buildapi.BuildConfig) (*buildapi.Build, error) {
	req := &buildapi.BuildRequest{}
	req.Name = buildConfig.Name
	return buildConfigInterface.Instantiate(req)
}

func updateBuildConfig(t *testing.T, buildConfigInterface client.BuildConfigInterface, buildConfig *buildapi.BuildConfig) (*buildapi.BuildConfig, error) {
	buildConfig.Labels = map[string]string{"updated": "true"}
	return buildConfigInterface.Update(buildConfig)
}
