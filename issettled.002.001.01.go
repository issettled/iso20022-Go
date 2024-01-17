// IsSettled iso20022

package schema

// Document
type Document *Document

// Max35Text
type Max35Text string

// ISODateTime
type ISODateTime string

// BatchBookingIndicator
type BatchBookingIndicator bool

// Max15NumericText
type Max15NumericText string

// DecimalNumber
type DecimalNumber float64

// ActiveCurrencyAndAmountSimpleType
type ActiveCurrencyAndAmountSimpleType float64

// ActiveCurrencyAndAmount
type ActiveCurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr"`
	Value   float64 `xml:",chardata"`
}

// ActiveCurrencyCode
type ActiveCurrencyCode string

// ISODate
type ISODate string

// SettlementInstruction4
type SettlementInstruction4 struct {
	SttlmMtd             string                                        `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount24                                `xml:"SttlmAcct"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt"`
	InstgRmbrsmntAgtAcct *CashAccount24                                `xml:"InstgRmbrsmntAgtAcct"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt"`
	InstdRmbrsmntAgtAcct *CashAccount24                                `xml:"InstdRmbrsmntAgtAcct"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt"`
	ThrdRmbrsmntAgtAcct  *CashAccount24                                `xml:"ThrdRmbrsmntAgtAcct"`
}

// SettlementMethod1Code
type SettlementMethod1Code string

// ClearingSystemIdentification3Choice
type ClearingSystemIdentification3Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalCashClearingSystem1Code
type ExternalCashClearingSystem1Code string

// CashAccount24
type CashAccount24 struct {
	Id  *AccountIdentification4Choice `xml:"Id"`
	Tp  *CashAccountType2Choice       `xml:"Tp"`
	Ccy string                        `xml:"Ccy"`
	Nm  string                        `xml:"Nm"`
}

// AccountIdentification4Choice
type AccountIdentification4Choice struct {
	IBAN string                         `xml:"IBAN"`
	Othr *GenericAccountIdentification1 `xml:"Othr"`
}

// IBAN2007Identifier
type IBAN2007Identifier string

// GenericAccountIdentification1
type GenericAccountIdentification1 struct {
	Id      string                    `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                    `xml:"Issr"`
}

// Max34Text
type Max34Text string

// AccountSchemeName1Choice
type AccountSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalAccountIdentification1Code
type ExternalAccountIdentification1Code string

// CashAccountType2Choice
type CashAccountType2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalCashAccountType1Code
type ExternalCashAccountType1Code string

// ActiveOrHistoricCurrencyCode
type ActiveOrHistoricCurrencyCode string

// Max70Text
type Max70Text string

// BranchAndFinancialInstitutionIdentification5
type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId *FinancialInstitutionIdentification8 `xml:"FinInstnId"`
	BrnchId    *BranchData2                         `xml:"BrnchId"`
}

// FinancialInstitutionIdentification8
type FinancialInstitutionIdentification8 struct {
	BICFI       string                               `xml:"BICFI"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Nm          string                               `xml:"Nm"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr"`
}

// BICFIIdentifier
type BICFIIdentifier string

// ClearingSystemMemberIdentification2
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId"`
	MmbId    string                               `xml:"MmbId"`
}

// ClearingSystemIdentification2Choice
type ClearingSystemIdentification2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalClearingSystemIdentification1Code
type ExternalClearingSystemIdentification1Code string

// Max140Text
type Max140Text string

// PostalAddress6
type PostalAddress6 struct {
	AdrTp       string   `xml:"AdrTp"`
	Dept        string   `xml:"Dept"`
	SubDept     string   `xml:"SubDept"`
	StrtNm      string   `xml:"StrtNm"`
	BldgNb      string   `xml:"BldgNb"`
	PstCd       string   `xml:"PstCd"`
	TwnNm       string   `xml:"TwnNm"`
	CtrySubDvsn string   `xml:"CtrySubDvsn"`
	Ctry        string   `xml:"Ctry"`
	AdrLine     []string `xml:"AdrLine"`
}

// AddressType2Code
type AddressType2Code string

// Max16Text
type Max16Text string

// CountryCode
type CountryCode string

