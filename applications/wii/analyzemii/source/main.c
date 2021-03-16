#include "libwiisdk_libraries/core/core_fl.h"
int main(int argc, char ** argv) {
  VIDEO_Init();
  WPAD_Init();
  KEYBOARD_Init();
  GuiWindow * mainWindow = NULL;
}