package testfeed

// ORIGIN: http://www.networkworld.com/rss/netflash.xml
var RDF = []byte(`<?xml version="1.0" encoding="utf-8"?>
<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns="http://purl.org/rss/1.0/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:rss="http://purl.org/rss/1.0/" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/">
	<channel rdf:about="http://www.networkworld.com/columnists/gibbs.html">
		<title>Netflash</title>
		<link>http://www.networkworld.com/rss/netflash.xml</link>
		<description>Network World News</description>
		<dc:publisher>Network World, Inc.</dc:publisher>
		<dc:rights>Copyright(C) 1994 - 2012 Network World, Inc.</dc:rights>
		<image rdf:resource="http://www.networkworld.com/redesign2/logorss.gif"></image>
		<items>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-microsoft-to-address-yammer-integration-264178.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-verizon-droid-264180.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-microsoft-surface-rt-tablet-touch-264167.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-cray-bumps-ibm-from-top500-264168.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-windows-8-security-unshaken-by-264166.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-firefox-turns-8-and-gets-264165.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-nvidia-amd-release-graphics-processors-264169.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-rim-blackberry10-264160.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-pirate-bay-co-founder-now-suspected-264159.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-virnetx-targets-ipad-mini-iphone-264158.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-ebay-returns-to-china-with-264157.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-microsoft-likely-to-leash-ios-264156.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-skype-for-windows-phone-8-264155.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-patent-settlement-a-win-for-264153.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-intel-ships-60-core-xeon-phi-264152.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111212-un39s-civil-aviation-body-recommends-264151.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111112-apple-htc-settle-patent-suits-264150.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111012-google-access-returns-to-china-264149.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111012-space-station-commander-controls-lego-robot-264148.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/111012-manufacturer-sues-ibm-over-sap-264140.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-wall-street-beat-tech-mampa-264146.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-google-blocked-in-china-by-264144.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-three-simple-rules-for-buying-264143.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-staff-emails-are-not-owned-264142.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-cray-chases-clusters-with-appro-264138.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-nyc-corporate-partnership-seeks-fresh-264139.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-symantec-encryption-264137.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-crossbeam-thoma-bravo-264135.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-fatal-half-measures-in-incident-264128.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-iphone6-rumors-264127.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-adobe-reader-x-sandbox-bypassed-264129.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-ransomware-crooks-make-millions-from-264130.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-cisco-recommends-mcafee-switch-for-264121.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-china-could-be-behind-twitter-264123.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-shareholders-kept-in-the-dark-264122.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-judge-to-consider-samsung39s-questions-264124.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-iranian-minister-faces-us-sanctions-264118.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110912-michigan-man-charged-with-selling-264119.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-how-to-secure-big-data-264120.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-microsoft-launches-skype-centered-hub-for-264116.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-windows8-security-patch-264111.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-atampt-reverses-facetime-blocking-264117.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-csa-huawei-264110.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-android-264108.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-microsoft-slates-first-windows-8-264109.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-vmware-cloud-paas-264107.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-ipad-tips-and-tricks-for-264104.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-windows8-264106.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-smartphone-and-tablet-users-helped-264105.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-godaddy-apology-264100.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-iphone-foxconn-264099.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-hp-urges-consumer-customers-not-264102.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-isoc-facebook-comcast-ipv6-264096.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-twitter-asks-many-users-to-264097.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-tech-apologies-264095.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-samsung-galaxy-smartphone-264094.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-cisco-collaboration-264093.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-why-isnt-microsofts-answer-to-264092.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-the-cloud-as-data-center-264090.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-apple-university-264091.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-five-must-have-business-apps-for-264087.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-oracle-buys-instantis-for-project-264085.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-security-cloud-computing-264086.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-citrix-and-netapp-simplify-on-premises-264083.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-zero-day-pdf-exploit-reportedly-defeats-264084.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-social-nets-create-election39s-biggest-264082.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-hitachi-releases-new-16tb-flash-264081.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-verizon-expands-lync-service-to-264080.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-heist-once-again-highlights-e-banking-264077.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-siemens-industrial-software-targeted-by-264071.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-us-commission-fingers-china-as-264078.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-lenovo-expects-convertible-pcs-will-264079.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-joyent-264069.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-chinese-ex-hacker-says-working-for-264074.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-cray39s-next-supercomputer-has-speedy-264075.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-lenovo-sees-profit-growth-sag-264076.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-blackberry-10-is-fips-certified-264072.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110812-obama-tech-adviser-says-re-election-264073.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-design-supply-component-issues-may-264068.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-microsoft-office-ios-android-264067.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-oracle-hit-with-patent-lawsuit-264065.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-can-windows-8-give-developers-264064.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-it-salaries-2013-264063.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-amd-open-source-264062.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-twitter-a-big-winner-in-264059.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-aclu-eff-challenge-law-targeting-264060.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-us-election-could-mean-movement-264057.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-gartner-magic-quadrant-264058.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-samsung-laying-groundwork-for-server-264053.html?source=nww_rss" />
			</rdf:Seq>
			<rdf:Seq>
				<rdf:li resource="http://www.networkworld.com/news/2012/110712-cloud-security-lawyers-264055.html?source=nww_rss" />
			</rdf:Seq>
		</items>
	</channel>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-microsoft-to-address-yammer-integration-264178.html?source=nww_rss">
		<title>Microsoft to address Yammer integration plans at SharePoint conference</title>
		<link>http://www.networkworld.com/news/2012/111212-microsoft-to-address-yammer-integration-264178.html?source=nww_rss</link>
		<description>Microsoft will shine the spotlight this week on SharePoint's new 2013 version at a conference devoted to the popular collaboration server, but recently acquired Yammer may grab substantial attention.</description>
		<dc:creator>Juan Carlos Perez</dc:creator>
		<dc:date>2012-11-12T12:36:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-verizon-droid-264180.html?source=nww_rss">
		<title>Verizon to unbox new HTC Droid DNA live on Google+</title>
		<link>http://www.networkworld.com/news/2012/111212-verizon-droid-264180.html?source=nww_rss</link>
		<description>Verizon announced today that it is planning a live unboxing of the latest "Droid"-labeled smartphone on Monday, Nov. 19, at noon EST, to be streamed live on Google+. The new device is widely believed to be the HTC Droid DNA, a U.S. version of the Taiwanese company's cutting-edge Butterfly J, currently available only in Japan.</description>
		<dc:creator>Jon Gold</dc:creator>
		<dc:date>2012-11-12T12:34:25-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-microsoft-surface-rt-tablet-touch-264167.html?source=nww_rss">
		<title>Microsoft Surface RT tablet touch cover keyboard fuels complaints</title>
		<link>http://www.networkworld.com/news/2012/111212-microsoft-surface-rt-tablet-touch-264167.html?source=nww_rss</link>
		<description>Microsoft's hip marketing may be luring people to buy its Surface RT tablets, but it hasn't prevented some customers from complaining that the touch cover keyboard falls apart.</description>
		<dc:creator>Christina Desmarais</dc:creator>
		<dc:date>2012-11-12T10:55:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-cray-bumps-ibm-from-top500-264168.html?source=nww_rss">
		<title>Cray bumps IBM from Top500 supercomputer top spot</title>
		<link>http://www.networkworld.com/news/2012/111212-cray-bumps-ibm-from-top500-264168.html?source=nww_rss</link>
		<description>The U.S. Department of Energy Oak Ridge National Laboratory's newly installed Titan system, a Cray XK7, has been anointed as the world's fastest supercomputer in the newly released 40th edition of the Top500 compilation of the world's fastest supercomputers.</description>
		<dc:creator>Joab Jackson</dc:creator>
		<dc:date>2012-11-12T10:41:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-windows-8-security-unshaken-by-264166.html?source=nww_rss">
		<title>Windows 8 security unshaken by antivirus vendor's claims</title>
		<link>http://www.networkworld.com/news/2012/111212-windows-8-security-unshaken-by-264166.html?source=nww_rss</link>
		<description>Bitdefender raises worry over trusting included antivirus software, but one analyst said Windows 8's core security picks up the slack</description>
		<dc:creator>Antone Gonsalves</dc:creator>
		<dc:date>2012-11-12T10:31:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-firefox-turns-8-and-gets-264165.html?source=nww_rss">
		<title>Firefox turns 8 and gets a key security boost</title>
		<link>http://www.networkworld.com/news/2012/111212-firefox-turns-8-and-gets-264165.html?source=nww_rss</link>
		<description>Between Android and Firefox, this has been a big week for milestones in the world of free and open source software.</description>
		<dc:creator>Katherine Noyes</dc:creator>
		<dc:date>2012-11-12T10:20:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-nvidia-amd-release-graphics-processors-264169.html?source=nww_rss">
		<title>Nvidia, AMD release graphics processors for supercomputing</title>
		<link>http://www.networkworld.com/news/2012/111212-nvidia-amd-release-graphics-processors-264169.html?source=nww_rss</link>
		<description>Nvidia and Advanced Micro Devices on Monday announced high-performance graphics chips for supercomputers.</description>
		<dc:creator>Agam Shah</dc:creator>
		<dc:date>2012-11-12T10:07:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-rim-blackberry10-264160.html?source=nww_rss">
		<title>RIM plans to unveil BlackBerry 10 and two smartphones Jan. 30</title>
		<link>http://www.networkworld.com/news/2012/111212-rim-blackberry10-264160.html?source=nww_rss</link>
		<description>Research in Motion will officially launch its new mobile operating system, BlackBerry 10, on Jan. 30, along with the first two RIM-designed smartphones to run it, the company announced today. The company has bet its future on the platform's success.</description>
		<dc:creator>John Cox</dc:creator>
		<dc:date>2012-11-12T09:10:18-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-pirate-bay-co-founder-now-suspected-264159.html?source=nww_rss">
		<title>Pirate Bay co-founder now suspected of serious fraud and another data intrusion</title>
		<link>http://www.networkworld.com/news/2012/111212-pirate-bay-co-founder-now-suspected-264159.html?source=nww_rss</link>
		<description>Swedish authorities now suspects Pirate Bay co-founder Gottfrid Svartholm Warg of serious fraud and another data intrusion in addition to the alleged hacking of IT company Logica that led to his arrest, public prosecutor Henrik Olin said Monday.</description>
		<dc:creator>Loek Essers</dc:creator>
		<dc:date>2012-11-12T08:35:50-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-virnetx-targets-ipad-mini-iphone-264158.html?source=nww_rss">
		<title>VirnetX targets iPad Mini, iPhone 5 in new lawsuit after winning $368M judgment against Apple</title>
		<link>http://www.networkworld.com/news/2012/111212-virnetx-targets-ipad-mini-iphone-264158.html?source=nww_rss</link>
		<description>The same day it won a $368 million verdict in a patent infringement case against Apple, VirnetX filed a new lawsuit, alleging that the iPad Mini and iPhone 5 violate the same patents, according to court documents.</description>
		<dc:creator>Gregg Keizer</dc:creator>
		<dc:date>2012-11-12T07:56:03-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-ebay-returns-to-china-with-264157.html?source=nww_rss">
		<title>EBay returns to China with new site focused on selling clothing</title>
		<link>http://www.networkworld.com/news/2012/111212-ebay-returns-to-china-with-264157.html?source=nww_rss</link>
		<description>EBay launched a new e-commerce business in China with a local partner on Monday, marking the company&amp;#39;s re-entry into an ultra-competitive market that once eluded its grasp. But this time, the company is focusing on a smaller market segment with a site geared at selling popular clothing brands to consumers.</description>
		<dc:creator>Michael Kan</dc:creator>
		<dc:date>2012-11-12T07:39:57-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-microsoft-likely-to-leash-ios-264156.html?source=nww_rss">
		<title>Microsoft likely to leash iOS Office apps to Office 365, say analysts</title>
		<link>http://www.networkworld.com/news/2012/111212-microsoft-likely-to-leash-ios-264156.html?source=nww_rss</link>
		<description>Microsoft will probably tie Office apps for the iPhone and iPad to its Office 365 "rental" subscription plans to prevent the mobile apps from cannibalizing sales and to skirt the "Apple tax," analysts said today.</description>
		<dc:creator>Gregg Keizer</dc:creator>
		<dc:date>2012-11-12T07:24:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-skype-for-windows-phone-8-264155.html?source=nww_rss">
		<title>Skype for Windows Phone 8 preview available for download</title>
		<link>http://www.networkworld.com/news/2012/111212-skype-for-windows-phone-8-264155.html?source=nww_rss</link>
		<description>Microsoft has made a preview version of Skype for Windows Phone 8 available for download, which promises better integration with the operating system&amp;#39; user interface and more features.</description>
		<dc:creator>Mikael Rickna$?s</dc:creator>
		<dc:date>2012-11-12T07:09:31-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-patent-settlement-a-win-for-264153.html?source=nww_rss">
		<title>Patent settlement a win for HTC, but company still needs better smartphones, analysts say</title>
		<link>http://www.networkworld.com/news/2012/111212-patent-settlement-a-win-for-264153.html?source=nww_rss</link>
		<description>Analysts say HTC still faces an uphill battle to rebuild its smartphone business amid heated competition, despite reaching a deal with Apple to settle their patent disputes.</description>
		<dc:creator>Michael Kan</dc:creator>
		<dc:date>2012-11-12T05:46:13-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-intel-ships-60-core-xeon-phi-264152.html?source=nww_rss">
		<title>Intel ships 60-core Xeon Phi processor</title>
		<link>http://www.networkworld.com/news/2012/111212-intel-ships-60-core-xeon-phi-264152.html?source=nww_rss</link>
		<description>Intel hopes to deliver performance and power-efficiency breakthroughs to servers with the new Xeon Phi family of processors, the first model of which is now shipping to customers, the company said on Monday.</description>
		<dc:creator>Agam Shah</dc:creator>
		<dc:date>2012-11-12T01:10:05-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111212-un39s-civil-aviation-body-recommends-264151.html?source=nww_rss">
		<title>UN's civil aviation body recommends cybersecurity task force</title>
		<link>http://www.networkworld.com/news/2012/111212-un39s-civil-aviation-body-recommends-264151.html?source=nww_rss</link>
		<description>The U.N.'s civil aviation body will recommend creating a cybersecurity task force at a meeting next week in Canada, as new technologies introduced into aviation systems are increasing the risk of cyberattacks.</description>
		<dc:creator>Jeremy Kirk</dc:creator>
		<dc:date>2012-11-11T10:29:42-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111112-apple-htc-settle-patent-suits-264150.html?source=nww_rss">
		<title>Apple, HTC settle patent suits worldwide</title>
		<link>http://www.networkworld.com/news/2012/111112-apple-htc-settle-patent-suits-264150.html?source=nww_rss</link>
		<description>Apple and Taiwanese smartphone maker HTC have settled all of their outstanding patent disputes, they said late Saturday.</description>
		<dc:creator>Martyn Williams</dc:creator>
		<dc:date>2012-11-10T09:48:05-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111012-google-access-returns-to-china-264149.html?source=nww_rss">
		<title>Google access returns to China after brief blocking</title>
		<link>http://www.networkworld.com/news/2012/111012-google-access-returns-to-china-264149.html?source=nww_rss</link>
		<description>Access to Google services in China appeared to return Saturday morning after they were blocked briefly as the country prepares to appoint new leadership.</description>
		<dc:creator>Michael Kan</dc:creator>
		<dc:date>2012-11-10T12:44:04-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111012-space-station-commander-controls-lego-robot-264148.html?source=nww_rss">
		<title>Space-station commander controls Lego robot on Earth with futuristic network</title>
		<link>http://www.networkworld.com/news/2012/111012-space-station-commander-controls-lego-robot-264148.html?source=nww_rss</link>
		<description>Late last month in Germany, a robot made its first moves on Earth under commands from an orbiting spacecraft.</description>
		<dc:creator>Stephen Lawson</dc:creator>
		<dc:date>2012-11-09T09:26:05-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/111012-manufacturer-sues-ibm-over-sap-264140.html?source=nww_rss">
		<title>Manufacturer sues IBM over SAP project &amp;#39;disaster&amp;#39;</title>
		<link>http://www.networkworld.com/news/2012/111012-manufacturer-sues-ibm-over-sap-264140.html?source=nww_rss</link>
		<description>IBM has been slapped with a multimillion dollar lawsuit by chemical products manufacturer Avantor Performance Materials, which alleges that IBM lied about the suitability of a SAP-based software package it sells in order to win Avantor&amp;#39;s business.</description>
		<dc:creator>Chris Kanaracus</dc:creator>
		<dc:date>2012-11-09T08:06:28-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-wall-street-beat-tech-mampa-264146.html?source=nww_rss">
		<title>Wall Street Beat: Tech M&amp;amp;A flattens as economic uncertainty reigns, report says</title>
		<link>http://www.networkworld.com/news/2012/110912-wall-street-beat-tech-mampa-264146.html?source=nww_rss</link>
		<description>The number and value of technology mergers and acquisitions is flattening out, mainly as a result of economic uncertainty, according to a PricewaterhouseCoopers report.</description>
		<dc:creator>Marc Ferranti</dc:creator>
		<dc:date>2012-11-09T06:17:25-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-google-blocked-in-china-by-264144.html?source=nww_rss">
		<title>Google blocked in China by censors, unclear how long it will last</title>
		<link>http://www.networkworld.com/news/2012/110912-google-blocked-in-china-by-264144.html?source=nww_rss</link>
		<description>China on Friday started blocking access to Google in what's likely a move to squelch any controversial content on the nation's Internet as its government prepares to change leadership.</description>
		<dc:creator>Michael Kan</dc:creator>
		<dc:date>2012-11-09T05:15:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-three-simple-rules-for-buying-264143.html?source=nww_rss">
		<title>Three simple rules for buying a new laptop</title>
		<link>http://www.networkworld.com/news/2012/110912-three-simple-rules-for-buying-264143.html?source=nww_rss</link>
		<description>This is the time of year when friends, family members, casual acquaintances, and people in the street stop me to ask about buying a new PC.</description>
		<dc:creator>Rick Broida</dc:creator>
		<dc:date>2012-11-09T03:29:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-staff-emails-are-not-owned-264142.html?source=nww_rss">
		<title>Staff emails are not owned by firms, rules judge</title>
		<link>http://www.networkworld.com/news/2012/110912-staff-emails-are-not-owned-264142.html?source=nww_rss</link>
		<description>A high court judge has ruled that companies do not have a general claim of ownership of the content contained in staff emails.</description>
		<dc:creator>Antony Savvas</dc:creator>
		<dc:date>2012-11-09T03:09:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-cray-chases-clusters-with-appro-264138.html?source=nww_rss">
		<title>Cray chases clusters with Appro acquisition</title>
		<link>http://www.networkworld.com/news/2012/110912-cray-chases-clusters-with-appro-264138.html?source=nww_rss</link>
		<description>Cray on Friday agreed to acquire server maker Appro International for US$25 million in cash as it looks to strengthen its high-performance computing product portfolio.</description>
		<dc:creator>Agam Shah</dc:creator>
		<dc:date>2012-11-09T03:05:53-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-nyc-corporate-partnership-seeks-fresh-264139.html?source=nww_rss">
		<title>NYC corporate partnership seeks fresh financial tech</title>
		<link>http://www.networkworld.com/news/2012/110912-nyc-corporate-partnership-seeks-fresh-264139.html?source=nww_rss</link>
		<description>A partnership of New York City businesses has put out a call for innovative financial technologies.</description>
		<dc:creator>Joab Jackson</dc:creator>
		<dc:date>2012-11-09T02:50:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-symantec-encryption-264137.html?source=nww_rss">
		<title>Symantec broadens encryption options for Apple mobile devices, Dropbox file-sharing</title>
		<link>http://www.networkworld.com/news/2012/110912-symantec-encryption-264137.html?source=nww_rss</link>
		<description>Symantec this week announced two new data encryption options, one that is file-based encryption for use with Dropbox, and the second that works as an extension to the Apple iOS mail client to encrypt and decrypt documents.</description>
		<dc:creator>Ellen Messmer</dc:creator>
		<dc:date>2012-11-09T02:02:38-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-crossbeam-thoma-bravo-264135.html?source=nww_rss">
		<title>Cloud security vendor Crossbeam bought by private equity firm</title>
		<link>http://www.networkworld.com/news/2012/110912-crossbeam-thoma-bravo-264135.html?source=nww_rss</link>
		<description>Crossbeam has been bought by private equity firm Thoma Bravo, opening up the door for the security company to expand its product line both in-house and via acquisitions.</description>
		<dc:creator>Tim Greene</dc:creator>
		<dc:date>2012-11-09T01:44:47-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-fatal-half-measures-in-incident-264128.html?source=nww_rss">
		<title>Fatal half-measures in incident response</title>
		<link>http://www.networkworld.com/news/2012/110912-fatal-half-measures-in-incident-264128.html?source=nww_rss</link>
		<description>It&amp;#39;s not a matter of if, but when, you are breached. So what&amp;#39;s your plan?</description>
		<category>IDG Insider</category>
		<dc:creator>George V. Hulme</dc:creator>
		<dc:date>2012-11-09T11:43:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-iphone6-rumors-264127.html?source=nww_rss">
		<title>iPhone 6 rumor rollup for the week ending Nov. 9</title>
		<link>http://www.networkworld.com/news/2012/110912-iphone6-rumors-264127.html?source=nww_rss</link>
		<description>The U.S. presidential election hangover seems to have had a dampening effect on iOSphere rumors for iPhone 6.</description>
		<dc:creator>John Cox</dc:creator>
		<dc:date>2012-11-09T11:26:10-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-adobe-reader-x-sandbox-bypassed-264129.html?source=nww_rss">
		<title>Adobe Reader X sandbox bypassed by zero-day flaw</title>
		<link>http://www.networkworld.com/news/2012/110912-adobe-reader-x-sandbox-bypassed-264129.html?source=nww_rss</link>
		<description>Criminals have gained access to a newly discovered flaw in Adobe's Reader X program that can beat its sandboxing security isolation technology, Russian security firm Group-IB has claimed.</description>
		<dc:creator>John E Dunn</dc:creator>
		<dc:date>2012-11-09T10:36:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-ransomware-crooks-make-millions-from-264130.html?source=nww_rss">
		<title>Ransomware crooks make millions from porn-shaming scams</title>
		<link>http://www.networkworld.com/news/2012/110912-ransomware-crooks-make-millions-from-264130.html?source=nww_rss</link>
		<description>Ransomware is a growth industry that puts at least $5 million a year into criminals&amp;#39; coffers, according to Symantec.</description>
		<dc:creator>Gregg Keizer</dc:creator>
		<dc:date>2012-11-09T10:33:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-cisco-recommends-mcafee-switch-for-264121.html?source=nww_rss">
		<title>Cisco recommends McAfee switch for IronPort customers hit by Sophos flaws</title>
		<link>http://www.networkworld.com/news/2012/110912-cisco-recommends-mcafee-switch-for-264121.html?source=nww_rss</link>
		<description>Cisco Systems has warned customers about critical vulnerabilities in the Sophos antivirus engine included in its Cisco IronPort email and Web security appliances.</description>
		<dc:creator>Lucian Constantin</dc:creator>
		<dc:date>2012-11-09T10:26:32-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-china-could-be-behind-twitter-264123.html?source=nww_rss">
		<title>China could be behind Twitter password reset</title>
		<link>http://www.networkworld.com/news/2012/110912-china-could-be-behind-twitter-264123.html?source=nww_rss</link>
		<description>Twitter sent notices of an attempted hacking to China-based foreign journalists and analysts just hours before apologizing for resetting the passwords of more users than necessary in a recent break-in of accounts.</description>
		<dc:creator>Antone Gonsalves</dc:creator>
		<dc:date>2012-11-09T07:00:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-shareholders-kept-in-the-dark-264122.html?source=nww_rss">
		<title>Shareholders kept in the dark on data breaches</title>
		<link>http://www.networkworld.com/news/2012/110912-shareholders-kept-in-the-dark-264122.html?source=nww_rss</link>
		<description>It happened more than three and a half years ago. So it presumably would be old news that Chinese hackers broke into soft drink behemoth Coca-Cola's computer systems and stole confidential files relating to its effort to acquire the China Huiyuan Juice Group for $2.4 billion. But it is just coming to light now.</description>
		<dc:creator>Taylor Armerding</dc:creator>
		<dc:date>2012-11-09T07:00:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-judge-to-consider-samsung39s-questions-264124.html?source=nww_rss">
		<title>Judge to consider Samsung&amp;#39;s questions about jury foreman</title>
		<link>http://www.networkworld.com/news/2012/110912-judge-to-consider-samsung39s-questions-264124.html?source=nww_rss</link>
		<description>A court in California said Thursday that it would consider Samsung Electronics&amp;#39; concern that the foreman of the jury deciding a patent infringement lawsuit between Apple and Samsung had concealed information.</description>
		<dc:creator>John Ribeiro</dc:creator>
		<dc:date>2012-11-09T06:41:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-iranian-minister-faces-us-sanctions-264118.html?source=nww_rss">
		<title>Iranian minister faces US sanctions for Internet censorship</title>
		<link>http://www.networkworld.com/news/2012/110912-iranian-minister-faces-us-sanctions-264118.html?source=nww_rss</link>
		<description>The U.S. on Thursday said it ordered sanctions against Iran&amp;#39;s Minister of Communication and Information Technology, Reza Taghipour, and other entities and persons responsible for engaging in censorship in their country.</description>
		<dc:creator>John Ribeiro</dc:creator>
		<dc:date>2012-11-09T03:17:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110912-michigan-man-charged-with-selling-264119.html?source=nww_rss">
		<title>Michigan man charged with selling counterfeit Microsoft software</title>
		<link>http://www.networkworld.com/news/2012/110912-michigan-man-charged-with-selling-264119.html?source=nww_rss</link>
		<description>A man from Michigan was arraigned in a U.S. federal court on Thursday on charges of mail fraud and selling counterfeit software worth over US$1.2 million that he purchased from China and Singapore, the U.S Department of Justice said Thursday.</description>
		<dc:creator>John Ribeiro</dc:creator>
		<dc:date>2012-11-08T11:52:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-how-to-secure-big-data-264120.html?source=nww_rss">
		<title>How to Secure Big Data in Hadoop</title>
		<link>http://www.networkworld.com/news/2012/110812-how-to-secure-big-data-264120.html?source=nww_rss</link>
		<description>The promise of big data is enormous, but it can also become an albatross around your neck if you don&amp;#39;t make security of both your data and your infrastructure a key part of your big data project from the beginning. Here are some steps you can take to avoid big data pitfalls.</description>
		<dc:creator>Thor Olavsrud</dc:creator>
		<dc:date>2012-11-08T07:08:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-microsoft-launches-skype-centered-hub-for-264116.html?source=nww_rss">
		<title>Microsoft launches Skype-centered hub for small businesses</title>
		<link>http://www.networkworld.com/news/2012/110812-microsoft-launches-skype-centered-hub-for-264116.html?source=nww_rss</link>
		<description>Microsoft has unveiled a commercial networking site where small businesses can promote their products and services and interact with potential customers and partners primarily using Skype.</description>
		<dc:creator>Juan Carlos Perez</dc:creator>
		<dc:date>2012-11-08T06:26:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-windows8-security-patch-264111.html?source=nww_rss">
		<title>Windows 8 gets first critical Patch Tuesday security bulletins</title>
		<link>http://www.networkworld.com/news/2012/110812-windows8-security-patch-264111.html?source=nww_rss</link>
		<description>Windows 8 hasn't even been on sale for a month yet but is already the recipient of three critical security updates via Microsoft's monthly Patch Tuesday security bulletins, each of which will block flaws that allow remote execution of code on targeted machines.</description>
		<dc:creator>Tim Greene</dc:creator>
		<dc:date>2012-11-08T05:16:10-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-atampt-reverses-facetime-blocking-264117.html?source=nww_rss">
		<title>AT&amp;amp;T reverses FaceTime blocking decision</title>
		<link>http://www.networkworld.com/news/2012/110812-atampt-reverses-facetime-blocking-264117.html?source=nww_rss</link>
		<description>AT&amp;amp;T has reversed its decision to allow Apple iPhone and iPad owners to use Apple&amp;#39;s FaceTime videoconferencing application only on the carrier&amp;#39;s most expensive data plans or if they are connected to Wi-Fi.</description>
		<dc:creator>Grant Gross</dc:creator>
		<dc:date>2012-11-08T05:13:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-csa-huawei-264110.html?source=nww_rss">
		<title>Huawei security chief: We can help keep U.S. safe from 'Net threats</title>
		<link>http://www.networkworld.com/news/2012/110812-csa-huawei-264110.html?source=nww_rss</link>
		<description>The chief security officer of Huawei, the Chinese company recently flagged by Congress as a national security threat, says the network equipment maker could actually help the United States defend itself against malicious Internet traffic.</description>
		<dc:creator>Ellen Messmer</dc:creator>
		<dc:date>2012-11-08T05:11:57-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-android-264108.html?source=nww_rss">
		<title>Hottest Android news and rumors for the week ending November 9</title>
		<link>http://www.networkworld.com/news/2012/110812-android-264108.html?source=nww_rss</link>
		<description>It's been a busy week for Android news and gossip, with interesting information flying around about both hardware and software alike. Arguably the biggest news is the HTC Droid DNA - or possibly DLX - which is a potentially groundbreaking new device, and the latest in the long-running series called "HTC is terrible at naming things."</description>
		<dc:creator>Jon Gold</dc:creator>
		<dc:date>2012-11-08T04:33:45-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-microsoft-slates-first-windows-8-264109.html?source=nww_rss">
		<title>Microsoft slates first Windows 8, RT patches since launch</title>
		<link>http://www.networkworld.com/news/2012/110812-microsoft-slates-first-windows-8-264109.html?source=nww_rss</link>
		<description>Microsoft will issue six security updates next week, including three for Windows 8 and its tablet spin-off Windows RT.</description>
		<dc:creator>Gregg Keizer</dc:creator>
		<dc:date>2012-11-08T04:32:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-vmware-cloud-paas-264107.html?source=nww_rss">
		<title>VMware updates micro version of Cloud Foundry PaaS</title>
		<link>http://www.networkworld.com/news/2012/110812-vmware-cloud-paas-264107.html?source=nww_rss</link>
		<description>Everything in the cloud seems to be getting bigger or smaller. VMware today went the small route, releasing an updated micro version of the company's popular open source platform as a service (PaaS), Cloud Foundry.</description>
		<dc:creator>Brandon Butler</dc:creator>
		<dc:date>2012-11-08T04:16:13-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-ipad-tips-and-tricks-for-264104.html?source=nww_rss">
		<title>iPad Tips and Tricks for Creating Content</title>
		<link>http://www.networkworld.com/news/2012/110812-ipad-tips-and-tricks-for-264104.html?source=nww_rss</link>
		<description>The iPad is known for consuming content, not creating it. But there are hidden features and short-cuts for the keyboard, within email and in Safari that make creating content on an iPad faster and easier than you think.</description>
		<dc:creator>Tom Kaneshige</dc:creator>
		<dc:date>2012-11-08T04:05:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-windows8-264106.html?source=nww_rss">
		<title>Windows 8 Update: Microsoft's Surface RT has a higher profit margin than the iPad</title>
		<link>http://www.networkworld.com/news/2012/110812-windows8-264106.html?source=nww_rss</link>
		<description>Microsoft stands to make $30.85 more per Surface RT tablet with keyboard than Apple does with a similar 32G iPad without a keyboard, according to iSuppli teardown analyses.</description>
		<dc:creator>Tim Greene</dc:creator>
		<dc:date>2012-11-08T03:34:01-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-smartphone-and-tablet-users-helped-264105.html?source=nww_rss">
		<title>Smartphone and tablet users helped Obama win</title>
		<link>http://www.networkworld.com/news/2012/110812-smartphone-and-tablet-users-helped-264105.html?source=nww_rss</link>
		<description>Smartphone and tablet users helped both presidential candidates raise funds and support in 2012, while mobile computing contributed directly to President Barack Obama's edge in Tuesday's presidential election.</description>
		<dc:creator>Matt Hamblen</dc:creator>
		<dc:date>2012-11-08T03:06:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-godaddy-apology-264100.html?source=nww_rss">
		<title>Go Daddy CEO's outage apology</title>
		<link>http://www.networkworld.com/news/2012/110812-godaddy-apology-264100.html?source=nww_rss</link>
		<description>Go Daddy Site Outage Investigation Completed (Sept. 11, 2012)</description>
		<dc:creator>Network World Staff</dc:creator>
		<dc:date>2012-11-08T02:24:18-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-iphone-foxconn-264099.html?source=nww_rss">
		<title>iPhone 5 design, quality demands slows manufacturing</title>
		<link>http://www.networkworld.com/news/2012/110812-iphone-foxconn-264099.html?source=nww_rss</link>
		<description>You're facing a 3-4 week wait if you order an iPhone 5 today, according to Apple's Website. Partly, that's because Apple's manufacturing partner is having trouble keeping up with high demand for the new phone.</description>
		<dc:creator>John Cox</dc:creator>
		<dc:date>2012-11-08T01:48:05-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-hp-urges-consumer-customers-not-264102.html?source=nww_rss">
		<title>HP urges consumer customers not to downgrade new PCs to Windows 7</title>
		<link>http://www.networkworld.com/news/2012/110812-hp-urges-consumer-customers-not-264102.html?source=nww_rss</link>
		<description>Hewlett-Packard (HP) has advised consumer customers not to downgrade new PCs equipped with Windows 8 to the earlier Windows 7.</description>
		<dc:creator>Gregg Keizer</dc:creator>
		<dc:date>2012-11-08T01:42:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-isoc-facebook-comcast-ipv6-264096.html?source=nww_rss">
		<title>ISOC honors Facebook, Comcast execs for IPv6 work</title>
		<link>http://www.networkworld.com/news/2012/110812-isoc-facebook-comcast-ipv6-264096.html?source=nww_rss</link>
		<description>The Internet Society last night awarded its highest honor for work related to IPv6, the next generation Internet Protocol, to executives from Facebook and Comcast. Recipients of the award were Paul Saab and Donn Lee, software engineers at Facebook, and John Brzozowski, Distinguished Engineer and Chief Architect for IPv6 at Comcast.</description>
		<dc:creator>Carolyn Duffy Marsan</dc:creator>
		<dc:date>2012-11-08T12:17:15-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-twitter-asks-many-users-to-264097.html?source=nww_rss">
		<title>Twitter asks many users to change passwords</title>
		<link>http://www.networkworld.com/news/2012/110812-twitter-asks-many-users-to-264097.html?source=nww_rss</link>
		<description>Twitter appears to have reset the passwords for an undetermined portion of its user base because of a possible security breach.</description>
		<dc:creator>Joab Jackson</dc:creator>
		<dc:date>2012-11-08T12:16:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-tech-apologies-264095.html?source=nww_rss">
		<title>Most memorable tech industry apologies of 2012: From Apple to Google to Microsoft</title>
		<link>http://www.networkworld.com/news/2012/110812-tech-apologies-264095.html?source=nww_rss</link>
		<description>Tech vendors have been as bombastic as ever promoting the magical and amazing things their latest smartphones, cloud computing wares and network gear can do. When things go wrong, they're naturally a little less visible, but plenty of companies have sucked it up and done the right thing this year (perhaps with a little legal prodding here and there) and publicly apologized for minor and major customers inconveniences.</description>
		<dc:creator>Bob Brown</dc:creator>
		<dc:date>2012-11-08T11:59:20-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-samsung-galaxy-smartphone-264094.html?source=nww_rss">
		<title>Samsung Galaxy S III is smartphone top seller - for the moment</title>
		<link>http://www.networkworld.com/news/2012/110812-samsung-galaxy-smartphone-264094.html?source=nww_rss</link>
		<description>The Samsung Galaxy S III, for the first time, is the top-selling smartphone in the world, beating out both the iPhone 4S and iPhone 5, according to Strategy Analytics data for the third quarter of 2012.</description>
		<dc:creator>Jon Gold</dc:creator>
		<dc:date>2012-11-08T11:58:20-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-cisco-collaboration-264093.html?source=nww_rss">
		<title>Cisco replaces collaboration group head again</title>
		<link>http://www.networkworld.com/news/2012/110812-cisco-collaboration-264093.html?source=nww_rss</link>
		<description>Cisco has named a former Symantec executive to head its struggling collaboration group, which saw three different leaders in less than a year.</description>
		<dc:creator>Jim Duffy</dc:creator>
		<dc:date>2012-11-08T11:28:32-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-why-isnt-microsofts-answer-to-264092.html?source=nww_rss">
		<title>Why isn't Microsoft's answer to Siri built into Windows 8?</title>
		<link>http://www.networkworld.com/news/2012/110812-why-isnt-microsofts-answer-to-264092.html?source=nww_rss</link>
		<description>Windows 8 is supposed to be Microsoft's majestic OS reset--a dramatic overhaul designed to usher the Windows platform into the age of mobility. And Windows 8 is also Microsoft's bid to achieve feature parity with iOS and Android, the other two OS powerhouses in the mobile universe.</description>
		<dc:creator>Mark Sullivan</dc:creator>
		<dc:date>2012-11-08T10:59:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-the-cloud-as-data-center-264090.html?source=nww_rss">
		<title>The cloud as data-center extension</title>
		<link>http://www.networkworld.com/news/2012/110812-the-cloud-as-data-center-264090.html?source=nww_rss</link>
		<description>A year after Oregon's Multnomah County deployed an on-premises portfolio management application, the two IT staffers dedicated to it resigned. Other staff struggled to maintain the specialized server environment. Left with no other option to guarantee support of the mission-critical tool, the county leapt into the cloud.</description>
		<dc:creator>Sandra Gittlen</dc:creator>
		<dc:date>2012-11-08T10:58:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-apple-university-264091.html?source=nww_rss">
		<title>Apple seeks standard to appease angry university net managers</title>
		<link>http://www.networkworld.com/news/2012/110812-apple-university-264091.html?source=nww_rss</link>
		<description>Under fire from its customers in the higher education market, Apple has proposed creating a new industry standard that would fix problems with its Bonjour zero configuration networking technology that is causing scalability and security problems on campus networks.</description>
		<dc:creator>Carolyn Duffy Marsan</dc:creator>
		<dc:date>2012-11-08T10:32:53-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-five-must-have-business-apps-for-264087.html?source=nww_rss">
		<title>Five must-have business apps for the iPad Mini</title>
		<link>http://www.networkworld.com/news/2012/110812-five-must-have-business-apps-for-264087.html?source=nww_rss</link>
		<description>How does the iPad Mini rate as a business tablet? PC World's Tony Bradley recently tackled that topic in "Apple iPad Mini: All the iPad at (nearly) half the cost."</description>
		<dc:creator>Rick Broida</dc:creator>
		<dc:date>2012-11-08T10:05:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-oracle-buys-instantis-for-project-264085.html?source=nww_rss">
		<title>Oracle buys Instantis for project portfolio management software</title>
		<link>http://www.networkworld.com/news/2012/110812-oracle-buys-instantis-for-project-264085.html?source=nww_rss</link>
		<description>Oracle on Thursday said it has agreed to acquire PPM (project portfolio management) software vendor Instantis, in a move that will build upon its past acquisition of Primavera. Terms of the deal, which is expected to be completed this year, were not disclosed.</description>
		<dc:creator>Chris Kanaracus</dc:creator>
		<dc:date>2012-11-08T09:47:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-security-cloud-computing-264086.html?source=nww_rss">
		<title>Evolving security standards a challenge for cloud computing, expert says</title>
		<link>http://www.networkworld.com/news/2012/110812-security-cloud-computing-264086.html?source=nww_rss</link>
		<description>Any enterprise looking to use cloud computing services will also be digging into what laws and regulations might hold in terms of security and privacy of data stored in the cloud. At the Cloud Security Alliance Congress in Orlando this week, discussion centered on two important regulatory frameworks now being put in place in Europe and the U.S.</description>
		<dc:creator>Ellen Messmer</dc:creator>
		<dc:date>2012-11-08T09:21:09-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-citrix-and-netapp-simplify-on-premises-264083.html?source=nww_rss">
		<title>Citrix and NetApp simplify on-premises data sharing</title>
		<link>http://www.networkworld.com/news/2012/110812-citrix-and-netapp-simplify-on-premises-264083.html?source=nww_rss</link>
		<description>Citrix Systems and NetApp have jointly developed a software and hardware package optimized for Citrix&amp;#39;s ShareFile with StorageZones.</description>
		<dc:creator>Mikael Rickna$?s</dc:creator>
		<dc:date>2012-11-08T08:10:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-zero-day-pdf-exploit-reportedly-defeats-264084.html?source=nww_rss">
		<title>Zero-day PDF exploit reportedly defeats Adobe Reader sandbox protection</title>
		<link>http://www.networkworld.com/news/2012/110812-zero-day-pdf-exploit-reportedly-defeats-264084.html?source=nww_rss</link>
		<description>Cybercriminals are using a new PDF exploit that bypasses the sandbox security features in Adobe Reader X and XI, in order to install banking malware on computers, according to researchers from Russian security firm Group-IB.</description>
		<dc:creator>Lucian Constantin</dc:creator>
		<dc:date>2012-11-08T08:03:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-social-nets-create-election39s-biggest-264082.html?source=nww_rss">
		<title>Social nets create election&amp;#39;s biggest memes</title>
		<link>http://www.networkworld.com/news/2012/110812-social-nets-create-election39s-biggest-264082.html?source=nww_rss</link>
		<description>The 2012 presidential campaign was focused on serious stuff, but that doesn&amp;#39;t mean there wasn&amp;#39;t room for some fun. That&amp;#39;s where social networks came in, and their users quickly seized on verbal slip-ups, comical photos and missteps by the candidates.</description>
		<dc:creator>Sharon Gaudin</dc:creator>
		<dc:date>2012-11-08T07:14:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-hitachi-releases-new-16tb-flash-264081.html?source=nww_rss">
		<title>Hitachi releases new 1.6TB flash modules</title>
		<link>http://www.networkworld.com/news/2012/110812-hitachi-releases-new-16tb-flash-264081.html?source=nww_rss</link>
		<description>HDS&amp;#39;s new flash module is built specifically for enterprise-class workloads. The 1.6TB module fits in an 8U flash chassis. Each enclosure can scale from 6.4TB up to 76.8TB.</description>
		<dc:creator>Lucas Mearian</dc:creator>
		<dc:date>2012-11-08T07:14:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-verizon-expands-lync-service-to-264080.html?source=nww_rss">
		<title>Verizon expands Lync service to hosted operations, management of UC</title>
		<link>http://www.networkworld.com/news/2012/110812-verizon-expands-lync-service-to-264080.html?source=nww_rss</link>
		<description>Verizon said it is extending its managed services offering for Microsoft Lync Server to its business customers by adding the ability to operate, monitor and manage unified communications and collaboration servers and functions.</description>
		<dc:creator>Matt Hamblen</dc:creator>
		<dc:date>2012-11-08T07:14:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-heist-once-again-highlights-e-banking-264077.html?source=nww_rss">
		<title>Heist once again highlights e-banking vulnerabilities</title>
		<link>http://www.networkworld.com/news/2012/110812-heist-once-again-highlights-e-banking-264077.html?source=nww_rss</link>
		<description>The CFO of a Missouri firm who discover that cyber thieves had withdrawn $180,000 from the company's account overnight called it "a helluva wake-up call." But he might have avoided the mess if he has paid better attention to the risks of electronic banking.</description>
		<dc:creator>Taylor Armerding</dc:creator>
		<dc:date>2012-11-08T07:12:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-siemens-industrial-software-targeted-by-264071.html?source=nww_rss">
		<title>Siemens industrial software targeted by Stuxnet is still full of holes</title>
		<link>http://www.networkworld.com/news/2012/110812-siemens-industrial-software-targeted-by-264071.html?source=nww_rss</link>
		<description>Software made by Siemens and targeted by the Stuxnet malware is still full of other dangerous vulnerabilities, according to Russian researchers whose presentation at the Defcon security conference earlier this year was cancelled following a request from the company.</description>
		<dc:creator>Jeremy Kirk</dc:creator>
		<dc:date>2012-11-08T07:06:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-us-commission-fingers-china-as-264078.html?source=nww_rss">
		<title>U.S. commission fingers China as biggest cyberthreat</title>
		<link>http://www.networkworld.com/news/2012/110812-us-commission-fingers-china-as-264078.html?source=nww_rss</link>
		<description>A U.S. commission has confirmed what many experts already believed: China has become "the most threatening actor in cyberspace," due to a persistent bombardment of U.S. military systems and defense contractors.</description>
		<dc:creator>Antone Gonsalves</dc:creator>
		<dc:date>2012-11-08T07:00:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-lenovo-expects-convertible-pcs-will-264079.html?source=nww_rss">
		<title>Lenovo expects convertible PCs will rise past tablets</title>
		<link>http://www.networkworld.com/news/2012/110812-lenovo-expects-convertible-pcs-will-264079.html?source=nww_rss</link>
		<description>Lenovo&amp;#39;s CEO said on Thursday he expects the market will gradually move away from entertainment-focused tablets in favor for convertible PCs, which he said can strike a balance between the functions of touch-based tablets and the productivity of a laptop.</description>
		<dc:creator>Michael Kan</dc:creator>
		<dc:date>2012-11-08T06:09:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-joyent-264069.html?source=nww_rss">
		<title>Amazon, Microsoft and Google targeted by cloud provider Joyent</title>
		<link>http://www.networkworld.com/news/2012/110812-joyent-264069.html?source=nww_rss</link>
		<description>Joyent may be the biggest cloud provider you haven't heard of.</description>
		<dc:creator>Brandon Butler</dc:creator>
		<dc:date>2012-11-08T06:00:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-chinese-ex-hacker-says-working-for-264074.html?source=nww_rss">
		<title>Chinese ex-hacker says working for the government would be too boring</title>
		<link>http://www.networkworld.com/news/2012/110812-chinese-ex-hacker-says-working-for-264074.html?source=nww_rss</link>
		<description>Tao Wan now works at IBM, but said he was once an angry young man, a veteran of the hacking scene that burgeoned in China in the late 1990s.</description>
		<dc:creator>Jeremy Kirk</dc:creator>
		<dc:date>2012-11-08T05:50:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-cray39s-next-supercomputer-has-speedy-264075.html?source=nww_rss">
		<title>Cray&amp;#39;s next supercomputer has speedy interconnect</title>
		<link>http://www.networkworld.com/news/2012/110812-cray39s-next-supercomputer-has-speedy-264075.html?source=nww_rss</link>
		<description>For its next generation of supercomputers, Cray has focused on radically improving the I/O (input/output) of individual nodes. The new XC30 supercomputer will feature a new interconnect, called Aries, and a new routing topology that together promise to dramatically improve internal bandwidth.</description>
		<dc:creator>Joab Jackson</dc:creator>
		<dc:date>2012-11-08T04:01:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-lenovo-sees-profit-growth-sag-264076.html?source=nww_rss">
		<title>Lenovo sees profit growth sag in fiscal Q2</title>
		<link>http://www.networkworld.com/news/2012/110812-lenovo-sees-profit-growth-sag-264076.html?source=nww_rss</link>
		<description>Lenovo said Thursday its net profit for the fiscal second quarter increased by only 13 percent year-over-year, marking a shift from the high profit growth the company has previously seen.</description>
		<dc:creator>Michael Kan</dc:creator>
		<dc:date>2012-11-08T03:06:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-blackberry-10-is-fips-certified-264072.html?source=nww_rss">
		<title>BlackBerry 10 is FIPS certified in advance of platform&amp;#39;s release</title>
		<link>http://www.networkworld.com/news/2012/110812-blackberry-10-is-fips-certified-264072.html?source=nww_rss</link>
		<description>Research In Motion took the unusual step of announcing a tough security certification for BlackBerry 10 in advance of the device&amp;#39;s launch next quarter.</description>
		<dc:creator>Matt Hamblen</dc:creator>
		<dc:date>2012-11-08T01:09:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110812-obama-tech-adviser-says-re-election-264073.html?source=nww_rss">
		<title>Obama tech adviser says re-election sets stage for progress on spectrum</title>
		<link>http://www.networkworld.com/news/2012/110812-obama-tech-adviser-says-re-election-264073.html?source=nww_rss</link>
		<description>Federal initiatives to make more spectrum available for mobile services are likely to take off running after President Barack Obama&amp;#39;s re-election on Tuesday, a member of a presidential technology commission said.</description>
		<dc:creator>Stephen Lawson</dc:creator>
		<dc:date>2012-11-07T09:33:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-design-supply-component-issues-may-264068.html?source=nww_rss">
		<title>Design, supply component issues may be hurting iPhone 5 production</title>
		<link>http://www.networkworld.com/news/2012/110712-design-supply-component-issues-may-264068.html?source=nww_rss</link>
		<description>Stringent iPhone 5 production specifications established by Apple and supply issues with new components like the Lightning port and larger screen could be responsible for Foxconn's delays of the handset, analysts said on Wednesday.</description>
		<dc:creator>Agam Shah</dc:creator>
		<dc:date>2012-11-07T05:29:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-microsoft-office-ios-android-264067.html?source=nww_rss">
		<title>Yawns may greet Microsoft Office port to iOS and Android</title>
		<link>http://www.networkworld.com/news/2012/110712-microsoft-office-ios-android-264067.html?source=nww_rss</link>
		<description>Microsoft's most iconic application suite -- Microsoft Office -- will be moving off the PC and the Windows OS to two mobile platforms, iOS and Android, in early 2013. But will customers in the enterprise, where Office has been a PC standard for years, really care?</description>
		<dc:creator>John Cox</dc:creator>
		<dc:date>2012-11-07T04:51:41-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-oracle-hit-with-patent-lawsuit-264065.html?source=nww_rss">
		<title>Oracle hit with patent lawsuit over WebLogic Server</title>
		<link>http://www.networkworld.com/news/2012/110712-oracle-hit-with-patent-lawsuit-264065.html?source=nww_rss</link>
		<description>Oracle is finding itself caught up in another Java-related patent lawsuit, but this time it's the one getting sued.</description>
		<dc:creator>Chris Kanaracus</dc:creator>
		<dc:date>2012-11-07T04:24:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-can-windows-8-give-developers-264064.html?source=nww_rss">
		<title>Can Windows 8 Give Developers What iOS and Android Lack?</title>
		<link>http://www.networkworld.com/news/2012/110712-can-windows-8-give-developers-264064.html?source=nww_rss</link>
		<description>Microsoft's future hinges on attracting developers to build Windows 8 apps. But by offering financial incentives, supporting a range of programming languages and allowing developers to write code once for multiple devices, those developers may soon follow.</description>
		<dc:creator>Shane O'neill</dc:creator>
		<dc:date>2012-11-07T04:10:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-it-salaries-2013-264063.html?source=nww_rss">
		<title>2013 IT salaries: 15 titles getting the biggest pay raises</title>
		<link>http://www.networkworld.com/news/2012/110712-it-salaries-2013-264063.html?source=nww_rss</link>
		<description>Mobile application developers and wireless network engineers can expect a 9% and 7.9% increase in starting salaries, respectively, says recruiting and staffing specialist Robert Half Technology.</description>
		<dc:creator>Ann Bednarz</dc:creator>
		<dc:date>2012-11-07T03:51:42-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-amd-open-source-264062.html?source=nww_rss">
		<title>Did AMD shoot itself in the foot by laying off open-source talent?</title>
		<link>http://www.networkworld.com/news/2012/110712-amd-open-source-264062.html?source=nww_rss</link>
		<description>The staff of Advanced Micro Devices' Operating System Research Center has been laid off, according to a report from The H Online, dramatically reducing the company's ability to contribute to the Linux community and support its own hardware on the platform.</description>
		<dc:creator>Jon Gold</dc:creator>
		<dc:date>2012-11-07T03:44:09-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-twitter-a-big-winner-in-264059.html?source=nww_rss">
		<title>Twitter a big winner in 2012 presidential election</title>
		<link>http://www.networkworld.com/news/2012/110712-twitter-a-big-winner-in-264059.html?source=nww_rss</link>
		<description>As presidential election day unfolded on Tuesday, people spent time posting photos of long lines at polling places, tweeting about casting a vote and commenting on a viral video of a malfunctioning voting machine.</description>
		<dc:creator>Sharon Gaudin</dc:creator>
		<dc:date>2012-11-07T03:41:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-aclu-eff-challenge-law-targeting-264060.html?source=nww_rss">
		<title>ACLU, EFF challenge law targeting online activities of sex offenders</title>
		<link>http://www.networkworld.com/news/2012/110712-aclu-eff-challenge-law-targeting-264060.html?source=nww_rss</link>
		<description>Two civil rights groups have filed a lawsuit challenging parts of a California ballot measure that requires registered sex offenders to turn over their Internet identities and service providers to police.</description>
		<dc:creator>Grant Gross</dc:creator>
		<dc:date>2012-11-07T03:37:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-us-election-could-mean-movement-264057.html?source=nww_rss">
		<title>US election could mean movement on high-skill immigration, copyright</title>
		<link>http://www.networkworld.com/news/2012/110712-us-election-could-mean-movement-264057.html?source=nww_rss</link>
		<description>Tuesday's election in the U.S. leaves President Barack Obama in the White House and maintains the balance of power in Congress. In many longstanding technology debates, policy experts see little movement forward, although lawmakers may look for compromises on a handful of issues.</description>
		<dc:creator>Grant Gross</dc:creator>
		<dc:date>2012-11-07T03:06:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-gartner-magic-quadrant-264058.html?source=nww_rss">
		<title>Gartner's IaaS Magic Quadrant: a who's who of cloud market</title>
		<link>http://www.networkworld.com/news/2012/110712-gartner-magic-quadrant-264058.html?source=nww_rss</link>
		<description>The cloud market can be a big, daunting place. Seemingly every tech vendor has a cloud strategy, with new products and services dubbed "cloud" coming out every week. But who are the real market leaders? Research firm Gartner's answer lies in its Magic Quadrant report for the infrastructure as a service (IaaS) market.</description>
		<dc:creator>Brandon Butler</dc:creator>
		<dc:date>2012-11-07T02:52:12-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-samsung-laying-groundwork-for-server-264053.html?source=nww_rss">
		<title>Samsung laying groundwork for server chips, analysts say</title>
		<link>http://www.networkworld.com/news/2012/110712-samsung-laying-groundwork-for-server-264053.html?source=nww_rss</link>
		<description>Samsung's recent licensing of 64-bit processor designs from ARM suggests that the chip maker may expand from smartphones and tablets into the server market, analysts said this week.</description>
		<dc:creator>Agam Shah</dc:creator>
		<dc:date>2012-11-07T01:51:00-04:00</dc:date>
	</item>
	<item rdf:about="http://www.networkworld.com/news/2012/110712-cloud-security-lawyers-264055.html?source=nww_rss">
		<title>Are lawyers getting in the way of cloud-based security?</title>
		<link>http://www.networkworld.com/news/2012/110712-cloud-security-lawyers-264055.html?source=nww_rss</link>
		<description>In an age where enterprises and their employees are being relentlessly targeted with malware-based phishing, denial-of-service and other attacks, the ability of the IT security staff to defend their networks and valuable corporate data faces yet one more obstacle, according to some: their own company lawyers.</description>
		<dc:creator>Ellen Messmer</dc:creator>
		<dc:date>2012-11-07T01:49:37-04:00</dc:date>
	</item>
</rdf:RDF>`)
