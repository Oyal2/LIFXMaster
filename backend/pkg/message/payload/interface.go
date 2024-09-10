package payload

import (
	"fmt"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/state"
)

type PayloadCoder interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

func HandleResponsePayload(responseType ResponseType, payload []byte) (PayloadCoder, error) {
	switch responseType {
	case Acknowledgement:
		fmt.Println("Handle Acknowledgement")
	case StateService:
		s := &state.Service{}
		return s, s.Decode(payload)
	case StateHostFirmware:
		hf := &state.HostFirmware{}
		return hf, hf.Decode(payload)
	case StateWifiInfo:
		wi := &state.WifiInfo{}
		return wi, wi.Decode(payload)
	case StateWifiFirmware:
		wf := &state.WifiFirmware{}
		return wf, wf.Decode(payload)
	case StatePower:
		p := &state.Power{}
		return p, p.Decode(payload)
	case StateLabel:
		l := &state.Label{}
		return l, l.Decode(payload)
	case StateVersion:
		v := &state.Version{}
		return v, v.Decode(payload)
	case StateInfo:
		i := &state.Info{}
		return i, i.Decode(payload)
	case StateLocation:
		l := &state.Location{}
		return l, l.Decode(payload)
	case StateGroup:
		g := &state.Group{}
		return g, g.Decode(payload)
	case EchoResponse:
		er := &state.EchoResponse{}
		return er, er.Decode(payload)
	case StateUnhandled:
		u := &state.Unhandled{}
		return u, u.Decode(payload)
	case LightState:
		l := &state.Light{}
		return l, l.Decode(payload)
	case StateLightPower:
		p := &state.Power{}
		return p, p.Decode(payload)
	case StateInfrared:
		i := &state.Infrared{}
		return i, i.Decode(payload)
	case StateHevCycle:
		hc := &state.HEVCycle{}
		return hc, hc.Decode(payload)
	case StateHevCycleConfiguration:
		hcc := &state.HEVCycleConfig{}
		return hcc, hcc.Decode(payload)
	case StateLastHevCycleResult:
		lhcr := &state.LastHEVCycleResult{}
		return lhcr, lhcr.Decode(payload)
	case StateZone:
		z := &state.Zone{}
		return z, z.Decode(payload)
	case StateMultiZone:
		mz := &state.Zone{}
		return mz, mz.Decode(payload)
	case StateMultiZoneEffect:
		mze := &state.MultiZoneEffect{}
		return mze, mze.Decode(payload)
	case StateExtendedColorZones:
		ecz := &state.ExtendedColorZone{}
		return ecz, ecz.Decode(payload)
	case StateRPower:
		rp := &state.RPower{}
		return rp, rp.Decode(payload)
	case StateDeviceChain:
		dc := &state.DeviceChain{}
		return dc, dc.Decode(payload)
	case State64:
		t := &state.Tile64{}
		return t, t.Decode(payload)
	case StateTileEffect:
		te := &state.TileEffect{}
		return te, te.Decode(payload)
	case SensorStateAmbientLight:
		sal := &state.TileEffect{}
		return sal, sal.Decode(payload)
	default:
		fmt.Println("Unknown ResponseType:", responseType)
	}

	return nil, nil
}
