// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package snap7

// #cgo CFLAGS: -I ./
// #cgo LDFLAGS: -L . -lsnap7 -lstdc++
// #include "snap7.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// S7Object S7API Cli_Create();
func Cli_Create() C.S7Object {
	return C.Cli_Create()
}

// void S7API Cli_Destroy(S7Object *Client);
func (Client C.S7Object) Cli_Destroy() {
	C.Cli_Destroy(&Client)
}

// int S7API Cli_ConnectTo(S7Object Client, const char *Address, int Rack, int Slot);
func (Client C.S7Object) Cli_ConnectTo(addr string, Rack, Slot int) {
	C.Cli_ConnectTo(Client, C.CString(addr), C.int(Rack), C.int(Slot))
}

// int S7API Cli_SetConnectionParams(S7Object Client, const char *Address, word LocalTSAP, word RemoteTSAP);
func (Client C.S7Object) Cli_SetConnectionParams(addr string, LocalTSAP, RemoteTSAP uint16) {
	C.Cli_SetConnectionParams(Client, C.CString(addr), C.uint16_t(LocalTSAP), C.uint16_t(RemoteTSAP))
}

// int S7API Cli_SetConnectionType(S7Object Client, word ConnectionType);
func (Client C.S7Object) Cli_SetConnectionType(addr string, ConnectionType uint16) {
	C.Cli_SetConnectionType(Client, C.uint16_t(ConnectionType))
}

// int S7API Cli_Connect(S7Object Client);
func (Client C.S7Object) Cli_Connect() {
	C.Cli_Connect(Client)
}

// int S7API Cli_Disconnect(S7Object Client);
func (Client C.S7Object) Cli_Disconnect() {
	C.Cli_Disconnect(Client)
}

// int S7API Cli_GetParam(S7Object Client, int ParamNumber, void *pValue);
func (Client C.S7Object) Cli_GetParam(ParamNumber int, pValue unsafe.Pointer) {
	C.Cli_GetParam(Client, C.int(ParamNumber), pValue)
}

// int S7API Cli_SetParam(S7Object Client, int ParamNumber, void *pValue);
func (Client C.S7Object) Cli_SetParam(ParamNumber int, pValue unsafe.Pointer) {
	C.Cli_SetParam(Client, C.int(ParamNumber), pValue)
}

// int S7API Cli_SetAsCallback(S7Object Client, pfn_CliCompletion pCompletion, void *usrPtr);

// Data I/O main functions
// int S7API Cli_ReadArea(S7Object Client, int Area, int DBNumber, int Start, int Amount, int WordLen, void *pUsrData);

// int S7API Cli_WriteArea(S7Object Client, int Area, int DBNumber, int Start, int Amount, int WordLen, void *pUsrData);
// int S7API Cli_ReadMultiVars(S7Object Client, PS7DataItem Item, int ItemsCount);
// int S7API Cli_WriteMultiVars(S7Object Client, PS7DataItem Item, int ItemsCount);

// Data I/O Lean functions
// int S7API Cli_DBRead(S7Object Client, int DBNumber, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_DBRead(DBNumber, Start, Size int) []byte {
	var data [2048]byte
	pUsrData := unsafe.Pointer(&data)
	C.Cli_DBRead(Client, C.int(DBNumber), C.int(Start), C.int(Size), pUsrData)
	return data[0:Size]
}

// int S7API Cli_DBWrite(S7Object Client, int DBNumber, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_DBWrite(DBNumber, Start, Size int, wr []byte) {
	var wdata [2048]byte
	for i := 0; i < int(Size); i++ {
		wdata[i] = wr[i]
	}
	pUsrData := unsafe.Pointer(&wdata)
	C.Cli_DBWrite(Client, C.int(DBNumber), C.int(Start), C.int(Size), pUsrData)
}

// int S7API Cli_MBRead(S7Object Client, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_MBRead(Start, Size int) []byte {
	var data [2048]byte
	pUsrData := unsafe.Pointer(&data)
	C.Cli_MBRead(Client, C.int(Start), C.int(Size), pUsrData)
	return data[0:Size]
}

