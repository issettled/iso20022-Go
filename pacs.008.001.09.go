// IsSettled iso20022

package schema

import (
	"time"
)

// Document
type Document *Document

// AccountIdentification4Choice
type AccountIdentification4Choice struct {
	IBAN string                         `xml:"IBAN"`
	Othr *GenericAccountIdentification1 `xml:"Othr"`
}

// AccountSchemeName1Choice
type AccountSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ActiveCurrencyAndAmountSimpleType
type ActiveCurrencyAndAmountSimpleType float64

// ActiveCurrencyAndAmount
type ActiveCurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr"`
	Value   float64 `xml:",chardata"`
}

// ActiveCurrencyCode
type ActiveCurrencyCode string

// ActiveOrHistoricCurrencyAndAmountSimpleType
type ActiveOrHistoricCurrencyAndAmountSimpleType float64

// ActiveOrHistoricCurrencyAndAmount
type ActiveOrHistoricCurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr"`
	Value   float64 `xml:",chardata"`
}

// ActiveOrHistoricCurrencyCode
type ActiveOrHistoricCurrencyCode string

// AddressType2Code
type AddressType2Code string

// AddressType3Choice
type AddressType3Choice struct {
	Cd    string                   `xml:"Cd"`
	Prtry *GenericIdentification30 `xml:"Prtry"`
}

// AnyBICDec2014Identifier
type AnyBICDec2014Identifier string

// BICFIDec2014Identifier
type BICFIDec2014Identifier string

// BaseOneRate
type BaseOneRate float64

// BatchBookingIndicator
type BatchBookingIndicator bool

// BranchAndFinancialInstitutionIdentification6
type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId *FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                          `xml:"BrnchId"`
}

// BranchData3
type BranchData3 struct {
	Id      string           `xml:"Id"`
	LEI     string           `xml:"LEI"`
	Nm      string           `xml:"Nm"`
	PstlAdr *PostalAddress24 `xml:"PstlAdr"`
}

// CashAccount38
type CashAccount38 struct {
	Id   *AccountIdentification4Choice `xml:"Id"`
	Tp   *CashAccountType2Choice       `xml:"Tp"`
	Ccy  string                        `xml:"Ccy"`
	Nm   string                        `xml:"Nm"`
	Prxy *ProxyAccountIdentification1  `xml:"Prxy"`
}

// CashAccountType2Choice
type CashAccountType2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// CategoryPurpose1Choice
type CategoryPurpose1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ChargeBearerType1Code
type ChargeBearerType1Code string

// Charges7
type Charges7 struct {
	Amt *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// ClearingChannel2Code
type ClearingChannel2Code string

// ClearingSystemIdentification2Choice
type ClearingSystemIdentification2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ClearingSystemIdentification3Choice
type ClearingSystemIdentification3Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ClearingSystemMemberIdentification2
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId"`
	MmbId    string                               `xml:"MmbId"`
}

// Contact4
type Contact4 struct {
	NmPrfx    string           `xml:"NmPrfx"`
	Nm        string           `xml:"Nm"`
	PhneNb    string           `xml:"PhneNb"`
	MobNb     string           `xml:"MobNb"`
	FaxNb     string           `xml:"FaxNb"`
	EmailAdr  string           `xml:"EmailAdr"`
	EmailPurp string           `xml:"EmailPurp"`
	JobTitl   string           `xml:"JobTitl"`
	Rspnsblty string           `xml:"Rspnsblty"`
	Dept      string           `xml:"Dept"`
	Othr      []*OtherContact1 `xml:"Othr"`
	PrefrdMtd string           `xml:"PrefrdMtd"`
}

// CountryCode
type CountryCode string

// CreditDebitCode
type CreditDebitCode string

