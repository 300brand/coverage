package search

import (
	"testing"
)

func TestBooleanTree(t *testing.T) {
	tests := []struct {
		Search string
		Tree   [][]string
	}{
		{
			"phrase",
			[][]string{[]string{"phrase"}},
		},
		{
			"some phrase OR another phrase",
			[][]string{
				[]string{"some phrase"},
				[]string{"another phrase"},
			},
		},
		{
			"some phrase AND another phrase",
			[][]string{[]string{"some phrase", "another phrase"}},
		},
		{
			"a AND b OR c",
			[][]string{[]string{"a", "b"}, []string{"c"}},
		},
	}
	for n, test := range tests {
		b := NewBoolean(test.Search)
		if len(test.Tree) != len(b.Tree) {
			t.Errorf("[%d] Tree length mismatch", n)
			t.Logf("test %+v", test.Tree)
			t.Logf("bool %+v", b.Tree)
			continue
		}
		for i := range test.Tree {
			if len(test.Tree[i]) != len(b.Tree[i]) {
				t.Errorf("[%d][%d] Subtree length mismatch", n, i)
				t.Logf("test %+v", test.Tree)
				t.Logf("bool %+v", b.Tree)
				continue
			}
			for j := range test.Tree[i] {
				if x, y := test.Tree[i][j], b.Tree[i][j].String(); x != y {
					t.Errorf("[%d][%d] Inequal values '%s' != '%s'", i, j, x, y)
				}
			}
		}
	}
}

