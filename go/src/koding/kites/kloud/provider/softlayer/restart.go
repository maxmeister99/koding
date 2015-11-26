package softlayer

import (
	"errors"
	"koding/kites/kloud/machinestate"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"golang.org/x/net/context"
)

func (m *Machine) Restart(ctx context.Context) (err error) {
	if err := m.UpdateState(machinestate.Rebooting); err != nil {
		return err
	}

	//Get the SoftLayer virtual guest service
	svc, err := m.Session.SLClient.GetSoftLayer_Virtual_Guest_Service()
	if err != nil {
		return err
	}

	ok, err := svc.RebootSoft(m.Meta.Id)
	if err != nil {
		return err
	}

	if !ok {
		m.Log.Warning("softlayer rebooting returned false instead of true")
	}

	if err := waitState(svc, m.Meta.Id, "RUNNING"); err != nil {
		return err
	}

	m.push("Checking remote machine", 90, machinestate.Starting)
	if !m.IsKlientReady() {
		return errors.New("klient is not ready")
	}

	return m.Session.DB.Run("jMachines", func(c *mgo.Collection) error {
		return c.UpdateId(
			m.Id,
			bson.M{"$set": bson.M{
				"status.state":      machinestate.Running.String(),
				"status.modifiedAt": time.Now().UTC(),
				"status.reason":     "Machine is running",
			}},
		)
	})
}