// GenericFinancialIdentification1
type GenericFinancialIdentification1 struct {
	Id      string                                    `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                    `xml:"Issr"`
}

// FinancialIdentificationSchemeName1Choice
type FinancialIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalFinancialInstitutionIdentification1Code
type ExternalFinancialInstitutionIdentification1Code string

// BranchData2
type BranchData2 struct {
	Id      string          `xml:"Id"`
	Nm      string          `xml:"Nm"`
	PstlAdr *PostalAddress6 `xml:"PstlAdr"`
}

// PaymentTypeInformation21
type PaymentTypeInformation21 struct {
	InstrPrty string                  `xml:"InstrPrty"`
	ClrChanl  string                  `xml:"ClrChanl"`
	SvcLvl    *ServiceLevel8Choice    `xml:"SvcLvl"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp"`
}

// Priority2Code
type Priority2Code string

// ClearingChannel2Code
type ClearingChannel2Code string

// ServiceLevel8Choice
type ServiceLevel8Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalServiceLevel1Code
type ExternalServiceLevel1Code string

// LocalInstrument2Choice
type LocalInstrument2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalLocalInstrument1Code
type ExternalLocalInstrument1Code string

// CategoryPurpose1Choice
type CategoryPurpose1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalCategoryPurpose1Code
type ExternalCategoryPurpose1Code string

// GroupHeader70
type GroupHeader70 struct {
	MsgId             string                                        `xml:"MsgId"`
	CreDtTm           string                                        `xml:"CreDtTm"`
	BtchBookg         bool                                          `xml:"BtchBookg"`
	NbOfTxs           string                                        `xml:"NbOfTxs"`
	CtrlSum           float64                                       `xml:"CtrlSum"`
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount                      `xml:"TtlIntrBkSttlmAmt"`
	IntrBkSttlmDt     string                                        `xml:"IntrBkSttlmDt"`
	SttlmInf          *SettlementInstruction4                       `xml:"SttlmInf"`
	PmtTpInf          *PaymentTypeInformation21                     `xml:"PmtTpInf"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt"`
}

// TransactionSettlementNotificationV01
type TransactionSettlementNotificationV01 struct {
	GrpHdr     *GroupHeader70                       `xml:"GrpHdr"`
	TxSttlmInf []*TransactionSettlementInformation1 `xml:"TxSttlmInf"`
}

// TransactionSettlementInformation1
type TransactionSettlementInformation1 struct {
	OrgnlGrpInf *OriginalGroupInformation3        `xml:"OrgnlGrpInf"`
	TxInf       []*PaymentTransactionInformation2 `xml:"TxInf"`
}

// PaymentTransactionInformation2
type PaymentTransactionInformation2 struct {
	SttlmId         string                         `xml:"SttlmId"`
	OrgnlInstrId    string                         `xml:"OrgnlInstrId"`
	OrgnlEndToEndId string                         `xml:"OrgnlEndToEndId"`
	OrgnlTxId       string                         `xml:"OrgnlTxId"`
	ActlSttlmAmt    *CurrencyAndAmount             `xml:"ActlSttlmAmt"`
	OrgnlTxRef      *OriginalTransactionReference1 `xml:"OrgnlTxRef"`
}

// CurrencyAndAmount
type CurrencyAndAmount struct {
	CcyAttr string `xml:"Ccy,attr"`
	*CurrencyAndAmountSimpleType
}

// OriginalGroupInformation3
type OriginalGroupInformation3 struct {
	OrgnlMsgId   string `xml:"OrgnlMsgId"`
	OrgnlMsgNmId string `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm string `xml:"OrgnlCreDtTm"`
}

// CurrencyAndAmountSimpleType
type CurrencyAndAmountSimpleType float64

// CurrencyCode
type CurrencyCode string

// OriginalTransactionReference1
type OriginalTransactionReference1 struct {
	IntrBkSttlmAmt *CurrencyAndAmount                            `xml:"IntrBkSttlmAmt"`
	Amt            *AmountType2Choice                            `xml:"Amt"`
	IntrBkSttlmDt  string                                        `xml:"IntrBkSttlmDt"`
	ReqdExctnDt    string                                        `xml:"ReqdExctnDt"`
	ReqdColltnDt   string                                        `xml:"ReqdColltnDt"`
	CdtrSchmeId    *PartyIdentification8                         `xml:"CdtrSchmeId"`
	SttlmInf       *SettlementInformation3                       `xml:"SttlmInf"`
	PmtTpInf       *PaymentTypeInformation6                      `xml:"PmtTpInf"`
	PmtMtd         string                                        `xml:"PmtMtd"`
	MndtRltdInf    *MandateRelatedInformation1                   `xml:"MndtRltdInf"`
	RmtInf         *RemittanceInformation1                       `xml:"RmtInf"`
	UltmtDbtr      *PartyIdentification8                         `xml:"UltmtDbtr"`
	Dbtr           *PartyIdentification8                         `xml:"Dbtr"`
	DbtrAcct       *CashAccount7                                 `xml:"DbtrAcct"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`
	DbtrAgtAcct    *CashAccount7                                 `xml:"DbtrAgtAcct"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt"`
	CdtrAgtAcct    *CashAccount7                                 `xml:"CdtrAgtAcct"`
	Cdtr           *PartyIdentification8                         `xml:"Cdtr"`
	CdtrAcct       *CashAccount7                                 `xml:"CdtrAcct"`
	UltmtCdtr      *PartyIdentification8                         `xml:"UltmtCdtr"`
}

