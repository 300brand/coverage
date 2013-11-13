package validurl

import (
	"net/url"
)

var invalidDomains = []string{
	"www.auntminnie.com",
	"cdn.auntminnie.com",
	"auntminnie.com",
	"accounts.google.com",
	"www.google.com",
	"login.oracle.com",
	"apex.oracle.com",
	"www.aviationweek.com",
	"aviationweek.com",
	"www.businesswire.com",
	"businesswire.com",
	"blogs.ft.com",
	"ftalphaville.ft.com",
	"ft.com",
	"www.ft.com",
	"www.complianceweek.com",
	"complianceweek.com",
	"app.cvent.com",
	"www.cvent.com",
	"cvent.com",
	"www.gartner.com",
	"gartner.com",
	"globalgoodgroup.com",
	"items.infoedglobal.com",
	"spin.infoedglobal.com",
	"www.infoedglobal.com",
	"infoedglobal.com",
	"www.infoed.org",
	"infoed.org",
	"www.yammer.com",
	"yammer.com",
	"insidehealth.com",
	"www.insidehealth.com",
	"investor.techtarget.com",
	"www.lohud.com",
	"lohud.com",
	"www.medtechwebinars.com",
	"medtechwebinars.com",
	"www.mightydeals.com",
	"mightydeals.com",
	"news.cincinnati.com",
	"oraclemeetings.webex.com",
	"oukc.oracle.com",
	"education.oracle.com",
	"oraclecse.webex.com",
	"oracleus.activeevents.com",
	"oracletalk.webex.com",
	"oracle.6connex.com",
	"reg.accelacomm.com",
	"solutioncenters.computerworld.com",
	"search.itunes.apple.com",
	"searchcloudprovider.techtarget.com",
	"www.shrm.org",
	"shrm.org",
	"searchconsumerization.techtarget.com",
	"searchitchannel.techtarget.com",
	"us.rd.yahoo.com",
	"www.yahoo.com",
	"yahoo.com",
	"www.windowsitpro.com",
	"windowsitpro.com",
	"www-nc.nytimes.com",
	"myaccount.nytimes.com",
	"www3.ambest.com",
	"ambest.com",
	"www.ambest.com",
}

func IsValid(u *url.URL) bool {
	for _, domain := range invalidDomains {
		if u.Host == domain {
			return false
		}
	}
	return true
}
