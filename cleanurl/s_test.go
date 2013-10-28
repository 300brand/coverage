package cleanurl

import "testing"

func Test_searchnetworking_techtarget_com(t *testing.T) {
	urls := map[string]string{
		"http://searchnetworking.techtarget.com/ehandbook/Assessing-your-options-for-increasing-wireless-LAN-capacity?asrc=RSS_BP_BLATWIRELESS": "http://searchnetworking.techtarget.com/ehandbook/Assessing-your-options-for-increasing-wireless-LAN-capacity",
		"http://searchnetworking.techtarget.com/ehandbook/Converged-infrastructure-The-future-of-network-efficiency?asrc=RSS_BP_BLATTELECOMM": "http://searchnetworking.techtarget.com/ehandbook/Converged-infrastructure-The-future-of-network-efficiency",
		"http://searchnetworking.techtarget.com/ehandbook/Converged-infrastructure-The-future-of-network-efficiency?asrc=RSS_BP_KABPETHERNET": "http://searchnetworking.techtarget.com/ehandbook/Converged-infrastructure-The-future-of-network-efficiency",
		"http://searchnetworking.techtarget.com/ehandbook/Converged-infrastructure-The-future-of-network-efficiency?asrc=RSS_BP_KABPINTERNETWORKING": "http://searchnetworking.techtarget.com/ehandbook/Converged-infrastructure-The-future-of-network-efficiency",
	}
	testURLs(t, urls)
}
func Test_www_sportsnetwork_com(t *testing.T) {
	urls := map[string]string{
		"http://www.sportsnetwork.com/merge/tsnform.aspx?c=orlandosentinel&page=nascar%2Fstat%2Fupdateresults.aspx%3Fsportcode%3DBF%2Cgameid%3D": "http://www.sportsnetwork.com/merge/tsnform.aspx?c=orlandosentinel&page=nascar%2Fstat%2Fupdateresults.aspx%3Fsportcode%3DBF%2Cgameid%3D",
	}
	testURLs(t, urls)
}
func Test_www_securitymagazine_com(t *testing.T) {
	urls := map[string]string{
		"http://www.securitymagazine.com/articles/84171-how-to-mitigate-the-risks-of-atm-skimming?v=preview": "http://www.securitymagazine.com/articles/84171-how-to-mitigate-the-risks-of-atm-skimming?v=preview",
		"http://www.securitymagazine.com/articles/84395-steps-to-better-mail-screening-and-ricin-detection?v=preview": "http://www.securitymagazine.com/articles/84395-steps-to-better-mail-screening-and-ricin-detection?v=preview",
		"http://www.securitymagazine.com/articles/84443-how-gtri-practices-could-protect-schools-from-violence?v=preview": "http://www.securitymagazine.com/articles/84443-how-gtri-practices-could-protect-schools-from-violence?v=preview",
		"http://www.securitymagazine.com/articles/84511-how-low-retention-costs-enterprises?v=preview": "http://www.securitymagazine.com/articles/84511-how-low-retention-costs-enterprises?v=preview",
		"http://www.securitymagazine.com/articles/84570-access-control-helps-decrease-crime-near-blue-green-diamond-condos?v=preview": "http://www.securitymagazine.com/articles/84570-access-control-helps-decrease-crime-near-blue-green-diamond-condos?v=preview",
		"http://www.securitymagazine.com/articles/84577-how-the-latest-insurance-trends-help-in-disaster-planning?v=preview": "http://www.securitymagazine.com/articles/84577-how-the-latest-insurance-trends-help-in-disaster-planning?v=preview",
		"http://www.securitymagazine.com/articles/84348-how-bostons-hotels-maintained-security-after-the-marathon-bombings?v=preview": "http://www.securitymagazine.com/articles/84348-how-bostons-hotels-maintained-security-after-the-marathon-bombings?v=preview",
		"http://www.securitymagazine.com/articles/84351-how-to-converge-psim-and-piam-for-total-situational-awareness?v=preview": "http://www.securitymagazine.com/articles/84351-how-to-converge-psim-and-piam-for-total-situational-awareness?v=preview",
		"http://www.securitymagazine.com/articles/84444-how-to-sort-through-big-data-for-threat-detection?v=preview": "http://www.securitymagazine.com/articles/84444-how-to-sort-through-big-data-for-threat-detection?v=preview",
		"http://www.securitymagazine.com/articles/84480-notifying-the-masses-in-k-12?v=preview": "http://www.securitymagazine.com/articles/84480-notifying-the-masses-in-k-12?v=preview",
		"http://www.securitymagazine.com/articles/84512-why-ethical-hacking-is-the-new-face-of-cyber-security?v=preview": "http://www.securitymagazine.com/articles/84512-why-ethical-hacking-is-the-new-face-of-cyber-security?v=preview",
		"http://www.securitymagazine.com/articles/84626-how-technology-intelligence-fuels-security-strategy?v=preview": "http://www.securitymagazine.com/articles/84626-how-technology-intelligence-fuels-security-strategy?v=preview",
		"http://www.securitymagazine.com/articles/84627-major-uses-for-hospital-visitor-management-systems?v=preview": "http://www.securitymagazine.com/articles/84627-major-uses-for-hospital-visitor-management-systems?v=preview",
		"http://www.securitymagazine.com/articles/84375-the-security-leader-the-business-leader?v=preview": "http://www.securitymagazine.com/articles/84375-the-security-leader-the-business-leader?v=preview",
		"http://www.securitymagazine.com/articles/84669-points-to-securing-the-global-enterprise?v=preview": "http://www.securitymagazine.com/articles/84669-points-to-securing-the-global-enterprise?v=preview",
		"http://www.securitymagazine.com/articles/84670-technology-intelligence-how-to-leverage-incident-management-data?v=preview": "http://www.securitymagazine.com/articles/84670-technology-intelligence-how-to-leverage-incident-management-data?v=preview",
		"http://www.securitymagazine.com/articles/83738-managing-disaster-relief-notification-after-hurricane-sandy?v=preview": "http://www.securitymagazine.com/articles/83738-managing-disaster-relief-notification-after-hurricane-sandy?v=preview",
		"http://www.securitymagazine.com/articles/84265-vendor-system-integrations-accelerate-in-the-security-industry?v=preview": "http://www.securitymagazine.com/articles/84265-vendor-system-integrations-accelerate-in-the-security-industry?v=preview",
		"http://www.securitymagazine.com/articles/84421-how-business-intelligence-becomes-an-eco-research-tool?v=preview": "http://www.securitymagazine.com/articles/84421-how-business-intelligence-becomes-an-eco-research-tool?v=preview",
		"http://www.securitymagazine.com/articles/84477-the-risk-is-real-the-timing-is-critical?v=preview": "http://www.securitymagazine.com/articles/84477-the-risk-is-real-the-timing-is-critical?v=preview",
		"http://www.securitymagazine.com/articles/84571-are-you-a-boss-driving-a-culture-of-safety-and-security-at-bp-westlake?v=preview": "http://www.securitymagazine.com/articles/84571-are-you-a-boss-driving-a-culture-of-safety-and-security-at-bp-westlake?v=preview",
		"http://www.securitymagazine.com/articles/84668-top-5-pitfalls-to-avoid-in-byod-security?v=preview": "http://www.securitymagazine.com/articles/84668-top-5-pitfalls-to-avoid-in-byod-security?v=preview",
		"http://www.securitymagazine.com/articles/84771-keys-to-creating-a-byod-program?v=preview": "http://www.securitymagazine.com/articles/84771-keys-to-creating-a-byod-program?v=preview",
		"http://www.securitymagazine.com/articles/84782-the-nsa-cyber-espionage-hackers-and-more?v=preview": "http://www.securitymagazine.com/articles/84782-the-nsa-cyber-espionage-hackers-and-more?v=preview",
		"http://www.securitymagazine.com/articles/84221-how-long-term-partnerships-benefit-end-users-and-integrators?v=preview": "http://www.securitymagazine.com/articles/84221-how-long-term-partnerships-benefit-end-users-and-integrators?v=preview",
		"http://www.securitymagazine.com/articles/84396-how-outsourcing-can-increase-badging-efficiency-and-reduce-cost?v=preview": "http://www.securitymagazine.com/articles/84396-how-outsourcing-can-increase-badging-efficiency-and-reduce-cost?v=preview",
		"http://www.securitymagazine.com/articles/84397-why-global-supply-chains-need-global-shared-security-standards?v=preview": "http://www.securitymagazine.com/articles/84397-why-global-supply-chains-need-global-shared-security-standards?v=preview",
		"http://www.securitymagazine.com/articles/84465-th-century-courthouse-goes-digital-with-cloud-based-access-control?v=preview": "http://www.securitymagazine.com/articles/84465-th-century-courthouse-goes-digital-with-cloud-based-access-control?v=preview",
		"http://www.securitymagazine.com/articles/84725-how-tour-management-expands-security-officers-role?v=preview": "http://www.securitymagazine.com/articles/84725-how-tour-management-expands-security-officers-role?v=preview",
		"http://www.securitymagazine.com/articles/83740-four-risks-in-manual-identity-enrollment?v=preview": "http://www.securitymagazine.com/articles/83740-four-risks-in-manual-identity-enrollment?v=preview",
		"http://www.securitymagazine.com/articles/84472-trading-up?v=preview": "http://www.securitymagazine.com/articles/84472-trading-up?v=preview",
		"http://www.securitymagazine.com/articles/84551-one-third-of-event-goers-fear-for-their-safety?v=preview": "http://www.securitymagazine.com/articles/84551-one-third-of-event-goers-fear-for-their-safety?v=preview",
		"http://www.securitymagazine.com/articles/84576-how-four-young-professionals-chose-the-security-profession?v=preview": "http://www.securitymagazine.com/articles/84576-how-four-young-professionals-chose-the-security-profession?v=preview",
		"http://www.securitymagazine.com/articles/84687-the-most-influential-people-in-security-2013?v=preview": "http://www.securitymagazine.com/articles/84687-the-most-influential-people-in-security-2013?v=preview",
		"http://www.securitymagazine.com/articles/84693-how-small-enterprises-security-officer-investments-pay-off?v=preview": "http://www.securitymagazine.com/articles/84693-how-small-enterprises-security-officer-investments-pay-off?v=preview",
		"http://www.securitymagazine.com/articles/84697-business-not-tourism-how-the-willis-tower-handles-visitor-management?v=preview": "http://www.securitymagazine.com/articles/84697-business-not-tourism-how-the-willis-tower-handles-visitor-management?v=preview",
		"http://www.securitymagazine.com/articles/84783-how-will-big-data-change-security?v=preview": "http://www.securitymagazine.com/articles/84783-how-will-big-data-change-security?v=preview",
		"http://www.securitymagazine.com/articles/84787-high-tech-goes-personal-identity-management?v=preview": "http://www.securitymagazine.com/articles/84787-high-tech-goes-personal-identity-management?v=preview",
		"http://www.securitymagazine.com/articles/84796-building-perimeter-security-to-increase-service-safety?v=preview": "http://www.securitymagazine.com/articles/84796-building-perimeter-security-to-increase-service-safety?v=preview",
		"http://www.securitymagazine.com/articles/84220-using-integrated-security-solutions-to-combat-theft?v=preview": "http://www.securitymagazine.com/articles/84220-using-integrated-security-solutions-to-combat-theft?v=preview",
		"http://www.securitymagazine.com/articles/84476-why-two-factor-authentication-is-a-statistical-necessity?v=preview": "http://www.securitymagazine.com/articles/84476-why-two-factor-authentication-is-a-statistical-necessity?v=preview",
		"http://www.securitymagazine.com/articles/84574-asis-2013-product-preview?v=preview": "http://www.securitymagazine.com/articles/84574-asis-2013-product-preview?v=preview",
		"http://www.securitymagazine.com/articles/84770-key-factors-for-effective-event-security-checkpoints?v=preview": "http://www.securitymagazine.com/articles/84770-key-factors-for-effective-event-security-checkpoints?v=preview",
		"http://www.securitymagazine.com/articles/84802-its-not-always-difficult-to-uncover-the-reason-behind-workplace-violence?v=preview": "http://www.securitymagazine.com/articles/84802-its-not-always-difficult-to-uncover-the-reason-behind-workplace-violence?v=preview",
		"http://www.securitymagazine.com/articles/84349-how-to-detect-and-thwart-recent-fraud-trends?v=preview": "http://www.securitymagazine.com/articles/84349-how-to-detect-and-thwart-recent-fraud-trends?v=preview",
		"http://www.securitymagazine.com/articles/84350-securing-mobile-access-in-the-age-of-consumerized-it?v=preview": "http://www.securitymagazine.com/articles/84350-securing-mobile-access-in-the-age-of-consumerized-it?v=preview",
		"http://www.securitymagazine.com/articles/84552-how-biometrics-help-manage-amusement-park-access?v=preview": "http://www.securitymagazine.com/articles/84552-how-biometrics-help-manage-amusement-park-access?v=preview",
		"http://www.securitymagazine.com/articles/84573-public-private-partnerships-that-push-beyond-information-sharing?v=preview": "http://www.securitymagazine.com/articles/84573-public-private-partnerships-that-push-beyond-information-sharing?v=preview",
		"http://www.securitymagazine.com/articles/84674-why-csos-lose-their-influence-and-get-fired?v=preview": "http://www.securitymagazine.com/articles/84674-why-csos-lose-their-influence-and-get-fired?v=preview",
		"http://www.securitymagazine.com/articles/84797-reducing-the-shrink?v=preview": "http://www.securitymagazine.com/articles/84797-reducing-the-shrink?v=preview",
	}
	testURLs(t, urls)
}
func Test_searchcio_techtarget_com(t *testing.T) {
	urls := map[string]string{
		"http://searchcio.techtarget.com/ezine/CIO-Decisions/Architecting-the-agile-enterprise?asrc=RSS_BP_BLATTELECOMM": "http://searchcio.techtarget.com/ezine/CIO-Decisions/Architecting-the-agile-enterprise",
		"http://searchcio.techtarget.com/ezine/CIO-Decisions/Architecting-the-agile-enterprise?asrc=RSS_BP_KABPINTERNETWORKING": "http://searchcio.techtarget.com/ezine/CIO-Decisions/Architecting-the-agile-enterprise",
	}
	testURLs(t, urls)
}
func Test_www_sap_com(t *testing.T) {
	urls := map[string]string{
		"http://www.sap.com/asset/index.epx?+goback=.gde_4419409_member_216995331&id=f0043fa1-5dd9-443c-9e86-8595c829adda": "http://www.sap.com/asset/index.epx?+goback=.gde_4419409_member_216995331&id=f0043fa1-5dd9-443c-9e86-8595c829adda",
	}
	testURLs(t, urls)
}
func Test_searchsecurity_techtarget_com(t *testing.T) {
	urls := map[string]string{
		"http://searchsecurity.techtarget.com/ezine/Information-Security-magazine/Unlock-new-pathways-to-network-security-architecture?asrc=RSS_BP_KABPSECURITY": "http://searchsecurity.techtarget.com/ezine/Information-Security-magazine/Unlock-new-pathways-to-network-security-architecture",
	}
	testURLs(t, urls)
}
func Test_solutioncenters_computerworld_com(t *testing.T) {
	urls := map[string]string{
		"http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11270.html?source=00767350094468CTWM0WT071O06__ctwlib%3FSOURCE%3D00767350094468CTWM0WT071O06": "http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11270.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11284.html?source=00767350094470CTW6V0J78XO5J__ctwlib%3FSOURCE%3D00767350094470CTW6V0J78XO5J": "http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11284.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11249.html?source=00767350094473CTW2ZEXZ61DF0__ctwlib%3FSOURCE%3D00767350094473CTW2ZEXZ61DF0": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11249.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11285.html?source=00767350094471CTWQBWQOY9T4P__ctwlib%3FSOURCE%3D00767350094471CTWQBWQOY9T4P": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11285.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11286.html?source=00767350094474CTWATAGMBKJM2__ctwlib%3FSOURCE%3D00767350094474CTWATAGMBKJM2": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11286.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11288.html?source=00767350094476CTWP6466W35QL__ctwlib%3FSOURCE%3D00767350094476CTWP6466W35QL": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11288.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11290.html?source=00767350094478CTW5L0K1H7E0G__ctwlib%3FSOURCE%3D00767350094478CTW5L0K1H7E0G": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11290.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11291.html?source=00767350094480CTWR9KC8ZY8LM__ctwlib%3FSOURCE%3D00767350094480CTWR9KC8ZY8LM": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11291.html",
		"http://solutioncenters.computerworld.com/citrix_xenmobile/registration/11811.html?source=00843190101349CTWCKGRAPGL0Y__ctwlib%3FSOURCE%3D00843190101349CTWCKGRAPGL0Y": "http://solutioncenters.computerworld.com/citrix_xenmobile/registration/11811.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11293.html?source=00767350094472CTW5LR0DAXULQ__ctwlib%3FSOURCE%3D00767350094472CTW5LR0DAXULQ": "http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11293.html",
		"http://solutioncenters.computerworld.com/hp_enterprise_mobile_app_solution_center/registration/11548.html?source=00678410100742CTWNVAEFFJQID__ctwlib%3FSOURCE%3D00678410100742CTWNVAEFFJQID": "http://solutioncenters.computerworld.com/hp_enterprise_mobile_app_solution_center/registration/11548.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11283.html?source=00767350094469CTWUWLL91HB31__ctwlib%3FSOURCE%3D00767350094469CTWUWLL91HB31": "http://solutioncenters.computerworld.com/internet_security_resource_center/downloads/11283.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11289.html?source=00767350094477CTWPYP6PT4JYN__ctwlib%3FSOURCE%3D00767350094477CTWPYP6PT4JYN": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11289.html",
		"http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11287.html?source=00767350094475CTWBOC4RIVEUS__ctwlib%3FSOURCE%3D00767350094475CTWBOC4RIVEUS": "http://solutioncenters.computerworld.com/internet_security_resource_center/registration/11287.html",
	}
	testURLs(t, urls)
}
func Test_www_sifma_org(t *testing.T) {
	urls := map[string]string{
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589934694": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589934694",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589935163": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589935163",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589942450": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589942450",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589944338": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589944338",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589944676": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589944676",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945342": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945342",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945398": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945398",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945473": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945473",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945525": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945525",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589942556": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589942556",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589942558": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589942558",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589944555": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589944555",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945503": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945503",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589939671": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589939671",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589944430": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589944430",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589944483": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589944483",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945273": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945273",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945468": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945468",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945522": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945522",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589935162": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589935162",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589935364": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589935364",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589944333": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589944333",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589944400": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589944400",
		"http://www.sifma.org//workarea/downloadasset.aspx?id=8589945506": "http://www.sifma.org//workarea/downloadasset.aspx?id=8589945506",
	}
	testURLs(t, urls)
}
func Test_www_sdmmag_com(t *testing.T) {
	urls := map[string]string{
		"http://www.sdmmag.com/articles/89421-glasaa-meeting-showcases-the-future?v=preview": "http://www.sdmmag.com/articles/89421-glasaa-meeting-showcases-the-future?v=preview",
		"http://www.sdmmag.com/articles/89430-tennessee-combats-menacing-hogs-usinghd-cameras?v=preview": "http://www.sdmmag.com/articles/89430-tennessee-combats-menacing-hogs-usinghd-cameras?v=preview",
		"http://www.sdmmag.com/articles/89689-wireless-residential-access-control-closer-to-mainstream-adoption?v=preview": "http://www.sdmmag.com/articles/89689-wireless-residential-access-control-closer-to-mainstream-adoption?v=preview",
		"http://www.sdmmag.com/articles/89690-factors-in-video-transmission?v=preview": "http://www.sdmmag.com/articles/89690-factors-in-video-transmission?v=preview",
		"http://www.sdmmag.com/articles/89367-the-right-tools-for-the-job?v=preview": "http://www.sdmmag.com/articles/89367-the-right-tools-for-the-job?v=preview",
		"http://www.sdmmag.com/articles/89405-city-surveillance-market-to-more-than-double-by-2017?v=preview": "http://www.sdmmag.com/articles/89405-city-surveillance-market-to-more-than-double-by-2017?v=preview",
		"http://www.sdmmag.com/articles/89466-purpose-not-product?v=preview": "http://www.sdmmag.com/articles/89466-purpose-not-product?v=preview",
		"http://www.sdmmag.com/articles/89687-digital-access-control-credentials?v=preview": "http://www.sdmmag.com/articles/89687-digital-access-control-credentials?v=preview",
		"http://www.sdmmag.com/articles/89688-integration-of-co-smoke-detectors?v=preview": "http://www.sdmmag.com/articles/89688-integration-of-co-smoke-detectors?v=preview",
	}
	testURLs(t, urls)
}
func Test_www_securitybistro_com(t *testing.T) {
	urls := map[string]string{
		"http://www.securitybistro.com/?p=7990": "http://www.securitybistro.com/?p=7990",
		"http://www.securitybistro.com/?p=8018": "http://www.securitybistro.com/?p=8018",
		"http://www.securitybistro.com/?p=8023": "http://www.securitybistro.com/?p=8023",
		"http://www.securitybistro.com/?p=8030": "http://www.securitybistro.com/?p=8030",
		"http://www.securitybistro.com/?p=8036": "http://www.securitybistro.com/?p=8036",
		"http://www.securitybistro.com/?p=7979": "http://www.securitybistro.com/?p=7979",
		"http://www.securitybistro.com/?p=8000": "http://www.securitybistro.com/?p=8000",
		"http://www.securitybistro.com/?p=8006": "http://www.securitybistro.com/?p=8006",
		"http://www.securitybistro.com/?p=8009": "http://www.securitybistro.com/?p=8009",
		"http://www.securitybistro.com/?p=8041": "http://www.securitybistro.com/?p=8041",
	}
	testURLs(t, urls)
}
func Test_search_itunes_apple_com(t *testing.T) {
	urls := map[string]string{
		"http://search.itunes.apple.com/WebObjects/MZContentLink.woa/wa/link?path=festival": "http://search.itunes.apple.com/WebObjects/MZContentLink.woa/wa/link?path=festival",
		"https://search.itunes.apple.com/WebObjects/MZContentLink.woa/wa/link?mt=8&path=apps%2Fautoadjust": "https://search.itunes.apple.com/WebObjects/MZContentLink.woa/wa/link?mt=8&path=apps%2Fautoadjust",
		"https://search.itunes.apple.com/WebObjects/MZContentLink.woa/wa/link?mt=8&path=apps%2FTheLoopMagazine": "https://search.itunes.apple.com/WebObjects/MZContentLink.woa/wa/link?mt=8&path=apps%2FTheLoopMagazine",
	}
	testURLs(t, urls)
}
func Test_searchcloudprovider_techtarget_com(t *testing.T) {
	urls := map[string]string{
		"http://searchcloudprovider.techtarget.com/news/1526056/Storage-execs-see-more-cloud-storage-data-reduction-SSDs-in-2011?track=sy483": "http://searchcloudprovider.techtarget.com/news/1526056/Storage-execs-see-more-cloud-storage-data-reduction-SSDs-in-2011?track=sy483",
	}
	testURLs(t, urls)
}
func Test_www_shrm_org(t *testing.T) {
	urls := map[string]string{
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNigerian-bereavement-absence.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNigerian-bereavement-absence.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FAla-Mere-Possibility-Not-Establish-Medical-Causation-Disability.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fdiversity%2B%2528SHRM%2BOnline%2BDiversity%2BNews%2529": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FAla-Mere-Possibility-Not-Establish-Medical-Causation-Disability.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fdiversity%2B%2528SHRM%2BOnline%2BDiversity%2BNews%2529",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FDel-Gender-Identity.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FDel-Gender-Identity.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FBike-to-Work-Good-for-Health-Wallet.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FBike-to-Work-Good-for-Health-Wallet.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fstaffingmanagement%2FArticles%2FPages%2FProtections-for-Interns.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fstaffingmanagement%2FArticles%2FPages%2FProtections-for-Interns.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FOFCCP-rules.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FOFCCP-rules.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FMinn-High-Court-Rules-Tips-Wage-State-Law.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FMinn-High-Court-Rules-Tips-Wage-State-Law.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fconsultants%2FArticles%2FPages%2FGenerating-Positive-Online-Reviews.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fconsultants%2FArticles%2FPages%2FGenerating-Positive-Online-Reviews.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FCEOs-Get-Little-Training.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FCEOs-Get-Little-Training.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FEmployees-Resign-Notice.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FEmployees-Resign-Notice.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FOregon-Bereavement-Leave.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FOregon-Bereavement-Leave.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FHow-Reliable-Are-Personality-Tests.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FHow-Reliable-Are-Personality-Tests.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FPersonality-Clashes-Cause-Workplace-Discord.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FPersonality-Clashes-Cause-Workplace-Discord.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FSingapore-Guidelines-Services-Sector-Job-Flexibility.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FSingapore-Guidelines-Services-Sector-Job-Flexibility.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FDOMA-subsequent-decision.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FDOMA-subsequent-decision.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNLRB-shutdown.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fhr%2B%2528SHRM%2BOnline%253A%2BHR%2BNews%2529": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNLRB-shutdown.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fhr%2B%2528SHRM%2BOnline%253A%2BHR%2BNews%2529",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNo-Biaso-Retaliation-Negative-Reviews.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fdiversity%2B%2528SHRM%2BOnline%2BDiversity%2BNews%2529": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNo-Biaso-Retaliation-Negative-Reviews.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fdiversity%2B%2528SHRM%2BOnline%2BDiversity%2BNews%2529",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FOFCCP-Back-Pay-Relief.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FOFCCP-Back-Pay-Relief.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FOnsite-FMLA-investigations.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FOnsite-FMLA-investigations.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Flegalissues%2Fstateandlocalresources%2Fpages%2Fva-employer-not-responsible-hostile-workplace.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Flegalissues%2Fstateandlocalresources%2Fpages%2Fva-employer-not-responsible-hostile-workplace.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FPublications%2FHRNews%2FPages%2FIWon-Lottery-KeepWorking.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FPublications%2FHRNews%2FPages%2FIWon-Lottery-KeepWorking.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FFBI-Track-Hate-Crimes-Sikhs.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FFBI-Track-Hate-Crimes-Sikhs.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FL-1-visa.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FL-1-visa.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FWellness-smokers.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FWellness-smokers.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalifornia-Health-Exchange.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalifornia-Health-Exchange.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fbenefits%2FArticles%2FPages%2FBoost-Retirement-Income.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fbenefits%2FArticles%2FPages%2FBoost-Retirement-Income.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FRhode-Island-Ban-the-Box.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FRhode-Island-Ban-the-Box.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FPublications%2FHRNews%2FPages%2FWorkers-CaringBosses-EmotionalIssues.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FPublications%2FHRNews%2FPages%2FWorkers-CaringBosses-EmotionalIssues.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fconsultants%2FArticles%2FPages%2FService-Recovery.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fconsultants%2FArticles%2FPages%2FService-Recovery.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FSit-Stand-Treadmill-Desks-Heart-Attack-Stroke.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FSit-Stand-Treadmill-Desks-Heart-Attack-Stroke.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FL-1-visa.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fglobal%2B%2528SHRM%2BOnline%2BGlobal%2BHR%2BNews%2529": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FL-1-visa.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fglobal%2B%2528SHRM%2BOnline%2BGlobal%2BHR%2BNews%2529",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Successful-Employers-Fees.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Successful-Employers-Fees.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FConnecticut-Minimum-Wage.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FConnecticut-Minimum-Wage.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FAir-Force-Academy-Cadets-Abroad.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FAir-Force-Academy-Cadets-Abroad.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Flaborrelations%2Farticles%2Fpages%2FUFCW-AFL-CIO.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Flaborrelations%2Farticles%2Fpages%2FUFCW-AFL-CIO.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FBlog-fired.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FBlog-fired.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FVeterans-final-rule.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fstaffing%2B%2528SHRM%2BOnline%2BRecruiting%2B%2526%2BStaffing%2BNews%2529": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FVeterans-final-rule.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fstaffing%2B%2528SHRM%2BOnline%2BRecruiting%2B%2526%2BStaffing%2BNews%2529",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FDecade-Decline-in-Teen-Jobs-Persists-This-Summer.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FDecade-Decline-in-Teen-Jobs-Persists-This-Summer.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FPages%2FCost-Replace-Millennials.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FPages%2FCost-Replace-Millennials.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FAttorneys-General-Criminal-Background-Check.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FAttorneys-General-Criminal-Background-Check.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FCEO-Personally-Liable-FLSA.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FCEO-Personally-Liable-FLSA.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Sexual-Harassment-Law-Amended-Sexual-Desire-Not-Required.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Sexual-Harassment-Law-Amended-Sexual-Desire-Not-Required.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FMich-Right-to-Work-Law-Applies-Classified-Civil-Service-Employees.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FMich-Right-to-Work-Law-Applies-Classified-Civil-Service-Employees.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FNJ-Punitive-Damages-Employment-Retaliation.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FNJ-Punitive-Damages-Employment-Retaliation.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FPublications%2FHRNews%2FPages%2FHousePanel-HealthCareReform-Hearing-Kentucky.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FPublications%2FHRNews%2FPages%2FHousePanel-HealthCareReform-Hearing-Kentucky.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FUrging-Employee-Vacations.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FUrging-Employee-Vacations.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fstaffingmanagement%2Farticles%2Fpages%2Fmillennials-drive-change-in-workplace.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fstaffingmanagement%2Farticles%2Fpages%2Fmillennials-drive-change-in-workplace.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FMexico-Employee-Productivity-Training.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FMexico-Employee-Productivity-Training.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FEEOC-criminal-background-credit-history-lawsuit.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FEEOC-criminal-background-credit-history-lawsuit.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FVeterans-final-rule.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Flaw%2B%2528SHRM%2BOnline%2BEmployment%2BLaws%2Band%2BRegulations%2BNews%2529": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FVeterans-final-rule.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Flaw%2B%2528SHRM%2BOnline%2BEmployment%2BLaws%2Band%2BRegulations%2BNews%2529",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Leave-Law-Emergency-Rescue-Personnel.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Leave-Law-Emergency-Rescue-Personnel.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FGenerationY-Distracted-Demanding-Lazy-or-No-Different-from-Others.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FGenerationY-Distracted-Demanding-Lazy-or-No-Different-from-Others.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fstaffingmanagement%2FArticles%2FPages%2FChanges-in-Recruitment-Methods.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fstaffingmanagement%2FArticles%2FPages%2FChanges-in-Recruitment-Methods.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Flegalissues%2Femploymentlawareas%2Fpages%2Feu-publishes-human-rights-guidance-csr.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Flegalissues%2Femploymentlawareas%2Fpages%2Feu-publishes-human-rights-guidance-csr.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FNJ-Restricts-Employers-Use-Social-Media.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FNJ-Restricts-Employers-Use-Social-Media.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fcompensation%2FArticles%2FPages%2FAutomatic-Tipping.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fcompensation%2FArticles%2FPages%2FAutomatic-Tipping.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FGallup-New-Workers-Disengaged-After-Honeymoon.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FGallup-New-Workers-Disengaged-After-Honeymoon.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2Fpages%2Fwhiner-know-it-all-naysayer-make-workplace-drag.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2Fpages%2Fwhiner-know-it-all-naysayer-make-workplace-drag.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FWorkers-Shocking-Horrible-News.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Femployeerelations%2Farticles%2FPages%2FWorkers-Shocking-Horrible-News.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FlBangladesh-Factory-Collapse.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FlBangladesh-Factory-Collapse.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNLRB-social-media-policy-photographs.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FNLRB-social-media-policy-photographs.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-2013-Training-Year-Sexual-Harassment-Prevention.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-2013-Training-Year-Sexual-Harassment-Prevention.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalifornia-marriage-licenses.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalifornia-marriage-licenses.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FNJ-Law-Banning-Salary-Secrecy.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FNJ-Law-Banning-Salary-Secrecy.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FBoomers-Leaving-Millennials-Job-Hopping-How-HR-Keeps-Up.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2FDiversity%2FArticles%2FPages%2FBoomers-Leaving-Millennials-Job-Hopping-How-HR-Keeps-Up.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Employment-Legislation-Update.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-Employment-Legislation-Update.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FMinimum-wage-Perez.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FMinimum-wage-Perez.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-DIR-Warns-Business-Misleading-Notices.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FCalif-DIR-Warns-Business-Misleading-Notices.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FIll-Did-Harassment-Lead-Constructive-Discharge.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fdiversity%2B%2528SHRM%2BOnline%2BDiversity%2BNews%2529": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FIll-Did-Harassment-Lead-Constructive-Discharge.aspx%3Futm_source%3Dfeedburner%26utm_medium%3Dfeed%26utm_campaign%3DFeed%253A%2Bshrm%252Fnews%252Fdiversity%2B%2528SHRM%2BOnline%2BDiversity%2BNews%2529",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FMinnesota-Expands-Sick-Leave.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FMinnesota-Expands-Sick-Leave.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fglobal%2Farticles%2Fpages%2Flanonymous-job-applications-canada.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Fglobal%2Farticles%2Fpages%2Flanonymous-job-applications-canada.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Ftechnology%2FArticles%2FPages%2FTrends-Lead-to-Job-Cuts-in-Tech-Sector.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Fhrdisciplines%2Ftechnology%2FArticles%2FPages%2FTrends-Lead-to-Job-Cuts-in-Tech-Sector.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FAppointment-NLRB-acting-general-counsel.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FAppointment-NLRB-acting-general-counsel.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FFLSA-interpretation-struck-down.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FFLSA-interpretation-struck-down.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Flegalissues%2Fstateandlocalresources%2Fpages%2Fcalif-leave-law-emergency-rescue-personnel.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2Flegalissues%2Fstateandlocalresources%2Fpages%2Fcalif-leave-law-emergency-rescue-personnel.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FPhilippines-Universal-Health-Care-Law.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FEmploymentLawAreas%2FPages%2FPhilippines-Universal-Health-Care-Law.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FImmigration-comprehensive-reform.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FFederalResources%2FPages%2FImmigration-comprehensive-reform.aspx",
		"http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FWindsor-wake.aspx": "http://www.shrm.org/Pages/login.aspx?ReturnUrl=%2FLegalIssues%2FStateandLocalResources%2FPages%2FWindsor-wake.aspx",
	}
	testURLs(t, urls)
}
func Test_searchconsumerization_techtarget_com(t *testing.T) {
	urls := map[string]string{
		"http://searchconsumerization.techtarget.com/ehandbook/Consumerization-security-and-compliance?asrc=RSS_BP_KABPNETHWSW": "http://searchconsumerization.techtarget.com/ehandbook/Consumerization-security-and-compliance",
	}
	testURLs(t, urls)
}
func Test_searchitchannel_techtarget_com(t *testing.T) {
	urls := map[string]string{
		"http://searchitchannel.techtarget.com/feature/Finalists-2010-data-storage-Products-of-the-Year?track=sy483": "http://searchitchannel.techtarget.com/feature/Finalists-2010-data-storage-Products-of-the-Year?track=sy483",
		"http://searchitchannel.techtarget.com/news/1525755/Thin-provisioning-driven-by-technology-advances-best-practices?track=sy483": "http://searchitchannel.techtarget.com/news/1525755/Thin-provisioning-driven-by-technology-advances-best-practices?track=sy483",
		"http://searchitchannel.techtarget.com/news/1526972/Xiotech-launches-Hybrid-ISE-solid-state-drives-system?track=sy483": "http://searchitchannel.techtarget.com/news/1526972/Xiotech-launches-Hybrid-ISE-solid-state-drives-system?track=sy483",
		"http://searchitchannel.techtarget.com/tip/Cloud-gateways-for-primary-storage-Benefits-and-challenges?track=sy483": "http://searchitchannel.techtarget.com/tip/Cloud-gateways-for-primary-storage-Benefits-and-challenges?track=sy483",
		"http://searchitchannel.techtarget.com/tip/Impact-of-storage-acquisitions-on-the-channel?track=sy483": "http://searchitchannel.techtarget.com/tip/Impact-of-storage-acquisitions-on-the-channel?track=sy483",
		"http://searchitchannel.techtarget.com/tip/Top-five-storage-channel-tips-of-2010?track=sy483": "http://searchitchannel.techtarget.com/tip/Top-five-storage-channel-tips-of-2010?track=sy483",
		"http://searchitchannel.techtarget.com/news/1525446/LTO-5-tape-cant-stop-backup-hardware-trend-toward-disk?track=sy483": "http://searchitchannel.techtarget.com/news/1525446/LTO-5-tape-cant-stop-backup-hardware-trend-toward-disk?track=sy483",
		"http://searchitchannel.techtarget.com/news/1526264/Enterprise-data-storage-2010-Products-of-the-Year-finalists?track=sy483": "http://searchitchannel.techtarget.com/news/1526264/Enterprise-data-storage-2010-Products-of-the-Year-finalists?track=sy483",
		"http://searchitchannel.techtarget.com/tip/Vendor-reseller-agreement-Avoid-midproject-partner-conflict?track=sy483": "http://searchitchannel.techtarget.com/tip/Vendor-reseller-agreement-Avoid-midproject-partner-conflict?track=sy483",
	}
	testURLs(t, urls)
}
