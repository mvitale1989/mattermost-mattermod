// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package server

import (
	// "bytes"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	// "path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/route53"
	jenkins "github.com/cpanato/golang-jenkins"
	"github.com/mattermost/mattermost-mattermod/model"
	"github.com/mattermost/mattermost-server/mlog"
	// "github.com/mattermost/mattermost-load-test/ltops"
	// "github.com/mattermost/mattermost-load-test/ltparse"
	// "github.com/mattermost/mattermost-load-test/terraform"
)

// TODO FIXME
// func waitForBuildAndSetupLoadtest(pr *model.PullRequest) {
// 	repo, ok := Config.GetRepository(pr.RepoOwner, pr.RepoName)
// 	if !ok || repo.JenkinsServer == "" {
// 		LogError("Unable to set up loadtest for PR %v in %v/%v without Jenkins configured for server", pr.Number, pr.RepoOwner, pr.RepoName)
// 		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to setup loadtest")
// 		return
// 	}

// 	credentials, ok := Config.JenkinsCredentials[repo.JenkinsServer]
// 	if !ok {
// 		LogError("No Jenkins credentials for server %v required for PR %v in %v/%v", repo.JenkinsServer, pr.Number, pr.RepoOwner, pr.RepoName)
// 		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to setup loadtest")
// 		return
// 	}

// 	client := jenkins.NewJenkins(&jenkins.Auth{
// 		Username: credentials.Username,
// 		ApiToken: credentials.ApiToken,
// 	}, credentials.URL)

// 	LogInfo("Waiting for Jenkins to build to set up loadtest for PR %v in %v/%v", pr.Number, pr.RepoOwner, pr.RepoName)

// 	pr = waitForBuild(client, pr)

// 	config := &ltops.ClusterConfig{
// 		Name:                  fmt.Sprintf("pr-%v", pr.Number),
// 		AppInstanceType:       "m4.xlarge",
// 		AppInstanceCount:      4,
// 		DBInstanceType:        "db.r4.xlarge",
// 		DBInstanceCount:       4,
// 		LoadtestInstanceCount: 1,
// 	}
// 	config.WorkingDirectory = filepath.Join("./clusters/", config.Name)

// 	LogInfo("Creating terraform cluster for loadtest for PR %v in %v/%v", pr.Number, pr.RepoOwner, pr.RepoName)
// 	cluster, err := terraform.CreateCluster(config)
// 	if err != nil {
// 		LogError("Unable to setup cluster: " + err.Error())
// 		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to setup loadtest")
// 	}
// 	// Wait for the cluster to init
// 	time.Sleep(time.Minute)

// 	results := bytes.NewBuffer(nil)

// 	LogInfo("Deploying to cluster for loadtest for PR %v in %v/%v", pr.Number, pr.RepoOwner, pr.RepoName)
// 	if err := cluster.Deploy("https://releases.mattermost.com/mattermost-platform-pr/"+strconv.Itoa(pr.Number)+"/mattermost-enterprise-linux-amd64.tar.gz", "mattermod.mattermost-license"); err != nil {
// 		LogError("Unable to deploy cluster: " + err.Error())
// 		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to setup loadtest")
// 		return
// 	}
// 	if err := cluster.Loadtest("https://releases.mattermost.com/mattermost-load-test/mattermost-load-test.tar.gz"); err != nil {
// 		LogError("Unable to deploy loadtests to cluster: " + err.Error())
// 		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to setup loadtest")
// 		return
// 	}

// 	// Wait for the cluster restart after deploy
// 	time.Sleep(time.Minute)

// 	LogInfo("Runing loadtest for PR %v in %v/%v", pr.Number, pr.RepoOwner, pr.RepoName)
// 	if err := cluster.Loadtest(results); err != nil {
// 		LogError("Unable to loadtest cluster: " + err.Error())
// 		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to setup loadtest")
// 		return
// 	}
// 	LogInfo("Destroying cluster for PR %v in %v/%v", pr.Number, pr.RepoOwner, pr.RepoName)
// 	if err := cluster.Destroy(); err != nil {
// 		LogError("Unable to destroy cluster: " + err.Error())
// 		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to setup loadtest")
// 		return
// 	}