// CreditTransferMandateData1
type CreditTransferMandateData1 struct {
	MndtId       string                     `xml:"MndtId"`
	Tp           *MandateTypeInformation2   `xml:"Tp"`
	DtOfSgntr    string                     `xml:"DtOfSgntr"`
	DtOfVrfctn   string                     `xml:"DtOfVrfctn"`
	ElctrncSgntr []byte                     `xml:"ElctrncSgntr"`
	FrstPmtDt    string                     `xml:"FrstPmtDt"`
	FnlPmtDt     string                     `xml:"FnlPmtDt"`
	Frqcy        *Frequency36Choice         `xml:"Frqcy"`
	Rsn          *MandateSetupReason1Choice `xml:"Rsn"`
}

// CreditTransferTransaction43
type CreditTransferTransaction43 struct {
	PmtId             *PaymentIdentification13                      `xml:"PmtId"`
	PmtTpInf          *PaymentTypeInformation28                     `xml:"PmtTpInf"`
	IntrBkSttlmAmt    *ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt     string                                        `xml:"IntrBkSttlmDt"`
	SttlmPrty         string                                        `xml:"SttlmPrty"`
	SttlmTmIndctn     *SettlementDateTimeIndication1                `xml:"SttlmTmIndctn"`
	SttlmTmReq        *SettlementTimeRequest2                       `xml:"SttlmTmReq"`
	AccptncDtTm       string                                        `xml:"AccptncDtTm"`
	PoolgAdjstmntDt   string                                        `xml:"PoolgAdjstmntDt"`
	InstdAmt          *ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt"`
	XchgRate          float64                                       `xml:"XchgRate"`
	ChrgBr            string                                        `xml:"ChrgBr"`
	ChrgsInf          []*Charges7                                   `xml:"ChrgsInf"`
	MndtRltdInf       *CreditTransferMandateData1                   `xml:"MndtRltdInf"`
	PrvsInstgAgt1     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1"`
	PrvsInstgAgt1Acct *CashAccount38                                `xml:"PrvsInstgAgt1Acct"`
	PrvsInstgAgt2     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2"`
	PrvsInstgAgt2Acct *CashAccount38                                `xml:"PrvsInstgAgt2Acct"`
	PrvsInstgAgt3     *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3"`
	PrvsInstgAgt3Acct *CashAccount38                                `xml:"PrvsInstgAgt3Acct"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`
	IntrmyAgt1        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1"`
	IntrmyAgt1Acct    *CashAccount38                                `xml:"IntrmyAgt1Acct"`
	IntrmyAgt2        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2"`
	IntrmyAgt2Acct    *CashAccount38                                `xml:"IntrmyAgt2Acct"`
	IntrmyAgt3        *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3"`
	IntrmyAgt3Acct    *CashAccount38                                `xml:"IntrmyAgt3Acct"`
	UltmtDbtr         *PartyIdentification135                       `xml:"UltmtDbtr"`
	InitgPty          *PartyIdentification135                       `xml:"InitgPty"`
	Dbtr              *PartyIdentification135                       `xml:"Dbtr"`
	DbtrAcct          *CashAccount38                                `xml:"DbtrAcct"`
	DbtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	DbtrAgtAcct       *CashAccount38                                `xml:"DbtrAgtAcct"`
	CdtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	CdtrAgtAcct       *CashAccount38                                `xml:"CdtrAgtAcct"`
	Cdtr              *PartyIdentification135                       `xml:"Cdtr"`
	CdtrAcct          *CashAccount38                                `xml:"CdtrAcct"`
	UltmtCdtr         *PartyIdentification135                       `xml:"UltmtCdtr"`
	InstrForCdtrAgt   []*InstructionForCreditorAgent3               `xml:"InstrForCdtrAgt"`
	InstrForNxtAgt    []*InstructionForNextAgent1                   `xml:"InstrForNxtAgt"`
	Purp              *Purpose2Choice                               `xml:"Purp"`
	RgltryRptg        []*RegulatoryReporting3                       `xml:"RgltryRptg"`
	Tax               *TaxInformation8                              `xml:"Tax"`
	RltdRmtInf        []*RemittanceLocation7                        `xml:"RltdRmtInf"`
	RmtInf            *RemittanceInformation16                      `xml:"RmtInf"`
	SplmtryData       []*SupplementaryData1                         `xml:"SplmtryData"`
}

