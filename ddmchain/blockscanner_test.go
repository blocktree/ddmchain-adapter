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

package ddmchain

import (
	"github.com/blocktree/openwallet/v2/log"
	"testing"
)

func TestWalletManager_DDMGetTransactionByHash(t *testing.T) {
	wm := testNewWalletManager()
	txid := "0x5500f79be928fd2ccdaeb69438c38e0304c3a4f73a0e4c77357ca0abc789871f"
	tx, err := wm.WalletClient.DdmGetTransactionByHash(txid)
	if err != nil {
		t.Errorf("get transaction by has failed, err=%v", err)
		return
	}
	log.Infof("tx: %+v", tx)
}


func TestWalletManager_GetGlobalMaxBlockHeight(t *testing.T) {
	//wm := testNewWalletManager()
	//log.Infof("tx: %+v", tx)
}
