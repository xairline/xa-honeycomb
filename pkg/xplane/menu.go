package xplane

import (
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/menus"
)

func (s *xplaneService) menuHandler(menuRef interface{}, itemRef interface{}) {
	s.debug = !s.debug

	if s.debug {
		logging.MinLevel = logging.Debug_Level
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Checked)
	} else {
		logging.MinLevel = logging.Info_Level
		menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Unchecked)
	}

	s.Logger.Infof("menu clicked: %v", itemRef)
}