// 	githubOutput := bytes.NewBuffer(nil)
// 	cfg := ltparse.ResultsConfig{
// 		Input:     results,
// 		Output:    githubOutput,
// 		Display:   "markdown",
// 		Aggregate: false,
// 	}

// 	ltparse.ParseResults(&cfg)
// 	LogInfo("Loadtest results for PR %v in %v/%v\n%v", pr.Number, pr.RepoOwner, pr.RepoName, githubOutput.String())
// 	commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, githubOutput.String())
// }
func waitForBuildAndSetupLoadtest(pr *model.PullRequest) {
	return
}

func waitForBuildAndSetupSpinmint(pr *model.PullRequest, upgradeServer bool) {
	repo, ok := Config.GetRepository(pr.RepoOwner, pr.RepoName)
	if !ok || repo.JenkinsServer == "" {
		mlog.Error("Unable to set up spintmint for PR without Jenkins configured for server", mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
		return
	}

	credentials, ok := Config.JenkinsCredentials[repo.JenkinsServer]
	if !ok {
		mlog.Error("No Jenkins credentials for server required for PR", mlog.String("jenkins", repo.JenkinsServer), mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
		return
	}

	client := jenkins.NewJenkins(&jenkins.Auth{
		Username: credentials.Username,
		ApiToken: credentials.ApiToken,
	}, credentials.URL)

	mlog.Info("Waiting for Jenkins to build to set up spinmint for PR", mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))

	pr, errr := waitForBuild(client, pr)
	if errr == false || pr == nil {
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
		return
	}

	var instance *ec2.Instance
	if result := <-Srv.Store.Spinmint().Get(pr.Number); result.Err != nil {
		mlog.Error("Unable to get the spinmint information. Will not build the spinmint", mlog.String("pr_error", result.Err.Error()))
	} else if result.Data == nil {
		mlog.Error("No spinmint for this PR in the Database. will start a fresh one.")
		var errInstance error
		instance, errInstance = setupSpinmint(pr.Number, pr.Ref, pr.RepoName, repo, upgradeServer)
		if errInstance != nil {
			LogErrorToMattermost("Unable to set up spinmint for PR %v in %v/%v: %v", pr.Number, pr.RepoOwner, pr.RepoName, errInstance.Error())
			commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
			return
		}
		spinmint := &model.Spinmint{
			InstanceId: *instance.InstanceId,
			RepoOwner:  pr.RepoOwner,
			RepoName:   pr.RepoName,
			Number:     pr.Number,
			CreatedAt:  time.Now().UTC().Unix(),
		}
		storeSpinmintInfo(spinmint)
	} else {
		spinmint := result.Data.(*model.Spinmint)
		instance.InstanceId = aws.String(spinmint.InstanceId)
	}

	mlog.Info("Waiting for instance to come up.")
	time.Sleep(time.Minute * 2)
	publicDNS, internalIP := getIPsForInstance(*instance.InstanceId)

	if err := updateRoute53Subdomain(*instance.InstanceId, publicDNS, "CREATE"); err != nil {
		LogErrorToMattermost("Unable to set up S3 subdomain for PR %v in %v/%v with instance %v: %v", pr.Number, pr.RepoOwner, pr.RepoName, *instance.InstanceId, err.Error())
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
		return
	}

	smLink := fmt.Sprintf("%v.%v", *instance.InstanceId, Config.AWSDnsSuffix)
	if Config.SpinmintsUseHttps {
		smLink = "https://" + smLink
	} else {
		smLink = "http://" + smLink
	}

	var message string
	if upgradeServer {
		message = Config.SetupSpinmintUpgradeDoneMessage
	} else {
		message = Config.SetupSpinmintDoneMessage
	}

	message = strings.Replace(message, SPINMINT_LINK, smLink, 1)
	message = strings.Replace(message, INSTANCE_ID, INSTANCE_ID_MESSAGE+*instance.InstanceId, 1)
	message = strings.Replace(message, INTERNAL_IP, internalIP, 1)

	commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, message)
}

