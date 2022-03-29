package provider

const (
	ACTION_CREATE = 257
	ACTION_REMOVE = 258
)

type nsdOrderRequest struct {
	Flags                          int            `json:"flags"`
	Description                    string         `json:"description"`
	DateDesiredDelivery            interface{}    `json:"dateDesiredDelivery"`
	ExceptionDiaryIdentifier       string         `json:"exceptionDiaryIdentifier"`
	InvoicingApplicationIdentifier string         `json:"invoicingApplicationIdentifier"`
	StepIdentifier                 string         `json:"stepIdentifier"`
	ContactList                    []contactList  `json:"contactList"`
	OrderDetails                   []orderDetails `json:"orderDetails"`
	RelatedOrderKey                string         `json:"relatedOrderKey"`
}

type contactList struct {
	ID             string `json:"$id"`
	Username       string `json:"username"`
	Fullname       string `json:"fullname"`
	MailAddress    string `json:"mailAddress"`
	PhoneNumber    string `json:"phoneNumber"`
	CellularNumber string `json:"cellularNumber"`
	Role           string `json:"role"`
}

type orderDetails struct {
	SourceSpecifications      string `json:"sourceSpecifications"`
	DestinationSpecifications string `json:"destinationSpecifications"`
	PortSpecifications        string `json:"portSpecifications"`
	Comment                   string `json:"comment"`
}
