UNAME := $(shell uname)
DO_ERROR := yes
EXE_EXTENSION =
ifeq ($(UNAME), Linux)
  SO := so
  DO_ERROR := no
endif
ifneq (,findstring ($(UNAME), MSYS))
  SO := so
  DO_ERROR := no
endif
ifneq (,findstring ($(UNAME), MINGW))
  SO := dll
  DO_ERROR := no
endif

#ifneq (,findstring ($(UNAME), Revolution))
#  $(error How the heck did you manage to get GNU/Make *and* uname ported to Wii?)
#endif

ifeq ($(DO_ERROR), yes)
  $(error Unsupported platform $(UNAME))
endif
LIBS := 
SOS :=
SOURCES := core/seagull/source/main.cpp
OBJS := $(addsuffix .o, $(notdir $(SOURCES)))
CXXFLAGS :=
CXX := g++
GO := go
GOFLAGS :=
CYTHONFLAGS :=
ARFLAGS :=
PYTHON := python3
TARGETS := $(OBJS)
LD := ld
LDFLAGS :=

%.cpp.o:%.cpp
	$(CXX) $(CXXFLAGS) -c -o $@ $<

%.so: %.so.cpp.o ; $(LD) $(LDFLAGS) -shared $^ -o $@

%.a:%.a.cpp.o ; $(AR) rvs $< $@ $(ARFLAGS)

%.py.cpp:%.py2.py ; $(PYTHON) -m cython $@ --embed $(CYTHONFLAGS)

main:$(TARGETS)
	$(CXX) -o $@ $^ $(CXXFLAGS) $(LIBS)