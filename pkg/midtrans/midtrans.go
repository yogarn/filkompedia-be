package midtrans

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type IMidtrans interface {
	NewTransactionToken(orderId string, amount int64) (*snap.Response, error)
}

type Midtrans struct {
	Client snap.Client
}

func NewMidtrans() IMidtrans {
	client := snap.Client{}
	client.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	return &Midtrans{
		Client: client,
	}
}

func (m *Midtrans) NewTransactionToken(orderId string, amount int64) (*snap.Response, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: amount,
		},
		//add more details for transaction
	}

	snapResp, err := m.Client.CreateTransaction(req)
	return snapResp, err
}
