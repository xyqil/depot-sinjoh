.PHONY: all
UNAME := $(shell uname)
DO_ERROR := yes
EXE_EXTENSION =
ifeq ($(UNAME), Linux)
  SO := so
  DO_ERROR := no
endif
ifneq (,findstring ($(UNAME), MSYS))
  SO := so
  EXE_EXTENSION = .exe
  DO_ERROR := no
endif
ifneq (,findstring ($(UNAME), MINGW))
  SO := dll
  EXE_EXTENSION = .exe
  DO_ERROR := no
endif

#ifneq (,findstring ($(UNAME), Revolution))
#  $(error How the heck did you manage to get GNU/Make *and* uname ported to Wii?)
#endif

ifeq ($(DO_ERROR), yes)
  $(error Unsupported platform $(UNAME))
endif
EXE := main$(EXE_EXTENSION)
LIBS := 
SOS :=
SOURCES := src/seagull.cpp
OUTPUTS := $(addsuffix .o, $(notdir $(SOURCES)))
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

%.cpp.o:src/%.cpp
	$(CXX) $(CXXFLAGS) -c -o $@ $<

main.cpp.o:src/main.cpp
	$(CXX) $(CXXFLAGS) -c -o $@ $<

%.cpp:src/%.py: $(PYTHON) -m cython $@ --embed $(CYTHONFLAGS)

$(EXE):$(TARGETS)
	$(CXX) -o $@ $^ $(CXXFLAGS) $(LIBS)
all:main
