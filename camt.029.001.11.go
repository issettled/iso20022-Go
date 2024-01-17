// IsSettled iso20022

package schema

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

// AmendmentInformationDetails14
type AmendmentInformationDetails14 struct {
	OrgnlMndtId      string                                        `xml:"OrgnlMndtId"`
	OrgnlCdtrSchmeId *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt"`
	OrgnlCdtrAgtAcct *CashAccount40                                `xml:"OrgnlCdtrAgtAcct"`
	OrgnlDbtr        *PartyIdentification135                       `xml:"OrgnlDbtr"`
	OrgnlDbtrAcct    *CashAccount40                                `xml:"OrgnlDbtrAcct"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt"`
	OrgnlDbtrAgtAcct *CashAccount40                                `xml:"OrgnlDbtrAgtAcct"`
	OrgnlFnlColltnDt string                                        `xml:"OrgnlFnlColltnDt"`
	OrgnlFrqcy       *Frequency36Choice                            `xml:"OrgnlFrqcy"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn"`
	OrgnlTrckgDays   string                                        `xml:"OrgnlTrckgDays"`
}

// AmountType4Choice
type AmountType4Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  *EquivalentAmount2                 `xml:"EqvtAmt"`
}

// AnyBICDec2014Identifier
type AnyBICDec2014Identifier string

// BICFIDec2014Identifier
type BICFIDec2014Identifier string

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

// CancellationIndividualStatus1Code
type CancellationIndividualStatus1Code string

// CancellationStatusReason3Choice
type CancellationStatusReason3Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// CancellationStatusReason4
type CancellationStatusReason4 struct {
	Orgtr    *PartyIdentification135          `xml:"Orgtr"`
	Rsn      *CancellationStatusReason3Choice `xml:"Rsn"`
	AddtlInf []string                         `xml:"AddtlInf"`
}

// Case5
type Case5 struct {
	Id             string         `xml:"Id"`
	Cretr          *Party40Choice `xml:"Cretr"`
	ReopCaseIndctn bool           `xml:"ReopCaseIndctn"`
}

// CaseAssignment5
type CaseAssignment5 struct {
	Id      string         `xml:"Id"`
	Assgnr  *Party40Choice `xml:"Assgnr"`
	Assgne  *Party40Choice `xml:"Assgne"`
	CreDtTm string         `xml:"CreDtTm"`
}