func TestBooleanBadMatches(t *testing.T) {
	tests := []struct {
		Haystack string
		Needle   string
	}{
		{
			"a b c",
			"d",
		},
		{
			"a b c",
			"a AND d",
		},
		{
			"a b c",
			"d OR e",
		},
		{
			"a b c",
			"e OR d OR f",
		},
		{
			"a b c",
			"d OR a AND e",
		},
		{
			"a b c",
			"e OR d OR f AND h AND g",
		},
		{
			"a b c d",
			"d c AND b a",
		},
		// Addition of NOT
		{
			"a b c",
			"a NOT b",
		},
		{
			"a b c",
			"a OR b NOT c",
		},
		{
			"a b c",
			"a NOT c OR d OR e",
		},
		// NOT tests from Jamie
		{
			"emcee",
			"EMC NOT emcee OR emceeing",
		},
		{
			"collision damage waiver",
			"CDW NOT collision damage waiver",
		},
		// Ugghhhh...
		{
			`WASHINGTON, March 6 (Reuters) - Three nominees to head the U.S. derivatives regulator met little pushback in the Senate on Thursday as the agency heads into quieter waters after the departure of Gary Gensler, the former chief who was deemed overly aggressive by big banks and other critics.
Timothy Massad, a corporate lawyer who also headed the $700 billion U.S. bank bailout program, has been nominated by President Barack Obama to take the helm of the Commodity Futures Trading Commission.
Massad and two other nominees must be confirmed by the Senate if they are to join the five-member Commission.
Once a little-known agency that regulated agriculture futures, the CFTC was given authority over the $630 trillion global swaps market in 2010, a task Gensler took at hand vigorously, angering banks in the process.
Most of the rule-writing is done, and Massad promised a more modest agenda of enforcing the new rules, fine-tuning them if needed, and cooperating to a greater extent with foreign regulators, with whom the CFTC has argued in the past.
"We must aggressively pursue wrongdoers whatever their position or size and we must deter and prevent unlawful practices," Massad said at a hearing of the Senate Agriculture Committee, which oversees the CFTC.
Sharon Bowen, a partner at law firm Latham & Watkins in New York, and Chris Giancarlo, an industry veteran at swaps broker GFI Group Inc, have also been nominated to the CFTC, and were also questioned during the hearing.
Appearing for a group of lawmakers heavily backed by agricultural interests, the nominees pledged to listen to concerns from farmers and ranchers, the original users of the futures markets the agency oversees.
Such groups have expressed concern that tight new rules written to tame Wall Street will make using the financial instruments prohibitively expensive.
Unusually, the Commission in its new set-up will no longer have a dedicated agriculture expert member. But lobby groups for the sector have said they can live with the three, as long as they are granted some form of access.
Massad also said he would work closely with authorities abroad, alluding to last year's public dispute with foreign regulators. At issue was the degree to which the CFTC wants non-U.S. banks to abide by its rules.
Giancarlo, a derivatives industry veteran from GFI and a nominee to a Republican spot on the Commission, warned of an overly hasty rule-writing process, a frequent complaint against the CFTC by banks.
Bowen, a second Democratic nominee to the Commission, may still face questions later in the process over the $7 billion Allen Stanford Ponzi scheme, because of her role in a decision that left investors out of pocket. Stanford is serving a 110-year prison sentence.
Bowen heads the Securities Investor Protection Corporation, which is embroiled in a legal battle with the Securities and Exchange Commission over SIPC's decision not to pay out to people who had lost money in the scandal.
U.S. Senator Thad Cochran, the highest ranking Republican on the Agriculture Committee, has publicly supported the SEC in its fight against SIPC. While he mentioned the scandal during his question, he did not pursue the issue.
`,
			"NetApp OR Red Hat OR EMC OR Lenovo OR Lexmark OR Symantec OR Tripp-Lite OR VMware OR Cisco OR Epson OR HP OR IBM OR Citrix OR Samsung OR Google OR 3M OR AirWatch OR Amazon OR Aruba Networks OR Asus OR Attachmate OR Autodesk OR Avaya OR Aver Information Inc. OR Avocent OR Barracuda Networks OR Belkin OR Blue Coat OR BMC Software OR Box.com OR Brocade OR Brother OR C2G OR Canon OR Check Point Software OR Code Scanners OR CommVault OR Drobo OR EDGE Tech Corp OR Elo Touchsystems OR Enovate IT OR Ergotron OR F5 Networks OR Fortinet OR Fujitsu OR Fusion IO OR Honeywell OR Imation OR InFocus OR Intel OR Isilon OR Juniper Networks OR Kaspersky Lab OR Kensington OR Kingston OR Kodak Scanners OR LANDesk Software OR LG Electronics OR Liebert OR LifeSize Communications OR McAfee OR Meraki OR Microsoft OR Mitsubishi OR Mobile Iron OR Motion Computing OR Motorola Enterprise Mobility OR NComputing OR NEC OR NETGEAR OR Nimble Storage OR Nuance Communications OR OKI OR Optoma OR Oracle Targus OR Panasonic OR Peerless OR Planar OR Plantronics OR Polycom OR PolyVision OR Quantum OR Quest Software OR Raritan OR Riverbed OR RSA Security OR Rubbermaid OR Salesforce.com OR SAP America OR Seagate Technology OR Sharp Electronics OR ShoreTel OR SonicWALL OR Sony OR Sophos OR Splunk Software OR StarTech.com OR Trend Micro OR Varonis OR Veeam OR ViewSonic OR Vision Solutions OR WatchGuard OR Websense OR West Point Products OR Western Digital OR Wyse Technology OR Accenture OR Best Buy OR Buy.com OR Rakuten OR Dell OR Dimension Data OR ePlus OR Ingram Micro OR Insight Enterprises OR Newegg OR Office Depot OR Office Max OR PC Connection OR PCM OR Presidio OR SHI OR Softchoice OR Staples OR Tech Data OR TigerDirect OR World Wide Technology OR Computer Associates",
		},
		{
			`A little more than twenty years ago, as a rising junior at the University of Northern Iowa,I moved in with my first gay flatmate. There would be two others in the coming years, during grad school in Maryland. It was an interesting time to live on the edge of gay men's lives. In 1993, I don't know that I'd even imagined the possibility of same-sex couples enjoying the privileges of state-sponsored matrimony. It did not seem to be a possibility my flatmate was interested in, probably for the same reason neither of us were much interested in booking a round a trip to Mars. I remember reading Andrew Sullivan's "Virtually Normal" a few years later, and how nice but fanciful the idea of legal gay marriage seemed to all of us then. If someone had told us that same-sex marriage would become legal in our state before we hit our forties, we wouldn't have believed it. We would have laughed. I don't think we even imagined it, yet it happened all the same.
We certainly didn't imagine that in 2014 Colorado bakers and New Mexico photographers would be subject to lawsuits for refusing to bake cakes for and shoot photos of same-sex weddings. On one level, it's deliriously gratifying that it has come to this. Not only can lesbians get hitched in New Mexico, which is an incredible fact all by itself, but Christian photographers who decline to work gay weddings can get sued for it andlose. Amazing! On another level, however, it's clear something has gone awry. The aim of legal same-sex marriage is equality under the law, not the criminalisation of a certain popular strain of Christian doctrine. The freedom to run a business in accordance with religious convictions that were recently all-but-universal is, like the freedom of same-sex couples to marry, a freedom worth having. Can't we have legally-binding gay weddings and photographers who won't shoot them? That seems nice for everybody. Let's do that.
Doing that seems to me to have been point of laws like Arizona's strangely controversial SB 1062, which was vetoed last week by Jan Brewer, Arizona's governor. Douglas Laycock, a professor of law at the University of Virginia, recently noted that the thrust of the bill was simply to refine existing state and federal religious-freedom protections."These laws", Mr Laycock writes, "enact a uniform standard-substantial burden and compelling interest-to be interpreted and applied to individual cases by courts. They rest on the sound premise that we should not punish people for practicing their religion unless we have a very good reason". The point of SB 1062 in particular was to clarify "that people are covered when state or local government requires them to violate their religion in the conduct of their business, and that people are covered when sued by a private citizen invoking state or local law to demand that they violate their religion." Mr Laycock goes on to emphasise, and this is very important:
But nothing in the amendment would have said who wins in either of these cases.SB1062 did not say that businesses can discriminate for religious reasons. It said that business people could assert a claim or defense under RFRA, ... that they would have to prove a substantial burden on a sincere religious practice, that the government or the person suing them would then have the burden of proof on compelling government interest, and that the state courts in Arizona would make the final decision.
It is incorrect to claim, as my colleague did last week, that SB 1062 was "in effect, an exemption from anti-discrimination laws for the pious". It was not. It was an attempt to calibrate the law so that worthy new legal rights don't infringe on worthy old ones. If forcing conservative Christian photographers to shoot gay weddings can be shown to promote a "compelling interest" of the state, and if the photographer fails to show that doing so would place a "substantial burden" on her sincere religious beliefs, then refusing to work a gay weddings would remain a violation of existing anti-discrimination law. That seems reasonable to me. As Mr Laycock says, "we should not punish people for practicing their religion unless we have a very good reason". When we do have a very good reason, we can go right ahead.
Ross Douthat of the New York Timesobserves that laws such as SB 1062 "have been seen, in the past, as a way for religious conservatives to negotiate surrender-to accept same-sex marriage's inevitability while carving out protections for dissent". But Mr Douthat worries that progressives are pressing their case too hard. "[N]ow, apparently, the official line is that you bigots don't get to negotiate anymore".Mark David Stern, writing at Slate, rather underscores Mr Douthat's point by refusing to see a difference between the religious conservative's wish to negotiate an honourable surrender and an apology for institutionalised racism:
At the core of Douthat's argument is a tacit shrug that, well, obviously anti-gay discrimination isn't as bad as racism: The Bible's hostility toward gays is a good deal clearer than its distaste for blacks. But Times readers would have no truck with such base bigotry, and so Douthat slips it between the lines, embedding it in the scaffolding that holds up his central premise. By the internal logic of Douthat's piece, homophobia is simply more defensible than racism. Nothing else could explain why denying gay customers is OK while denying black customers isn't.
That is to say, we can't have both gay weddings and photographers who won't shoot them because that would be like Jim Crow.But it wouldn't!
Mr Stern ignores a number of important distinctions here. First, there is the crucial difference between refusing to do business with someone simply because he is gay and refusing to sign a contract to play a part in a marriage ceremony that violates one's own dearly-held religious convictions about the function and meaning of marriage. Second, the religious belief that marriage ought to be reserved for heterosexual couples does not entail "homophobia", even if it is no accident that these are highly correlated atttitudes. Third, there are good reasons why "anti-gay discrimination isn't as bad as racism", and these reasons are relevant to anti-discrimination law and the justification of its scope. 
Racism is baked into America's DNA. The horror of America'sbrutal history of slavery and racial apartheid, it's centrality to American history and culture, and it's comprehensive distortion of American institutions demands redress. Redress requires, and therefore justifies, reasonable restrictions on otherwise sacrosanct liberal rights of free association.Institutionalised homophobia is also terrible, but not that terrible. Walt Whitman didn't have it as bad as Solomon Northrup, the subject of the Oscar-winning "Twelve Years a Slave". Whitmancouldn't have had it that bad, because we was white. Finally and belatedly guaranteeing equal legal rights to blacks after 400 years of monstrous racist oppression would not have been enough to begin rectify the vast injustice at the heart of American history. It was necessary, and remains necessary, to go further than formal equality. But finally and belatedly guaranteeing equal rights and equal legal protections to gays is enough. Punishing people for adhering to the dictates of their heteronormative faith traditions just isn't needed for gays to get a fair shake in America. Making religious conservatives feel persecuted, refusing them the dignity of toleration, is unnecessary and probably counterproductive. It's okay not to punish people for practicing their religion when it doesn't really hurt anyone. There are other photographers. There are others cakes. Everything's going to be all right. 
In 1993 I could not have imagined that we would be having a conversation about whether Christian vendors ought to be allowed to refuse to participate in gay weddings, for I did not realise it was possible that attitudes toward homosexuality might evolve so rapidly. Our moral culture has evolved with truly bewildering speed. It did not take some monumental Civil Rights Act to make it happen, and it's not necessary to curb-stomp religious conservatives to keep our evolution on track. "Virtually Normal" seemed like lemonade oceans in 1995; now the same-sex family unit in "Modern Family" seems more hackneyed than risque. The way public opinion is trending, 20 years hence it will be hard to convince our children that this ridiculous conversation was one we needed to have. 
`,
			"NetApp OR Red Hat OR EMC OR Lenovo OR Lexmark OR Symantec OR Tripp-Lite OR VMware OR Cisco OR Epson OR HP OR IBM OR Citrix OR Samsung OR Google OR 3M OR AirWatch OR Amazon OR Aruba Networks OR Asus OR Attachmate OR Autodesk OR Avaya OR Aver Information Inc. OR Avocent OR Barracuda Networks OR Belkin OR Blue Coat OR BMC Software OR Box.com OR Brocade OR Brother OR C2G OR Canon OR Check Point Software OR Code Scanners OR CommVault OR Drobo OR EDGE Tech Corp OR Elo Touchsystems OR Enovate IT OR Ergotron OR F5 Networks OR Fortinet OR Fujitsu OR Fusion IO OR Honeywell OR Imation OR InFocus OR Intel OR Isilon OR Juniper Networks OR Kaspersky Lab OR Kensington OR Kingston OR Kodak Scanners OR LANDesk Software OR LG Electronics OR Liebert OR LifeSize Communications OR McAfee OR Meraki OR Microsoft OR Mitsubishi OR Mobile Iron OR Motion Computing OR Motorola Enterprise Mobility OR NComputing OR NEC OR NETGEAR OR Nimble Storage OR Nuance Communications OR OKI OR Optoma OR Oracle Targus OR Panasonic OR Peerless OR Planar OR Plantronics OR Polycom OR PolyVision OR Quantum OR Quest Software OR Raritan OR Riverbed OR RSA Security OR Rubbermaid OR Salesforce.com OR SAP America OR Seagate Technology OR Sharp Electronics OR ShoreTel OR SonicWALL OR Sony OR Sophos OR Splunk Software OR StarTech.com OR Trend Micro OR Varonis OR Veeam OR ViewSonic OR Vision Solutions OR WatchGuard OR Websense OR West Point Products OR Western Digital OR Wyse Technology OR Accenture OR Best Buy OR Buy.com OR Rakuten OR Dell OR Dimension Data OR ePlus OR Ingram Micro OR Insight Enterprises OR Newegg OR Office Depot OR Office Max OR PC Connection OR PCM OR Presidio OR SHI OR Softchoice OR Staples OR Tech Data OR TigerDirect OR World Wide Technology OR Computer Associates",
		},
	}
	for i, test := range tests {
		b := NewBoolean(test.Needle)
		if b.Match([]byte(test.Haystack)) {
			t.Errorf("[%d] '%s' should not be found in '%s'", i, test.Needle, test.Haystack)
		}
	}
}

