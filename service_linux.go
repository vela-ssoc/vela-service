package service

import (
	"context"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/mitchellh/mapstructure"
	cond "github.com/vela-ssoc/vela-cond"
)

var (
	Conn     *dbus.Conn
	UnitList []unitFetcher
)

func Connection() *dbus.Conn {
	if Conn != nil {
		return Conn
	}

	var err error
	Conn, err = dbus.NewWithContext(context.Background())
	if err != nil {
		xEnv.Errorf("connect to dbus error: %v", err)
	}

	UnitList = []unitFetcher{listUnitsByPatternWrapper, listUnitsFilteredWrapper, listUnitsWrapper}
	return Conn
}

func (su *Summary) Units() []dbus.UnitStatus {
	var state []string
	var pat []string
	conn := Connection()
	if conn == nil {
		return nil
	}

	for _, unit := range UnitList {
		us, err := unit(conn, state, pat)
		if err != nil {
			xEnv.Debugf("linux dbus load unit by %v error %v", unit, err)
			continue
		}

		xEnv.Debugf("linux dbus load unit success by %v", unit)
		return us
	}

	xEnv.Error("dbus unit error by all methods")
	return nil
}

func (su *Summary) collect(cnd *cond.Cond) {
	if Conn = Connection(); Conn == nil {
		xEnv.Errorf("no conn to dbus")
		return
	}

	us := su.Units()
	if us == nil {
		return
	}

	for _, unit := range us {
		if unit.LoadState == "not-found" {
			xEnv.Debugf("linux dbus unit %s state not found", unit.Name)
		}
		sv := u2s(unit)
		if cnd.Match(sv) {
			su.append(sv)
		}
	}
}

func u2s(unit dbus.UnitStatus) *Service {
	s := &Service{
		Name:        unit.Name,
		StartType:   unit.JobType,
		ExecPath:    string(unit.Path),
		DisplayName: unit.Name,
		Description: unit.Description,
		State:       unit.ActiveState,
	}

	conn := Connection()
	if conn == nil {
		xEnv.Errorf("linux dbus got nil dbus conn")
		return s
	}

	rpp, err := conn.GetAllPropertiesContext(context.Background(), unit.Name)
	if err != nil {
		xEnv.Errorf("linux dbus got unit prop fail %v", err)
		return s
	}

	pps := Properties{}
	err = mapstructure.Decode(rpp, &pps)
	if err != nil {
		xEnv.Errorf("linux dbus unit read prop  decode fail %v", err)
		return s
	}

	if pps.ExecMainPID > 0 {
		s.ExitCode = uint32(pps.ExecMainCode)
		s.Pid = uint32(pps.ExecMainPID)
	}

	return s
}
