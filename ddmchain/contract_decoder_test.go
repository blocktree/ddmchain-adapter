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

import "testing"

func TestWalletManager_GetTokenBalanceByAddress(t *testing.T) {
	tm := NewWalletManager()
	baseAPI := "http://47.106.102.2:20008"
	client := &Client{BaseURL: baseAPI, Debug: true}
	tm.WalletClient = client
	tm.Config.ChainID = 1

	addrs := []AddrBalanceInf{
		&AddrBalance{Address: "0xeb826dbf0996ca8a71b1aa2d3acb1f59952b87ca", Index: 0},
	}

	err := tm.GetTokenBalanceByAddress("0x61468dfac6d242d7ee539e1d950c28b21292b30f", addrs...)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
}
