package illumioapi_test

import (
	"testing"

	"github.com/brian1917/illumioapi"
)

func TestWorkloads(t *testing.T) {

	// Clear previous test runs
	allWLs, api, _ := illumioapi.GetAllWorkloads(pce)
	if api.StatusCode != 200 {
		t.Errorf("GetAllWorkloads is returning a status code of %d", api.StatusCode)
	}

	for _, workload := range allWLs {
		if workload.Name == "go_test_workload" || workload.Name == "updated_go_test_workload" {
			illumioapi.DeleteHref(pce, workload.Href)
		}
	}

	// Create a new workload
	newWL, api, _ := illumioapi.CreateWorkload(pce, illumioapi.Workload{Name: "go_test_workload"})
	if api.StatusCode != 201 {
		t.Errorf("CreateWorkload is returning a status code of %d", api.StatusCode)
	}

	// Get all workloads
	allWLs, api, _ = illumioapi.GetAllWorkloads(pce)
	if api.StatusCode != 200 {
		t.Errorf("GetAllWorkloads is returning a status code of %d", api.StatusCode)
	}
	if len(allWLs) < 1 {
		t.Errorf("GetAllWorkloads is not returning a populated array")
	}

	// Update a workload
	newWL.Name = "updated_go_test_workload"
	api, err := illumioapi.UpdateWorkload(pce, newWL)
	if api.StatusCode != 204 {
		t.Errorf("UpdateWorkload is returning a status code of %d", api.StatusCode)
	}

	// Delete the updated workload
	deleteAPI, _ := illumioapi.DeleteHref(pce, newWL.Href)
	if api.StatusCode != 204 {
		t.Errorf("Delete workload is returning a status code of %d", deleteAPI.StatusCode)
	}

	// Bulk upload workloads
	wls := []illumioapi.Workload{illumioapi.Workload{Name: "Bulk 1"}, illumioapi.Workload{Name: "Bulk 2"}}
	createWLsAPI, err := illumioapi.BulkWorkload(pce, wls, "create")
	for _, a := range createWLsAPI {
		if a.StatusCode != 200 {
			t.Errorf("Create bulk workloads is returning a status code of %d", a.StatusCode)
		}
	}

	// Bulk update workloads
	wls, api, err = illumioapi.GetAllWorkloads(pce)
	updatedWLs := []illumioapi.Workload{}
	for _, wl := range wls {
		wl.Name = wl.Name + "-Updated"
		updatedWLs = append(updatedWLs, wl)
	}
	updateAPIResps, err := illumioapi.BulkWorkload(pce, updatedWLs, "update")
	if err != nil {
		t.Errorf("Update bulk workloads error - %s", err)
	}
	for _, r := range updateAPIResps {
		if r.StatusCode != 200 {
			t.Errorf("Update bulk workloads is returning a status code of %d", r.StatusCode)
		}
	}

}
