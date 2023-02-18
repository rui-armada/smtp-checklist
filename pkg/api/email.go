package api

import (
	"checklist/pkg/broker"
	"fmt"
	"log"
	"net/http"
	"time"
)

func sendEmail(w http.ResponseWriter, r *http.Request) {
	to := []string{
		"rui.armada@jumia.com",
	}
	now := time.Now()
	zabbix := r.FormValue("zabbix")
	haproxy := r.FormValue("haproxy")
	tabularjobs := r.FormValue("tabularjobs")
	lastcell := r.FormValue("lastcell")
	lastkobraphilips := r.FormValue("lastkobraphilips")
	sessioncpu := r.FormValue("sessioncpu")
	grafana := r.FormValue("grafana")
	jira := r.FormValue("jira")
	dell := r.FormValue("dell")
	body := fmt.Sprintf(`
		<style>
		table {
			border-collapse: collapse;
			margin: 18px 0;
			font-size: 0.9em;
			font-family: sans-serif;
			min-width: 400px;
			box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
			background-color: #009879;
			color: #ffffff;
			text-align: left;
		}
		</style>
		<table>
			<tr>
			<th style="border: 1px solid black;">Zabbix</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Haproxy</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Tabular Jobs</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Last refresh on Cell</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Last refresh on Kobra and Philips</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Session CPU time ms</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Grafana tabular</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Jira tickets</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
			<tr>
			<th style="border: 1px solid black;">Dell Storage Jobs</th>
			<td style="border: 1px solid black;">%s</td>
			</tr>
		</table>
  	`, zabbix, haproxy, tabularjobs, lastcell, lastkobraphilips, sessioncpu, grafana, jira, dell)

	err := broker.SendMail("10.90.100.64:25", "morningcheck@jumia.com", fmt.Sprintf("Morning Check %v", now), body, to)
	if err != nil {
		log.Printf("email broker failed: %s", err)
	}
}