// CreditorReferenceInformation2
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp"`
	Ref string                  `xml:"Ref"`
}

// CreditorReferenceType1Choice
type CreditorReferenceType1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// CreditorReferenceType2
type CreditorReferenceType2 struct {
	CdOrPrtry *CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      string                        `xml:"Issr"`
}

// DateAndPlaceOfBirth1
type DateAndPlaceOfBirth1 struct {
	BirthDt     string `xml:"BirthDt"`
	PrvcOfBirth string `xml:"PrvcOfBirth"`
	CityOfBirth string `xml:"CityOfBirth"`
	CtryOfBirth string `xml:"CtryOfBirth"`
}

// DatePeriod2
type DatePeriod2 struct {
	FrDt string `xml:"FrDt"`
	ToDt string `xml:"ToDt"`
}

// DecimalNumber
type DecimalNumber float64

// DiscountAmountAndType1
type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice         `xml:"Tp"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// DiscountAmountType1Choice
type DiscountAmountType1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// DocumentAdjustment1
type DocumentAdjustment1 struct {
	Amt       *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd string                             `xml:"CdtDbtInd"`
	Rsn       string                             `xml:"Rsn"`
	AddtlInf  string                             `xml:"AddtlInf"`
}

// DocumentLineIdentification1
type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp"`
	Nb     string             `xml:"Nb"`
	RltdDt string             `xml:"RltdDt"`
}

// DocumentLineInformation1
type DocumentLineInformation1 struct {
	Id   []*DocumentLineIdentification1 `xml:"Id"`
	Desc string                         `xml:"Desc"`
	Amt  *RemittanceAmount3             `xml:"Amt"`
}

// DocumentLineType1
type DocumentLineType1 struct {
	CdOrPrtry *DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      string                   `xml:"Issr"`
}

// DocumentLineType1Choice
type DocumentLineType1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// DocumentType3Code
type DocumentType3Code string

// DocumentType6Code
type DocumentType6Code string

// Exact2NumericText
type Exact2NumericText string

// Exact4AlphaNumericText
type Exact4AlphaNumericText string

// ExternalAccountIdentification1Code
type ExternalAccountIdentification1Code string

// ExternalCashAccountType1Code
type ExternalCashAccountType1Code string

// ExternalCashClearingSystem1Code
type ExternalCashClearingSystem1Code string

// ExternalCategoryPurpose1Code
type ExternalCategoryPurpose1Code string

// ExternalClearingSystemIdentification1Code
type ExternalClearingSystemIdentification1Code string

// ExternalCreditorAgentInstruction1Code
type ExternalCreditorAgentInstruction1Code string

// ExternalDiscountAmountType1Code
type ExternalDiscountAmountType1Code string

// ExternalDocumentLineType1Code
type ExternalDocumentLineType1Code string

// ExternalFinancialInstitutionIdentification1Code
type ExternalFinancialInstitutionIdentification1Code string

// ExternalGarnishmentType1Code
type ExternalGarnishmentType1Code string

// ExternalLocalInstrument1Code
type ExternalLocalInstrument1Code string

// ExternalMandateSetupReason1Code
type ExternalMandateSetupReason1Code string

// ExternalOrganisationIdentification1Code
type ExternalOrganisationIdentification1Code string

// ExternalPersonIdentification1Code
type ExternalPersonIdentification1Code string

// ExternalProxyAccountType1Code
type ExternalProxyAccountType1Code string

// ExternalPurpose1Code
type ExternalPurpose1Code string

// ExternalServiceLevel1Code
type ExternalServiceLevel1Code string

// ExternalTaxAmountType1Code
type ExternalTaxAmountType1Code string