func waitForBuildAndSetupSpinmintExperimental(pr *model.PullRequest) {
	repo, ok := Config.GetRepository(pr.RepoOwner, pr.RepoName)
	if !ok || repo.JenkinsServer == "" {
		mlog.Error("Unable to set up spintmint for PR without Jenkins configured for server", mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
		return
	}

	credentials, ok := Config.JenkinsCredentials[repo.JenkinsServer]
	if !ok {
		mlog.Error("No Jenkins credentials for server required for PR", mlog.String("jenkins", repo.JenkinsServer), mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
		return
	}

	client := jenkins.NewJenkins(&jenkins.Auth{
		Username: credentials.Username,
		ApiToken: credentials.ApiToken,
	}, credentials.URL)

	mlog.Info("Waiting for Jenkins to build to set up spinmint for PR", mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))

	pr, errr := waitForBuild(client, pr)
	if errr == false || pr == nil {
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
		return
	}

	var installation string
	if result := <-Srv.Store.Spinmint().Get(pr.Number); result.Err != nil {
		mlog.Error("Unable to get the spinmint information. Will not build the spinmint", mlog.String("pr_error", result.Err.Error()))
	} else if result.Data == nil {
		mlog.Error("No spinmint for this PR in the Database. will start a fresh one.")
		var errInstance error
		installation, errInstance = setupSpinmintExperimental(pr)
		if errInstance != nil {
			LogErrorToMattermost("Unable to set up spinmint for PR %v in %v/%v: %v", pr.Number, pr.RepoOwner, pr.RepoName, errInstance.Error())
			commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.SetupSpinmintFailedMessage)
			return
		}
		spinmint := &model.Spinmint{
			InstanceId: installation,
			RepoOwner:  pr.RepoOwner,
			RepoName:   pr.RepoName,
			Number:     pr.Number,
			CreatedAt:  time.Now().UTC().Unix(),
		}
		storeSpinmintInfo(spinmint)
	}
}

func waitForMobileAppsBuild(pr *model.PullRequest) {
	repo, ok := Config.GetRepository(pr.RepoOwner, pr.RepoName)
	if !ok || repo.JenkinsServer == "" {
		mlog.Error("Unable to build the mobile app for PR %v in %v/%v without Jenkins configured for server", mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.BuildMobileAppFailedMessage)
		return
	}

	credentials, ok := Config.JenkinsCredentials[repo.JenkinsServer]
	if !ok {
		mlog.Error("No Jenkins credentials for server required for PR", mlog.String("jenkins", repo.JenkinsServer), mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.BuildMobileAppFailedMessage)
		return
	}

	client := jenkins.NewJenkins(&jenkins.Auth{
		Username: credentials.Username,
		ApiToken: credentials.ApiToken,
	}, credentials.URL)

	mlog.Info("Waiting for Jenkins to build to start build the mobile app for PR", mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))

	pr, errr := waitForBuild(client, pr)
	if errr == false || pr == nil {
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.BuildMobileAppFailedMessage)
		return
	}

	//Job that will build the apps for a PR
	jobName := fmt.Sprintf("mm/job/%s", repo.JobName)
	job, err := client.GetJob(jobName)
	if err != nil {
		mlog.Error("Failed to get Jenkins job", mlog.String("job", jobName), mlog.Err(err))
		return
	}

	mlog.Info("Will start the job", mlog.String("job", jobName))
	parameters := url.Values{}
	parameters.Add("PR_NUMBER", strconv.Itoa(pr.Number))
	err = client.Build(jobName, parameters)
	if err != nil {
		mlog.Error("Failed to build Jenkins job", mlog.String("job", jobName), mlog.Err(err))
		return
	}

	job.Name = jobName
	for {
		build, err := client.GetLastBuild(job)
		if err != nil {
			mlog.Error("Failed to get the build Jenkins job", mlog.String("job", jobName), mlog.Err(err))
			return
		}
		if !build.Building && build.Result == "SUCCESS" {
			mlog.Info("build mobile app for PR succeeded!", mlog.Int("build", build.Number), mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
			break
		} else if build.Result == "FAILURE" {
			mlog.Error("build has status FAILURE aborting.", mlog.Int("build", build.Number), mlog.String("result", build.Result))
			commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, Config.BuildMobileAppFailedMessage)
			return
		} else {
			mlog.Info("build is running", mlog.Int("build", build.Number), mlog.Bool("building", build.Building))
		}
		time.Sleep(60 * time.Second)
	}

	prNumberStr := fmt.Sprintf("PR-%d", pr.Number)
	msgMobile := Config.BuildMobileAppDoneMessage
	msgMobile = strings.Replace(msgMobile, "PR_NUMBER", prNumberStr, 2)
	msgMobile = strings.Replace(msgMobile, "ANDROID_APP", prNumberStr, 1)
	msgMobile = strings.Replace(msgMobile, "IOS_APP", prNumberStr, 1)
	commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, msgMobile)
	return
}

