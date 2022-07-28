/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/blocktree/openwallet/v2/openw"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
	log.Info("balance:", balance.ConfirmBalance)
	log.Info("balance:", balance.Balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract,nil)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract,
	feeSupportAccount *openwallet.FeesSupportAccount) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, feeSupportAccount)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}


func TestAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WCDPdNNVkUvzHJmFDCTxsPeb5oCq9X2yb2"
	accountID := "FYxhRfWSR6AXqCfnn94uftwKtvRCiz73opchkzwVATTd"
	//to := "0xb722ea468a210918c95cf7a02a44d0e28f614e1f"

	testGetAssetsAccountBalance(tm, walletID, accountID)
}

func TestTransfer(t *testing.T) {

	//0x337dab81e24d916383409a7861e2ea9f46e3c94c
	//0x05159cfa87d8f0ad32b8589680598fe5a7884e78
	//0xa6e2e4f52851864bc2e783103f1067a0bf7d13e3
	//0x9f8d6babb0b28448027c9485c9e4f8b2049f6573
	//0x99b44df54026cf960ee51541b5e648418fda7290

	tm := testInitWalletManager()
	walletID := "WAQKBtF9AqAHM8ifFN4AnqQXYrE4QMyQ3t"
	accountID := "EG8kqqHtQLt4KBMog6mSPDt2i5fKirJ7sB8bduwu3ke5"
	to := "0x99b44df54026cf960ee51541b5e648418fda7290"

	//accountID := "2itC8V6c3J1KoWiqPFN9sHyo7WzLKYoT1fakzAF1e2NR"
	//to := "0x5be7747b0752ba0e96d38634df72647db01c3e94"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.04", "", nil)
	if err != nil {
		return
	}

	log.Std.Info("rawTx: %+v", rawTx)

	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

}

func TestTransfer_ERC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WCDPdNNVkUvzHJmFDCTxsPeb5oCq9X2yb2"
	accountID := "FYxhRfWSR6AXqCfnn94uftwKtvRCiz73opchkzwVATTd"
	to := "0xb722ea468a210918c95cf7a02a44d0e28f614e1f"

	contract := openwallet.SmartContract{
		Address:  "4092678e4E78230F46A1534C0fbc8fA39780892B",
		Symbol:   "ETH",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1", "", &contract)
	if err != nil {
		return
	}

	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

}

func TestSummary(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WAQKBtF9AqAHM8ifFN4AnqQXYrE4QMyQ3t"
	//accountID := "EG8kqqHtQLt4KBMog6mSPDt2i5fKirJ7sB8bduwu3ke5"
	//summaryAddress := "0x2623de7edc1b7da05e50999b5aec8daccddc8c63"

	accountID := "2itC8V6c3J1KoWiqPFN9sHyo7WzLKYoT1fakzAF1e2NR"
	summaryAddress := "0x5be7747b0752ba0e96d38634df72647db01c3e94"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}

func TestSummary_ERC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WHSh1o6SavEya1Y3t1nGPUiEDbr7WBMffm"
	accountID := "Axm3ihzw9Rmpf1L7ciHjdyET3uyHexSNzFq3s7p8n2uS"
	summaryAddress := "0xb722ea468a210918c95cf7a02a44d0e28f614e1f"

	feesSupport := openwallet.FeesSupportAccount{
		AccountID: "3csEgf2TcxwNeoFSTsePXFVmzcyNhHAS49jsTv99n1Nv",
		//FixSupportAmount: "0.01",
		FeesSupportScale: "1.3",
	}

	contract := openwallet.SmartContract{
		Address:  "4092678e4E78230F46A1534C0fbc8fA39780892B",
		Symbol:   "ETH",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, &contract, &feesSupport)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}