// FIToFICustomerCreditTransferV09
type FIToFICustomerCreditTransferV09 struct {
	GrpHdr      *GroupHeader93                 `xml:"GrpHdr"`
	CdtTrfTxInf []*CreditTransferTransaction43 `xml:"CdtTrfTxInf"`
	SplmtryData []*SupplementaryData1          `xml:"SplmtryData"`
}

// FinancialIdentificationSchemeName1Choice
type FinancialIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// FinancialInstitutionIdentification18
type FinancialInstitutionIdentification18 struct {
	BICFI       string                               `xml:"BICFI"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	LEI         string                               `xml:"LEI"`
	Nm          string                               `xml:"Nm"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr"`
}

// Frequency36Choice
type Frequency36Choice struct {
	Tp     string               `xml:"Tp"`
	Prd    *FrequencyPeriod1    `xml:"Prd"`
	PtInTm *FrequencyAndMoment1 `xml:"PtInTm"`
}

// Frequency6Code
type Frequency6Code string

// FrequencyAndMoment1
type FrequencyAndMoment1 struct {
	Tp     string `xml:"Tp"`
	PtInTm string `xml:"PtInTm"`
}

// FrequencyPeriod1
type FrequencyPeriod1 struct {
	Tp        string  `xml:"Tp"`
	CntPerPrd float64 `xml:"CntPerPrd"`
}

// Garnishment3
type Garnishment3 struct {
	Tp                *GarnishmentType1                  `xml:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr"`
	RefNb             string                             `xml:"RefNb"`
	Dt                string                             `xml:"Dt"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd"`
}

// GarnishmentType1
type GarnishmentType1 struct {
	CdOrPrtry *GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      string                  `xml:"Issr"`
}

// GarnishmentType1Choice
type GarnishmentType1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// GenericAccountIdentification1
type GenericAccountIdentification1 struct {
	Id      string                    `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                    `xml:"Issr"`
}

// GenericFinancialIdentification1
type GenericFinancialIdentification1 struct {
	Id      string                                    `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                    `xml:"Issr"`
}

// GenericIdentification30
type GenericIdentification30 struct {
	Id      string `xml:"Id"`
	Issr    string `xml:"Issr"`
	SchmeNm string `xml:"SchmeNm"`
}

// GenericOrganisationIdentification1
type GenericOrganisationIdentification1 struct {
	Id      string                                       `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                       `xml:"Issr"`
}

// GenericPersonIdentification1
type GenericPersonIdentification1 struct {
	Id      string                                 `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                 `xml:"Issr"`
}

// GroupHeader93
type GroupHeader93 struct {
	MsgId             string                                        `xml:"MsgId"`
	CreDtTm           string                                        `xml:"CreDtTm"`
	BtchBookg         bool                                          `xml:"BtchBookg"`
	NbOfTxs           string                                        `xml:"NbOfTxs"`
	CtrlSum           float64                                       `xml:"CtrlSum"`
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount                      `xml:"TtlIntrBkSttlmAmt"`
	IntrBkSttlmDt     string                                        `xml:"IntrBkSttlmDt"`
	SttlmInf          *SettlementInstruction7                       `xml:"SttlmInf"`
	PmtTpInf          *PaymentTypeInformation28                     `xml:"PmtTpInf"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`
}

// IBAN2007Identifier
type IBAN2007Identifier string

// ISODate
type ISODate string

// ISODateTime
type ISODateTime string

// ISOTime
type ISOTime time.Time

// Instruction4Code
type Instruction4Code string

// InstructionForCreditorAgent3
type InstructionForCreditorAgent3 struct {
	Cd       string `xml:"Cd"`
	InstrInf string `xml:"InstrInf"`
}

// InstructionForNextAgent1
type InstructionForNextAgent1 struct {
	Cd       string `xml:"Cd"`
	InstrInf string `xml:"InstrInf"`
}

// LEIIdentifier
type LEIIdentifier string

// LocalInstrument2Choice
type LocalInstrument2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// MandateClassification1Choice
type MandateClassification1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// MandateClassification1Code
type MandateClassification1Code string