func waitForBuild(client *jenkins.Jenkins, pr *model.PullRequest) (*model.PullRequest, bool) {
	for {
		if result := <-Srv.Store.PullRequest().Get(pr.RepoOwner, pr.RepoName, pr.Number); result.Err != nil {
			mlog.Error("Unable to get updated PR while waiting for spinmint", mlog.String("build_error", result.Err.Error()))
			return nil, false
		} else {
			// Update the PR in case the build link has changed because of a new commit
			pr = result.Data.(*model.PullRequest)
		}

		if pr.RepoName == "mattermost-webapp" {
			if pr.BuildStatus == "in_progress" {
				mlog.Info("Build in CircleCI in progress will get an update check...")
				prUpdate, err := GetUpdateChecks(pr.RepoOwner, pr.RepoName, pr.Number)
				if err != nil {
					mlog.Error("Unable to get checks while waiting for spinmint", mlog.String("githubError", err.Error()))
					return nil, false
				}
				if prUpdate.BuildStatus == "in_progress" {
					mlog.Info("Build in CircleCI running will wait to conclusion")
				} else if prUpdate.BuildStatus == "completed" && prUpdate.BuildConclusion == "success" {
					mlog.Info("Build in CircleCI succeed")
					return prUpdate, true
				} else if prUpdate.BuildStatus == "completed" && prUpdate.BuildConclusion == "failure" {
					mlog.Info("Build in CircleCI failed")
					return prUpdate, false
				}
			} else if pr.BuildStatus == "completed" && pr.BuildConclusion == "success" {
				mlog.Info("Build in CircleCI succeed")
				return pr, true
			} else if pr.BuildStatus == "completed" && pr.BuildConclusion == "failure" {
				mlog.Info("Build in CircleCI failed")
				return pr, false
			} else {
				mlog.Info("Have no info about the Checks")
				return pr, false
			}

		} else {
			if pr.BuildLink != "" {
				mlog.Info("BuildLink for PR", mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName), mlog.String("buildlink", pr.BuildLink))
				// Doing this because the lib we are using does not support folders :(
				var jobNumber int64
				var jobName string

				parts := strings.Split(pr.BuildLink, "/")
				// Doing this because the lib we are using does not support folders :(
				if pr.RepoName == "mattermost-server" {
					jobNumber, _ = strconv.ParseInt(parts[len(parts)-3], 10, 32)
					jobName = parts[len(parts)-6]     //mattermost-server
					subJobName := parts[len(parts)-4] //PR-XXXX

					jobName = "mp/job/" + jobName + "/job/" + subJobName
					mlog.Info("Job name for server", mlog.String("job", jobName), mlog.String("buidlink", pr.BuildLink))
				} else if pr.RepoName == "mattermost-mobile" {
					jobNumber, _ = strconv.ParseInt(parts[len(parts)-2], 10, 32)
					jobName = parts[len(parts)-3] //mattermost-mobile
					jobName = "mm/job/" + jobName
					mlog.Info("Job name for mobile", mlog.String("job", jobName), mlog.String("buidlink", pr.BuildLink))
				} else {
					mlog.Error("Did not know this repository. Aborting.", mlog.String("repo_name", pr.RepoName))
					return pr, false
				}

				job, err := client.GetJob(jobName)
				if err != nil {
					mlog.Error("Failed to get Jenkins job", mlog.String("job", jobName), mlog.Err(err))
					return pr, false
				}

				// Doing this because the lib we are using does not support folders :(
				// This time is in the Jenkins job Name because it returns just the name
				job.Name = jobName

				build, err := client.GetBuild(job, int(jobNumber))
				if err != nil {
					LogErrorToMattermost("Failed to get build %v for PR %v in %v/%v: %v", jobNumber, pr.Number, pr.RepoOwner, pr.RepoName, err)
					return pr, false
				}

				if !build.Building && build.Result == "SUCCESS" {
					mlog.Info("build for PR succeeded!", mlog.Int64("jobnumber", jobNumber), mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))
					return pr, true
				} else if build.Result == "FAILURE" {
					mlog.Error("build has status FAILURE. Aborting.", mlog.Int("build", build.Number), mlog.String("build_error", build.Result))
					return pr, false
				} else {
					mlog.Info("build is running", mlog.Int("build", build.Number), mlog.Bool("building", build.Building))
				}
			} else {
				mlog.Error("Unable to find build link for PR", mlog.Int("pr", pr.Number))
				return pr, false
			}
		}

		mlog.Info("Sleeping a bit....Will re-check the Jenkins Build or CircleCI checks...")
		time.Sleep(30 * time.Second)
	}
}

