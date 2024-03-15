package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	wx "github.com/matwachich/fynex-widgets"
	"github.com/ncruces/zenity"
)

type GlobalConfig struct {
	widget.BaseWidget

	reset, load, save *widget.Button

	el  *wx.NumEntry // event length
	elt *widget.Select
	ib  *wx.NumEntry // include before
	ibt *widget.Select
	ia  *wx.NumEntry // include after
	iat *widget.Select

	t  *wx.NumEntry // threshold
	k  *wx.NumEntry // kernel size
	df *wx.NumEntry // downscale factor
	fs *wx.NumEntry // frame skip

	bb *widget.Check // bounding box
	tm *widget.Check // timecode
	fm *widget.Check // frame metrics

	r   *widget.Check  // show region editor
	rf  *widget.Entry  // region file
	rfb *widget.Button // region file browse
}

func NewGlobalConfig(app *App) *GlobalConfig {
	w := &GlobalConfig{}
	w.ExtendBaseWidget(w)

	w.reset = &widget.Button{
		Text: "Reset",
		Icon: theme.MediaReplayIcon(),
		OnTapped: func() {

		},
	}
	w.save = &widget.Button{
		Text: "Save",
		Icon: theme.DocumentSaveIcon(),
		OnTapped: func() {

		},
	}
	w.load = &widget.Button{
		Text: "Load",
		Icon: theme.FolderOpenIcon(),
		OnTapped: func() {

		},
	}

	// event lenght
	w.el = wx.NewNumEntry()
	w.el.Float = true
	w.elt = widget.NewSelect([]string{"Seconds", "Frames"}, nil)

	// include before
	w.ib = wx.NewNumEntry()
	w.ib.Float = true
	w.ibt = widget.NewSelect([]string{"Seconds", "Frames"}, nil)

	// include after
	w.ia = wx.NewNumEntry()
	w.ia.Float = true
	w.iat = widget.NewSelect([]string{"Seconds", "Frames"}, nil)

	// threshold
	w.t = wx.NewNumEntry()
	w.t.Float = true

	// kernel size
	w.k = wx.NewNumEntry()

	// downscale factor
	w.df = wx.NewNumEntry()

	// frame skip
	w.fs = wx.NewNumEntry()

	// overlays
	w.bb = widget.NewCheck("Bounding box", nil)
	w.tm = widget.NewCheck("Timecode", nil)
	w.fm = widget.NewCheck("Frame metrics", nil)

	// region // TODO support multiple regions
	w.r = widget.NewCheck("Show editor", func(b bool) {
		if b {
			w.rf.Disable()
			w.rfb.Disable()
		} else {
			w.rf.Enable()
			w.rfb.Enable()
		}
	})

	w.rf = widget.NewEntry()
	w.rfb = &widget.Button{
		Icon: theme.FolderOpenIcon(),
		OnTapped: func() {
			if file, err := zenity.SelectFile(zenity.Title("Select region file"), zenity.Modal()); err == nil {
				w.rf.SetText(file)
			}
		},
	}

	return w
}

func (w *GlobalConfig) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(nil)
}