// MandateSetupReason1Choice
type MandateSetupReason1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// MandateTypeInformation2
type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"SvcLvl"`
	LclInstrm *LocalInstrument2Choice       `xml:"LclInstrm"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"CtgyPurp"`
	Clssfctn  *MandateClassification1Choice `xml:"Clssfctn"`
}

// Max10KBinary
type Max10KBinary []byte

// Max10Text
type Max10Text string

// Max128Text
type Max128Text string

// Max140Text
type Max140Text string

// Max15NumericText
type Max15NumericText string

// Max16Text
type Max16Text string

// Max2048Text
type Max2048Text string

// Max34Text
type Max34Text string

// Max350Text
type Max350Text string

// Max35Text
type Max35Text string

// Max4Text
type Max4Text string

// Max70Text
type Max70Text string

// NameAndAddress16
type NameAndAddress16 struct {
	Nm  string           `xml:"Nm"`
	Adr *PostalAddress24 `xml:"Adr"`
}

// NamePrefix2Code
type NamePrefix2Code string

// Number
type Number float64

// OrganisationIdentification29
type OrganisationIdentification29 struct {
	AnyBIC string                                `xml:"AnyBIC"`
	LEI    string                                `xml:"LEI"`
	Othr   []*GenericOrganisationIdentification1 `xml:"Othr"`
}

// OrganisationIdentificationSchemeName1Choice
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// OtherContact1
type OtherContact1 struct {
	ChanlTp string `xml:"ChanlTp"`
	Id      string `xml:"Id"`
}

// Party38Choice
type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"OrgId"`
	PrvtId *PersonIdentification13       `xml:"PrvtId"`
}

// PartyIdentification135
type PartyIdentification135 struct {
	Nm        string           `xml:"Nm"`
	PstlAdr   *PostalAddress24 `xml:"PstlAdr"`
	Id        *Party38Choice   `xml:"Id"`
	CtryOfRes string           `xml:"CtryOfRes"`
	CtctDtls  *Contact4        `xml:"CtctDtls"`
}

// PaymentIdentification13
type PaymentIdentification13 struct {
	InstrId    string `xml:"InstrId"`
	EndToEndId string `xml:"EndToEndId"`
	TxId       string `xml:"TxId"`
	UETR       string `xml:"UETR"`
	ClrSysRef  string `xml:"ClrSysRef"`
}

// PaymentTypeInformation28
type PaymentTypeInformation28 struct {
	InstrPrty string                  `xml:"InstrPrty"`
	ClrChanl  string                  `xml:"ClrChanl"`
	SvcLvl    []*ServiceLevel8Choice  `xml:"SvcLvl"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp"`
}

// PercentageRate
type PercentageRate float64

// PersonIdentification13
type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth"`
	Othr            []*GenericPersonIdentification1 `xml:"Othr"`
}

// PersonIdentificationSchemeName1Choice
type PersonIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// PhoneNumber
type PhoneNumber string

// PostalAddress24
type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp"`
	Dept        string              `xml:"Dept"`
	SubDept     string              `xml:"SubDept"`
	StrtNm      string              `xml:"StrtNm"`
	BldgNb      string              `xml:"BldgNb"`
	BldgNm      string              `xml:"BldgNm"`
	Flr         string              `xml:"Flr"`
	PstBx       string              `xml:"PstBx"`
	Room        string              `xml:"Room"`
	PstCd       string              `xml:"PstCd"`
	TwnNm       string              `xml:"TwnNm"`
	TwnLctnNm   string              `xml:"TwnLctnNm"`
	DstrctNm    string              `xml:"DstrctNm"`
	CtrySubDvsn string              `xml:"CtrySubDvsn"`
	Ctry        string              `xml:"Ctry"`
	AdrLine     []string            `xml:"AdrLine"`
}

// PreferredContactMethod1Code
type PreferredContactMethod1Code string

// Priority2Code
type Priority2Code string

// Priority3Code
type Priority3Code string

