package versio

type RecordType string
type TTLValue string

const (
	RecordTypeA RecordType		= "A"
	RecordTypeAAAA RecordType	= "AAAA"
	RecordTypeMX RecordType 	= "MX"
	RecordTypeCNAME RecordType 	= "CNAME"
	RecordTypePTR RecordType	= "PTR"
	RecordTypeSRV RecordType	= "SRV"
	RecordTypeTxt RecordType	= "TXT"
	RecordTypeSPF RecordType	= "SPF"
	TTLValue300 TTLValue		= "300"
	TTLValue3600 TTLValue		= "3600"
	TTLValue7200 TTLValue		= "7200"
	TTLValue14400 TTLValue		= "14400"
	TTLValue28800 TTLValue		= "28800"
	TTLValue86400 TTLValue		= "86400"
)