// CashAccount40
type CashAccount40 struct {
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

// ChargeIncludedIndicator
type ChargeIncludedIndicator bool

// ChargeType3Choice
type ChargeType3Choice struct {
	Cd    string                  `xml:"Cd"`
	Prtry *GenericIdentification3 `xml:"Prtry"`
}

// Charges6
type Charges6 struct {
	TtlChrgsAndTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt"`
	Rcrd              []*ChargesRecord3                  `xml:"Rcrd"`
}

// Charges9
type Charges9 struct {
	Amt     *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt     *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
	AgtAcct *CashAccount40                                `xml:"AgtAcct"`
}

// ChargesRecord3
type ChargesRecord3 struct {
	Amt         *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	CdtDbtInd   string                                        `xml:"CdtDbtInd"`
	ChrgInclInd bool                                          `xml:"ChrgInclInd"`
	Tp          *ChargeType3Choice                            `xml:"Tp"`
	Rate        float64                                       `xml:"Rate"`
	Br          string                                        `xml:"Br"`
	Agt         *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
	Tax         *TaxCharges2                                  `xml:"Tax"`
}

// ClaimNonReceipt2
type ClaimNonReceipt2 struct {
	DtPrcd      string                                        `xml:"DtPrcd"`
	OrgnlNxtAgt *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlNxtAgt"`
}

// ClaimNonReceipt2Choice
type ClaimNonReceipt2Choice struct {
	Accptd *ClaimNonReceipt2                   `xml:"Accptd"`
	Rjctd  *ClaimNonReceiptRejectReason1Choice `xml:"Rjctd"`
}

// ClaimNonReceiptRejectReason1Choice
type ClaimNonReceiptRejectReason1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
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

// Compensation4
type Compensation4 struct {
	Amt         *ActiveCurrencyAndAmount                      `xml:"Amt"`
	DbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	DbtrAgtAcct *CashAccount40                                `xml:"DbtrAgtAcct"`
	CdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	CdtrAgtAcct *CashAccount40                                `xml:"CdtrAgtAcct"`
	Rsn         *CompensationReason1Choice                    `xml:"Rsn"`
}

// CompensationReason1Choice
type CompensationReason1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
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

// CorrectiveGroupInformation1
type CorrectiveGroupInformation1 struct {
	MsgId   string `xml:"MsgId"`
	MsgNmId string `xml:"MsgNmId"`
	CreDtTm string `xml:"CreDtTm"`
}

// CorrectiveInterbankTransaction3
type CorrectiveInterbankTransaction3 struct {
	GrpHdr         *CorrectiveGroupInformation1       `xml:"GrpHdr"`
	InstrId        string                             `xml:"InstrId"`
	EndToEndId     string                             `xml:"EndToEndId"`
	TxId           string                             `xml:"TxId"`
	UETR           string                             `xml:"UETR"`
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  string                             `xml:"IntrBkSttlmDt"`
}

// CorrectivePaymentInitiation5
type CorrectivePaymentInitiation5 struct {
	GrpHdr       *CorrectiveGroupInformation1       `xml:"GrpHdr"`
	PmtInfId     string                             `xml:"PmtInfId"`
	InstrId      string                             `xml:"InstrId"`
	EndToEndId   string                             `xml:"EndToEndId"`
	UETR         string                             `xml:"UETR"`
	InstdAmt     *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	ReqdExctnDt  *DateAndDateTime2Choice            `xml:"ReqdExctnDt"`
	ReqdColltnDt string                             `xml:"ReqdColltnDt"`
}

// CorrectiveTransaction5Choice
type CorrectiveTransaction5Choice struct {
	Initn  *CorrectivePaymentInitiation5    `xml:"Initn"`
	IntrBk *CorrectiveInterbankTransaction3 `xml:"IntrBk"`
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

// DateAndDateTime2Choice
type DateAndDateTime2Choice struct {
	Dt   string `xml:"Dt"`
	DtTm string `xml:"DtTm"`
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

// EquivalentAmount2
type EquivalentAmount2 struct {
	Amt      *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CcyOfTrf string                             `xml:"CcyOfTrf"`
}

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

// ExternalChargeType1Code
type ExternalChargeType1Code string

// ExternalClaimNonReceiptRejection1Code
type ExternalClaimNonReceiptRejection1Code string

// ExternalClearingSystemIdentification1Code
type ExternalClearingSystemIdentification1Code string

// ExternalDiscountAmountType1Code
type ExternalDiscountAmountType1Code string

// ExternalDocumentLineType1Code
type ExternalDocumentLineType1Code string

// ExternalFinancialInstitutionIdentification1Code
type ExternalFinancialInstitutionIdentification1Code string

// ExternalGarnishmentType1Code
type ExternalGarnishmentType1Code string

// ExternalInvestigationExecutionConfirmation1Code
type ExternalInvestigationExecutionConfirmation1Code string

// ExternalLocalInstrument1Code
type ExternalLocalInstrument1Code string

// ExternalMandateSetupReason1Code
type ExternalMandateSetupReason1Code string

// ExternalOrganisationIdentification1Code
type ExternalOrganisationIdentification1Code string

// ExternalPaymentCancellationRejection1Code
type ExternalPaymentCancellationRejection1Code string

// ExternalPaymentCompensationReason1Code
type ExternalPaymentCompensationReason1Code string

// ExternalPaymentModificationRejection1Code
type ExternalPaymentModificationRejection1Code string

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

// GenericIdentification3
type GenericIdentification3 struct {
	Id   string `xml:"Id"`
	Issr string `xml:"Issr"`
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

// GroupCancellationStatus1Code
type GroupCancellationStatus1Code string

// IBAN2007Identifier
type IBAN2007Identifier string

// ISODate
type ISODate string

// ISODateTime
type ISODateTime string

// ISOYear
type ISOYear string

// InvestigationStatus5Choice
type InvestigationStatus5Choice struct {
	Conf           string                             `xml:"Conf"`
	RjctdMod       []*ModificationStatusReason1Choice `xml:"RjctdMod"`
	DplctOf        *Case5                             `xml:"DplctOf"`
	AssgnmtCxlConf bool                               `xml:"AssgnmtCxlConf"`
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

// MandateRelatedData2Choice
type MandateRelatedData2Choice struct {
	DrctDbtMndt *MandateRelatedInformation15 `xml:"DrctDbtMndt"`
	CdtTrfMndt  *CreditTransferMandateData1  `xml:"CdtTrfMndt"`
}

// MandateRelatedInformation15
type MandateRelatedInformation15 struct {
	MndtId        string                         `xml:"MndtId"`
	DtOfSgntr     string                         `xml:"DtOfSgntr"`
	AmdmntInd     bool                           `xml:"AmdmntInd"`
	AmdmntInfDtls *AmendmentInformationDetails14 `xml:"AmdmntInfDtls"`
	ElctrncSgntr  string                         `xml:"ElctrncSgntr"`
	FrstColltnDt  string                         `xml:"FrstColltnDt"`
	FnlColltnDt   string                         `xml:"FnlColltnDt"`
	Frqcy         *Frequency36Choice             `xml:"Frqcy"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn"`
	TrckgDays     string                         `xml:"TrckgDays"`
}

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

// Max1025Text
type Max1025Text string

// Max105Text
type Max105Text string

// Max10KBinary
type Max10KBinary []byte

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

// ModificationStatusReason1Choice
type ModificationStatusReason1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ModificationStatusReason2
type ModificationStatusReason2 struct {
	Orgtr    *PartyIdentification135          `xml:"Orgtr"`
	Rsn      *ModificationStatusReason1Choice `xml:"Rsn"`
	AddtlInf []string                         `xml:"AddtlInf"`
}

// NamePrefix2Code
type NamePrefix2Code string

// Number
type Number float64

// NumberOfCancellationsPerStatus1
type NumberOfCancellationsPerStatus1 struct {
	DtldNbOfTxs string  `xml:"DtldNbOfTxs"`
	DtldSts     string  `xml:"DtldSts"`
	DtldCtrlSum float64 `xml:"DtldCtrlSum"`
}

// NumberOfTransactionsPerStatus1
type NumberOfTransactionsPerStatus1 struct {
	DtldNbOfTxs string  `xml:"DtldNbOfTxs"`
	DtldSts     string  `xml:"DtldSts"`
	DtldCtrlSum float64 `xml:"DtldCtrlSum"`
}

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

// OriginalGroupHeader14
type OriginalGroupHeader14 struct {
	OrgnlGrpCxlId    string                            `xml:"OrgnlGrpCxlId"`
	RslvdCase        *Case5                            `xml:"RslvdCase"`
	OrgnlMsgId       string                            `xml:"OrgnlMsgId"`
	OrgnlMsgNmId     string                            `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm     string                            `xml:"OrgnlCreDtTm"`
	OrgnlNbOfTxs     string                            `xml:"OrgnlNbOfTxs"`
	OrgnlCtrlSum     float64                           `xml:"OrgnlCtrlSum"`
	GrpCxlSts        string                            `xml:"GrpCxlSts"`
	CxlStsRsnInf     []*CancellationStatusReason4      `xml:"CxlStsRsnInf"`
	NbOfTxsPerCxlSts []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts"`
}

// OriginalGroupInformation29
type OriginalGroupInformation29 struct {
	OrgnlMsgId   string `xml:"OrgnlMsgId"`
	OrgnlMsgNmId string `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm string `xml:"OrgnlCreDtTm"`
}

// OriginalPaymentInstruction43
type OriginalPaymentInstruction43 struct {
	OrgnlPmtInfCxlId string                             `xml:"OrgnlPmtInfCxlId"`
	RslvdCase        *Case5                             `xml:"RslvdCase"`
	OrgnlPmtInfId    string                             `xml:"OrgnlPmtInfId"`
	OrgnlGrpInf      *OriginalGroupInformation29        `xml:"OrgnlGrpInf"`
	OrgnlNbOfTxs     string                             `xml:"OrgnlNbOfTxs"`
	OrgnlCtrlSum     float64                            `xml:"OrgnlCtrlSum"`
	PmtInfCxlSts     string                             `xml:"PmtInfCxlSts"`
	CxlStsRsnInf     []*CancellationStatusReason4       `xml:"CxlStsRsnInf"`
	NbOfTxsPerCxlSts []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts"`
	TxInfAndSts      []*PaymentTransaction139           `xml:"TxInfAndSts"`
}

// OriginalTransactionReference35
type OriginalTransactionReference35 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt"`
	Amt            *AmountType4Choice                            `xml:"Amt"`
	IntrBkSttlmDt  string                                        `xml:"IntrBkSttlmDt"`
	ReqdColltnDt   string                                        `xml:"ReqdColltnDt"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId"`
	SttlmInf       *SettlementInstruction11                      `xml:"SttlmInf"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf"`
	PmtMtd         string                                        `xml:"PmtMtd"`
	MndtRltdInf    *MandateRelatedData2Choice                    `xml:"MndtRltdInf"`
	RmtInf         *RemittanceInformation21                      `xml:"RmtInf"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr"`
	Dbtr           *Party40Choice                                `xml:"Dbtr"`
	DbtrAcct       *CashAccount40                                `xml:"DbtrAcct"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	DbtrAgtAcct    *CashAccount40                                `xml:"DbtrAgtAcct"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	CdtrAgtAcct    *CashAccount40                                `xml:"CdtrAgtAcct"`
	Cdtr           *Party40Choice                                `xml:"Cdtr"`
	CdtrAcct       *CashAccount40                                `xml:"CdtrAcct"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr"`
	Purp           *Purpose2Choice                               `xml:"Purp"`
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

// Party40Choice
type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"Pty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// PartyIdentification135
type PartyIdentification135 struct {
	Nm        string           `xml:"Nm"`
	PstlAdr   *PostalAddress24 `xml:"PstlAdr"`
	Id        *Party38Choice   `xml:"Id"`
	CtryOfRes string           `xml:"CtryOfRes"`
	CtctDtls  *Contact4        `xml:"CtctDtls"`
}

// PaymentMethod4Code
type PaymentMethod4Code string

// PaymentTransaction132
type PaymentTransaction132 struct {
	ModStsId            string                             `xml:"ModStsId"`
	RslvdCase           *Case5                             `xml:"RslvdCase"`
	OrgnlGrpInf         *OriginalGroupInformation29        `xml:"OrgnlGrpInf"`
	OrgnlPmtInfId       string                             `xml:"OrgnlPmtInfId"`
	OrgnlInstrId        string                             `xml:"OrgnlInstrId"`
	OrgnlEndToEndId     string                             `xml:"OrgnlEndToEndId"`
	OrgnlTxId           string                             `xml:"OrgnlTxId"`
	OrgnlClrSysRef      string                             `xml:"OrgnlClrSysRef"`
	OrgnlUETR           string                             `xml:"OrgnlUETR"`
	ModStsRsnInf        []*ModificationStatusReason2       `xml:"ModStsRsnInf"`
	RsltnRltdInf        *ResolutionData3                   `xml:"RsltnRltdInf"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  string                             `xml:"OrgnlIntrBkSttlmDt"`
	Assgnr              *Party40Choice                     `xml:"Assgnr"`
	Assgne              *Party40Choice                     `xml:"Assgne"`
	OrgnlTxRef          *OriginalTransactionReference35    `xml:"OrgnlTxRef"`
}

// PaymentTransaction138
type PaymentTransaction138 struct {
	CxlStsId            string                             `xml:"CxlStsId"`
	RslvdCase           *Case5                             `xml:"RslvdCase"`
	OrgnlGrpInf         *OriginalGroupInformation29        `xml:"OrgnlGrpInf"`
	OrgnlInstrId        string                             `xml:"OrgnlInstrId"`
	OrgnlEndToEndId     string                             `xml:"OrgnlEndToEndId"`
	OrgnlTxId           string                             `xml:"OrgnlTxId"`
	OrgnlClrSysRef      string                             `xml:"OrgnlClrSysRef"`
	OrgnlUETR           string                             `xml:"OrgnlUETR"`
	TxCxlSts            string                             `xml:"TxCxlSts"`
	CxlStsRsnInf        []*CancellationStatusReason4       `xml:"CxlStsRsnInf"`
	RsltnRltdInf        *ResolutionData3                   `xml:"RsltnRltdInf"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  string                             `xml:"OrgnlIntrBkSttlmDt"`
	Assgnr              *Party40Choice                     `xml:"Assgnr"`
	Assgne              *Party40Choice                     `xml:"Assgne"`
	OrgnlTxRef          *OriginalTransactionReference35    `xml:"OrgnlTxRef"`
}

// PaymentTransaction139
type PaymentTransaction139 struct {
	CxlStsId          string                             `xml:"CxlStsId"`
	RslvdCase         *Case5                             `xml:"RslvdCase"`
	OrgnlInstrId      string                             `xml:"OrgnlInstrId"`
	OrgnlEndToEndId   string                             `xml:"OrgnlEndToEndId"`
	UETR              string                             `xml:"UETR"`
	TxCxlSts          string                             `xml:"TxCxlSts"`
	CxlStsRsnInf      []*CancellationStatusReason4       `xml:"CxlStsRsnInf"`
	OrgnlInstdAmt     *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`
	OrgnlReqdExctnDt  *DateAndDateTime2Choice            `xml:"OrgnlReqdExctnDt"`
	OrgnlReqdColltnDt string                             `xml:"OrgnlReqdColltnDt"`
	OrgnlTxRef        *OriginalTransactionReference35    `xml:"OrgnlTxRef"`
}

// PaymentTypeInformation27
type PaymentTypeInformation27 struct {
	InstrPrty string                  `xml:"InstrPrty"`
	ClrChanl  string                  `xml:"ClrChanl"`
	SvcLvl    []*ServiceLevel8Choice  `xml:"SvcLvl"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm"`
	SeqTp     string                  `xml:"SeqTp"`
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

// RemittanceInformation21
type RemittanceInformation21 struct {
	Ustrd []string                             `xml:"Ustrd"`
	Strd  []*StructuredRemittanceInformation17 `xml:"Strd"`
}

// ResolutionData3
type ResolutionData3 struct {
	EndToEndId     string                             `xml:"EndToEndId"`
	TxId           string                             `xml:"TxId"`
	UETR           string                             `xml:"UETR"`
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  string                             `xml:"IntrBkSttlmDt"`
	ClrChanl       string                             `xml:"ClrChanl"`
	Compstn        *Compensation4                     `xml:"Compstn"`
	Chrgs          []*Charges9                        `xml:"Chrgs"`
}

// ResolutionOfInvestigationV11
type ResolutionOfInvestigationV11 struct {
	Assgnmt       *CaseAssignment5              `xml:"Assgnmt"`
	RslvdCase     *Case5                        `xml:"RslvdCase"`
	Sts           *InvestigationStatus5Choice   `xml:"Sts"`
	CxlDtls       []*UnderlyingTransaction29    `xml:"CxlDtls"`
	ModDtls       *PaymentTransaction132        `xml:"ModDtls"`
	ClmNonRctDtls *ClaimNonReceipt2Choice       `xml:"ClmNonRctDtls"`
	StmtDtls      *StatementResolutionEntry4    `xml:"StmtDtls"`
	CrrctnTx      *CorrectiveTransaction5Choice `xml:"CrrctnTx"`
	RsltnRltdInf  *ResolutionData3              `xml:"RsltnRltdInf"`
	SplmtryData   []*SupplementaryData1         `xml:"SplmtryData"`
}

// SequenceType3Code
type SequenceType3Code string

// ServiceLevel8Choice
type ServiceLevel8Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// SettlementInstruction11
type SettlementInstruction11 struct {
	SttlmMtd             string                                        `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount40                                `xml:"SttlmAcct"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt"`
	InstgRmbrsmntAgtAcct *CashAccount40                                `xml:"InstgRmbrsmntAgtAcct"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt"`
	InstdRmbrsmntAgtAcct *CashAccount40                                `xml:"InstdRmbrsmntAgtAcct"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt"`
	ThrdRmbrsmntAgtAcct  *CashAccount40                                `xml:"ThrdRmbrsmntAgtAcct"`
}

// SettlementMethod1Code
type SettlementMethod1Code string

// StatementResolutionEntry4
type StatementResolutionEntry4 struct {
	OrgnlGrpInf *OriginalGroupInformation29        `xml:"OrgnlGrpInf"`
	OrgnlStmtId string                             `xml:"OrgnlStmtId"`
	UETR        string                             `xml:"UETR"`
	AcctSvcrRef string                             `xml:"AcctSvcrRef"`
	CrrctdAmt   *ActiveOrHistoricCurrencyAndAmount `xml:"CrrctdAmt"`
	Chrgs       []*Charges6                        `xml:"Chrgs"`
	Purp        *Purpose2Choice                    `xml:"Purp"`
}

// StructuredRemittanceInformation17
type StructuredRemittanceInformation17 struct {
	RfrdDocInf  []*ReferredDocumentInformation7 `xml:"RfrdDocInf"`
	RfrdDocAmt  *RemittanceAmount2              `xml:"RfrdDocAmt"`
	CdtrRefInf  *CreditorReferenceInformation2  `xml:"CdtrRefInf"`
	Invcr       *PartyIdentification135         `xml:"Invcr"`
	Invcee      *PartyIdentification135         `xml:"Invcee"`
	TaxRmt      *TaxData1                       `xml:"TaxRmt"`
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

// TaxAmount3
type TaxAmount3 struct {
	Rate         float64                            `xml:"Rate"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt"`
	Dtls         []*TaxRecordDetails3               `xml:"Dtls"`
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

// TaxCharges2
type TaxCharges2 struct {
	Id   string                             `xml:"Id"`
	Rate float64                            `xml:"Rate"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// TaxData1
type TaxData1 struct {
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
	Rcrd            []*TaxRecord3                      `xml:"Rcrd"`
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

// TaxPeriod3
type TaxPeriod3 struct {
	Yr     string       `xml:"Yr"`
	Tp     string       `xml:"Tp"`
	FrToDt *DatePeriod2 `xml:"FrToDt"`
}

// TaxRecord3
type TaxRecord3 struct {
	Tp       string      `xml:"Tp"`
	Ctgy     string      `xml:"Ctgy"`
	CtgyDtls string      `xml:"CtgyDtls"`
	DbtrSts  string      `xml:"DbtrSts"`
	CertId   string      `xml:"CertId"`
	FrmsCd   string      `xml:"FrmsCd"`
	Prd      *TaxPeriod3 `xml:"Prd"`
	TaxAmt   *TaxAmount3 `xml:"TaxAmt"`
	AddtlInf string      `xml:"AddtlInf"`
}

// TaxRecordDetails3
type TaxRecordDetails3 struct {
	Prd *TaxPeriod3                        `xml:"Prd"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// TaxRecordPeriod1Code
type TaxRecordPeriod1Code string

// TransactionIndividualStatus1Code
type TransactionIndividualStatus1Code string

// TrueFalseIndicator
type TrueFalseIndicator bool

// UUIDv4Identifier
type UUIDv4Identifier string

// UnderlyingTransaction29
type UnderlyingTransaction29 struct {
	OrgnlGrpInfAndSts *OriginalGroupHeader14          `xml:"OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []*OriginalPaymentInstruction43 `xml:"OrgnlPmtInfAndSts"`
	TxInfAndSts       []*PaymentTransaction138        `xml:"TxInfAndSts"`
}

// YesNoIndicator
type YesNoIndicator bool
