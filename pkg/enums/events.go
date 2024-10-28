package enums

type Events string

const (
	NO  Events = "NO"
	RR  Events = "RR"
	CR  Events = "CR"
	UNK Events = "UNKNOW"
)

var eventsDescription = map[Events]string{
	NO:  "New Order Request",
	RR:  "Replace Request",
	CR:  "Cancel Request",
	UNK: "Unknown",
}

func GetDescription(event Events) string {
	return eventsDescription[event]
}
