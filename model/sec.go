package model

import (
    "fmt"
    "regexp"

    "github.com/shopspring/decimal"

    "go.tradeforge.dev/fmp/pkg/types"
)

type ListInsiderTradesParams struct {
    Date *types.Date `query:"date,omitempty"`
    Page *uint       `query:"page,omitempty"`
}

type ListInsiderTradesResponse []InsiderTrade

type InsiderTrade struct {
    Symbol                   string          `json:"symbol"`
    FormType                 SECFormType     `json:"formType"`
    CompanyCIK               string          `json:"companyCik"`
    ReportingCIK             string          `json:"reportingCik"`
    TransactionDate          types.Date      `json:"transactionDate"`
    TransactionType          string          `json:"transactionType"`
    SecuritiesOwned          decimal.Decimal `json:"securitiesOwned"`
    SecuritiesTransacted     decimal.Decimal `json:"securitiesTransacted"`
    OwnerName                string          `json:"reportingName"`
    OwnerType                string          `json:"typeOfOwner"`
    AcquisitionOrDisposition string          `json:"acquisitionOrDisposition"`
    Price                    decimal.Decimal `json:"price"`
    FilingDate               types.Date      `json:"filingDate"`
}

type ListSECFilingsRSSFeedParams struct {
    Type  SECFormType `query:"type" validate:"required"`
    Since *types.Date `query:"from,omitempty"`
    Until *types.Date `query:"to,omitempty"`
    Limit *uint       `query:"limit,omitempty"`
}

type ListSECFilingsRSSFeedResponse []SECFiling

type SECFiling struct {
    Symbol     string         `json:"symbol"`
    CIK        string         `json:"cik"`
    Type       SECFilingType  `json:"type"`
    Link       string         `json:"link"`
    FinalLink  string         `json:"finalLink"`
    FiledAt    types.DateTime `json:"filingDate"`
    AcceptedAt types.DateTime `json:"acceptedDate"`
}

type SECFilingType struct {
    Form          SECFormType            `json:"type"`
    Specification SECFilingSpecification `json:"specification"`
}

func (t *SECFilingType) UnmarshalJSON(data []byte) error {
    formType := SECFormType(data)
    if err := formType.Validate(); err != nil {
        return fmt.Errorf("invalid SEC filing type: %w", err)
    }
    *t = SECFilingType{
        Form:          formType.Name(),
        Specification: formType.Specification(),
    }
    return nil
}

// SECFormType represents an SEC filing type.
type SECFormType string

const (
    // Form10K is the annual report with comprehensive financial data.
    Form10K SECFormType = "10-K"
    // Form10Q is the quarterly report with unaudited financial data.
    Form10Q SECFormType = "10-Q"
    // Form8K is the report of significant events or changes.
    Form8K SECFormType = "8-K"
    // Form6K is the report of foreign private issuers.
    Form6K SECFormType = "6-K"

    // FormS1 is the initial registration of securities (e.g., IPO).
    FormS1 SECFormType = "S-1"
    // FormS3 is the simplified securities registration for eligible companies.
    FormS3 SECFormType = "S-3"
    // FormS4 is the registration of securities for mergers or acquisitions.
    FormS4 SECFormType = "S-4"
    // FormS8 is the registration of securities for employee stock or benefit plans.
    FormS8 SECFormType = "S-8"

    // Schedule13D is the report filed by investors owning more than 5% of a company’s stock.
    Schedule13D SECFormType = "SCHEDULE 13D"
    // Schedule13G is the report for passive investors owning more than 5% of a company’s stock.
    Schedule13G SECFormType = "SCHEDULE 13G"
    // Form3 is the initial statement of beneficial ownership for corporate insiders.
    Form3 SECFormType = "3"
    // Form4 is the statement of changes in beneficial ownership for corporate insiders.
    Form4 SECFormType = "4"
    // Form5 is the annual statement of changes in beneficial ownership.
    Form5 SECFormType = "5"

    // FormDEF14A is the definitive proxy statement for shareholder meetings.
    FormDEF14A SECFormType = "DEF 14A"
    // FormDEF14C is the definitive information statement for shareholder meetings without proxies.
    FormDEF14C SECFormType = "DEF 14C"
    // FormPRE14A is the preliminary proxy statement for shareholder meetings.
    FormPRE14A SECFormType = "PRE 14A"
    // FormPRE14C is the preliminary information statement for shareholder meetings without proxies.
    FormPRE14C SECFormType = "PRE 14C"
    // Schedule14D9 is the recommendation statement in response to a tender offer.
    Schedule14D9 SECFormType = "SCHEDULE 14D-9"
    // Form10 is the registration of a class of securities under the Exchange Act.
    Form10 SECFormType = "10"

    // Form497 is the filing for definitive prospectus materials.
    Form497 SECFormType = "497"
    // Form497K is the key information summary of a mutual fund prospectus.
    Form497K SECFormType = "497K"
    // Form497J is the certification of compliance with SEC requirements.
    Form497J SECFormType = "497J"
    // Form424B1 is the preliminary prospectus with pricing information.
    Form424B1 SECFormType = "424"
)