// Returns instance ID of instance created
func setupSpinmint(prNumber int, prRef, prRepo string, repo *Repository, upgrade bool) (*ec2.Instance, error) {
	mlog.Info("Setting up spinmint for PR", mlog.Int("pr", prNumber))

	svc := ec2.New(session.New(), Config.GetAwsConfig())

	var setupScript string
	if upgrade {
		setupScript = repo.InstanceSetupUpgradeScript
	} else {
		setupScript = repo.InstanceSetupScript
	}

	data, err := ioutil.ReadFile(path.Join("config", setupScript))
	if err != nil {
		return nil, err
	}
	sdata := string(data)
	sdata = strings.Replace(sdata, "BUILD_NUMBER", strconv.Itoa(prNumber), -1)
	sdata = strings.Replace(sdata, "BRANCH_NAME", prRef, -1)
	bsdata := []byte(sdata)
	sdata = base64.StdEncoding.EncodeToString(bsdata)

	var one int64 = 1
	params := &ec2.RunInstancesInput{
		ImageId:          &Config.AWSImageId,
		MaxCount:         &one,
		MinCount:         &one,
		InstanceType:     &Config.AWSInstanceType,
		UserData:         &sdata,
		SecurityGroupIds: []*string{&Config.AWSSecurityGroup},
		SubnetId:         &Config.AWSSubNetId,
	}

	resp, err := svc.RunInstances(params)
	if err != nil {
		return nil, err
	}

	// Add tags to the created instance
	time.Sleep(time.Second * 10)
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{resp.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("Spinmint-" + prRef),
			},
			{
				Key:   aws.String("Created"),
				Value: aws.String(time.Now().Format("2006-01-02/15:04:05")),
			},
			{
				Key:   aws.String("PRNumber"),
				Value: aws.String("PR-" + strconv.Itoa(prNumber)),
			},
			{
				Key:   aws.String("RepoName"),
				Value: aws.String(prRepo),
			},
		},
	})
	if errtag != nil {
		mlog.Error("Could not create tags for instance", mlog.String("instance", *resp.Instances[0].InstanceId), mlog.String("tag_error", errtag.Error()))
	}

	return resp.Instances[0], nil
}

