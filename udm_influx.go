package unifi

import (
	"time"

	influx "github.com/influxdata/influxdb1-client/v2"
)

// Points generates Unifi Gateway datapoints for InfluxDB.
// These points can be passed directly to influx.
func (u *UDM) Points() ([]*influx.Point, error) {
	return u.PointsAt(time.Now())
}

// PointsAt generates Unifi Gateway datapoints for InfluxDB.
// These points can be passed directly to influx.
// This is just like Points(), but specify when points were created.
func (u *UDM) PointsAt(now time.Time) ([]*influx.Point, error) {
	if u.Stat.gw == nil {
		// Disabled devices lack stats.
		u.Stat.gw = &gw{}
	}
	if u.Stat.sw == nil {
		// Disabled devices lack stats.
		u.Stat.sw = &sw{}
	}
	tags := map[string]string{
		"id":                     u.ID,
		"mac":                    u.Mac,
		"device_oid":             u.Stat.gw.Oid,
		"site_id":                u.SiteID,
		"site_name":              u.SiteName,
		"adopted":                u.Adopted.Txt,
		"name":                   u.Name,
		"cfgversion":             u.Cfgversion,
		"config_network_ip":      u.ConfigNetwork.IP,
		"config_network_type":    u.ConfigNetwork.Type,
		"connect_request_ip":     u.ConnectRequestIP,
		"connect_request_port":   u.ConnectRequestPort,
		"device_id":              u.DeviceID,
		"guest_token":            u.GuestToken,
		"inform_ip":              u.InformIP,
		"known_cfgversion":       u.KnownCfgversion,
		"model":                  u.Model,
		"serial":                 u.Serial,
		"type":                   u.Type,
		"usg_caps":               u.UsgCaps.Txt,
		"speedtest-status-saved": u.SpeedtestStatusSaved.Txt,
		"wan1_up":                u.Wan1.Up.Txt,
		"wan2_up":                u.Wan2.Up.Txt,
	}
	fields := map[string]interface{}{
		"ip":                             u.IP,
		"bytes":                          u.Bytes.Val,
		"last_seen":                      u.LastSeen.Val,
		"license_state":                  u.LicenseState,
		"fw_caps":                        u.FwCaps.Val,
		"guest-num_sta":                  u.GuestNumSta.Val,
		"rx_bytes":                       u.RxBytes.Val,
		"tx_bytes":                       u.TxBytes.Val,
		"uptime":                         u.Uptime.Val,
		"state":                          u.State.Val,
		"user-num_sta":                   u.UserNumSta.Val,
		"version":                        u.Version,
		"num_desktop":                    u.NumDesktop.Val,
		"num_handheld":                   u.NumHandheld.Val,
		"num_mobile":                     u.NumMobile.Val,
		"speedtest-status_latency":       u.SpeedtestStatus.Latency.Val,
		"speedtest-status_rundate":       u.SpeedtestStatus.Rundate.Val,
		"speedtest-status_runtime":       u.SpeedtestStatus.Runtime.Val,
		"speedtest-status_download":      u.SpeedtestStatus.StatusDownload.Val,
		"speedtest-status_ping":          u.SpeedtestStatus.StatusPing.Val,
		"speedtest-status_summary":       u.SpeedtestStatus.StatusSummary.Val,
		"speedtest-status_upload":        u.SpeedtestStatus.StatusUpload.Val,
		"speedtest-status_xput_download": u.SpeedtestStatus.XputDownload.Val,
		"speedtest-status_xput_upload":   u.SpeedtestStatus.XputUpload.Val,
		"config_network_wan_type":        u.ConfigNetwork.Type,
		"wan1_bytes-r":                   u.Wan1.BytesR.Val,
		"wan1_enable":                    u.Wan1.Enable.Val,
		"wan1_full_duplex":               u.Wan1.FullDuplex.Val,
		"wan1_gateway":                   u.Wan1.Gateway,
		"wan1_ifname":                    u.Wan1.Ifname,
		"wan1_ip":                        u.Wan1.IP,
		"wan1_mac":                       u.Wan1.Mac,
		"wan1_max_speed":                 u.Wan1.MaxSpeed.Val,
		"wan1_name":                      u.Wan1.Name,
		"wan1_netmask":                   u.Wan1.Netmask,
		"wan1_rx_bytes":                  u.Wan1.RxBytes.Val,
		"wan1_rx_bytes-r":                u.Wan1.RxBytesR.Val,
		"wan1_rx_dropped":                u.Wan1.RxDropped.Val,
		"wan1_rx_errors":                 u.Wan1.RxErrors.Val,
		"wan1_rx_multicast":              u.Wan1.RxMulticast.Val,
		"wan1_rx_packets":                u.Wan1.RxPackets.Val,
		"wan1_type":                      u.Wan1.Type,
		"wan1_speed":                     u.Wan1.Speed.Val,
		"wan1_up":                        u.Wan1.Up.Val,
		"wan1_tx_bytes":                  u.Wan1.TxBytes.Val,
		"wan1_tx_bytes-r":                u.Wan1.TxBytesR.Val,
		"wan1_tx_dropped":                u.Wan1.TxDropped.Val,
		"wan1_tx_errors":                 u.Wan1.TxErrors.Val,
		"wan1_tx_packets":                u.Wan1.TxPackets.Val,
		"wan2_bytes-r":                   u.Wan2.BytesR.Val,
		"wan2_enable":                    u.Wan2.Enable.Val,
		"wan2_full_duplex":               u.Wan2.FullDuplex.Val,
		"wan2_gateway":                   u.Wan2.Gateway,
		"wan2_ifname":                    u.Wan2.Ifname,
		"wan2_ip":                        u.Wan2.IP,
		"wan2_mac":                       u.Wan2.Mac,
		"wan2_max_speed":                 u.Wan2.MaxSpeed.Val,
		"wan2_name":                      u.Wan2.Name,
		"wan2_netmask":                   u.Wan2.Netmask,
		"wan2_rx_bytes":                  u.Wan2.RxBytes.Val,
		"wan2_rx_bytes-r":                u.Wan2.RxBytesR.Val,
		"wan2_rx_dropped":                u.Wan2.RxDropped.Val,
		"wan2_rx_errors":                 u.Wan2.RxErrors.Val,
		"wan2_rx_multicast":              u.Wan2.RxMulticast.Val,
		"wan2_rx_packets":                u.Wan2.RxPackets.Val,
		"wan2_type":                      u.Wan2.Type,
		"wan2_speed":                     u.Wan2.Speed.Val,
		"wan2_up":                        u.Wan2.Up.Val,
		"wan2_tx_bytes":                  u.Wan2.TxBytes.Val,
		"wan2_tx_bytes-r":                u.Wan2.TxBytesR.Val,
		"wan2_tx_dropped":                u.Wan2.TxDropped.Val,
		"wan2_tx_errors":                 u.Wan2.TxErrors.Val,
		"wan2_tx_packets":                u.Wan2.TxPackets.Val,
		"loadavg_1":                      u.SysStats.Loadavg1.Val,
		"loadavg_5":                      u.SysStats.Loadavg5.Val,
		"loadavg_15":                     u.SysStats.Loadavg15.Val,
		"mem_used":                       u.SysStats.MemUsed.Val,
		"mem_buffer":                     u.SysStats.MemBuffer.Val,
		"mem_total":                      u.SysStats.MemTotal.Val,
		"cpu":                            u.SystemStats.CPU.Val,
		"mem":                            u.SystemStats.Mem.Val,
		"system_uptime":                  u.SystemStats.Uptime.Val,
		"gw":                             u.Stat.Gw,
		"lan-rx_bytes":                   u.Stat.LanRxBytes.Val,
		"lan-rx_packets":                 u.Stat.LanRxPackets.Val,
		"lan-tx_bytes":                   u.Stat.LanTxBytes.Val,
		"lan-tx_packets":                 u.Stat.LanTxPackets.Val,
		"wan-rx_bytes":                   u.Stat.WanRxBytes.Val,
		"wan-rx_dropped":                 u.Stat.WanRxDropped.Val,
		"wan-rx_packets":                 u.Stat.WanRxPackets.Val,
		"wan-tx_bytes":                   u.Stat.WanTxBytes.Val,
		"wan-tx_packets":                 u.Stat.WanTxPackets.Val,
		"uplink_name":                    u.Uplink.Name,
		"uplink_latency":                 u.Uplink.Latency.Val,
		"uplink_speed":                   u.Uplink.Speed.Val,
		"uplink_num_ports":               u.Uplink.NumPort.Val,
		"uplink_max_speed":               u.Uplink.MaxSpeed.Val,
	}
	pt, err := influx.NewPoint("usg", tags, fields, now)
	if err != nil {
		return nil, err
	}
	points := []*influx.Point{pt}
	tags = map[string]string{
		"id":                     u.ID,
		"mac":                    u.Mac,
		"device_oid":             u.Stat.sw.Oid,
		"site_id":                u.SiteID,
		"site_name":              u.SiteName,
		"name":                   u.Name,
		"adopted":                u.Adopted.Txt,
		"cfgversion":             u.Cfgversion,
		"config_network_ip":      u.ConfigNetwork.IP,
		"config_network_type":    u.ConfigNetwork.Type,
		"device_id":              u.DeviceID,
		"inform_ip":              u.InformIP,
		"known_cfgversion":       u.KnownCfgversion,
		"locating":               u.Locating.Txt,
		"model":                  u.Model,
		"serial":                 u.Serial,
		"type":                   u.Type,
		"dot1x_portctrl_enabled": u.Dot1XPortctrlEnabled.Txt,
		"flowctrl_enabled":       u.FlowctrlEnabled.Txt,
		"has_fan":                u.HasFan.Txt,
		"has_temperature":        u.HasTemperature.Txt,
		"jumboframe_enabled":     u.JumboframeEnabled.Txt,
		"stp_priority":           u.StpPriority,
		"stp_version":            u.StpVersion,
	}
	fields = map[string]interface{}{
		"fw_caps":             u.FwCaps.Val,
		"guest-num_sta":       u.GuestNumSta.Val,
		"ip":                  u.IP,
		"bytes":               u.Bytes.Val,
		"fan_level":           float64(0),
		"general_temperature": float64(0),
		"last_seen":           u.LastSeen.Val,
		"license_state":       u.LicenseState,
		"overheating":         u.Overheating.Val,
		"rx_bytes":            u.RxBytes.Val,
		"tx_bytes":            u.TxBytes.Val,
		"uptime":              u.Uptime.Val,
		"state":               u.State.Val,
		"user-num_sta":        u.UserNumSta.Val,
		"version":             u.Version,
		"loadavg_1":           u.SysStats.Loadavg1.Val,
		"loadavg_5":           u.SysStats.Loadavg5.Val,
		"loadavg_15":          u.SysStats.Loadavg15.Val,
		"mem_buffer":          u.SysStats.MemBuffer.Val,
		"mem_used":            u.SysStats.MemUsed.Val,
		"mem_total":           u.SysStats.MemTotal.Val,
		"cpu":                 u.SystemStats.CPU.Val,
		"mem":                 u.SystemStats.Mem.Val,
		"system_uptime":       u.SystemStats.Uptime.Val,
		"stat_bytes":          u.Stat.Bytes.Val,
		"stat_rx_bytes":       u.Stat.RxBytes.Val,
		"stat_rx_crypts":      u.Stat.RxCrypts.Val,
		"stat_rx_dropped":     u.Stat.RxDropped.Val,
		"stat_rx_errors":      u.Stat.RxErrors.Val,
		"stat_rx_frags":       u.Stat.RxFrags.Val,
		"stat_rx_packets":     u.Stat.TxPackets.Val,
		"stat_tx_bytes":       u.Stat.TxBytes.Val,
		"stat_tx_dropped":     u.Stat.TxDropped.Val,
		"stat_tx_errors":      u.Stat.TxErrors.Val,
		"stat_tx_packets":     u.Stat.TxPackets.Val,
		"stat_tx_retries":     u.Stat.TxRetries.Val,
		"uplink_depth":        "0",
	}
	pt, err = influx.NewPoint("usw", tags, fields, now)
	if err != nil {
		return nil, err
	}
	points = append(points, pt)

	for _, p := range u.NetworkTable {
		tags := map[string]string{
			"device_name":               u.Name,
			"device_id":                 u.ID,
			"device_mac":                u.Mac,
			"site_name":                 u.SiteName,
			"up":                        p.Up.Txt,
			"dhcpd_dns_enabled":         p.DhcpdDNSEnabled.Txt,
			"dhcpd_enabled":             p.DhcpdEnabled.Txt,
			"dhcpd_time_offset_enabled": p.DhcpdTimeOffsetEnabled.Txt,
			"dhcp_relay_enabledy":       p.DhcpRelayEnabled.Txt,
			"dhcpd_gateway_enabled":     p.DhcpdGatewayEnabled.Txt,
			"enabled":                   p.Enabled.Txt,
			"vlan_enabled":              p.VlanEnabled.Txt,
			"attr_no_delete":            p.AttrNoDelete.Txt,
			"is_guest":                  p.IsGuest.Txt,
			"is_nat":                    p.IsNat.Txt,
			"networkgroup":              p.Networkgroup,
			"site_id":                   p.SiteID,
		}
		fields := map[string]interface{}{
			"domain_name":         p.DomainName,
			"dhcpd_start":         p.DhcpdStart,
			"dhcpd_stop":          p.DhcpdStop,
			"ip":                  p.IP,
			"ip_subnet":           p.IPSubnet,
			"mac":                 p.Mac,
			"name":                p.Name,
			"num_sta":             p.NumSta.Val,
			"purpose":             p.Purpose,
			"rx_bytes":            p.RxBytes.Val,
			"rx_packets":          p.RxPackets.Val,
			"tx_bytes":            p.TxBytes.Val,
			"tx_packets":          p.TxPackets.Val,
			"ipv6_interface_type": p.Ipv6InterfaceType,
			"attr_hidden_id":      p.AttrHiddenID,
		}
		pt, err = influx.NewPoint("usg_networks", tags, fields, now)
		if err != nil {
			return points, err
		}
		points = append(points, pt)
	}

	for _, p := range u.PortTable {
		tags := map[string]string{
			"site_id":       u.SiteID,
			"site_name":     u.SiteName,
			"device_name":   u.Name,
			"name":          p.Name,
			"enable":        p.Enable.Txt,
			"is_uplink":     p.IsUplink.Txt,
			"up":            p.Up.Txt,
			"portconf_id":   p.PortconfID,
			"dot1x_mode":    p.Dot1XMode,
			"dot1x_status":  p.Dot1XStatus,
			"stp_state":     p.StpState,
			"sfp_found":     p.SfpFound.Txt,
			"op_mode":       p.OpMode,
			"poe_mode":      p.PoeMode,
			"port_poe":      p.PortPoe.Txt,
			"port_idx":      p.PortIdx.Txt,
			"port_id":       u.Name + " Port " + p.PortIdx.Txt,
			"poe_enable":    p.PoeEnable.Txt,
			"flowctrl_rx":   p.FlowctrlRx.Txt,
			"flowctrl_tx":   p.FlowctrlTx.Txt,
			"autoneg":       p.Autoneg.Txt,
			"full_duplex":   p.FullDuplex.Txt,
			"jumbo":         p.Jumbo.Txt,
			"masked":        p.Masked.Txt,
			"poe_good":      p.PoeGood.Txt,
			"media":         p.Media,
			"poe_class":     p.PoeClass,
			"poe_caps":      p.PoeCaps.Txt,
			"aggregated_by": p.AggregatedBy.Txt,
		}
		fields := map[string]interface{}{
			"dbytes_r":     p.BytesR.Val,
			"rx_broadcast": p.RxBroadcast.Val,
			"rx_bytes":     p.RxBytes.Val,
			"rx_bytes-r":   p.RxBytesR.Val,
			"rx_dropped":   p.RxDropped.Val,
			"rx_errors":    p.RxErrors.Val,
			"rx_multicast": p.RxMulticast.Val,
			"rx_packets":   p.RxPackets.Val,
			"speed":        p.Speed.Val,
			"stp_pathcost": p.StpPathcost.Val,
			"tx_broadcast": p.TxBroadcast.Val,
			"tx_bytes":     p.TxBytes.Val,
			"tx_bytes-r":   p.TxBytesR.Val,
			"tx_dropped":   p.TxDropped.Val,
			"tx_errors":    p.TxErrors.Val,
			"tx_multicast": p.TxMulticast.Val,
			"tx_packets":   p.TxPackets.Val,
			"poe_current":  p.PoeCurrent.Val,
			"poe_power":    p.PoePower.Val,
			"poe_voltage":  p.PoeVoltage.Val,
			"full_duplex":  p.FullDuplex.Val,
		}
		pt, err = influx.NewPoint("usw_ports", tags, fields, now)
		if err != nil {
			return points, err
		}
		points = append(points, pt)
	}
	return points, nil
}