// AmountType2Choice
type AmountType2Choice struct {
	InstdAmt *CurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  *EquivalentAmount  `xml:"EqvtAmt"`
}

// EquivalentAmount
type EquivalentAmount struct {
	Amt      *CurrencyAndAmount `xml:"Amt"`
	CcyOfTrf string             `xml:"CcyOfTrf"`
}

// PartyIdentification8
type PartyIdentification8 struct {
	Nm        string          `xml:"Nm"`
	PstlAdr   *PostalAddress1 `xml:"PstlAdr"`
	Id        *Party2Choice   `xml:"Id"`
	CtryOfRes string          `xml:"CtryOfRes"`
}

// PostalAddress1
type PostalAddress1 struct {
	AdrTp       string   `xml:"AdrTp"`
	AdrLine     []string `xml:"AdrLine"`
	StrtNm      string   `xml:"StrtNm"`
	BldgNb      string   `xml:"BldgNb"`
	PstCd       string   `xml:"PstCd"`
	TwnNm       string   `xml:"TwnNm"`
	CtrySubDvsn string   `xml:"CtrySubDvsn"`
	Ctry        string   `xml:"Ctry"`
}

// Party2Choice
type Party2Choice struct {
	OrgId  *OrganisationIdentification2 `xml:"OrgId"`
	PrvtId []*PersonIdentification3     `xml:"PrvtId"`
}

// OrganisationIdentification2
type OrganisationIdentification2 struct {
	BIC     string                  `xml:"BIC"`
	IBEI    string                  `xml:"IBEI"`
	BEI     string                  `xml:"BEI"`
	EANGLN  string                  `xml:"EANGLN"`
	USCHU   string                  `xml:"USCHU"`
	DUNS    string                  `xml:"DUNS"`
	BkPtyId string                  `xml:"BkPtyId"`
	TaxIdNb string                  `xml:"TaxIdNb"`
	PrtryId *GenericIdentification3 `xml:"PrtryId"`
}

// BICIdentifier
type BICIdentifier string

// IBEIIdentifier
type IBEIIdentifier string

// BEIIdentifier
type BEIIdentifier string

// EANGLNIdentifier
type EANGLNIdentifier string

// CHIPSUniversalIdentifier
type CHIPSUniversalIdentifier string

// DunsIdentifier
type DunsIdentifier string

// GenericIdentification3
type GenericIdentification3 struct {
	Id   string `xml:"Id"`
	Issr string `xml:"Issr"`
}

