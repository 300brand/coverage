package validurl

import (
	"github.com/300brand/logger"
	"net/url"
	"strings"
)

var invalidDomains = []string{
	"http://www.auntminnie.com",
	"https://accounts.google.com",
	"https://login.oracle.com",
	"https://apex.oracle.com",
	"https://www.aviationweek.com",
	"http://www.businesswire.com",
	"http://blogs.ft.com",
	"http://www.complianceweek.com",
	"http://www.cvent.com",
	"http://ftalphaville.ft.com",
	"http://www.ft.com",
	"http://www.gartner.com",
	"http://globalgoodgroup.com",
	"http://infoedglobal.com",
	"http://insidehealth.com",
	"http://investor.techtarget.com",
	"http://www.lohud.com",
	"http://www.medtechwebinars.com",
	"http://www.mightydeals.com",
	"http://news.cincinnati.com/needlogin",
	"https://oraclemeetings.webex.com",
	"http://oukc.oracle.com",
	"https://oraclecse.webex.com",
	"https://oracleus.activeevents.com",
	"https://oracletalk.webex.com",
	"https://oracle.6connex.com",
	"http://reg.accelacomm.com",
	"http://solutioncenters.computerworld.com",
	"http://search.itunes.apple.com",
	"https://search.itunes.apple.com",
	"http://searchcloudprovider.techtarget.com",
	"http://www.shrm.org",
	"http://searchconsumerization.techtarget.com",
	"http://searchitchannel.techtarget.com",
	"http://us.rd.yahoo.com",
	"http://windowsitpro.com",
	"http://www-nc.nytimes.com",
	"http://www3.ambest.com",
}

func IsValid(u *net.URL) bool {
	for _, domain := range invalidDomains {
		if strings.Contains(u.String(), domain) {
			return false
		}
	}
	return true
}