// int S7API Cli_MBWrite(S7Object Client, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_MBWrite(Start, Size int, wr []byte) {
	var wdata [2048]byte
	for i := 0; i < int(Size); i++ {
		wdata[i] = wr[i]
	}
	pUsrData := unsafe.Pointer(&wdata)
	C.Cli_MBWrite(Client, C.int(Start), C.int(Size), pUsrData)
}

// int S7API Cli_EBRead(S7Object Client, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_EBRead(Start, Size int, pUsrData unsafe.Pointer) {
	C.Cli_EBRead(Client, C.int(Start), C.int(Size), pUsrData)
}

// int S7API Cli_EBWrite(S7Object Client, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_EBWrite(Start, Size int, pUsrData unsafe.Pointer) {
	C.Cli_EBWrite(Client, C.int(Start), C.int(Size), pUsrData)
}

// int S7API Cli_ABRead(S7Object Client, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_ABRead(Start, Size int, pUsrData unsafe.Pointer) {
	C.Cli_ABRead(Client, C.int(Start), C.int(Size), pUsrData)
}

// int S7API Cli_ABWrite(S7Object Client, int Start, int Size, void *pUsrData);
func (Client C.S7Object) Cli_ABWrite(Start, Size int, pUsrData unsafe.Pointer) {
	C.Cli_ABWrite(Client, C.int(Start), C.int(Size), pUsrData)
}

// int S7API Cli_TMRead(S7Object Client, int Start, int Amount, void *pUsrData);
func (Client C.S7Object) Cli_TMRead(Start, Amount int, pUsrData unsafe.Pointer) {
	C.Cli_TMRead(Client, C.int(Start), C.int(Amount), pUsrData)
}

// int S7API Cli_TMWrite(S7Object Client, int Start, int Amount, void *pUsrData);
func (Client C.S7Object) Cli_TMWrite(Start, Amount int, pUsrData unsafe.Pointer) {
	C.Cli_TMWrite(Client, C.int(Start), C.int(Amount), pUsrData)
}

// int S7API Cli_CTRead(S7Object Client, int Start, int Amount, void *pUsrData);
func (Client C.S7Object) Cli_CTRead(Start, Amount int, pUsrData unsafe.Pointer) {
	C.Cli_CTRead(Client, C.int(Start), C.int(Amount), pUsrData)
}

// int S7API Cli_CTWrite(S7Object Client, int Start, int Amount, void *pUsrData);
func (Client C.S7Object) Cli_CTWrite(Start, Amount int, pUsrData unsafe.Pointer) {
	C.Cli_CTWrite(Client, C.int(Start), C.int(Amount), pUsrData)
}

// Directory functions
// int S7API Cli_ListBlocks(S7Object Client, TS7BlocksList *pUsrData);
// int S7API Cli_GetAgBlockInfo(S7Object Client, int BlockType, int BlockNum, TS7BlockInfo *pUsrData);
// int S7API Cli_GetPgBlockInfo(S7Object Client, void *pBlock, TS7BlockInfo *pUsrData, int Size);
// int S7API Cli_ListBlocksOfType(S7Object Client, int BlockType, TS7BlocksOfType *pUsrData, int *ItemsCount);

// Blocks functions
// int S7API Cli_Upload(S7Object Client, int BlockType, int BlockNum, void *pUsrData, int *Size);
// int S7API Cli_FullUpload(S7Object Client, int BlockType, int BlockNum, void *pUsrData, int *Size);
// int S7API Cli_Download(S7Object Client, int BlockNum, void *pUsrData, int Size);
// int S7API Cli_Delete(S7Object Client, int BlockType, int BlockNum);
// int S7API Cli_DBGet(S7Object Client, int DBNumber, void *pUsrData, int *Size);
// int S7API Cli_DBFill(S7Object Client, int DBNumber, int FillChar);

// Date/Time functions
// int S7API Cli_GetPlcDateTime(S7Object Client, tm *DateTime);
// int S7API Cli_SetPlcDateTime(S7Object Client, tm *DateTime);
// int S7API Cli_SetPlcSystemDateTime(S7Object Client);