// PersonIdentification3
type PersonIdentification3 struct {
	DrvrsLicNb      string                  `xml:"DrvrsLicNb"`
	CstmrNb         string                  `xml:"CstmrNb"`
	SclSctyNb       string                  `xml:"SclSctyNb"`
	AlnRegnNb       string                  `xml:"AlnRegnNb"`
	PsptNb          string                  `xml:"PsptNb"`
	TaxIdNb         string                  `xml:"TaxIdNb"`
	IdntyCardNb     string                  `xml:"IdntyCardNb"`
	MplyrIdNb       string                  `xml:"MplyrIdNb"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth    `xml:"DtAndPlcOfBirth"`
	OthrId          *GenericIdentification4 `xml:"OthrId"`
	Issr            string                  `xml:"Issr"`
}

// DateAndPlaceOfBirth
type DateAndPlaceOfBirth struct {
	BirthDt     string `xml:"BirthDt"`
	PrvcOfBirth string `xml:"PrvcOfBirth"`
	CityOfBirth string `xml:"CityOfBirth"`
	CtryOfBirth string `xml:"CtryOfBirth"`
}

// GenericIdentification4
type GenericIdentification4 struct {
	Id   string `xml:"Id"`
	IdTp string `xml:"IdTp"`
}

// SettlementInformation3
type SettlementInformation3 struct {
	SttlmMtd             string                                        `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount7                                 `xml:"SttlmAcct"`
	ClrSys               *ClearingSystemIdentification1Choice          `xml:"ClrSys"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"InstgRmbrsmntAgt"`
	InstgRmbrsmntAgtAcct *CashAccount7                                 `xml:"InstgRmbrsmntAgtAcct"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"InstdRmbrsmntAgt"`
	InstdRmbrsmntAgtAcct *CashAccount7                                 `xml:"InstdRmbrsmntAgtAcct"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification3 `xml:"ThrdRmbrsmntAgt"`
	ThrdRmbrsmntAgtAcct  *CashAccount7                                 `xml:"ThrdRmbrsmntAgtAcct"`
}

// CashAccount7
type CashAccount7 struct {
	Id  *AccountIdentification3Choice `xml:"Id"`
	Tp  *CashAccountType2             `xml:"Tp"`
	Ccy string                        `xml:"Ccy"`
	Nm  string                        `xml:"Nm"`
}

// AccountIdentification3Choice
type AccountIdentification3Choice struct {
	IBAN      string                            `xml:"IBAN"`
	BBAN      string                            `xml:"BBAN"`
	UPIC      string                            `xml:"UPIC"`
	PrtryAcct *SimpleIdentificationInformation2 `xml:"PrtryAcct"`
}

// IBANIdentifier
type IBANIdentifier string

// BBANIdentifier
type BBANIdentifier string

// UPICIdentifier
type UPICIdentifier string

// SimpleIdentificationInformation2
type SimpleIdentificationInformation2 struct {
	Id string `xml:"Id"`
}

// CashAccountType2
type CashAccountType2 struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// CashAccountType4Code
type CashAccountType4Code string

// ClearingSystemIdentification1Choice
type ClearingSystemIdentification1Choice struct {
	ClrSysId string `xml:"ClrSysId"`
	Prtry    string `xml:"Prtry"`
}

// CashClearingSystem3Code
type CashClearingSystem3Code string

// BranchAndFinancialInstitutionIdentification3
type BranchAndFinancialInstitutionIdentification3 struct {
	FinInstnId *FinancialInstitutionIdentification5Choice `xml:"FinInstnId"`
	BrnchId    *BranchData                                `xml:"BrnchId"`
}

// FinancialInstitutionIdentification5Choice
type FinancialInstitutionIdentification5Choice struct {
	BIC         string                                     `xml:"BIC"`
	ClrSysMmbId *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId"`
	NmAndAdr    *NameAndAddress7                           `xml:"NmAndAdr"`
	PrtryId     *GenericIdentification3                    `xml:"PrtryId"`
	CmbndId     *FinancialInstitutionIdentification3       `xml:"CmbndId"`
}

// ClearingSystemMemberIdentification3Choice
type ClearingSystemMemberIdentification3Choice struct {
	Id    string `xml:"Id"`
	Prtry string `xml:"Prtry"`
}

// ExternalClearingSystemMemberCode
type ExternalClearingSystemMemberCode string

// NameAndAddress7
type NameAndAddress7 struct {
	Nm      string          `xml:"Nm"`
	PstlAdr *PostalAddress1 `xml:"PstlAdr"`
}

// FinancialInstitutionIdentification3
type FinancialInstitutionIdentification3 struct {
	BIC         string                                     `xml:"BIC"`
	ClrSysMmbId *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId"`
	Nm          string                                     `xml:"Nm"`
	PstlAdr     *PostalAddress1                            `xml:"PstlAdr"`
	PrtryId     *GenericIdentification3                    `xml:"PrtryId"`
}

// BranchData
type BranchData struct {
	Id      string          `xml:"Id"`
	Nm      string          `xml:"Nm"`
	PstlAdr *PostalAddress1 `xml:"PstlAdr"`
}

// PaymentTypeInformation6
type PaymentTypeInformation6 struct {
	InstrPrty string                  `xml:"InstrPrty"`
	SvcLvl    *ServiceLevel2Choice    `xml:"SvcLvl"`
	ClrChanl  string                  `xml:"ClrChanl"`
	LclInstrm *LocalInstrument1Choice `xml:"LclInstrm"`
	SeqTp     string                  `xml:"SeqTp"`
	CtgyPurp  string                  `xml:"CtgyPurp"`
}

// ServiceLevel2Choice
type ServiceLevel2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ServiceLevel1Code
type ServiceLevel1Code string