func setupSpinmintExperimental(pr *model.PullRequest) (string, error) {
	mlog.Info("Setting up spinmint experimental for PR", mlog.Int("pr", pr.Number))
	url := fmt.Sprintf("%s/api/clusters", Config.ProvisionerServer)
	mlog.Info("Provisioner Server ", mlog.String("Server", url))
	client := &http.Client{}

	// Get cluster list
	mlog.Info("Provisioner Server getting clusters")
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		mlog.Error("Error making the post request to check the k8s cluster", mlog.Err(err))
		return "", err
	}
	defer resp.Body.Close()

	var createClusterList []Cluster
	err = json.NewDecoder(resp.Body).Decode(&createClusterList)
	if err != nil && err != io.EOF {
		mlog.Error("Error decoding", mlog.Err(err))
		return "", err
	}

	clusterCount := 0
	for _, cluster := range createClusterList {
		if cluster.State == "stable" {
			clusterCount++
			mlog.Info("Provisioner Server counting", mlog.Int("clusterCount", clusterCount))
		}
	}

	// Get cluster list
	mlog.Info("Provisioner Server getting installations")
	urlInstallation := fmt.Sprintf("%s/api/installations", Config.ProvisionerServer)
	req, err = http.NewRequest("GET", urlInstallation, nil)
	resp, err = client.Do(req)
	if err != nil {
		mlog.Error("Error making the post request to check the installations", mlog.Err(err))
		return "", err
	}
	defer resp.Body.Close()

	var createInstallationList []Installation
	err = json.NewDecoder(resp.Body).Decode(&createClusterList)
	if err != nil && err != io.EOF {
		mlog.Error("Error decoding", mlog.Err(err))
		return "", err
	}

	installationCount := 0
	for _, installation := range createInstallationList {
		if installation.State == "stable" {
			installationCount++
			mlog.Info("Provisioner Server counting MM Installations", mlog.Int("installationCount", installationCount))
		}
	}

	mlog.Info("will check if need create a cluster")
	if clusterCount == 0 || installationCount/clusterCount > 5 {
		mlog.Info("Need to spin a new k8s cluster", mlog.Int("clusterCount", installationCount), mlog.Int("clusterCount", clusterCount))
		payloadCluster := fmt.Sprint("{\n\"size\":\"SizeAlef1000\"\n}")
		var jsonStr = []byte(payloadCluster)
		reqCluster, errCluster := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		reqCluster.Header.Set("Content-Type", "application/json")

		clientCluster := &http.Client{}
		respCluster, errCluster := clientCluster.Do(reqCluster)
		if errCluster != nil {
			mlog.Error("Error making the post request to create the k8s cluster", mlog.Err(errCluster))
			return "", err
		}
		defer respCluster.Body.Close()

		var createClusterRequest Cluster
		errCluster = json.NewDecoder(respCluster.Body).Decode(&createClusterRequest)
		if errCluster != nil && errCluster != io.EOF {
			mlog.Error("Error decoding", mlog.Err(errCluster))
			return "", err
		}
		mlog.Info("Provisioner Server - cluster request", mlog.String("ClusterID", createClusterRequest.ID))

		for {
			url2 := fmt.Sprintf("%s/api/cluster/%s", Config.ProvisionerServer, createClusterRequest.ID)
			req2, err2 := http.NewRequest("GET", url2, nil)
			client2 := &http.Client{}
			resp2, err2 := client2.Do(req2)
			if err2 != nil {
				mlog.Error("Error making the post request to create the k8s cluster", mlog.Err(err2))
				return "", err2
			}
			defer resp2.Body.Close()
			var clusterRequest Cluster
			err2 = json.NewDecoder(resp2.Body).Decode(&clusterRequest)
			if err2 != nil && err2 != io.EOF {
				mlog.Error("Error decoding", mlog.Err(err2))
			}
			if clusterRequest.State == "stable" {
				commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Kubernetes cluster created. Now will deploy Mattermost... Hang on!")
				break
			} else if clusterRequest.State == "creation-failed" {
				commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to create the k8s cluster.")
				return "", fmt.Errorf("error creating k8s cluster")
			}
			mlog.Info("Provisioner Server - cluster request creating... sleep", mlog.String("ClusterID", createClusterRequest.ID), mlog.String("State", createClusterRequest.State))
			time.Sleep(20 * time.Second)
		}

	} else {
		mlog.Info("not needed to create a cluster")
		commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "We don't need a new Kubernetes cluster, will reuse an existing one. Requesting to deploy Mattermost.")
	}

	mlog.Info("Provisioner Server - Installation request")
	payload := fmt.Sprintf("{\n\"ownerId\":\"PR-%d\",\n\"dns\": \"pr-%d.test.cloud.mattermost.com\",\n\"version\": \"PR-%d\"\n\"affinity\":\"shared\"}", pr.Number, pr.Number, pr.Number)
	var mmStr = []byte(payload)
	url = fmt.Sprintf("%s/api/installations", Config.ProvisionerServer)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(mmStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		mlog.Error("Error making the post request to create the mm cluster", mlog.Err(err))
		return "", err
	}
	defer resp.Body.Close()

	var createInstallationRequest Installation
	err = json.NewDecoder(resp.Body).Decode(&createInstallationRequest)
	if err != nil && err != io.EOF {
		mlog.Error("Error decoding installation request", mlog.Err(err))
		return "", err
	}
	mlog.Info("Provisioner Server - installation request", mlog.String("InstallationID", createInstallationRequest.ID))

	for {
		url = fmt.Sprintf("%s/api/installation/%s", Config.ProvisionerServer, createInstallationRequest.ID)
		req, err := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			mlog.Error("Error making the post request to create the mm installation", mlog.Err(err))
			return "", err
		}
		defer resp.Body.Close()
		var installationRequest Installation
		err = json.NewDecoder(resp.Body).Decode(&installationRequest)
		if err != nil && err != io.EOF {
			mlog.Error("Error decoding installation", mlog.Err(err))
		}
		if installationRequest.State == "stable" {
			msg := fmt.Sprintf("Mattermost test server created! :tada:\nAccess here: https://pr-%d.test.cloud.mattermost.com", pr.Number)
			commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, msg)
			break
		} else if installationRequest.State == "creation-failed" {
			commentOnIssue(pr.RepoOwner, pr.RepoName, pr.Number, "Failed to create mattermost installation.")
			return "", fmt.Errorf("error creating mattermost installation")
		}
		mlog.Info("Provisioner Server - installation request creating... sleep", mlog.String("InstallationID", installationRequest.ID), mlog.String("State", installationRequest.State))
		time.Sleep(10 * time.Second)
	}

	return createInstallationRequest.ID, nil
}