// ProxyAccountIdentification1
type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp"`
	Id string                   `xml:"Id"`
}

// ProxyAccountType1Choice
type ProxyAccountType1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// Purpose2Choice
type Purpose2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ReferredDocumentInformation7
type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4      `xml:"Tp"`
	Nb       string                      `xml:"Nb"`
	RltdDt   string                      `xml:"RltdDt"`
	LineDtls []*DocumentLineInformation1 `xml:"LineDtls"`
}

// ReferredDocumentType3Choice
type ReferredDocumentType3Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ReferredDocumentType4
type ReferredDocumentType4 struct {
	CdOrPrtry *ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      string                       `xml:"Issr"`
}

// RegulatoryAuthority2
type RegulatoryAuthority2 struct {
	Nm   string `xml:"Nm"`
	Ctry string `xml:"Ctry"`
}

// RegulatoryReporting3
type RegulatoryReporting3 struct {
	DbtCdtRptgInd string                            `xml:"DbtCdtRptgInd"`
	Authrty       *RegulatoryAuthority2             `xml:"Authrty"`
	Dtls          []*StructuredRegulatoryReporting3 `xml:"Dtls"`
}

// RegulatoryReportingType1Code
type RegulatoryReportingType1Code string

// RemittanceAmount2
type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt"`
	DscntApldAmt      []*DiscountAmountAndType1          `xml:"DscntApldAmt"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt"`
	TaxAmt            []*TaxAmountAndType1               `xml:"TaxAmt"`
	AdjstmntAmtAndRsn []*DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt"`
}

// RemittanceAmount3
type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt"`
	DscntApldAmt      []*DiscountAmountAndType1          `xml:"DscntApldAmt"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt"`
	TaxAmt            []*TaxAmountAndType1               `xml:"TaxAmt"`
	AdjstmntAmtAndRsn []*DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt"`
}

// RemittanceInformation16
type RemittanceInformation16 struct {
	Ustrd []string                             `xml:"Ustrd"`
	Strd  []*StructuredRemittanceInformation16 `xml:"Strd"`
}

// RemittanceLocation7
type RemittanceLocation7 struct {
	RmtId       string                     `xml:"RmtId"`
	RmtLctnDtls []*RemittanceLocationData1 `xml:"RmtLctnDtls"`
}

// RemittanceLocationData1
type RemittanceLocationData1 struct {
	Mtd        string            `xml:"Mtd"`
	ElctrncAdr string            `xml:"ElctrncAdr"`
	PstlAdr    *NameAndAddress16 `xml:"PstlAdr"`
}

// RemittanceLocationMethod2Code
type RemittanceLocationMethod2Code string

// ServiceLevel8Choice
type ServiceLevel8Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// SettlementDateTimeIndication1
type SettlementDateTimeIndication1 struct {
	DbtDtTm string `xml:"DbtDtTm"`
	CdtDtTm string `xml:"CdtDtTm"`
}

// SettlementInstruction7
type SettlementInstruction7 struct {
	SttlmMtd             string                                        `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount38                                `xml:"SttlmAcct"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct"`
}

// SettlementMethod1Code
type SettlementMethod1Code string

// SettlementTimeRequest2
type SettlementTimeRequest2 struct {
	CLSTm  time.Time `xml:"CLSTm"`
	TillTm time.Time `xml:"TillTm"`
	FrTm   time.Time `xml:"FrTm"`
	RjctTm time.Time `xml:"RjctTm"`
}