// System Info functions
// int S7API Cli_GetOrderCode(S7Object Client, TS7OrderCode *pUsrData);
// int S7API Cli_GetCpuInfo(S7Object Client, TS7CpuInfo *pUsrData);
type CpuInfo struct {
	ModuleTypeName string
	SerialNumber   string
	ASName         string
	Copyright      string
	ModuleName     string
}

func (Client C.S7Object) Cli_GetCpuInfo() CpuInfo {
	var Info C.TS7CpuInfo
	var ret CpuInfo
	C.Cli_GetCpuInfo(Client, &Info)
	// if (Check(res,"Unit Info"))

	type cpuInfo struct {
		ModuleTypeName [33]byte
		SerialNumber   [25]byte
		ASName         [25]byte
		Copyright      [27]byte
		ModuleName     [25]byte
	}
	info := (*cpuInfo)(unsafe.Pointer(&Info))

	var str []byte

	for _, v := range info.ModuleTypeName {
		str = append(str, v)
	}
	ret.ModuleTypeName = string(str)
	str = []byte{}

	for _, v := range info.SerialNumber {
		str = append(str, v)
	}
	ret.SerialNumber = string(str)
	str = []byte{}

	for _, v := range info.Copyright {
		str = append(str, v)
	}
	ret.Copyright = string(str)
	str = []byte{}

	for _, v := range info.ASName {
		str = append(str, v)
	}
	ret.ASName = string(str)
	str = []byte{}

	for _, v := range info.ModuleName {
		str = append(str, v)
	}
	ret.ModuleName = string(str)
	str = []byte{}

	fmt.Println("  Module Type Name :", ret.ModuleTypeName)
	fmt.Println("  Seriel Number    :", ret.SerialNumber)
	fmt.Println("  Copyright        :", ret.Copyright)
	fmt.Println("  AS Name          :", ret.ASName)
	fmt.Println("  Module Name      :", ret.ModuleName)

	return ret
}

// int S7API Cli_GetCpInfo(S7Object Client, TS7CpInfo *pUsrData);
func (Client C.S7Object) Cli_GetCpInfo() (Info C.TS7CpInfo) {
	// var Info C.TS7CpInfo
	C.Cli_GetCpInfo(Client, &Info)
	fmt.Printf("  Max PDU Length   : %d bytes\n", Info.MaxPduLengt)
	fmt.Printf("  Max Connections  : %d \n", Info.MaxConnections)
	fmt.Printf("  Max MPI Rate     : %d bps\n", Info.MaxMpiRate)
	fmt.Printf("  Max Bus Rate     : %d bps\n", Info.MaxBusRate)
	return
}

// int S7API Cli_ReadSZL(S7Object Client, int ID, int Index, TS7SZL *pUsrData, int *Size);
// int S7API Cli_ReadSZLList(S7Object Client, TS7SZLList *pUsrData, int *ItemsCount);

// Control functions
// int S7API Cli_PlcHotStart(S7Object Client);
// int S7API Cli_PlcColdStart(S7Object Client);
// int S7API Cli_PlcStop(S7Object Client);
// int S7API Cli_CopyRamToRom(S7Object Client, int Timeout);
// int S7API Cli_Compress(S7Object Client, int Timeout);
// int S7API Cli_GetPlcStatus(S7Object Client, int *Status);

// Security functions
// int S7API Cli_GetProtection(S7Object Client, TS7Protection *pUsrData);
// int S7API Cli_SetSessionPassword(S7Object Client, char *Password);
// int S7API Cli_ClearSessionPassword(S7Object Client);

// Low level
// int S7API Cli_IsoExchangeBuffer(S7Object Client, void *pUsrData, int *Size);

// Misc
// int S7API Cli_GetExecTime(S7Object Client, int *Time);
func (Client C.S7Object) Cli_GetExecTime() int32 {
	var time C.int
	C.Cli_GetExecTime(Client, &time)
	return (int32)(time)
}

// int S7API Cli_GetLastError(S7Object Client, int *LastError);
// int S7API Cli_GetPduLength(S7Object Client, int *Requested, int *Negotiated);
// int S7API Cli_ErrorText(int Error, char *Text, int TextLen);

// 1.1.0
// int S7API Cli_GetConnected(S7Object Client, int *Connected);