func destroySpinmint(pr *model.PullRequest, instanceID string) {
	mlog.Info("Destroying spinmint for PR", mlog.String("instance", instanceID), mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))

	svc := ec2.New(session.New(), Config.GetAwsConfig())

	params := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{
			&instanceID,
		},
	}

	_, err := svc.TerminateInstances(params)
	if err != nil {
		mlog.Error("Error terminating instances", mlog.Err(err))
		return
	}

	// Remove route53 entry
	err = updateRoute53Subdomain(instanceID, "", "DELETE")
	if err != nil {
		mlog.Error("Error removing the Route53 entry", mlog.Err(err))
		return
	}

	// Remove from the local db
	removeSpinmintInfo(instanceID)
}

func destroySpinmintExperimental(pr *model.PullRequest, instanceClusterID string) {
	mlog.Info("Destroying spinmint experimental for PR", mlog.String("instance", instanceClusterID), mlog.Int("pr", pr.Number), mlog.String("repo_owner", pr.RepoOwner), mlog.String("repo_name", pr.RepoName))

	url := fmt.Sprintf("%s/api/installation/%s", Config.ProvisionerServer, instanceClusterID)
	req, err := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		mlog.Error("Error deleting the installation", mlog.Err(err))
	}
	defer resp.Body.Close()

	// Remove from the local db
	removeSpinmintInfo(instanceClusterID)
}

