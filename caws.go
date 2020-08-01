package main

import (
	"github.com/rivo/tview"
	"github.com/seike460/caws/awsservice"
)

// Caws CawsCore
type Caws struct {
	App   *tview.Application
	Flex  *tview.Flex
	LList *tview.List
	MList *tview.List
	RList *tview.List
}

func newCaws() Caws {

	app := tview.NewApplication()

	serviceList := tview.NewList()
	serviceList.SetTitle("- AWS Service -").SetBorder(true)

	mlist := tview.NewList()
	mlist.SetTitle("- Target -").SetBorder(true)

	rlist := tview.NewList()
	rlist.SetTitle("- Objects -").SetBorder(true)

	flex := tview.NewFlex().
		AddItem(serviceList, 20, 0, true).
		AddItem(mlist, 0, 1, false).
		AddItem(rlist, 0, 2, false)

	return Caws{
		App:   app,
		Flex:  flex,
		LList: serviceList,
		MList: mlist,
		RList: rlist,
	}
}

func setS3Service(caws Caws) {

	s := awsservice.NewS3()
	s.SetBucketsList(caws.MList, caws.LList, caws.RList, caws.App)

}
