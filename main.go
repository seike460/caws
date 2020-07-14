package main

import (
	"os"

	"github.com/rivo/tview"
	"github.com/seike460/caws/awsservice"
)

func main() {
	os.Exit(run())
}

func run() int {

	app := tview.NewApplication()

	serviceList := tview.NewList()
	serviceList.SetTitle("- AWS Service -").SetBorder(true)

	bucketsList := tview.NewList()
	bucketsList.SetTitle("- Target -").SetBorder(true)

	objectList := tview.NewList()
	objectList.SetTitle("- Objects -").SetBorder(true)

	flex := tview.NewFlex().
		AddItem(serviceList, 20, 0, true).
		AddItem(bucketsList, 0, 1, false).
		AddItem(objectList, 0, 2, false)

	serviceList.AddItem(" S3 ", " Objects Storage ", 0, func() {
		s := awsservice.NewS3()
		buckets, _ := s.ListBuckets()
		bucketsList.Clear()
		for _, val := range buckets.Buckets {
			bucketsList.AddItem(*val.Name, *buckets.Owner.DisplayName, 0, func() {
				bucketFunc(objectList, bucketsList, s, app)
			})
		}
		app.SetFocus(bucketsList)
	})
	serviceList.AddItem(" Quit ", " Press to exit ", 0, func() {
		app.Stop()
	})
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	return 0
}

func bucketFunc(objectList *tview.List, bucketsList *tview.List, s awsservice.S3, app *tview.Application) {
	objectList.Clear()
	itemText, _ := bucketsList.GetItemText(bucketsList.GetCurrentItem())
	objects, _ := s.ListObjects(itemText)
	for _, v := range objects.Contents {
		objectList.AddItem(*v.Key, v.LastModified.Format("2006/01/02 15:04:05"), 0, nil)
	}
	app.SetFocus(objectList)
}