func getIPsForInstance(instance string) (publicIP string, privateIP string) {
	svc := ec2.New(session.New(), Config.GetAwsConfig())
	params := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			&instance,
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		mlog.Error("Problem getting instance ip", mlog.Err(err))
		return "", ""
	}

	return *resp.Reservations[0].Instances[0].PublicIpAddress, *resp.Reservations[0].Instances[0].PrivateIpAddress
}

func updateRoute53Subdomain(name, target, action string) error {
	svc := route53.New(session.New(), Config.GetAwsConfig())
	domainName := fmt.Sprintf("%v.%v", name, Config.AWSDnsSuffix)

	targetServer := target
	if target == "" && action == "DELETE" {
		targetServer, _ = getIPsForInstance(name)
	}

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String(action),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(domainName),
						TTL:  aws.Int64(30),
						Type: aws.String("A"),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(targetServer),
							},
						},
					},
				},
			},
		},
		HostedZoneId: &Config.AWSHostedZoneId,
	}

	_, err := svc.ChangeResourceRecordSets(params)
	if err != nil {
		return err
	}

	return nil
}

func CheckSpinmintLifeTime() {
	mlog.Info("Checking Spinmint lifetime...")
	spinmints := []*model.Spinmint{}
	if result := <-Srv.Store.Spinmint().List(); result.Err != nil {
		mlog.Error("Unable to get updated PR while waiting for spinmint", mlog.String("spinmint_error", result.Err.Error()))
	} else {
		spinmints = result.Data.([]*model.Spinmint)
	}

	for _, spinmint := range spinmints {
		mlog.Info("Check if need destroy spinmint for PR", mlog.String("instance", spinmint.InstanceId), mlog.Int("spinmint", spinmint.Number), mlog.String("repo_owner", spinmint.RepoOwner), mlog.String("repo_name", spinmint.RepoName))
		spinmintCreated := time.Unix(spinmint.CreatedAt, 0)
		duration := time.Since(spinmintCreated)
		if int(duration.Hours()) > Config.SpinmintExpirationHour {
			mlog.Info("Will destroy spinmint for PR", mlog.String("instance", spinmint.InstanceId), mlog.Int("spinmint", spinmint.Number), mlog.String("repo_owner", spinmint.RepoOwner), mlog.String("repo_name", spinmint.RepoName))
			pr := &model.PullRequest{
				RepoOwner: spinmint.RepoOwner,
				RepoName:  spinmint.RepoName,
				Number:    spinmint.Number,
			}
			go destroySpinmint(pr, spinmint.InstanceId)
			removeSpinmintInfo(spinmint.InstanceId)
			commentOnIssue(spinmint.RepoOwner, spinmint.RepoName, spinmint.Number, Config.DestroyedExpirationSpinmintMessage)
		}
	}

	mlog.Info("Done checking Spinmint lifetime.")
}

func storeSpinmintInfo(spinmint *model.Spinmint) {
	if result := <-Srv.Store.Spinmint().Save(spinmint); result.Err != nil {
		mlog.Error(result.Err.Error())
	}
}

func removeSpinmintInfo(instanceId string) {
	if result := <-Srv.Store.Spinmint().Delete(instanceId); result.Err != nil {
		mlog.Error(result.Err.Error())
	}
}

type Cluster struct {
	ID                  string
	Provider            string
	Provisioner         string
	ProviderMetadata    []byte `json:",omitempty"`
	ProvisionerMetadata []byte `json:",omitempty"`
	AllowInstallations  bool
	Size                string
	State               string
	CreateAt            int64
	DeleteAt            int64
	LockAcquiredBy      *string
	LockAcquiredAt      int64
}

type Installation struct {
	ID             string
	OwnerID        string
	Version        string
	DNS            string
	Affinity       string
	GroupID        *string
	State          string
	CreateAt       int64
	DeleteAt       int64
	LockAcquiredBy *string
	LockAcquiredAt int64
}