func (t SECFormType) Name() SECFormType {
    switch {
    case t.IsAmendment():
        r := regexp.MustCompile(`(/A)+`)
        return SECFormType(r.ReplaceAllString(string(t), ""))
    case t.IsProspectus():
        r := regexp.MustCompile(`[B-J]`)
        return SECFormType(r.ReplaceAllString(string(t), ""))
    default:
        return t
    }
}

func (t SECFormType) Validate() error {
    switch t {
    case Form10K, Form10Q, Form8K, Form6K, FormS1, FormS3, FormS4, FormS8, Schedule13D, Schedule13G, Form3, Form4, Form5, FormDEF14A, FormDEF14C, FormPRE14A, FormPRE14C, Schedule14D9, Form10, Form497, Form497K, Form497J, Form424B1:
        return nil
    default:
        return fmt.Errorf("invalid SEC filing type: %s", t)
    }
}

func (t SECFormType) String() string {
    return string(t)
}

func (t SECFormType) Specification() SECFilingSpecification {
    switch {
    case t.IsAmendment():
        return Amendment
    case t.IsRegistration():
        return Registration
    case t.IsProxy():
        return Proxy
    case t.IsOwnership():
        return Ownership
    case t.IsProspectus():
        return Prospectus
    case t.IsSchedule():
        return Schedule
    case t.IsEarnings():
        return Earnings
    default:
        return Other
    }
}

func (t SECFormType) IsAmendment() bool {
    m, err := regexp.Match(`.*(/A)+`, []byte(t))
    if err != nil {
        return false
    }
    return m
}

func (t SECFormType) IsRegistration() bool {
    m, err := regexp.Match(`S-\d`, []byte(t))
    if err != nil {
        return false
    }
    return m
}

func (t SECFormType) IsProxy() bool {
    m, err := regexp.Match(`DEF 14[A-C]`, []byte(t))
    if err != nil {
        return false
    }
    return m
}

func (t SECFormType) IsOwnership() bool {
    m, err := regexp.Match(`\d`, []byte(t))
    if err != nil {
        return false
    }
    return m
}

func (t SECFormType) IsProspectus() bool {
    m, err := regexp.Match(`424[B-J]`, []byte(t))
    if err != nil {
        return false
    }
    return m
}

func (t SECFormType) IsSchedule() bool {
    m, err := regexp.Match(`SCHEDULE 13[D-G]`, []byte(t))
    if err != nil {
        return false
    }
    return m
}

func (t SECFormType) IsEarnings() bool {
    m, err := regexp.Match(`10-[KQ]`, []byte(t))
    if err != nil {
        return false
    }
    return m
}

type SECFilingSpecification string

const (
    // Amendment is the first amendment to a filing.
    Amendment SECFilingSpecification = "Amendment"
    // Registration is the initial registration of securities.
    Registration SECFilingSpecification = "Registration"
    // Proxy is the definitive proxy statement for shareholder meetings.
    Proxy SECFilingSpecification = "Proxy"
    // Ownership is the statement of changes in beneficial ownership.
    Ownership SECFilingSpecification = "Ownership"
    // Prospectus is the definitive prospectus for securities offerings.
    Prospectus SECFilingSpecification = "Prospectus"
    // Schedule is the report filed by investors owning more than 5% of a company’s stock.
    Schedule SECFilingSpecification = "Schedule"
    // Earnings is the annual or quarterly report with financial data.
    Earnings SECFilingSpecification = "Earnings"
    // Other is a catch-all for other filing types.
    Other SECFilingSpecification = "Other"
)
