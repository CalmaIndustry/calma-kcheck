package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"regexp"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type HelmRelease struct {
	Name       string `json:"name"`
	Chart      string `json:"chart"`
	AppVersion string `json:"app_version"`
}

type HelmRepo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type HelmSearch struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	AppVersion string `json:"app_version"`
}

func init() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		log.Fatalf("Cannot find home directory")
	}

	// Load the kubeconfig from the specified path
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	// Get the current context
	context := config.CurrentContext
	fmt.Printf("Current context: %s\n", context)
}

func Helm() {
	cmd := exec.Command("helm", "list", "--all-namespaces", "--output", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Unlucky %v", err)
		return
	}

	var releases []HelmRelease
	if err := json.Unmarshal(output, &releases); err != nil {
		log.Fatalf("Error parsing JSON: %v\n", err)
	}

	fmt.Println("Helm Releases:")
	for _, release := range releases {
		fmt.Printf("Name: %s, App Version: %s, Chart: %s\n",
			release.Name, release.AppVersion, release.Chart)

		re := regexp.MustCompile(`^(.+)-\d+\.\d+\.\d+$`)
		match := re.FindStringSubmatch(release.Chart)
		c := exec.Command("helm", "search", "repo", match[1], "-l", "--output", "json")
		test, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf("unlucky %v\n", err)
		}
		var versions []HelmSearch
		if err := json.Unmarshal(test, &versions); err != nil {
			log.Fatalf("Error parsing JSON: %v\n", err)
		}
		for i, version := range versions {
			fmt.Printf("Index is %v, version is %v\n", i, version)

			x := fmt.Sprintf("%s-%s", release.Name, version.Version)
			if x == release.Chart {
				fmt.Println("We're good")
			} else {
				fmt.Printf("Need to upgrade chef")
			}
			break

		}

	}
	helmrepo()
}

func helmrepo() {
	cmd := exec.Command("helm", "repo", "list", "--output", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("unlucky %v\n", err)
	}
	var repolist []HelmRepo
	if err := json.Unmarshal(output, &repolist); err != nil {
		log.Fatalf("Error parsing JSON: %v\n", err)
	}
	fmt.Println("Helm repo:")
	for _, repolist := range repolist {
		fmt.Printf("Name: %s, Url: %s\n",
			repolist.Name, repolist.Url)
	}

}

func Version(version string) {
	fmt.Printf("Hello from cmd! Version is : %v\n", version)
}