func TestBooleanGoodMatches(t *testing.T) {
	tests := []struct {
		Haystack string
		Needle   string
	}{
		{
			"a b c",
			"a",
		},
		{
			"a b c",
			"a AND b",
		},
		{
			"a b c",
			"a AND c",
		},
		{
			"a b c",
			"a OR b",
		},
		{
			"a b c",
			"d OR b",
		},
		{
			"a b c",
			"e OR d OR c",
		},
		{
			"a b c",
			"d OR a AND c",
		},
		{
			"a b c",
			"e OR d OR b AND c AND a",
		},
		{
			"a b c",
			"a AND c OR d",
		},
		{
			"a b c d",
			"a b AND b c",
		},
		{
			"a b c d",
			"a b AND c d",
		},
		{
			"a b c",
			"a AND b OR c AND d",
		},
		// Addition of NOT
		{
			"a b c",
			"a NOT d",
		},
	}
	for i, test := range tests {
		b := NewBoolean(test.Needle)
		if !b.Match([]byte(test.Haystack)) {
			t.Errorf("[%d] '%s' not found in '%s'", i, test.Needle, test.Haystack)
		}
	}
}

func TestBooleanMinTerms(t *testing.T) {
	tests := []struct {
		Query string
		Terms int
	}{
		{"EMC OR Brocade", 1},
		{"PaaS AND EMC OR Brocade", 1},
		{"PaaS AND EMC OR PaaS AND Brocade", 2},
		{"Platform as a Service AND EMC OR Platform as a Service AND Brocade", 3},
	}
	for i, test := range tests {
		b := NewBoolean(test.Query)
		if n := b.MinTerms(); n != test.Terms {
			t.Errorf("[%d] %s: %d != %d", i, test.Query, n, test.Terms)
		}
	}
}

func TestBooleanTerms(t *testing.T) {
	tests := []struct {
		Query string
		Terms []string
	}{
		{"spacemonkey NOT NASA", []string{"spacemonkey"}},
	}
	for i, test := range tests {
		b := NewBoolean(test.Query)
		t.Logf("[%d] %v", i, b.Terms())
	}
}
