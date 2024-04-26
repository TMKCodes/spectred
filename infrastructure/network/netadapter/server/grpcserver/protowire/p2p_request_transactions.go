package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_RequestTransactions) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_RequestTransactions is nil")
	}
	return x.RequestTransactions.toAppMessage()
}

func (x *RequestTransactionsMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestTransactionsMessage is nil")
	}
	if len(x.Ids) > appmessage.MaxInvPerRequestTransactionsMsg {
		return nil, errors.Errorf("too many hashes for message "+
			"[count %d, max %d]", len(x.Ids), appmessage.MaxInvPerRequestTransactionsMsg)
	}
	ids, err := protoTransactionIDsToDomain(x.Ids)
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgRequestTransactions{IDs: ids}, nil
}

func (x *SpectredMessage_RequestTransactions) fromAppMessage(msgGetTransactions *appmessage.MsgRequestTransactions) error {
	if len(msgGetTransactions.IDs) > appmessage.MaxInvPerRequestTransactionsMsg {
		return errors.Errorf("too many hashes for message "+
			"[count %d, max %d]", len(x.RequestTransactions.Ids), appmessage.MaxInvPerRequestTransactionsMsg)
	}

	x.RequestTransactions = &RequestTransactionsMessage{
		Ids: wireTransactionIDsToProto(msgGetTransactions.IDs),
	}
	return nil
}