// LocalInstrument1Choice
type LocalInstrument1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ExternalLocalInstrumentCode
type ExternalLocalInstrumentCode string

// SequenceType1Code
type SequenceType1Code string

// PaymentCategoryPurpose1Code
type PaymentCategoryPurpose1Code string

// PaymentMethod4Code
type PaymentMethod4Code string

// MandateRelatedInformation1
type MandateRelatedInformation1 struct {
	MndtId        string                        `xml:"MndtId"`
	DtOfSgntr     string                        `xml:"DtOfSgntr"`
	AmdmntInd     bool                          `xml:"AmdmntInd"`
	AmdmntInfDtls *AmendmentInformationDetails1 `xml:"AmdmntInfDtls"`
	ElctrncSgntr  string                        `xml:"ElctrncSgntr"`
	FrstColltnDt  string                        `xml:"FrstColltnDt"`
	FnlColltnDt   string                        `xml:"FnlColltnDt"`
	Frqcy         string                        `xml:"Frqcy"`
}

// TrueFalseIndicator
type TrueFalseIndicator bool

// AmendmentInformationDetails1
type AmendmentInformationDetails1 struct {
	OrgnlMndtId      string                                        `xml:"OrgnlMndtId"`
	OrgnlCdtrSchmeId *PartyIdentification8                         `xml:"OrgnlCdtrSchmeId"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlCdtrAgt"`
	OrgnlCdtrAgtAcct *CashAccount7                                 `xml:"OrgnlCdtrAgtAcct"`
	OrgnlDbtr        *PartyIdentification8                         `xml:"OrgnlDbtr"`
	OrgnlDbtrAcct    *CashAccount7                                 `xml:"OrgnlDbtrAcct"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlDbtrAgt"`
	OrgnlDbtrAgtAcct *CashAccount7                                 `xml:"OrgnlDbtrAgtAcct"`
	OrgnlFnlColltnDt string                                        `xml:"OrgnlFnlColltnDt"`
	OrgnlFrqcy       string                                        `xml:"OrgnlFrqcy"`
}

// Frequency1Code
type Frequency1Code string

// Max1025Text
type Max1025Text string

// RemittanceInformation1
type RemittanceInformation1 struct {
	Ustrd []string                            `xml:"Ustrd"`
	Strd  []*StructuredRemittanceInformation6 `xml:"Strd"`
}

// StructuredRemittanceInformation6
type StructuredRemittanceInformation6 struct {
	RfrdDocInf    *ReferredDocumentInformation1    `xml:"RfrdDocInf"`
	RfrdDocRltdDt string                           `xml:"RfrdDocRltdDt"`
	RfrdDocAmt    []*ReferredDocumentAmount1Choice `xml:"RfrdDocAmt"`
	CdtrRefInf    *CreditorReferenceInformation1   `xml:"CdtrRefInf"`
	Invcr         *PartyIdentification8            `xml:"Invcr"`
	Invcee        *PartyIdentification8            `xml:"Invcee"`
	AddtlRmtInf   string                           `xml:"AddtlRmtInf"`
}

// ReferredDocumentInformation1
type ReferredDocumentInformation1 struct {
	RfrdDocTp *ReferredDocumentType1 `xml:"RfrdDocTp"`
	RfrdDocNb string                 `xml:"RfrdDocNb"`
}

// ReferredDocumentType1
type ReferredDocumentType1 struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
	Issr  string `xml:"Issr"`
}

// DocumentType2Code
type DocumentType2Code string

// ReferredDocumentAmount1Choice
type ReferredDocumentAmount1Choice struct {
	DuePyblAmt   *CurrencyAndAmount `xml:"DuePyblAmt"`
	DscntApldAmt *CurrencyAndAmount `xml:"DscntApldAmt"`
	RmtdAmt      *CurrencyAndAmount `xml:"RmtdAmt"`
	CdtNoteAmt   *CurrencyAndAmount `xml:"CdtNoteAmt"`
	TaxAmt       *CurrencyAndAmount `xml:"TaxAmt"`
}

// CreditorReferenceInformation1
type CreditorReferenceInformation1 struct {
	CdtrRefTp *CreditorReferenceType1 `xml:"CdtrRefTp"`
	CdtrRef   string                  `xml:"CdtrRef"`
}

// CreditorReferenceType1
type CreditorReferenceType1 struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
	Issr  string `xml:"Issr"`
}

// DocumentType3Code
type DocumentType3Code string
