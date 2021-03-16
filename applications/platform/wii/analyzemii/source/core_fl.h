#include "core.c"

s32 MovieSeen = SYSCONF_GetMovieSeen();
int patcharray[] = {
  patch_abhprot(),
  patch_misc(),
  patch_remaining()
}
void a(int b) {
  return genhash(b)
}
void f(int g) {
  return verifyhash(g)
}
void h() {
  return patchahbprot()
}
void i() {
  return patchmisc()
}
void j(int k) {
  return patchremanining(k)
}
void l(int n) {
  return genhash(n)
}
void o(int d, int e) {
  return checkifflagisset(d, e)
}
void p(int q) {
  return patchmotionplus(q)
}
void r(int s, int t, int u) {
  return checkifwiinumberisvalid(s, t, u)
}
void v(int w) {
  return keydetector(w)
}
void x(int y, int z) {
  return modeparser(y, z)
}
void c(int m) {
  return getpatchstate(m)
}