// StructuredRegulatoryReporting3
type StructuredRegulatoryReporting3 struct {
	Tp   string                             `xml:"Tp"`
	Dt   string                             `xml:"Dt"`
	Ctry string                             `xml:"Ctry"`
	Cd   string                             `xml:"Cd"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	Inf  []string                           `xml:"Inf"`
}

// StructuredRemittanceInformation16
type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []*ReferredDocumentInformation7 `xml:"RfrdDocInf"`
	RfrdDocAmt  *RemittanceAmount2              `xml:"RfrdDocAmt"`
	CdtrRefInf  *CreditorReferenceInformation2  `xml:"CdtrRefInf"`
	Invcr       *PartyIdentification135         `xml:"Invcr"`
	Invcee      *PartyIdentification135         `xml:"Invcee"`
	TaxRmt      *TaxInformation7                `xml:"TaxRmt"`
	GrnshmtRmt  *Garnishment3                   `xml:"GrnshmtRmt"`
	AddtlRmtInf []string                        `xml:"AddtlRmtInf"`
}

// SupplementaryData1
type SupplementaryData1 struct {
	PlcAndNm string                      `xml:"PlcAndNm"`
	Envlp    *SupplementaryDataEnvelope1 `xml:"Envlp"`
}

// SupplementaryDataEnvelope1
type SupplementaryDataEnvelope1 struct {
}

// TaxAmount2
type TaxAmount2 struct {
	Rate         float64                            `xml:"Rate"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt"`
	Dtls         []*TaxRecordDetails2               `xml:"Dtls"`
}

// TaxAmountAndType1
type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice              `xml:"Tp"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// TaxAmountType1Choice
type TaxAmountType1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// TaxAuthorisation1
type TaxAuthorisation1 struct {
	Titl string `xml:"Titl"`
	Nm   string `xml:"Nm"`
}

// TaxInformation7
type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr"`
	Dbtr            *TaxParty2                         `xml:"Dbtr"`
	UltmtDbtr       *TaxParty2                         `xml:"UltmtDbtr"`
	AdmstnZone      string                             `xml:"AdmstnZone"`
	RefNb           string                             `xml:"RefNb"`
	Mtd             string                             `xml:"Mtd"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt"`
	Dt              string                             `xml:"Dt"`
	SeqNb           float64                            `xml:"SeqNb"`
	Rcrd            []*TaxRecord2                      `xml:"Rcrd"`
}

// TaxInformation8
type TaxInformation8 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr"`
	Dbtr            *TaxParty2                         `xml:"Dbtr"`
	AdmstnZone      string                             `xml:"AdmstnZone"`
	RefNb           string                             `xml:"RefNb"`
	Mtd             string                             `xml:"Mtd"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt"`
	Dt              string                             `xml:"Dt"`
	SeqNb           float64                            `xml:"SeqNb"`
	Rcrd            []*TaxRecord2                      `xml:"Rcrd"`
}

// TaxParty1
type TaxParty1 struct {
	TaxId  string `xml:"TaxId"`
	RegnId string `xml:"RegnId"`
	TaxTp  string `xml:"TaxTp"`
}

// TaxParty2
type TaxParty2 struct {
	TaxId   string             `xml:"TaxId"`
	RegnId  string             `xml:"RegnId"`
	TaxTp   string             `xml:"TaxTp"`
	Authstn *TaxAuthorisation1 `xml:"Authstn"`
}

// TaxPeriod2
type TaxPeriod2 struct {
	Yr     string       `xml:"Yr"`
	Tp     string       `xml:"Tp"`
	FrToDt *DatePeriod2 `xml:"FrToDt"`
}

// TaxRecord2
type TaxRecord2 struct {
	Tp       string      `xml:"Tp"`
	Ctgy     string      `xml:"Ctgy"`
	CtgyDtls string      `xml:"CtgyDtls"`
	DbtrSts  string      `xml:"DbtrSts"`
	CertId   string      `xml:"CertId"`
	FrmsCd   string      `xml:"FrmsCd"`
	Prd      *TaxPeriod2 `xml:"Prd"`
	TaxAmt   *TaxAmount2 `xml:"TaxAmt"`
	AddtlInf string      `xml:"AddtlInf"`
}

// TaxRecordDetails2
type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                        `xml:"Prd"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// TaxRecordPeriod1Code
type TaxRecordPeriod1Code string

// TrueFalseIndicator
type TrueFalseIndicator bool

// UUIDv4Identifier
type UUIDv4Identifier string
