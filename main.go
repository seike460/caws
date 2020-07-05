package main

import (
	"github.com/rivo/tview"
	"github.com/seike460/caws/awsservice"
)

func main() {

	app := tview.NewApplication()
	s := awsservice.NewS3()
	buckets, err := s.ListBuckets()

	if err != nil {
		panic(err)
	}

	serviceList := tview.NewList()
	serviceList.SetTitle("- AWS Service -").SetBorder(true)

	serviceList.AddItem("S3", "Objects Storage", 's', func() {
		app.Stop()
	})

	serviceList.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	bucketsList := tview.NewList()
	bucketsList.SetTitle("- Buckets -").SetBorder(true)

	objectList := tview.NewList()
	objectList.SetTitle("- Objects -").SetBorder(true)

	for _, val := range buckets.Buckets {
		bucketsList.AddItem(*val.Name, *buckets.Owner.DisplayName, 0, nil)
		objects, err := s.ListObjects(*val.Name)
		if err != nil {
			break
			//panic(err)
			//continue
		}
		for _, v := range objects.Contents {
			objectList.AddItem(*v.Key, v.LastModified.Format("2006/1/2 3:04:05 pm"), 0, nil)
		}

	}
	flex := tview.NewFlex().
		AddItem(serviceList, 0, 1, false).
		AddItem(bucketsList, 0, 2, false).
		AddItem(objectList, 0, 4, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
