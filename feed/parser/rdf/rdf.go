package rdf

import (
	"encoding/xml"
	"github.com/300brand/coverage/feed/parser/charset"
	"github.com/300brand/coverage/feed/parser/decoder"
	"github.com/300brand/coverage/feed/parser/time"
)

type Doc struct {
	XMLName xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
	Channel Channel  `xml:"channel"`
	Item    []Item   `xml:"item"`
}

type Channel struct {
	// About                 string    `xml:"about,attr"`
	// Abstract              string    `xml:"abstract"`
	// AccessRights          string    `xml:"accessRights"`
	// AccrualMethod         string    `xml:"accrualMethod"`
	// AccrualPeriodicity    string    `xml:"accrualPeriodicity"`
	// AccrualPolicy         string    `xml:"accrualPolicy"`
	// Alternative           string    `xml:"alternative"`
	// Audience              string    `xml:"audience"`
	// Available             string    `xml:"available"`
	// BibliographicCitation string    `xml:"bibliographicCitation"`
	// ConformsTo            string    `xml:"conformsTo"`
	// Contributor           string    `xml:"contributor"`
	// Coverage              string    `xml:"coverage"`
	// Created               string    `xml:"created"`
	// Creator               string    `xml:"creator"`
	// Date                  time.Time `xml:"date"`
	// DateAccepted          string    `xml:"dateAccepted"`
	// DateCopyrighted       string    `xml:"dateCopyrighted"`
	// DateSubmitted         string    `xml:"dateSubmitted"`
	// Description           string    `xml:"description"`
	// EducationLevel        string    `xml:"educationLevel"`
	// Extent                string    `xml:"extent"`
	// Format                string    `xml:"format"`
	// HasFormat             string    `xml:"hasFormat"`
	// HasPart               string    `xml:"hasPart"`
	// HasVersion            string    `xml:"hasVersion"`
	// Identifier            string    `xml:"identifier"`
	// InstructionalMethod   string    `xml:"instructionalMethod"`
	// IsFormatOf            string    `xml:"isFormatOf"`
	// IsPartOf              string    `xml:"isPartOf"`
	// IsReferencedBy        string    `xml:"isReferencedBy"`
	// IsReplacedBy          string    `xml:"isReplacedBy"`
	// IsRequiredBy          string    `xml:"isRequiredBy"`
	// Issued                string    `xml:"issued"`
	// IsVersionOf           string    `xml:"isVersionOf"`
	// Items                 Items     `xml:"items"`
	// Language              string    `xml:"language"`
	// License               string    `xml:"license"`
	// Link                  string    `xml:"link"`
	// Mediator              string    `xml:"mediator"`
	// Medium                string    `xml:"medium"`
	// Modified              string    `xml:"modified"`
	// Provenance            string    `xml:"provenance"`
	// Publisher             string    `xml:"publisher"`
	// References            string    `xml:"references"`
	// Relation              string    `xml:"relation"`
	// Replaces              string    `xml:"replaces"`
	// Requires              string    `xml:"requires"`
	// Rights                string    `xml:"rights"`
	// RightsHolder          string    `xml:"rightsHolder"`
	// Source                string    `xml:"source"`
	// Spatial               string    `xml:"spatial"`
	// Subject               string    `xml:"subject"`
	// TableOfContents       string    `xml:"tableOfContents"`
	// Temporal              string    `xml:"temporal"`
	Title string `xml:"title"`
	// Type                  string    `xml:"type"`
	// Valid                 string    `xml:"valid"`
}

type Item struct {
	// About                 string    `xml:"about,attr"`
	// Abstract              string    `xml:"abstract"`
	// AccessRights          string    `xml:"accessRights"`
	// AccrualMethod         string    `xml:"accrualMethod"`
	// AccrualPeriodicity    string    `xml:"accrualPeriodicity"`
	// AccrualPolicy         string    `xml:"accrualPolicy"`
	// Alternative           string    `xml:"alternative"`
	// Audience              string    `xml:"audience"`
	// Available             string    `xml:"available"`
	// BibliographicCitation string    `xml:"bibliographicCitation"`
	// ConformsTo            string    `xml:"conformsTo"`
	// Contributor           string    `xml:"contributor"`
	// Coverage              string    `xml:"coverage"`
	// Created               string    `xml:"created"`
	// Creator               string    `xml:"creator"`
	Date time.Time `xml:"date"`
	// DateAccepted          string    `xml:"dateAccepted"`
	// DateCopyrighted       string    `xml:"dateCopyrighted"`
	// DateSubmitted         string    `xml:"dateSubmitted"`
	// Description           string    `xml:"description"`
	// EducationLevel        string    `xml:"educationLevel"`
	// Extent                string    `xml:"extent"`
	// Format                string    `xml:"format"`
	// HasFormat             string    `xml:"hasFormat"`
	// HasPart               string    `xml:"hasPart"`
	// HasVersion            string    `xml:"hasVersion"`
	// Identifier            string    `xml:"identifier"`
	// InstructionalMethod   string    `xml:"instructionalMethod"`
	// IsFormatOf            string    `xml:"isFormatOf"`
	// IsPartOf              string    `xml:"isPartOf"`
	// IsReferencedBy        string    `xml:"isReferencedBy"`
	// IsReplacedBy          string    `xml:"isReplacedBy"`
	// IsRequiredBy          string    `xml:"isRequiredBy"`
	// Issued                string    `xml:"issued"`
	// IsVersionOf           string    `xml:"isVersionOf"`
	// Language              string    `xml:"language"`
	// License               string    `xml:"license"`
	Link string `xml:"link"`
	// Mediator              string    `xml:"mediator"`
	// Medium                string    `xml:"medium"`
	// Modified              string    `xml:"modified"`
	// Provenance            string    `xml:"provenance"`
	// Publisher             string    `xml:"publisher"`
	// References            string    `xml:"references"`
	// Relation              string    `xml:"relation"`
	// Replaces              string    `xml:"replaces"`
	// Requires              string    `xml:"requires"`
	// Rights                string    `xml:"rights"`
	// RightsHolder          string    `xml:"rightsHolder"`
	// Source                string    `xml:"source"`
	// Spatial               string    `xml:"spatial"`
	// Subject               string    `xml:"subject"`
	// TableOfContents       string    `xml:"tableOfContents"`
	// Temporal              string    `xml:"temporal"`
	Title string `xml:"title"`
	// Type                  string    `xml:"type"`
	// Valid                 string    `xml:"valid"`
}

type Items struct {
	Seq Seq `xml:"Seq"`
}

type Seq struct {
	Li []Li `xml:"li"`
}

type Li struct {
	Resource string `xml:"resource,attr"`
}

func init() {
	decoder.RegisterDecoder("rdf", &Doc{})
}

func (doc Doc) New() decoder.Decoder {
	return &Doc{}
}

func (doc *Doc) Decode(data []byte) error {
	return charset.TryAll(doc, data)
}
