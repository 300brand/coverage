package cleanurl

import "testing"

func Test_www_lifescienceleader_com(t *testing.T) {
	urls := map[string]string{
		"http://www.lifescienceleader.com/magazine/past-issues3/item/4580-republicans-stop-allowing-vocal-ideologues-to-dictate-health-policy?list=n":                "http://www.lifescienceleader.com/magazine/past-issues3/item/4580-republicans-stop-allowing-vocal-ideologues-to-dictate-health-policy?list=n",
		"http://www.lifescienceleader.com/magazine/past-issues3/item/4585-adaptive-funding-survival-of-the-fittest-in-the-life-sciences-start-up-evolution?list=n":   "http://www.lifescienceleader.com/magazine/past-issues3/item/4585-adaptive-funding-survival-of-the-fittest-in-the-life-sciences-start-up-evolution?list=n",
		"http://www.lifescienceleader.com/magazine/past-issues3/item/4586-astrazeneca-translates-%E2%80%9Cwinning-with-science%E2%80%9D-from-words-to-action?list=n": "http://www.lifescienceleader.com/magazine/past-issues3/item/4586-astrazeneca-translates-%E2%80%9Cwinning-with-science%E2%80%9D-from-words-to-action?list=n",
	}
	testURLs(t, urls)
}
func Test_www_lohud_com(t *testing.T) {
	urls := map[string]string{
		"http://www.lohud.com/needlogin?redirecturl=http%3A%2F%2Fwww.lohud.com%2Farticle%2F20130802%2FNEWS%2F308020085%3Fnclick_check%3D1&type=login": "http://www.lohud.com/needlogin?redirecturl=http%3A%2F%2Fwww.lohud.com%2Farticle%2F20130802%2FNEWS%2F308020085%3Fnclick_check%3D1&type=login",
	}
	testURLs(t, urls)
}
func Test_live_wsj_com(t *testing.T) {
	urls := map[string]string{
		"http://live.wsj.com/video/paid-to-be-sick-in-seattle/ca2ea47d-8970-4a50-93a9-724f6af45c29.html?keywords=opinion+journal+live#%21ca2ea47d-8970-4a50-93a9-724f6af45c29": "http://live.wsj.com/video/paid-to-be-sick-in-seattle/ca2ea47d-8970-4a50-93a9-724f6af45c29.html?keywords=opinion+journal+live#%21ca2ea47d-8970-4a50-93a9-724f6af45c29",
	}
	testURLs(t, urls)
}
func Test_link_brightcove_com(t *testing.T) {
	urls := map[string]string{
		"http://link.brightcove.com/services/player/bcpid72925238001?bckey=AQ~~%2CAAAAAFcSbzI~%2COkyYKKfkn3x_w7fnuqEOoldofOzUPSQN&bctid=888603428001": "http://link.brightcove.com/services/player/bcpid72925238001?bckey=AQ~~%2CAAAAAFcSbzI~%2COkyYKKfkn3x_w7fnuqEOoldofOzUPSQN&bctid=888603428001",
	}
	testURLs(t, urls)
}
func Test_login_oracle_com(t *testing.T) {
	urls := map[string]string{
		"https://login.oracle.com/pls/orasso/orasso.wwsso_app_admin.ls_login?Site2pstoreToken=v1.4~49D3DFA3~3D65953B5404B2EB10CDBE7BF5B4B45A60BEB5ED87DCFEC7AD95FEDC849F3642E613E73F41A2AB0FCB6656EBC698E48A2A6A261148484B12D1B1B272EE13605268E2B6686DE1E59A8BEB4DAE93D5A38DE0CD29B3DDE74BCA8D9DF823673F7DEE28AF9C90A4D829CDA52DFE4D331FA55883F99A6985F3DA43D7C5A6A2D4FCDACE3B41BCB1CC8E2CD014D77DA2D2CC2954B44F39AA906770BB474DAAE4B762D378ACF2AA540F14CEA89D6987D9113AD30FA0BCDA4E57D2CF2D": "https://login.oracle.com/pls/orasso/orasso.wwsso_app_admin.ls_login?Site2pstoreToken=v1.4~49D3DFA3~3D65953B5404B2EB10CDBE7BF5B4B45A60BEB5ED87DCFEC7AD95FEDC849F3642E613E73F41A2AB0FCB6656EBC698E48A2A6A261148484B12D1B1B272EE13605268E2B6686DE1E59A8BEB4DAE93D5A38DE0CD29B3DDE74BCA8D9DF823673F7DEE28AF9C90A4D829CDA52DFE4D331FA55883F99A6985F3DA43D7C5A6A2D4FCDACE3B41BCB1CC8E2CD014D77DA2D2CC2954B44F39AA906770BB474DAAE4B762D378ACF2AA540F14CEA89D6987D9113AD30FA0BCDA4E57D2CF2D",
		"https://login.oracle.com/pls/orasso/orasso.wwsso_app_admin.ls_login?Site2pstoreToken=v1.4~49D3DFA3~C1B194AEE1479DCDA8075404369A87D3361AD9D74266DE9EFFC48B2CCDEF06E47CAE85F71908C825B55B2934708D91F37CA080BB44C7EA3D1CD1EF7499A10638BE636F89B2780284A7187DE6F3F00C8D7E3D0A205BEE41BE91C122543B3805B179204DB762DD1836602F3BE00B45545E9B0B3A700676CD13AF8A9C01088217B574C2D1EABC1D15B2671346E07FB465BFFB3D672039C20B9D758A1878A719E8E565423896CA6F9DCE73B0D1F0F2AE801732F03D51D5324990": "https://login.oracle.com/pls/orasso/orasso.wwsso_app_admin.ls_login?Site2pstoreToken=v1.4~49D3DFA3~C1B194AEE1479DCDA8075404369A87D3361AD9D74266DE9EFFC48B2CCDEF06E47CAE85F71908C825B55B2934708D91F37CA080BB44C7EA3D1CD1EF7499A10638BE636F89B2780284A7187DE6F3F00C8D7E3D0A205BEE41BE91C122543B3805B179204DB762DD1836602F3BE00B45545E9B0B3A700676CD13AF8A9C01088217B574C2D1EABC1D15B2671346E07FB465BFFB3D672039C20B9D758A1878A719E8E565423896CA6F9DCE73B0D1F0F2AE801732F03D51D5324990",
		"https://login.oracle.com/pls/orasso/orasso.wwsso_app_admin.ls_login?Site2pstoreToken=v1.4~49D3DFA3~FBD266C3BBFC7514889FC82CCBA0287C30EB13299908EF169594FF6271C772D6B6C73FD36D51F1DDB156BE1E7CC7E6B35891F764C386BCADD41E041A07F39E3114B8C81014EE51169C1390042E1BD842A54456B21F234A4917B138AD600208BEC67938AB8FFB0F4408FF542DC8D98855DBC278FDB0FD304F91C2B53AD2765A6770F5E1512AEB9028028828826E310A65FFDE8A07C0E6D6F110285D7F7F9858BDB4757FE622D59AC14AEEFD6C969B11FE":                 "https://login.oracle.com/pls/orasso/orasso.wwsso_app_admin.ls_login?Site2pstoreToken=v1.4~49D3DFA3~FBD266C3BBFC7514889FC82CCBA0287C30EB13299908EF169594FF6271C772D6B6C73FD36D51F1DDB156BE1E7CC7E6B35891F764C386BCADD41E041A07F39E3114B8C81014EE51169C1390042E1BD842A54456B21F234A4917B138AD600208BEC67938AB8FFB0F4408FF542DC8D98855DBC278FDB0FD304F91C2B53AD2765A6770F5E1512AEB9028028828826E310A65FFDE8A07C0E6D6F110285D7F7F9858BDB4757FE622D59AC14AEEFD6C969B11FE",
	}
	testURLs(t, urls)
}
func Test_live_chicagotribune_com(t *testing.T) {
	urls := map[string]string{
		"http://live.chicagotribune.com/Event/Entertainment_Chicago_Gourmet?Page=0": "http://live.chicagotribune.com/Event/Entertainment_Chicago_Gourmet?Page=0",
		"http://live.chicagotribune.com/Event/Week_1_NFL_Bears_vs_Bengals?Page=0":   "http://live.chicagotribune.com/Event/Week_1_NFL_Bears_vs_Bengals?Page=0",
	}
	testURLs(t, urls)
}
