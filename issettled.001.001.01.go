// IsSettled iso20022

package schema

// Document
type Document *Document

// ExternalFinancialInstitutionIdentification1Code
type ExternalFinancialInstitutionIdentification1Code string

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

// SettlementMethod1Code
type SettlementMethod1Code string

// ExternalCashClearingSystem1Code
type ExternalCashClearingSystem1Code string

// IBAN2007Identifier
type IBAN2007Identifier string

// Max34Text
type Max34Text string

// ExternalAccountIdentification1Code
type ExternalAccountIdentification1Code string

// ExternalCashAccountType1Code
type ExternalCashAccountType1Code string

// ActiveOrHistoricCurrencyCode
type ActiveOrHistoricCurrencyCode string

// Max70Text
type Max70Text string

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

// ClearingSystemIdentification3Choice
type ClearingSystemIdentification3Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

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

// GenericAccountIdentification1
type GenericAccountIdentification1 struct {
	Id      string                    `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                    `xml:"Issr"`
}

// AccountSchemeName1Choice
type AccountSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// CashAccountType2Choice
type CashAccountType2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

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

// FinancialIdentificationSchemeName1Choice
type FinancialIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// GenericFinancialIdentification1
type GenericFinancialIdentification1 struct {
	Id      string                                    `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                    `xml:"Issr"`
}

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

// CustomerIdentificationStatusNotificationV01
type CustomerIdentificationStatusNotificationV01 struct {
	GrpHdr          *GroupHeader70                                `xml:"GrpHdr"`
	AcctInfAndIdSts []*AccountInformationAndIdentificationStatus1 `xml:"AcctInfAndIdSts"`
}

// AccountInformationAndIdentificationStatus1
type AccountInformationAndIdentificationStatus1 struct {
	Id         string                              `xml:"Id"`
	AcctIdInf  *AccountIdentificationInformation1  `xml:"AcctIdInf"`
	CstmrIdInf *CustomerIdentificationInformation1 `xml:"CstmrIdInf"`
}

// AccountIdentificationInformation1
type AccountIdentificationInformation1 struct {
	AdrsngId       string                  `xml:"AdrsngId"`
	SttlmAcctId    string                  `xml:"SttlmAcctId"`
	NoteTp         string                  `xml:"NoteTp"`
	Note           string                  `xml:"Note"`
	CstmrAcctIdSts string                  `xml:"CstmrAcctIdSts"`
	PmtId          *PaymentIdentification3 `xml:"PmtId"`
}

// CustomerIdentificationInformation1
type CustomerIdentificationInformation1 struct {
	KYCSts string                  `xml:"KYCSts"`
	AMLSts string                  `xml:"AMLSts"`
	PIIInf string                  `xml:"PIIInf"`
	PmtId  *PaymentIdentification3 `xml:"PmtId"`
}

// PaymentIdentification3
type PaymentIdentification3 struct {
	InstrId    string `xml:"InstrId"`
	EndToEndId string `xml:"EndToEndId"`
	TxId       string `xml:"TxId"`
	ClrSysRef  string `xml:"ClrSysRef"`
}

// NoteType1
type NoteType1 string

// Note1
type Note1 string

// CustomerAccountIdentificationStatus1
type CustomerAccountIdentificationStatus1 string

// PersonallyIdentifiableInformationInfo1
type PersonallyIdentifiableInformationInfo1 string

// KYCStatus1
type KYCStatus1 string

// AMLStatus1
type AMLStatus1 string

// AddressingIdentification1
type AddressingIdentification1 string

// SettlementAccountIdentification1
type SettlementAccountIdentification1 string

// DateTimeString
type DateTimeString string

// DateString
type DateString string
