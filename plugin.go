package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	// Repo information.
	Repo struct {
		Owner string
		Name  string
	}

	// Build information.
	Build struct {
		Event   string
		Number  int
		Commit  string
		Message string
		Branch  string
		Author  string
		Status  string
		Link    string
	}

	// Config for the plugin.
	Config struct {
		URL         string
		Token       string
		Namespace   string
		Controller  []string
		Exclude     []string
		UpdateImage string
		User        string
		Message     string
	}

	// Plugin values.
	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
	}
)

// Exec executes the plugin.
func (p *Plugin) Exec() error {
	if p.Config.URL == "" {
		return fmt.Errorf("Error: base URL of the flux controller must be added")
	}

	if len(p.Config.Controller) == 0 {
		return fmt.Errorf("Error: list of controllers to release at least one")
	}

	if p.Config.User == "" {
		p.Config.User = p.Build.Author
	}

	if p.Config.Message == "" {
		p.Config.Message = p.Build.Message
	}

	args := []string{"--url " + p.Config.URL}
	if p.Config.Token != "" {
		args = append(args, "--token "+p.Config.Token)
	}

	if p.Config.Namespace != "" {
		args = append(args, "--namespace "+p.Config.Namespace)
	}

	args = append(args, "--controller ", strings.Join(p.Config.Controller, ","))

	if len(p.Config.Exclude) > 0 {
		args = append(args, "--exclude ", strings.Join(p.Config.Exclude, ","))
	}

	if p.Config.UpdateImage != "" {
		args = append(args, "--update-image "+p.Config.UpdateImage)
	}

	if p.Config.User != "" {
		args = append(args, "--user "+p.Config.User)
	}
	if p.Config.Message != "" {
		args = append(args, "--message "+p.Config.Message)
	}

	cmd := exec.Command("fluxctl", strings.Join(args, " "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("================================================")
	fmt.Println("Successfully deploy new images for kubernetes cluster.")
	fmt.Println("================================================")

	return nil
}
