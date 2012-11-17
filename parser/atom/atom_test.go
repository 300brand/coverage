package atom

import (
	"git.300brand.com/coverage/parser/testfeed"
	"testing"
)

func TestEntryLen(t *testing.T) {
	doc := Doc{}
	if err := doc.Decode(testfeed.Atom); err != nil {
		t.Error(err)
	}
	if len(doc.Entry) != 50 {
		t.Errorf("Invalid number of entries: %d", len(doc.Entry))
	}
}

func TestParserFail(t *testing.T) {
	doc := Doc{}
	if err := doc.Decode(testfeed.RSS); err == nil {
		t.Error("Expected error when parsing RSS feed")
	}
}

func TestTitle(t *testing.T) {
	doc := Doc{}
	doc.Decode(testfeed.Atom)
	if doc.Title == "" {
		t.Error("Blank title")
	}
	t.Logf("Title: %s", doc.Title)
}

func TestURLs(t *testing.T) {
	urls := []string{
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/uss_enterprise_scapped/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/something_for_the_weekend_we_dont_talk/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/debenhams_expresso/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/shat_app/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/quotw_ending_november_2/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/bond_villain/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/bond_quiz/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/wwii_enigma_machine_auction/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/11/02/1000s_of_lab_rats_are_drowning_victims_of_sandy_hurricane/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/31/eric_idle_brian_cox_galaxy_cox/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/31/bond_martini_deconstructed/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/30/apple_siri_escort_services_china_blocked/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/30/sgtk_japan_toto_toilet_football/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/29/alan_turing_ten_pound_note_epetition/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/29/james_bonds_ppk/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/26/bond_myths_busted/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/26/best_bond_film/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/26/london_cab_electric_emission_free_byd/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/25/bond_behind_the_scenes/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/24/bodyform_video/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/23/flossie_at_risk/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/23/police_robot_back_on_duty/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/23/laptop_seeker_escapes_hydraulic_ram_in_garbage_truck/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/22/sheep_need_twitter/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/22/mi5_gchq_apprenticeships/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/19/wi_fi_tumble/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/19/wikipedia_women_editathon/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/19/james_bond_cars/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/19/quotw_ending_october_19/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/19/best_bond/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/19/netease_pig_farm_china/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/19/red_backs_terrorize_japan/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/18/bodyform_video/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/18/dance_competition_explains_scientific_research/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/16/new_q_memo/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/15/cinderella_glass_slipper/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/15/sky_jibe/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/12/european_union_wins_the_2012_nobel_peace_prize/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/12/reg_logos/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/12/quotw_ending_october_12/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/12/winston_churchill_digital_archive/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/12/bond_villain_poll/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/12/james_bond_villain/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/11/foxconn_injured_worker_treatment/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/11/apple_logo_russian_christians/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/09/new_zealand_hobbit_money/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/08/baumgartner_jump/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/08/view_to_a_kill/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/05/world_of_warcraft_colleen_lachowicz_attacked_by_republicans/",
		"http://go.theregister.com/feed/www.theregister.co.uk/2012/10/05/steve_jobs_is_still_dead/",
	}
	doc := Doc{}
	doc.Decode(testfeed.Atom)
	if len(doc.Entry) == 0 {
		t.Error("No entries found")
	}
	for i, e := range doc.Entry {
		if e.Link[0].Href != urls[i] {
			t.Errorf("URL Mismatch:\nGOT: %s\nEXP: %s", e.Link[0].Href, urls[i])
		}
	}
}
