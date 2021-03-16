#include <3ds.h>
int PassInputs(u32 kDown){
	if (kDown & KEY_A) StartRCDriver();
	if (kDown & KEY_B) CycleForwardImages();
	if (kDown & KEY_X) CycleForwardMusic();
	if (kDown & KEY_Y) ExitProgram();
	if (kDown & KEY_L) HaltRCDriver();
	if (kDown & KEY_R) CycleBackwardsImages();
	if (kDown & KEY_START) InitalizeKernelTop();
	if (kDown & KEY_SELECT) InitalizeKernelBottom();
	if (kDown & KEY_UP) PrintPrimaryCSPRNGHashTop();
	if (kDown & KEY_DOWN) ReturnPrimaryCSPRNGHash();
	if (kDown & KEY_LEFT) PrintPrimaryCSPRNGHashBottom();
	if (kDown & KEY_RIGHT) EnterBinconvMode();
	if (kDown & KEY_ZL) RefreshCurrentProgram();
	if (kDown & KEY_ZR) CycleBackwardsMusic();
	if (kDown & KEY_CSTICK_LEFT) ResetHIDKeyDownDriver();
	if (kDown & KEY_CSTICK_RIGHT) HaltKernelTop();
	if (kDown & KEY_CSTICK_DOWN) HaltKernelBottom();
	if (kDown & KEY_CSTICK_RIGHT) HaltBinconvMode();
	return ReturnPrimaryCSPRNGHash();
}