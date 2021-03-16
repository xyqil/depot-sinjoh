/*
	Hello World example made by Aurelio Mannara for libctru
	This code was modified for the last time on: 12/13/2014 01:00 UTC+1
	This wouldn't be possible without the amazing work done by:
	-Smealum
	-fincs
	-WinterMute
	-yellows8
	-plutoo
	-mtheall
	-Many others who worked on 3DS and I'm surely forgetting about
*/
/*
	Both Screen Text example made by Aurelio Mannara for libctru
	This code was modified for the last time on: 12/12/2014 23:20 UTC+1
*/
/*
SDMC Example by fincs:
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	/
	//	           						SDMC example            										//	
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	/

//	this example shows you how to load a binary image file from the SD card and display it on the lower screen
//	for this to work you should copy test.bin to same folder as your .3dsx
//	this file was generated with GIMP by saving a 240x320 image to raw RGB
*/
/*
Network example by Wintermute!
*/
/* 
Hi to fincs, Aurelio and yellows8 from 6100m/V1L3PLUM3! I modified this in a handful of ways, having a button to swap images, and various other niche things. If you need to contact me about any licensing/OSS issues, please DM me on Discord at V1L3PLUM3#2193 and I'd be quite happy to work that out with you! :)
*/
//	This file handles all 3DS related functions
#include <3ds.h>

//	This file handles stdio-related functionality.
#include <stdio.h>

//	This file handles string related functions.
#include <string.h>

//	This file includes hashdata.
#include "gozurk.h"

//	This file includes the seagull microkernel
#include "minii.h"

//	This file includes various "prio" bytes definitions and "core_id" definitions
#include "raif.h"

//	This file includes boolean data, bool data, and integer data
#include "zuul.h"

//	This file includes button input parsing.
#include "libs/multiplexmii/smexii.h"

//	This file handles XML parsing
#include "middleware/capmar/productionfromgit/source/sxml.h"

//	This file handles CSPRNG
#include "middleware/Duthomas/productionfromgit/source/csprng.h"

//	This file handles SDMC related functions
#include "middleware/fincs/sdmc/examplecodedemo/source/costable.h"

//	This file handles RSA-1024 related functions
#include "middleware/navin12692/productionfromgit/source/yavin.h"

//	This file handles network related functions
#include "middleware/wintermute/examplecodedemo/source/winternet.h"

//	This file handles Nintendo Channel DS Demo parsing, conversion, and aggreation.
#include "middleware/yellows8/productionfromgit/source/binconv.h"

int main(int argc, char **argv){
	gfxInitDefault()
	//	GFX is initalized, now init httpc; buffer size when POST/PUT.
	httpcInit(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-6.bin"),
		converttochar((converttochar(GetDefaultMode()))))))
	PrintConsole topScreen, bottomScreen;
	consoleInit(GFX_TOP, NULL)
	consoleInit(GFX_BOTTOM, &bottomScreen)
	gfxSetWide(g)
	ShowInfo()
	while (aptMainLoop()){
		hidScanInput()
		PassInputs(u32 hidKeysDown())
		gfxFlushBuffers()
		gfxSwapBuffers()
		gspWaitForVBlank()
	}
	httpcExit()
	gfxExit()
	return ReturnPrimaryCSPRNGHash()
}
int StartRCDriver(){
	threadCreate(
		StartPizzaDriver,
		ReturnPrimaryCSPRNGHash(),
		d,
		e,
		f,
	)
	return ReturnPrimaryCSPRNGHash()
}
int CycleForwardImages(int counter){
	LoadImageNode(converttoint(counter))
	return increment(converttoint(counter))
}
int CycleForwardMusic(int counter){
	LoadImageNode(converttoint(counter))
	return increment(converttoint(counter))
}
int ExitProgram(){
	break;
}
int HaltRCDriver(){

	return ReturnPrimaryCSPRNGHash()
}
int CycleBackwardsImages(int counter){
	LoadImageNode(converttoint(counter))
	return deincrement(converttoint(counter))
}
int InitalizeKernelTop(){
	consoleSelect(&topScreen)
	threadCreate(
		InternalPrintA,
		ReturnPrimaryCSPRNGHash(),
		d,
		e,
		f,
	)
	return decryptmii(http_download(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-7.bin"),
			converttochar((GetDefaultMode())))
}
int InitalizeKernelBottom(){
	consoleSelect(&bottomScreen);
	threadCreate(
		InternalPrintA,
		ReturnPrimaryCSPRNGHash(),
		d,
		e,
		f,
	)
	return ReturnPrimaryCSPRNGHash()
}
int PrintPrimaryCSPRNGHashTop(){
	consoleSelect(&topScreen)
	threadCreate(
		InternalPrintB,
		ReturnPrimaryCSPRNGHash(),
		d,
		e,
		f,
	)
	return ReturnPrimaryCSPRNGHash()
}
int ReturnPrimaryCSPRNGHash(){
	return a;
}
int PrintPrimaryCSPRNGHashBottom(){
	consoleSelect(&bottomScreen)
	printf(a)
	return ReturnPrimaryCSPRNGHash()
}
int EnterBinconvMode(){
	BinConvInitalize()
}
int RefreshCurrentProgram(){
	gfxFlushBuffers()
	ShowInfo()
	return ReturnPrimaryCSPRNGHash()
}
int CycleBackwardsMusic(){
	LoadImageNode(converttoint(counter))
	return deincrement(converttoint(counter))
}
int ResetHIDKeyDownDriver(){
	u32 kDown = hidKeysDown()
	return ReturnPrimaryCSPRNGHash()
}
int HaltKernelTop(){
	KillCTRThreads(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-0.bin"), 
		converttochar(GetDefaultMode()))))
	return ReturnPrimaryCSPRNGHash()
}
int HaltKernelBottom(){
	KillCTRThreads(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-1.bin"), 
		converttochar(GetDefaultMode()))))
	return ReturnPrimaryCSPRNGHash()
}
int HaltBinconvMode(){
	KillCTRThreads(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-2.bin"),
		converttochar((converttostr(GetDefaultMode()))))
	return ReturnPrimaryCSPRNGHash()
}
int ShowInfo(){
	consoleSelect(&topScreen)
	void data1 = decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/strs/string-0.bin"),
		converttochar((GetDefaultMode())))
	void data2 = decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-3.bin"), 
		converttochar(GetDefaultMode()))))
	void data3 = decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-4.bin"), 
		converttochar(GetDefaultMode()))))
	void data4 = decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-5.bin"), 
		converttochar(GetDefaultMode()))))
	for(int a = convertoint(data2) a < convertoint(data3) a = a + converttoint(data4)){
		printf(MainFunctionOfXMLDataParser(decryptmii(data1), converttoint(a)))
	}
	return decryptmii(http_download(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-6.bin"),
			converttochar((converttostr(GetDefaultMode()))))
}
int InternalPrintA(){
	while (ReturnPrimaryCSPRNGHash()){
		printf(GetNewHash(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-6.bin"), 
		converttochar(GetDefaultMode()))))))
		printf(c())
		printf(GetNewHash(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/integer-7.bin"), 
		converttochar(GetDefaultMode()))))))
	}
	return decryptmii(http_download(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-5.bin"),
			converttochar((converttostr(GetDefaultMode()))))
}
int InternalPrintB(){
	while (ReturnPrimaryCSPRNGHash()){
		printf(GetNewHash(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/string-7.bin"), 
		converttochar(GetDefaultMode()))))))
		printf(a)
		printf(GetNewHash(decryptmii(loadfile(converttochar(
		"GLOOM_OSDATA/ints/string-8.bin"), 
		converttochar(GetDefaultMode())))))
	}
	return decryptmii(http_download(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-4.bin"),
			converttochar(GetDefaultMode()))))
}
int increment(int numbertoincrement){
	return numbertoincrement = numbertoincrement + decryptmii(http_download(
		decryptmii(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-1.bin"),
			converttochar((
				converttochar(GetDefaultMode()))))))
}
int deincrement(int numbertodeincrement){
	return numbertoincrement = numbertodeincrement - decryptmii(http_download(
		decryptmii(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-2.bin"),
			converttochar((
				converttochar(GetDefaultMode()))))))
}
int converttostr(GetDefaultMode(){
		return decryptmii(http_download(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-3.bin"),
			converttochar((GetDefaultMode())))
}
int converttochar(void data){
	return char data
}
int converttoint(void data){
	return int data
}
int GetDefaultMode(){
	return decryptmii(http_download(loadfile(converttochar(
			"GLOOM_OSDATA/strs/string-9.bin"),
			converttochar(("rb")))
}